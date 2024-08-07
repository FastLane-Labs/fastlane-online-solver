package pools

import (
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/contract/multicall"
	"github.com/FastLane-Labs/fastlane-online-solver/contract/uniswapV2Factory"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/FastLane-Labs/fastlane-online-solver/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GeneratePoolsStaticData_UniswapV2 generates static data for Uniswap V2 pools.
//
// This function connects to the Uniswap V2 factory contract and retrieves information about pools
// that include the specified tokens.
// The static data includes the pool address and token addresses.
// Reserves are also fethed for filtering zero reserves pools.
//
// Arguments:
//   - client: the *ethclient.Client instance
//   - factoryAddr: the address of the Uniswap V2 factory
//   - interestingTokens: a list of token addresses (from config) that are of interest for generating pool data
//   - maxConcurrentRpcCalls: the maximum number of concurrent RPC calls to use (to manage rate limits)
//
// Returns:
//   - []*UniswapV2Pool: a slice of Uniswap V2 pool objects containing static data about the pools
//   - error: if any error occurs during the data retrieval
func GeneratePoolsStaticData_UniswapV2(
	client *ethclient.Client,
	factoryAddr common.Address,
	interestingTokens []common.Address,
	maxConcurrentRpcCalls int,
) ([]*UniswapV2Pool, error) {

	log.Info("fetching uniswapV2 pools static data...", "factory", factoryAddr.Hex())
	fetchStartTime := time.Now()

	numPairs, err := numUniswapV2Pools(client, factoryAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get num of uniswap pools: %v", err)
	}
	log.Info("fetched number of uniswap pools",
		"numPools", numPairs,
		"factory", factoryAddr.Hex(),
		"elapsed", time.Since(fetchStartTime),
	)
	pools := make([]*UniswapV2Pool, numPairs.Int64())

	if err := fillUniswapV2Address(client, factoryAddr, pools, maxConcurrentRpcCalls); err != nil {
		return nil, fmt.Errorf("failed to fill uniswap pools addresses: %v", err)
	}
	log.Info("fetched addrs of uniswap pools",
		"numPools", len(pools),
		"factory", factoryAddr.Hex(),
		"elapsed", time.Since(fetchStartTime),
	)
	if err := fillUniswapV2StaticData(client, pools, maxConcurrentRpcCalls); err != nil {
		return nil, fmt.Errorf("failed to multicall: %v", err)
	}
	log.Info("fetched uniswap pools",
		"numPools", len(pools),
		"factory", factoryAddr.Hex(),
		"elapsed", time.Since(fetchStartTime),
	)

	filteredPools := []*UniswapV2Pool{}
	for _, pool := range pools {
		if pool.reserveA.Cmp(big.NewInt(0)) > 0 &&
			pool.reserveB.Cmp(big.NewInt(0)) > 0 &&
			slices.Contains(interestingTokens, pool.tokenA) &&
			slices.Contains(interestingTokens, pool.tokenB) {
			filteredPools = append(filteredPools, pool)
		}
	}

	log.Info("uniswap v2 pools static data fetched",
		"numPools", len(filteredPools),
		"factory", factoryAddr.Hex(),
		"took", time.Since(fetchStartTime),
	)

	return filteredPools, nil
}

// GeneratePoolsDynamicData_UniswapV2 generates dynamic data for Uniswap V2 pools.
//
// This function fetches real-time data for the provided Uniswap V2 pools using an *ethclient.Client.
// The dynamic data for UNiswapV2 means reserves.
//
// Arguments:
//   - client: the *ethclient.Client instance
//   - pools: a slice of Uniswap V2 pool objects for which to fetch reserves
//   - maxConcurrentRpcCalls: the maximum number of concurrent RPC calls to use (to manage rate limits)
//
// Returns:
//   - error: if any error occurs during the data retrieval
func GeneratePoolsDynamicData_UniswapV2(
	client *ethclient.Client,
	pools []*UniswapV2Pool,
	maxConcurrentRpcCalls int,
) error {
	log.Info("fetching uniswapV2 pools dynamic data ...", "numPools", len(pools))

	startTime := time.Now()

	indices := make([]int, 0)
	for i := 0; i < len(pools); i++ {
		indices = append(indices, i)
	}

	calldataBatchGenerator := func(index int) ([]multicall.Multicall3Call, error) {
		calldataReserves, err := UniswapV2PairAbi.Pack("getReserves")
		if err != nil {
			return nil, fmt.Errorf("failed to pack getReserves: %v", err)
		}
		return []multicall.Multicall3Call{
			{
				Target:   pools[index].address,
				CallData: calldataReserves,
			},
		}, nil
	}

	returnDataBatchHandler := func(index int, returnDatas [][]byte) error {
		if len(returnDatas) != 1 {
			return fmt.Errorf("unexpected return data length: %d, expected 1", len(returnDatas))
		}
		reserves, err := UniswapV2PairAbi.Unpack("getReserves", returnDatas[0])
		if err != nil {
			return fmt.Errorf("failed to unpack getReserves: %v", err)
		}
		pools[index].reserveA = reserves[0].(*big.Int)
		pools[index].reserveB = reserves[1].(*big.Int)
		return nil
	}

	if err := utils.Multicall(
		client,
		maxConcurrentRpcCalls,
		1500,
		calldataBatchGenerator,
		returnDataBatchHandler,
		indices,
	); err != nil {
		return fmt.Errorf("failed to multicall: %v", err)
	}

	log.Info("fetched uniswap pools reserves",
		"numPools", len(pools),
		"took", time.Since(startTime),
	)
	return nil
}

func numUniswapV2Pools(client *ethclient.Client, factoryAddr common.Address) (*big.Int, error) {
	factory, err := uniswapV2Factory.NewUniswapV2Factory(factoryAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to get factory: %v", err)
	}
	numPairs, err := factory.AllPairsLength(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all pairs length: %v", err)
	}
	return numPairs, nil
}

func fillUniswapV2Address(
	client *ethclient.Client,
	factoryAddr common.Address,
	pools []*UniswapV2Pool,
	maxConcurrentRpcCalls int,
) error {

	indices := make([]int, 0)
	for i := 0; i < len(pools); i++ {
		indices = append(indices, i)
	}

	calldataBatchGenerator := func(index int) ([]multicall.Multicall3Call, error) {
		calldata, err := UniswapV2FactoryAbi.Pack("allPairs", big.NewInt(int64(index)))
		if err != nil {
			return nil, fmt.Errorf("failed to pack allPairs: %v", err)
		}
		return []multicall.Multicall3Call{
			{
				Target:   factoryAddr,
				CallData: calldata,
			},
		}, nil
	}

	returnDataBatchHandler := func(index int, returnDatas [][]byte) error {
		if len(returnDatas) != 1 {
			return fmt.Errorf("unexpected return data length: %d, expected 1", len(returnDatas))
		}
		returnData := returnDatas[0]
		unpacked, err := UniswapV2FactoryAbi.Unpack("allPairs", returnData)
		if err != nil {
			return fmt.Errorf("failed to unpack allPairs: %v", err)
		}
		if len(unpacked) != 1 {
			return fmt.Errorf("unexpected unpacked length: %d", len(unpacked))
		}
		pairAddr, ok := unpacked[0].(common.Address)
		if !ok {
			return fmt.Errorf("failed to convert unpacked to common.Address")
		}
		pools[index] = NewUniswapV2Pool(pairAddr)
		return nil
	}

	return utils.Multicall(
		client,
		maxConcurrentRpcCalls,
		5000,
		calldataBatchGenerator,
		returnDataBatchHandler,
		indices,
	)
}

func fillUniswapV2StaticData(
	client *ethclient.Client,
	pools []*UniswapV2Pool,
	maxConcurrentRpcCalls int,
) error {
	indices := make([]int, 0)
	for i := 0; i < len(pools); i++ {
		indices = append(indices, i)
	}

	calldataBatchGenerator := func(index int) ([]multicall.Multicall3Call, error) {
		poolAddr := pools[index].Address()
		calldataToken0, err := UniswapV2PairAbi.Pack("token0")
		if err != nil {
			return nil, fmt.Errorf("failed to pack token0: %v", err)
		}
		calldataToken1, err := UniswapV2PairAbi.Pack("token1")
		if err != nil {
			return nil, fmt.Errorf("failed to pack token1: %v", err)
		}
		calldataReserves, err := UniswapV2PairAbi.Pack("getReserves")
		if err != nil {
			return nil, fmt.Errorf("failed to pack getReserves: %v", err)
		}
		return []multicall.Multicall3Call{
			{
				Target:   poolAddr,
				CallData: calldataToken0,
			},
			{
				Target:   poolAddr,
				CallData: calldataToken1,
			},
			{
				Target:   poolAddr,
				CallData: calldataReserves,
			},
		}, nil
	}

	returnDataBatchHandler := func(index int, returnDatas [][]byte) error {
		if len(returnDatas) != 3 {
			return fmt.Errorf("unexpected return data length: %d, expected 3", len(returnDatas))
		}
		token0, err := UniswapV2PairAbi.Unpack("token0", returnDatas[0])
		if err != nil {
			return fmt.Errorf("failed to unpack token0: %v", err)
		}
		token0Addr, ok := token0[0].(common.Address)
		if !ok {
			return fmt.Errorf("failed to convert token0 to common.Address, token0: %v", token0)
		}
		pools[index].tokenA = token0Addr

		token1, err := UniswapV2PairAbi.Unpack("token1", returnDatas[1])
		if err != nil {
			return fmt.Errorf("failed to unpack token1: %v", err)
		}
		token1Addr, ok := token1[0].(common.Address)
		if !ok {
			return fmt.Errorf("failed to convert token1 to common.Address, token1: %v", token1)
		}
		pools[index].tokenB = token1Addr

		reserves, err := UniswapV2PairAbi.Unpack("getReserves", returnDatas[2])
		if err != nil {
			return fmt.Errorf("failed to unpack getReserves: %v", err)
		}
		pools[index].reserveA = reserves[0].(*big.Int)
		pools[index].reserveB = reserves[1].(*big.Int)
		return nil
	}

	return utils.Multicall(
		client,
		maxConcurrentRpcCalls,
		1500,
		calldataBatchGenerator,
		returnDataBatchHandler,
		indices,
	)
}
