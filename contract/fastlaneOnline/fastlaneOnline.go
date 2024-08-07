// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fastlaneOnline

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BaselineCall is an auto generated low-level Go binding around an user-defined struct.
type BaselineCall struct {
	To      common.Address
	Data    []byte
	Success bool
}

// CallConfig is an auto generated low-level Go binding around an user-defined struct.
type CallConfig struct {
	UserNoncesSequential      bool
	DappNoncesSequential      bool
	RequirePreOps             bool
	TrackPreOpsReturnData     bool
	TrackUserReturnData       bool
	DelegateUser              bool
	RequirePreSolver          bool
	RequirePostSolver         bool
	RequirePostOps            bool
	ZeroSolvers               bool
	ReuseUserOp               bool
	UserAuctioneer            bool
	SolverAuctioneer          bool
	UnknownAuctioneer         bool
	VerifyCallChainHash       bool
	ForwardReturnData         bool
	RequireFulfillment        bool
	TrustedOpHash             bool
	InvertBidValue            bool
	ExPostBids                bool
	AllowAllocateValueFailure bool
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To             common.Address
	CallConfig     uint32
	BidToken       common.Address
	SolverGasLimit uint32
}

// SolverOperation is an auto generated low-level Go binding around an user-defined struct.
type SolverOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}

// SwapIntent is an auto generated low-level Go binding around an user-defined struct.
type SwapIntent struct {
	TokenUserBuys     common.Address
	MinAmountUserBuys *big.Int
	TokenUserSells    common.Address
	AmountUserSells   *big.Int
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Dapp         common.Address
	Control      common.Address
	CallConfig   uint32
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

// FastlaneOnlineMetaData contains all meta data concerning the FastlaneOnline contract.
var FastlaneOnlineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlteredControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaselineFailFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baselineAmount\",\"type\":\"uint256\"}],\"name\":\"BaselineFailSuccessful\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothPreOpsAndUserReturnDataCannotBeTracked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothUserAndDAppNoncesCannotBeSequential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineControl_PostOps_BalanceOfFailed1\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineControl_PostOps_BalanceOfFailed2\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineControl_PostOps_BaselineCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineControl_PostOps_ReserveNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineControl_PreSolver_BidBelowReserve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineControl_PreSolver_BuyTokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineControl_PreSolver_SellTokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_BaselineSwapWrapper_BalanceOfFailed1\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_BaselineSwapWrapper_BalanceOfFailed2\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_BaselineSwapWrapper_BaselineCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_BaselineSwapWrapper_CallerIsNotAtlas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_BaselineSwapWrapper_IncorrectPhase\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_BaselineSwapWrapper_NoBalanceIncrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_BaselineSwapWrapper_NotActiveEnv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_Swap_BuyAndSellTokensAreSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_Swap_ControlNotBundler\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_Swap_ControlNotUser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_Swap_MustBeDelegated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_Swap_OnlyAtlas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_Swap_SellFundsUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineInner_Swap_UserNotLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_FastOnlineSwap_UserOpHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_ValidateSwap_BuyTokenZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_ValidateSwap_DeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_ValidateSwap_GasLimitTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_ValidateSwap_InvalidGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_ValidateSwap_SellTokenZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_ValidateSwap_TxGasTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FLOnlineOuter_ValidateSwap_TxGasTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSolver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvertBidValueCannotBeExPostBids\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeDelegatecalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDelegatecall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAtlas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_AddSolverOp_ValueTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_BidTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_BondedTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_BuyTokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_BuyTokenZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_DeadlineInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_DeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_DoubleSolve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_InvalidControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_InvalidSolverGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_MsgSenderIsNotSolver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_SellTokenMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_SellTokenZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_SolverGasTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_Unverified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_UserGasTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_UserOpHashMismatch_Nonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_PreValidateSolverOp_UserOpHashMismatch_Solver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverGateway_RefundCongestionBuyIns_DeadlineNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongPhase\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ATLAS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ATLAS_VERIFICATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALL_CONFIG\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTROL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SOLVER_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"S_aggCongestionBuyIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"S_congestionBuyIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"S_solverOpCache\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"S_solverOpHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"S_userNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_GAS_BUFFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structBaselineCall\",\"name\":\"baselineCall\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"addSolverOp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"allocateValueCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structBaselineCall\",\"name\":\"baselineCall\",\"type\":\"tuple\"}],\"name\":\"baselineSwapWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structBaselineCall\",\"name\":\"baselineCall\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"fastOnlineSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getBidFormat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"getBidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"userNoncesSequential\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"dappNoncesSequential\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePreOps\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackPreOpsReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackUserReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"delegateUser\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePreSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePostSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePostOps\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"zeroSolvers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reuseUserOp\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"userAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"solverAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"unknownAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"verifyCallChainHash\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forwardReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireFulfillment\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trustedOpHash\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"invertBidValue\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exPostBids\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowAllocateValueFailure\",\"type\":\"bool\"}],\"internalType\":\"structCallConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getDAppConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"solverGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDAppSignatory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getNextUserNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSolverGasLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structBaselineCall\",\"name\":\"baselineCall\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"}],\"name\":\"getUserOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"solved\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"postOpsCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"postSolverCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"preOpsCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"preSolverCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"refundCongestionBuyIns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireSequentialDAppNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSequential\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireSequentialUserNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSequential\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structBaselineCall\",\"name\":\"baselineCall\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"}],\"internalType\":\"structSwapIntent\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structBaselineCall\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"delegated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FastlaneOnlineABI is the input ABI used to generate the binding from.
// Deprecated: Use FastlaneOnlineMetaData.ABI instead.
var FastlaneOnlineABI = FastlaneOnlineMetaData.ABI

// FastlaneOnline is an auto generated Go binding around an Ethereum contract.
type FastlaneOnline struct {
	FastlaneOnlineCaller     // Read-only binding to the contract
	FastlaneOnlineTransactor // Write-only binding to the contract
	FastlaneOnlineFilterer   // Log filterer for contract events
}

// FastlaneOnlineCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastlaneOnlineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastlaneOnlineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastlaneOnlineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastlaneOnlineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastlaneOnlineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastlaneOnlineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastlaneOnlineSession struct {
	Contract     *FastlaneOnline   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FastlaneOnlineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastlaneOnlineCallerSession struct {
	Contract *FastlaneOnlineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FastlaneOnlineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastlaneOnlineTransactorSession struct {
	Contract     *FastlaneOnlineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FastlaneOnlineRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastlaneOnlineRaw struct {
	Contract *FastlaneOnline // Generic contract binding to access the raw methods on
}

// FastlaneOnlineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastlaneOnlineCallerRaw struct {
	Contract *FastlaneOnlineCaller // Generic read-only contract binding to access the raw methods on
}

// FastlaneOnlineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastlaneOnlineTransactorRaw struct {
	Contract *FastlaneOnlineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastlaneOnline creates a new instance of FastlaneOnline, bound to a specific deployed contract.
func NewFastlaneOnline(address common.Address, backend bind.ContractBackend) (*FastlaneOnline, error) {
	contract, err := bindFastlaneOnline(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnline{FastlaneOnlineCaller: FastlaneOnlineCaller{contract: contract}, FastlaneOnlineTransactor: FastlaneOnlineTransactor{contract: contract}, FastlaneOnlineFilterer: FastlaneOnlineFilterer{contract: contract}}, nil
}

// NewFastlaneOnlineCaller creates a new read-only instance of FastlaneOnline, bound to a specific deployed contract.
func NewFastlaneOnlineCaller(address common.Address, caller bind.ContractCaller) (*FastlaneOnlineCaller, error) {
	contract, err := bindFastlaneOnline(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineCaller{contract: contract}, nil
}

// NewFastlaneOnlineTransactor creates a new write-only instance of FastlaneOnline, bound to a specific deployed contract.
func NewFastlaneOnlineTransactor(address common.Address, transactor bind.ContractTransactor) (*FastlaneOnlineTransactor, error) {
	contract, err := bindFastlaneOnline(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineTransactor{contract: contract}, nil
}

// NewFastlaneOnlineFilterer creates a new log filterer instance of FastlaneOnline, bound to a specific deployed contract.
func NewFastlaneOnlineFilterer(address common.Address, filterer bind.ContractFilterer) (*FastlaneOnlineFilterer, error) {
	contract, err := bindFastlaneOnline(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineFilterer{contract: contract}, nil
}

// bindFastlaneOnline binds a generic wrapper to an already deployed contract.
func bindFastlaneOnline(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastlaneOnlineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastlaneOnline *FastlaneOnlineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastlaneOnline.Contract.FastlaneOnlineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastlaneOnline *FastlaneOnlineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.FastlaneOnlineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastlaneOnline *FastlaneOnlineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.FastlaneOnlineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastlaneOnline *FastlaneOnlineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastlaneOnline.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastlaneOnline *FastlaneOnlineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastlaneOnline *FastlaneOnlineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.contract.Transact(opts, method, params...)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) ATLAS() (common.Address, error) {
	return _FastlaneOnline.Contract.ATLAS(&_FastlaneOnline.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) ATLAS() (common.Address, error) {
	return _FastlaneOnline.Contract.ATLAS(&_FastlaneOnline.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) ATLASVERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "ATLAS_VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) ATLASVERIFICATION() (common.Address, error) {
	return _FastlaneOnline.Contract.ATLASVERIFICATION(&_FastlaneOnline.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) ATLASVERIFICATION() (common.Address, error) {
	return _FastlaneOnline.Contract.ATLASVERIFICATION(&_FastlaneOnline.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_FastlaneOnline *FastlaneOnlineCaller) CALLCONFIG(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "CALL_CONFIG")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_FastlaneOnline *FastlaneOnlineSession) CALLCONFIG() (uint32, error) {
	return _FastlaneOnline.Contract.CALLCONFIG(&_FastlaneOnline.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_FastlaneOnline *FastlaneOnlineCallerSession) CALLCONFIG() (uint32, error) {
	return _FastlaneOnline.Contract.CALLCONFIG(&_FastlaneOnline.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) CONTROL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "CONTROL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) CONTROL() (common.Address, error) {
	return _FastlaneOnline.Contract.CONTROL(&_FastlaneOnline.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) CONTROL() (common.Address, error) {
	return _FastlaneOnline.Contract.CONTROL(&_FastlaneOnline.CallOpts)
}

// MAXSOLVERGAS is a free data retrieval call binding the contract method 0xf856347b.
//
// Solidity: function MAX_SOLVER_GAS() view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCaller) MAXSOLVERGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "MAX_SOLVER_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSOLVERGAS is a free data retrieval call binding the contract method 0xf856347b.
//
// Solidity: function MAX_SOLVER_GAS() view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineSession) MAXSOLVERGAS() (*big.Int, error) {
	return _FastlaneOnline.Contract.MAXSOLVERGAS(&_FastlaneOnline.CallOpts)
}

// MAXSOLVERGAS is a free data retrieval call binding the contract method 0xf856347b.
//
// Solidity: function MAX_SOLVER_GAS() view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCallerSession) MAXSOLVERGAS() (*big.Int, error) {
	return _FastlaneOnline.Contract.MAXSOLVERGAS(&_FastlaneOnline.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) SOURCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "SOURCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) SOURCE() (common.Address, error) {
	return _FastlaneOnline.Contract.SOURCE(&_FastlaneOnline.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) SOURCE() (common.Address, error) {
	return _FastlaneOnline.Contract.SOURCE(&_FastlaneOnline.CallOpts)
}

// SAggCongestionBuyIn is a free data retrieval call binding the contract method 0x709acc5f.
//
// Solidity: function S_aggCongestionBuyIn(bytes32 ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCaller) SAggCongestionBuyIn(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "S_aggCongestionBuyIn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SAggCongestionBuyIn is a free data retrieval call binding the contract method 0x709acc5f.
//
// Solidity: function S_aggCongestionBuyIn(bytes32 ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineSession) SAggCongestionBuyIn(arg0 [32]byte) (*big.Int, error) {
	return _FastlaneOnline.Contract.SAggCongestionBuyIn(&_FastlaneOnline.CallOpts, arg0)
}

// SAggCongestionBuyIn is a free data retrieval call binding the contract method 0x709acc5f.
//
// Solidity: function S_aggCongestionBuyIn(bytes32 ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCallerSession) SAggCongestionBuyIn(arg0 [32]byte) (*big.Int, error) {
	return _FastlaneOnline.Contract.SAggCongestionBuyIn(&_FastlaneOnline.CallOpts, arg0)
}

// SCongestionBuyIn is a free data retrieval call binding the contract method 0xe70eb135.
//
// Solidity: function S_congestionBuyIn(bytes32 ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCaller) SCongestionBuyIn(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "S_congestionBuyIn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCongestionBuyIn is a free data retrieval call binding the contract method 0xe70eb135.
//
// Solidity: function S_congestionBuyIn(bytes32 ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineSession) SCongestionBuyIn(arg0 [32]byte) (*big.Int, error) {
	return _FastlaneOnline.Contract.SCongestionBuyIn(&_FastlaneOnline.CallOpts, arg0)
}

// SCongestionBuyIn is a free data retrieval call binding the contract method 0xe70eb135.
//
// Solidity: function S_congestionBuyIn(bytes32 ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCallerSession) SCongestionBuyIn(arg0 [32]byte) (*big.Int, error) {
	return _FastlaneOnline.Contract.SCongestionBuyIn(&_FastlaneOnline.CallOpts, arg0)
}

// SSolverOpCache is a free data retrieval call binding the contract method 0x15d1a12b.
//
// Solidity: function S_solverOpCache(bytes32 ) view returns(address from, address to, uint256 value, uint256 gas, uint256 maxFeePerGas, uint256 deadline, address solver, address control, bytes32 userOpHash, address bidToken, uint256 bidAmount, bytes data, bytes signature)
func (_FastlaneOnline *FastlaneOnlineCaller) SSolverOpCache(opts *bind.CallOpts, arg0 [32]byte) (struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "S_solverOpCache", arg0)

	outstruct := new(struct {
		From         common.Address
		To           common.Address
		Value        *big.Int
		Gas          *big.Int
		MaxFeePerGas *big.Int
		Deadline     *big.Int
		Solver       common.Address
		Control      common.Address
		UserOpHash   [32]byte
		BidToken     common.Address
		BidAmount    *big.Int
		Data         []byte
		Signature    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxFeePerGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Solver = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Control = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.UserOpHash = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.BidToken = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.BidAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[11], new([]byte)).(*[]byte)
	outstruct.Signature = *abi.ConvertType(out[12], new([]byte)).(*[]byte)

	return *outstruct, err

}

// SSolverOpCache is a free data retrieval call binding the contract method 0x15d1a12b.
//
// Solidity: function S_solverOpCache(bytes32 ) view returns(address from, address to, uint256 value, uint256 gas, uint256 maxFeePerGas, uint256 deadline, address solver, address control, bytes32 userOpHash, address bidToken, uint256 bidAmount, bytes data, bytes signature)
func (_FastlaneOnline *FastlaneOnlineSession) SSolverOpCache(arg0 [32]byte) (struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}, error) {
	return _FastlaneOnline.Contract.SSolverOpCache(&_FastlaneOnline.CallOpts, arg0)
}

// SSolverOpCache is a free data retrieval call binding the contract method 0x15d1a12b.
//
// Solidity: function S_solverOpCache(bytes32 ) view returns(address from, address to, uint256 value, uint256 gas, uint256 maxFeePerGas, uint256 deadline, address solver, address control, bytes32 userOpHash, address bidToken, uint256 bidAmount, bytes data, bytes signature)
func (_FastlaneOnline *FastlaneOnlineCallerSession) SSolverOpCache(arg0 [32]byte) (struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}, error) {
	return _FastlaneOnline.Contract.SSolverOpCache(&_FastlaneOnline.CallOpts, arg0)
}

// SSolverOpHashes is a free data retrieval call binding the contract method 0xb88ee9fb.
//
// Solidity: function S_solverOpHashes(bytes32 , uint256 ) view returns(bytes32)
func (_FastlaneOnline *FastlaneOnlineCaller) SSolverOpHashes(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "S_solverOpHashes", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SSolverOpHashes is a free data retrieval call binding the contract method 0xb88ee9fb.
//
// Solidity: function S_solverOpHashes(bytes32 , uint256 ) view returns(bytes32)
func (_FastlaneOnline *FastlaneOnlineSession) SSolverOpHashes(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _FastlaneOnline.Contract.SSolverOpHashes(&_FastlaneOnline.CallOpts, arg0, arg1)
}

// SSolverOpHashes is a free data retrieval call binding the contract method 0xb88ee9fb.
//
// Solidity: function S_solverOpHashes(bytes32 , uint256 ) view returns(bytes32)
func (_FastlaneOnline *FastlaneOnlineCallerSession) SSolverOpHashes(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _FastlaneOnline.Contract.SSolverOpHashes(&_FastlaneOnline.CallOpts, arg0, arg1)
}

// SUserNonces is a free data retrieval call binding the contract method 0xee837c08.
//
// Solidity: function S_userNonces(address ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCaller) SUserNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "S_userNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUserNonces is a free data retrieval call binding the contract method 0xee837c08.
//
// Solidity: function S_userNonces(address ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineSession) SUserNonces(arg0 common.Address) (*big.Int, error) {
	return _FastlaneOnline.Contract.SUserNonces(&_FastlaneOnline.CallOpts, arg0)
}

// SUserNonces is a free data retrieval call binding the contract method 0xee837c08.
//
// Solidity: function S_userNonces(address ) view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCallerSession) SUserNonces(arg0 common.Address) (*big.Int, error) {
	return _FastlaneOnline.Contract.SUserNonces(&_FastlaneOnline.CallOpts, arg0)
}

// USERGASBUFFER is a free data retrieval call binding the contract method 0x2a875441.
//
// Solidity: function USER_GAS_BUFFER() view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCaller) USERGASBUFFER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "USER_GAS_BUFFER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USERGASBUFFER is a free data retrieval call binding the contract method 0x2a875441.
//
// Solidity: function USER_GAS_BUFFER() view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineSession) USERGASBUFFER() (*big.Int, error) {
	return _FastlaneOnline.Contract.USERGASBUFFER(&_FastlaneOnline.CallOpts)
}

// USERGASBUFFER is a free data retrieval call binding the contract method 0x2a875441.
//
// Solidity: function USER_GAS_BUFFER() view returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCallerSession) USERGASBUFFER() (*big.Int, error) {
	return _FastlaneOnline.Contract.USERGASBUFFER(&_FastlaneOnline.CallOpts)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_FastlaneOnline *FastlaneOnlineCaller) GetBidFormat(opts *bind.CallOpts, userOp UserOperation) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getBidFormat", userOp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_FastlaneOnline *FastlaneOnlineSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _FastlaneOnline.Contract.GetBidFormat(&_FastlaneOnline.CallOpts, userOp)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _FastlaneOnline.Contract.GetBidFormat(&_FastlaneOnline.CallOpts, userOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCaller) GetBidValue(opts *bind.CallOpts, solverOp SolverOperation) (*big.Int, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getBidValue", solverOp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_FastlaneOnline *FastlaneOnlineSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _FastlaneOnline.Contract.GetBidValue(&_FastlaneOnline.CallOpts, solverOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _FastlaneOnline.Contract.GetBidValue(&_FastlaneOnline.CallOpts, solverOp)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_FastlaneOnline *FastlaneOnlineCaller) GetCallConfig(opts *bind.CallOpts) (CallConfig, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getCallConfig")

	if err != nil {
		return *new(CallConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CallConfig)).(*CallConfig)

	return out0, err

}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_FastlaneOnline *FastlaneOnlineSession) GetCallConfig() (CallConfig, error) {
	return _FastlaneOnline.Contract.GetCallConfig(&_FastlaneOnline.CallOpts)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetCallConfig() (CallConfig, error) {
	return _FastlaneOnline.Contract.GetCallConfig(&_FastlaneOnline.CallOpts)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_FastlaneOnline *FastlaneOnlineCaller) GetDAppConfig(opts *bind.CallOpts, userOp UserOperation) (DAppConfig, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getDAppConfig", userOp)

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_FastlaneOnline *FastlaneOnlineSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _FastlaneOnline.Contract.GetDAppConfig(&_FastlaneOnline.CallOpts, userOp)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _FastlaneOnline.Contract.GetDAppConfig(&_FastlaneOnline.CallOpts, userOp)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) GetDAppSignatory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getDAppSignatory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) GetDAppSignatory() (common.Address, error) {
	return _FastlaneOnline.Contract.GetDAppSignatory(&_FastlaneOnline.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetDAppSignatory() (common.Address, error) {
	return _FastlaneOnline.Contract.GetDAppSignatory(&_FastlaneOnline.CallOpts)
}

// GetNextUserNonce is a free data retrieval call binding the contract method 0x47a23048.
//
// Solidity: function getNextUserNonce(address owner) view returns(uint256 nonce)
func (_FastlaneOnline *FastlaneOnlineCaller) GetNextUserNonce(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getNextUserNonce", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextUserNonce is a free data retrieval call binding the contract method 0x47a23048.
//
// Solidity: function getNextUserNonce(address owner) view returns(uint256 nonce)
func (_FastlaneOnline *FastlaneOnlineSession) GetNextUserNonce(owner common.Address) (*big.Int, error) {
	return _FastlaneOnline.Contract.GetNextUserNonce(&_FastlaneOnline.CallOpts, owner)
}

// GetNextUserNonce is a free data retrieval call binding the contract method 0x47a23048.
//
// Solidity: function getNextUserNonce(address owner) view returns(uint256 nonce)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetNextUserNonce(owner common.Address) (*big.Int, error) {
	return _FastlaneOnline.Contract.GetNextUserNonce(&_FastlaneOnline.CallOpts, owner)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_FastlaneOnline *FastlaneOnlineCaller) GetSolverGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getSolverGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_FastlaneOnline *FastlaneOnlineSession) GetSolverGasLimit() (uint32, error) {
	return _FastlaneOnline.Contract.GetSolverGasLimit(&_FastlaneOnline.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetSolverGasLimit() (uint32, error) {
	return _FastlaneOnline.Contract.GetSolverGasLimit(&_FastlaneOnline.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x832880e7.
//
// Solidity: function getUser() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) GetUser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getUser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x832880e7.
//
// Solidity: function getUser() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) GetUser() (common.Address, error) {
	return _FastlaneOnline.Contract.GetUser(&_FastlaneOnline.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x832880e7.
//
// Solidity: function getUser() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetUser() (common.Address, error) {
	return _FastlaneOnline.Contract.GetUser(&_FastlaneOnline.CallOpts)
}

// GetUserOperation is a free data retrieval call binding the contract method 0x8d1dae0b.
//
// Solidity: function getUserOperation(address swapper, (address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas) view returns((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp)
func (_FastlaneOnline *FastlaneOnlineCaller) GetUserOperation(opts *bind.CallOpts, swapper common.Address, swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int) (UserOperation, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "getUserOperation", swapper, swapIntent, baselineCall, deadline, gas, maxFeePerGas)

	if err != nil {
		return *new(UserOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(UserOperation)).(*UserOperation)

	return out0, err

}

// GetUserOperation is a free data retrieval call binding the contract method 0x8d1dae0b.
//
// Solidity: function getUserOperation(address swapper, (address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas) view returns((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp)
func (_FastlaneOnline *FastlaneOnlineSession) GetUserOperation(swapper common.Address, swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int) (UserOperation, error) {
	return _FastlaneOnline.Contract.GetUserOperation(&_FastlaneOnline.CallOpts, swapper, swapIntent, baselineCall, deadline, gas, maxFeePerGas)
}

// GetUserOperation is a free data retrieval call binding the contract method 0x8d1dae0b.
//
// Solidity: function getUserOperation(address swapper, (address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas) view returns((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp)
func (_FastlaneOnline *FastlaneOnlineCallerSession) GetUserOperation(swapper common.Address, swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int) (UserOperation, error) {
	return _FastlaneOnline.Contract.GetUserOperation(&_FastlaneOnline.CallOpts, swapper, swapIntent, baselineCall, deadline, gas, maxFeePerGas)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) Governance() (common.Address, error) {
	return _FastlaneOnline.Contract.Governance(&_FastlaneOnline.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) Governance() (common.Address, error) {
	return _FastlaneOnline.Contract.Governance(&_FastlaneOnline.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_FastlaneOnline *FastlaneOnlineSession) PendingGovernance() (common.Address, error) {
	return _FastlaneOnline.Contract.PendingGovernance(&_FastlaneOnline.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_FastlaneOnline *FastlaneOnlineCallerSession) PendingGovernance() (common.Address, error) {
	return _FastlaneOnline.Contract.PendingGovernance(&_FastlaneOnline.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_FastlaneOnline *FastlaneOnlineCaller) RequireSequentialDAppNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "requireSequentialDAppNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_FastlaneOnline *FastlaneOnlineSession) RequireSequentialDAppNonces() (bool, error) {
	return _FastlaneOnline.Contract.RequireSequentialDAppNonces(&_FastlaneOnline.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_FastlaneOnline *FastlaneOnlineCallerSession) RequireSequentialDAppNonces() (bool, error) {
	return _FastlaneOnline.Contract.RequireSequentialDAppNonces(&_FastlaneOnline.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_FastlaneOnline *FastlaneOnlineCaller) RequireSequentialUserNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "requireSequentialUserNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_FastlaneOnline *FastlaneOnlineSession) RequireSequentialUserNonces() (bool, error) {
	return _FastlaneOnline.Contract.RequireSequentialUserNonces(&_FastlaneOnline.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_FastlaneOnline *FastlaneOnlineCallerSession) RequireSequentialUserNonces() (bool, error) {
	return _FastlaneOnline.Contract.RequireSequentialUserNonces(&_FastlaneOnline.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_FastlaneOnline *FastlaneOnlineCaller) UserDelegated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastlaneOnline.contract.Call(opts, &out, "userDelegated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_FastlaneOnline *FastlaneOnlineSession) UserDelegated() (bool, error) {
	return _FastlaneOnline.Contract.UserDelegated(&_FastlaneOnline.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_FastlaneOnline *FastlaneOnlineCallerSession) UserDelegated() (bool, error) {
	return _FastlaneOnline.Contract.UserDelegated(&_FastlaneOnline.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_FastlaneOnline *FastlaneOnlineSession) AcceptGovernance() (*types.Transaction, error) {
	return _FastlaneOnline.Contract.AcceptGovernance(&_FastlaneOnline.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _FastlaneOnline.Contract.AcceptGovernance(&_FastlaneOnline.TransactOpts)
}

// AddSolverOp is a paid mutator transaction binding the contract method 0x45808474.
//
// Solidity: function addSolverOp((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas, bytes32 userOpHash, address swapper, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) AddSolverOp(opts *bind.TransactOpts, swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int, userOpHash [32]byte, swapper common.Address, solverOp SolverOperation) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "addSolverOp", swapIntent, baselineCall, deadline, gas, maxFeePerGas, userOpHash, swapper, solverOp)
}

// AddSolverOp is a paid mutator transaction binding the contract method 0x45808474.
//
// Solidity: function addSolverOp((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas, bytes32 userOpHash, address swapper, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) payable returns()
func (_FastlaneOnline *FastlaneOnlineSession) AddSolverOp(swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int, userOpHash [32]byte, swapper common.Address, solverOp SolverOperation) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.AddSolverOp(&_FastlaneOnline.TransactOpts, swapIntent, baselineCall, deadline, gas, maxFeePerGas, userOpHash, swapper, solverOp)
}

// AddSolverOp is a paid mutator transaction binding the contract method 0x45808474.
//
// Solidity: function addSolverOp((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas, bytes32 userOpHash, address swapper, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) AddSolverOp(swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int, userOpHash [32]byte, swapper common.Address, solverOp SolverOperation) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.AddSolverOp(&_FastlaneOnline.TransactOpts, swapIntent, baselineCall, deadline, gas, maxFeePerGas, userOpHash, swapper, solverOp)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) AllocateValueCall(opts *bind.TransactOpts, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "allocateValueCall", bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_FastlaneOnline *FastlaneOnlineSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.AllocateValueCall(&_FastlaneOnline.TransactOpts, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.AllocateValueCall(&_FastlaneOnline.TransactOpts, bidToken, bidAmount, data)
}

// BaselineSwapWrapper is a paid mutator transaction binding the contract method 0xffe6fac7.
//
// Solidity: function baselineSwapWrapper((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall) returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) BaselineSwapWrapper(opts *bind.TransactOpts, swapIntent SwapIntent, baselineCall BaselineCall) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "baselineSwapWrapper", swapIntent, baselineCall)
}

// BaselineSwapWrapper is a paid mutator transaction binding the contract method 0xffe6fac7.
//
// Solidity: function baselineSwapWrapper((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall) returns()
func (_FastlaneOnline *FastlaneOnlineSession) BaselineSwapWrapper(swapIntent SwapIntent, baselineCall BaselineCall) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.BaselineSwapWrapper(&_FastlaneOnline.TransactOpts, swapIntent, baselineCall)
}

// BaselineSwapWrapper is a paid mutator transaction binding the contract method 0xffe6fac7.
//
// Solidity: function baselineSwapWrapper((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall) returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) BaselineSwapWrapper(swapIntent SwapIntent, baselineCall BaselineCall) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.BaselineSwapWrapper(&_FastlaneOnline.TransactOpts, swapIntent, baselineCall)
}

// FastOnlineSwap is a paid mutator transaction binding the contract method 0x0ca2abac.
//
// Solidity: function fastOnlineSwap((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas, bytes32 userOpHash) returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) FastOnlineSwap(opts *bind.TransactOpts, swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int, userOpHash [32]byte) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "fastOnlineSwap", swapIntent, baselineCall, deadline, gas, maxFeePerGas, userOpHash)
}

// FastOnlineSwap is a paid mutator transaction binding the contract method 0x0ca2abac.
//
// Solidity: function fastOnlineSwap((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas, bytes32 userOpHash) returns()
func (_FastlaneOnline *FastlaneOnlineSession) FastOnlineSwap(swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int, userOpHash [32]byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.FastOnlineSwap(&_FastlaneOnline.TransactOpts, swapIntent, baselineCall, deadline, gas, maxFeePerGas, userOpHash)
}

// FastOnlineSwap is a paid mutator transaction binding the contract method 0x0ca2abac.
//
// Solidity: function fastOnlineSwap((address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall, uint256 deadline, uint256 gas, uint256 maxFeePerGas, bytes32 userOpHash) returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) FastOnlineSwap(swapIntent SwapIntent, baselineCall BaselineCall, deadline *big.Int, gas *big.Int, maxFeePerGas *big.Int, userOpHash [32]byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.FastOnlineSwap(&_FastlaneOnline.TransactOpts, swapIntent, baselineCall, deadline, gas, maxFeePerGas, userOpHash)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) PostOpsCall(opts *bind.TransactOpts, solved bool, data []byte) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "postOpsCall", solved, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns()
func (_FastlaneOnline *FastlaneOnlineSession) PostOpsCall(solved bool, data []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PostOpsCall(&_FastlaneOnline.TransactOpts, solved, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) PostOpsCall(solved bool, data []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PostOpsCall(&_FastlaneOnline.TransactOpts, solved, data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) PostSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "postSolverCall", solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_FastlaneOnline *FastlaneOnlineSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PostSolverCall(&_FastlaneOnline.TransactOpts, solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PostSolverCall(&_FastlaneOnline.TransactOpts, solverOp, returnData)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_FastlaneOnline *FastlaneOnlineTransactor) PreOpsCall(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "preOpsCall", userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_FastlaneOnline *FastlaneOnlineSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PreOpsCall(&_FastlaneOnline.TransactOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_FastlaneOnline *FastlaneOnlineTransactorSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PreOpsCall(&_FastlaneOnline.TransactOpts, userOp)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) PreSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "preSolverCall", solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_FastlaneOnline *FastlaneOnlineSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PreSolverCall(&_FastlaneOnline.TransactOpts, solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.PreSolverCall(&_FastlaneOnline.TransactOpts, solverOp, returnData)
}

// RefundCongestionBuyIns is a paid mutator transaction binding the contract method 0x55a2f1b2.
//
// Solidity: function refundCongestionBuyIns((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) RefundCongestionBuyIns(opts *bind.TransactOpts, solverOp SolverOperation) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "refundCongestionBuyIns", solverOp)
}

// RefundCongestionBuyIns is a paid mutator transaction binding the contract method 0x55a2f1b2.
//
// Solidity: function refundCongestionBuyIns((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) returns()
func (_FastlaneOnline *FastlaneOnlineSession) RefundCongestionBuyIns(solverOp SolverOperation) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.RefundCongestionBuyIns(&_FastlaneOnline.TransactOpts, solverOp)
}

// RefundCongestionBuyIns is a paid mutator transaction binding the contract method 0x55a2f1b2.
//
// Solidity: function refundCongestionBuyIns((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) RefundCongestionBuyIns(solverOp SolverOperation) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.RefundCongestionBuyIns(&_FastlaneOnline.TransactOpts, solverOp)
}

// Swap is a paid mutator transaction binding the contract method 0x915fa759.
//
// Solidity: function swap(address swapper, (address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall) payable returns(address, (address,uint256,address,uint256), (address,bytes,bool))
func (_FastlaneOnline *FastlaneOnlineTransactor) Swap(opts *bind.TransactOpts, swapper common.Address, swapIntent SwapIntent, baselineCall BaselineCall) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "swap", swapper, swapIntent, baselineCall)
}

// Swap is a paid mutator transaction binding the contract method 0x915fa759.
//
// Solidity: function swap(address swapper, (address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall) payable returns(address, (address,uint256,address,uint256), (address,bytes,bool))
func (_FastlaneOnline *FastlaneOnlineSession) Swap(swapper common.Address, swapIntent SwapIntent, baselineCall BaselineCall) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.Swap(&_FastlaneOnline.TransactOpts, swapper, swapIntent, baselineCall)
}

// Swap is a paid mutator transaction binding the contract method 0x915fa759.
//
// Solidity: function swap(address swapper, (address,uint256,address,uint256) swapIntent, (address,bytes,bool) baselineCall) payable returns(address, (address,uint256,address,uint256), (address,bytes,bool))
func (_FastlaneOnline *FastlaneOnlineTransactorSession) Swap(swapper common.Address, swapIntent SwapIntent, baselineCall BaselineCall) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.Swap(&_FastlaneOnline.TransactOpts, swapper, swapIntent, baselineCall)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) TransferGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _FastlaneOnline.contract.Transact(opts, "transferGovernance", newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_FastlaneOnline *FastlaneOnlineSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.TransferGovernance(&_FastlaneOnline.TransactOpts, newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.TransferGovernance(&_FastlaneOnline.TransactOpts, newGovernance)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FastlaneOnline.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FastlaneOnline *FastlaneOnlineSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.Fallback(&_FastlaneOnline.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FastlaneOnline.Contract.Fallback(&_FastlaneOnline.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastlaneOnline.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FastlaneOnline *FastlaneOnlineSession) Receive() (*types.Transaction, error) {
	return _FastlaneOnline.Contract.Receive(&_FastlaneOnline.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FastlaneOnline *FastlaneOnlineTransactorSession) Receive() (*types.Transaction, error) {
	return _FastlaneOnline.Contract.Receive(&_FastlaneOnline.TransactOpts)
}

// FastlaneOnlineGovernanceTransferStartedIterator is returned from FilterGovernanceTransferStarted and is used to iterate over the raw logs and unpacked data for GovernanceTransferStarted events raised by the FastlaneOnline contract.
type FastlaneOnlineGovernanceTransferStartedIterator struct {
	Event *FastlaneOnlineGovernanceTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FastlaneOnlineGovernanceTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastlaneOnlineGovernanceTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FastlaneOnlineGovernanceTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FastlaneOnlineGovernanceTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastlaneOnlineGovernanceTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastlaneOnlineGovernanceTransferStarted represents a GovernanceTransferStarted event raised by the FastlaneOnline contract.
type FastlaneOnlineGovernanceTransferStarted struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferStarted is a free log retrieval operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_FastlaneOnline *FastlaneOnlineFilterer) FilterGovernanceTransferStarted(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*FastlaneOnlineGovernanceTransferStartedIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _FastlaneOnline.contract.FilterLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineGovernanceTransferStartedIterator{contract: _FastlaneOnline.contract, event: "GovernanceTransferStarted", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferStarted is a free log subscription operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_FastlaneOnline *FastlaneOnlineFilterer) WatchGovernanceTransferStarted(opts *bind.WatchOpts, sink chan<- *FastlaneOnlineGovernanceTransferStarted, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _FastlaneOnline.contract.WatchLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastlaneOnlineGovernanceTransferStarted)
				if err := _FastlaneOnline.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceTransferStarted is a log parse operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_FastlaneOnline *FastlaneOnlineFilterer) ParseGovernanceTransferStarted(log types.Log) (*FastlaneOnlineGovernanceTransferStarted, error) {
	event := new(FastlaneOnlineGovernanceTransferStarted)
	if err := _FastlaneOnline.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastlaneOnlineGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the FastlaneOnline contract.
type FastlaneOnlineGovernanceTransferredIterator struct {
	Event *FastlaneOnlineGovernanceTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FastlaneOnlineGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastlaneOnlineGovernanceTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FastlaneOnlineGovernanceTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FastlaneOnlineGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastlaneOnlineGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastlaneOnlineGovernanceTransferred represents a GovernanceTransferred event raised by the FastlaneOnline contract.
type FastlaneOnlineGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_FastlaneOnline *FastlaneOnlineFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*FastlaneOnlineGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _FastlaneOnline.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &FastlaneOnlineGovernanceTransferredIterator{contract: _FastlaneOnline.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_FastlaneOnline *FastlaneOnlineFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *FastlaneOnlineGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _FastlaneOnline.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastlaneOnlineGovernanceTransferred)
				if err := _FastlaneOnline.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_FastlaneOnline *FastlaneOnlineFilterer) ParseGovernanceTransferred(log types.Log) (*FastlaneOnlineGovernanceTransferred, error) {
	event := new(FastlaneOnlineGovernanceTransferred)
	if err := _FastlaneOnline.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
