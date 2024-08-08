package bot

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	opsRelayUtils "github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/FastLane-Labs/fastlane-online-solver/config"
	"github.com/FastLane-Labs/fastlane-online-solver/contract/fastlaneOnline"
	"github.com/FastLane-Labs/fastlane-online-solver/contract/fastlaneOnlineSolver"
	"github.com/FastLane-Labs/fastlane-online-solver/events"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/FastLane-Labs/fastlane-online-solver/pools"
	"github.com/FastLane-Labs/fastlane-online-solver/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	FastlaneOnlineAbi, _       = fastlaneOnline.FastlaneOnlineMetaData.GetAbi()
	FastlaneOnlineSolverAbi, _ = fastlaneOnlineSolver.FastlaneOnlineSolverMetaData.GetAbi()
)

type Bot struct {
	config *config.Config

	ethClientConn *events.EthClientConnection
	mempool       events.Mempool

	poolsStore *pools.PoolsStore

	chainId                      *big.Int
	domain                       *apitypes.TypedDataDomain
	solverGasLimit               *big.Int
	fastlaneOnlineContract       *fastlaneOnline.FastlaneOnline
	fastlaneOnlineSolverContract *fastlaneOnlineSolver.FastlaneOnlineSolver
}

func NewBot(conf *config.Config) (*Bot, error) {
	ethClientConn, err := events.NewEthClientConnection(conf)
	if err != nil {
		return nil, err
	}

	chainId, err := ethClientConn.Client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain id: %v", err)
	}

	poolsStore, err := pools.NewPoolsStoreFromBinFiles(conf)
	if err != nil {
		return nil, err
	}

	dAppControlContract, err := fastlaneOnline.NewFastlaneOnline(common.HexToAddress(conf.FastlaneOnlineAddress), ethClientConn.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create dApp control contract: %v", err)
	}

	solverGasLimit, err := dAppControlContract.MAXSOLVERGAS(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get solver gas limit: %v", err)
	}
	solverGasLimit.Sub(solverGasLimit, big.NewInt(1))

	fastlaneOnlineSolverContract, err := fastlaneOnlineSolver.NewFastlaneOnlineSolver(common.HexToAddress(conf.SolverContractAddress), ethClientConn.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create solver contract: %v", err)
	}

	if config.SOLVER_PK == nil {
		return nil, fmt.Errorf("solver private key not set")
	}

	return &Bot{
		config:        conf,
		ethClientConn: ethClientConn,
		mempool:       ethClientConn, //mempool conn is same as ethClientConn

		poolsStore: poolsStore,

		chainId:                      chainId,
		solverGasLimit:               solverGasLimit,
		domain:                       utils.Domain(conf, chainId.Int64()),
		fastlaneOnlineContract:       dAppControlContract,
		fastlaneOnlineSolverContract: fastlaneOnlineSolverContract,
	}, nil
}

func (b *Bot) Run() {
	t := time.Now()
	log.Info("initializing dynamic state of pools")
	b.poolsStore.InitializeDynamicState(b.ethClientConn.Client, b.config.MaxConcurrentRpcCalls)
	log.Info("initialized dynamic state of pools", "took", time.Since(t))

	go func() {
		for {
			select {
			case block := <-b.ethClientConn.BlocksChan():
				b.handleBlock(block)
			case aLog := <-b.ethClientConn.LogsChan():
				b.handleLog(aLog)
			case tx := <-b.mempool.MempoolTxChan():
				b.handlePendingTx(tx)
			}
		}
	}()

	err := b.ethClientConn.Start()
	if err != nil {
		panic(err)
	}

	err = b.mempool.Start()
	if err != nil {
		panic(err)
	}
}

func (b *Bot) handleBlock(block *types.Header) {
	log.Info("block update", "block", block.Number)
}

func (b *Bot) handleLog(aLog *types.Log) {
	b.poolsStore.HandleLog(aLog)
}

func (b *Bot) handlePendingTx(tx *types.Transaction) {
	if *tx.To() != common.HexToAddress(b.config.FastlaneOnlineAddress) {
		return
	}
	if len(tx.Data()) < 4 {
		return
	}
	userMethodStr := "fastOnlineSwap"
	userMethod, exists := FastlaneOnlineAbi.Methods[userMethodStr]
	if !exists {
		panic("method not found in fastlane online abi - " + userMethodStr)
	}

	if !bytes.Equal(userMethod.ID, tx.Data()[:4]) {
		return
	}

	rawTx, err := tx.MarshalJSON()
	if err != nil {
		log.Error("failed to marshal tx to json", "err", err)
		return
	}

	log.Info("detected userOperation tx in mempool", "hash", tx.Hash().Hex(), "rawTx", hex.EncodeToString(rawTx))

	decodedUserInput, err := userMethod.Inputs.UnpackValues(tx.Data()[4:])
	if err != nil {
		log.Error("failed to unpack user method", "err", err)
		return
	}

	if len(decodedUserInput) != 1 {
		log.Error("unexpected decoded user input length", "len", len(decodedUserInput))
		return
	}

	userOperation, ok := decodedUserInput[0].(struct {
		From         common.Address "json:\"from\""
		To           common.Address "json:\"to\""
		Value        *big.Int       "json:\"value\""
		Gas          *big.Int       "json:\"gas\""
		MaxFeePerGas *big.Int       "json:\"maxFeePerGas\""
		Nonce        *big.Int       "json:\"nonce\""
		Deadline     *big.Int       "json:\"deadline\""
		Dapp         common.Address "json:\"dapp\""
		Control      common.Address "json:\"control\""
		CallConfig   uint32         "json:\"callConfig\""
		SessionKey   common.Address "json:\"sessionKey\""
		Data         []uint8        "json:\"data\""
		Signature    []uint8        "json:\"signature\""
	})
	if !ok {
		log.Error("failed to cast user operation, type of decodedUserInput[0] is not fastlaneOnline.UserOperation",
			"decodedUserInput[0]", decodedUserInput[0], "type",
			fmt.Sprintf("%T", decodedUserInput[0]),
		)
		return
	}

	userOperationDataMethodStr := "swap"
	userOperationDataMethod, exists := FastlaneOnlineAbi.Methods[userOperationDataMethodStr]
	if !exists {
		panic("method not found in user operation abi - " + userOperationDataMethodStr)
	}

	decodedUserOperationData, err := userOperationDataMethod.Inputs.UnpackValues(userOperation.Data[4:])
	if err != nil {
		log.Error("failed to unpack user operation data", "err", err)
		return
	}

	if len(decodedUserOperationData) != 2 {
		log.Error("unexpected decoded user operation data length", "len", len(decodedUserOperationData))
		return
	}

	swapIntent, ok := decodedUserOperationData[0].(struct {
		TokenUserBuys     common.Address "json:\"tokenUserBuys\""
		MinAmountUserBuys *big.Int       "json:\"minAmountUserBuys\""
		TokenUserSells    common.Address "json:\"tokenUserSells\""
		AmountUserSells   *big.Int       "json:\"amountUserSells\""
	})
	if !ok {
		log.Error("failed to cast swap intent, type of decodedUserOperationData[0] is not fastlaneOnline.SwapIntent",
			"decodedUserOperationData[0]", decodedUserOperationData[0],
			"type", fmt.Sprintf("%T", decodedUserOperationData[0]),
		)
		return
	}

	effectiveTokenUserSells := swapIntent.TokenUserSells
	if effectiveTokenUserSells == (common.Address{}) {
		effectiveTokenUserSells = common.HexToAddress(b.config.WethAddress)
	}
	effectiveTokenUserBuys := swapIntent.TokenUserBuys
	if effectiveTokenUserBuys == (common.Address{}) {
		effectiveTokenUserBuys = common.HexToAddress(b.config.WethAddress)
	}

	log.Info("detected swap intent in mempool",
		"tokenUserBuys", swapIntent.TokenUserBuys,
		"tokenUserSells", swapIntent.TokenUserSells,
		"amountUserSells", swapIntent.AmountUserSells,
		"minAmountUserBuys", swapIntent.MinAmountUserBuys,
	)

	bestPath, err := b.poolsStore.BestPath(effectiveTokenUserSells, effectiveTokenUserBuys, swapIntent.AmountUserSells)
	if err != nil {
		log.Error("error getting best path", "err", err)
		return
	}

	if bestPath.Output.Cmp(swapIntent.MinAmountUserBuys) < 0 {
		log.Info("best path output less than min amount user buys",
			"bestPathOutput", bestPath.Output,
			"minAmountUserBuys", swapIntent.MinAmountUserBuys,
		)
		return
	}

	solverProfit := big.NewInt(0).Sub(bestPath.Output, swapIntent.MinAmountUserBuys)
	solverProfit.Mul(solverProfit, big.NewInt(b.config.ProfitMarginx10000))
	solverProfit.Div(solverProfit, big.NewInt(10000))
	if solverProfit.Cmp(big.NewInt(0)) == 0 {
		log.Info("solver profit is zero, not sending solver op")
		return
	}
	bidAmount := big.NewInt(0).Sub(bestPath.Output, solverProfit)

	log.Info("best path found",
		"output", bestPath.Output,
		"minAmountUserBuys", swapIntent.MinAmountUserBuys,
		"solverProfit", solverProfit,
		"bidAmount", bidAmount,
		"tokenUserBuys", swapIntent.TokenUserBuys,
	)

	solverOpMethodStr := "execute"
	if _, exists := FastlaneOnlineSolverAbi.Methods[solverOpMethodStr]; !exists {
		panic("method not found in fastlane online solver abi - " + solverOpMethodStr)
	}

	fastlaneOnlineSolverSwaps := make([]fastlaneOnlineSolver.Swap, len(bestPath.Swaps))
	for i, swap := range bestPath.Swaps {
		fastlaneOnlineSolverSwaps[i] = fastlaneOnlineSolver.Swap{
			DexType:  uint8(swap.Pool.PoolType()),
			PoolAddr: swap.Pool.Address(),
			TokenIn:  swap.FromToken,
			TokenOut: swap.ToToken,
		}
	}

	solverOpData, err := FastlaneOnlineSolverAbi.Pack(solverOpMethodStr,
		fastlaneOnlineSolverSwaps,
		swapIntent.TokenUserSells,
		swapIntent.TokenUserBuys,
		swapIntent.TokenUserBuys,
		bidAmount,
	)
	if err != nil {
		log.Error("failed to pack solver op data", "err", err)
		return
	}

	userOpHash, err := utils.UserOpHash(userOperation, b.domain)
	if err != nil {
		log.Error("error hashing user op", "err", err)
		return
	}

	solverOp := &operation.SolverOperation{
		From:         config.SOLVER_EOA,
		To:           common.HexToAddress(b.config.AtlasAddress),
		Value:        big.NewInt(0),
		Gas:          b.solverGasLimit,
		MaxFeePerGas: userOperation.MaxFeePerGas,
		Deadline:     userOperation.Deadline,
		Solver:       common.HexToAddress(b.config.SolverContractAddress),
		Control:      common.HexToAddress(b.config.FastlaneOnlineAddress),
		UserOpHash:   userOpHash,
		BidToken:     swapIntent.TokenUserBuys,
		BidAmount:    bidAmount,
		Data:         solverOpData,
		Signature:    nil,
	}
	proofHash, relayErr := solverOp.Hash(b.domain)
	if relayErr != nil {
		log.Error("error hashing solver op", "err", err)
		return
	}
	sig, err := opsRelayUtils.SignMessage(proofHash.Bytes(), config.SOLVER_PK)
	if err != nil {
		log.Error("error signing solver op", "err", err)
		return
	}
	solverOp.Signature = sig

	frontrunTxOps := &bind.TransactOpts{
		From:  config.SOLVER_EOA,
		Nonce: nil,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.LatestSignerForChainID(b.chainId), config.SOLVER_PK)
		},
		Value:    big.NewInt(0),
		GasLimit: uint64(1000_000),
	}
	gasMargin := big.NewInt(0).Mul(big.NewInt(b.config.FrontrunGasMarginGwei), big.NewInt(1e9))
	if tx.Type() == types.LegacyTxType {
		frontrunTxOps.GasPrice = big.NewInt(0).Add(tx.GasPrice(), gasMargin)
	} else if tx.Type() == types.DynamicFeeTxType {
		frontrunTxOps.GasFeeCap = tx.GasFeeCap()
		frontrunTxOps.GasTipCap = big.NewInt(0).Add(tx.GasTipCap(), gasMargin)
	} else {
		log.Error("unknown user tx type", "txType", tx.Type())
		return
	}

	broadcastTx, err := b.fastlaneOnlineContract.AddSolverOp(
		frontrunTxOps,
		fastlaneOnline.UserOperation{
			From:         userOperation.From,
			To:           userOperation.To,
			Value:        userOperation.Value,
			Gas:          userOperation.Gas,
			MaxFeePerGas: userOperation.MaxFeePerGas,
			Nonce:        userOperation.Nonce,
			Deadline:     userOperation.Deadline,
			Dapp:         userOperation.Dapp,
			Control:      userOperation.Control,
			CallConfig:   userOperation.CallConfig,
			SessionKey:   userOperation.SessionKey,
			Data:         userOperation.Data,
			Signature:    userOperation.Signature,
		},
		fastlaneOnline.SolverOperation{
			From:         solverOp.From,
			To:           solverOp.To,
			Value:        solverOp.Value,
			Gas:          solverOp.Gas,
			MaxFeePerGas: solverOp.MaxFeePerGas,
			Deadline:     solverOp.Deadline,
			Solver:       solverOp.Solver,
			Control:      solverOp.Control,
			UserOpHash:   solverOp.UserOpHash,
			BidToken:     solverOp.BidToken,
			BidAmount:    solverOp.BidAmount,
			Data:         solverOp.Data,
			Signature:    solverOp.Signature,
		},
	)
	if err != nil {
		log.Error("error adding solver op", "err", err)
		return
	}

	log.Info("sent solver op tx", "tx", broadcastTx.Hash().Hex())
}
