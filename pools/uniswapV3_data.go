package pools

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/contract/multicall"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/FastLane-Labs/fastlane-online-solver/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	UniswapV3PoolCreatedHash = common.HexToHash("0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118")
)

// GeneratePoolsStaticData_UniswapV3 generates static data for Uniswap V3 pools.
//
// This function -
// 1. fetches `PoolCreated` historical logs for the factory contract to get V3 pool addresses.
// 2. fills the static data for the pools. (tokens, fee etc)
// 3. filters out pools with zero liquidity and tokens that are not of interest.
//
// Arguments:
//   - client: the *ethclient.Client instance
//   - factoryAddr: the address of the Uniswap V3 factory
//   - interestingTokens: a list of token addresses (from config) that are of interest for generating pool data
//   - maxConcurrentRpcCalls: the maximum number of concurrent RPC calls to use (to manage rate limits)
//
// Returns:
//   - []*UniswapV3Pool: a slice of Uniswap V3 pool objects containing static data about the pools
//   - error: if any error occurs during the data retrieval
func GeneratePoolsStaticData_UniswapV3(
	client *ethclient.Client,
	factoryAddr common.Address,
	interestingTokens []common.Address,
	maxConcurrentRpcCalls int,
) ([]*UniswapV3Pool, error) {

	log.Info("fetching uniswapV3 pools static data...", "factory", factoryAddr.Hex())
	startTime := time.Now()

	pools, err := fetchUniswapV3PoolAddresses(client, factoryAddr, maxConcurrentRpcCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch uniswap v3 pool addresses: %v", err)
	}
	log.Info("fetched uniswapV3 pool addresses",
		"numPools", len(pools),
		"factory", factoryAddr.Hex(),
		"elapsed", time.Since(startTime),
	)

	log.Info("filling uniswapV3 pools static data...",
		"numPools", len(pools),
		"factory", factoryAddr.Hex(),
	)
	err = fillUniswapV3StaticData(client, pools, maxConcurrentRpcCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to fill uniswap v3 static data: %v", err)
	}
	log.Info("fetched uniswapV3 pools static data",
		"numPools", len(pools),
		"elapsed", time.Since(startTime),
	)

	filteredPools := make([]*UniswapV3Pool, 0)
	for _, pool := range pools {
		if pool.liquidity.Cmp(big.NewInt(0)) > 0 &&
			slices.Contains(interestingTokens, pool.tokenA) &&
			slices.Contains(interestingTokens, pool.tokenB) {
			filteredPools = append(filteredPools, pool)
		}
	}

	log.Info("uniswap v3 pools static data fetched",
		"numPools", len(filteredPools),
		"factory", factoryAddr.Hex(),
		"took", time.Since(startTime),
	)

	return filteredPools, nil
}

// GeneratePoolsDynamicData_UniswapV3 generates dynamic data for Uniswap V3 pools.
//
// This function fetches real-time data for the provided Uniswap V3 pools using an *ethclient.Client.
// The dynamic data for UNiswapV3 includes sqrtPricex96, liquidity, tick,
// and liquidityDeltas (netLiquidity changes while crossing ticks).
//
// Arguments:
//   - client: the *ethclient.Client instance
//   - pools: a slice of Uniswap V2 pool objects for which to fetch reserves
//   - maxConcurrentRpcCalls: the maximum number of concurrent RPC calls to use (to manage rate limits)
//
// Returns:
//   - error: if any error occurs during the data retrieval
func GeneratePoolsDynamicData_UniswapV3(
	client *ethclient.Client,
	pools []*UniswapV3Pool,
	maxConcurrentRpcCalls int,
) error {
	log.Info("fetching uniswapV3 pools dynamic data ...", "numPools", len(pools))
	startTime := time.Now()

	indices := make([]int, 0)
	for i := 0; i < len(pools); i++ {
		indices = append(indices, i)
	}

	calldataBatchGenerator := func(index int) ([]multicall.Multicall3Call, error) {
		calldataSlot0, err := UniswapV3PoolAbi.Pack("slot0")
		if err != nil {
			return nil, fmt.Errorf("failed to pack slot0: %v", err)
		}
		calldataLiquidity, err := UniswapV3PoolAbi.Pack("liquidity")
		if err != nil {
			return nil, fmt.Errorf("failed to pack liquidity: %v", err)
		}
		return []multicall.Multicall3Call{
			{
				Target:   pools[index].address,
				CallData: calldataSlot0,
			}, {
				Target:   pools[index].address,
				CallData: calldataLiquidity,
			},
		}, nil
	}
	returnDataBatchHandler := func(index int, returnDatas [][]byte) error {
		if len(returnDatas) != 2 {
			return fmt.Errorf("unexpected return data length: %d", len(returnDatas))
		}
		slot0, err := UniswapV3PoolAbi.Unpack("slot0", returnDatas[0])
		if err != nil {
			return fmt.Errorf("failed to unpack slot0: %v", err)
		}
		pools[index].sqrtPrice96 = slot0[0].(*big.Int)
		unspacedTick := slot0[1].(*big.Int).Int64()
		pools[index].tick = pools[index].spaceTick(unspacedTick)

		liquidity, err := UniswapV3PoolAbi.Unpack("liquidity", returnDatas[1])
		if err != nil {
			return fmt.Errorf("failed to unpack liquidity: %v", err)
		}
		pools[index].liquidity = liquidity[0].(*big.Int)
		return nil
	}

	if err := utils.Multicall(
		client,
		maxConcurrentRpcCalls,
		2000,
		calldataBatchGenerator,
		returnDataBatchHandler,
		indices,
	); err != nil {
		return fmt.Errorf("failed to multicall for slot0 and liquidity: %v", err)
	}
	log.Info("fetched uniswapV3 pools liquidity and slot0",
		"numPools", len(pools),
		"elapsed", time.Since(startTime),
	)

	log.Info("fetching uniswapV3 pools liquidity deltas...",
		"numPools", len(pools),
		"numTicksPerPool", 2*Default_unsiwapV3_num_ticks+1,
	)

	calldataBatchGenerator = func(index int) ([]multicall.Multicall3Call, error) {
		p := pools[index]
		calldata := make([]multicall.Multicall3Call, 0)
		for tick := p.tick - Default_unsiwapV3_num_ticks*p.tickSpacing; tick <= p.tick+Default_unsiwapV3_num_ticks*p.tickSpacing; tick += p.tickSpacing {
			tickCalldata, err := UniswapV3PoolAbi.Pack("ticks", big.NewInt(tick))
			if err != nil {
				return nil, fmt.Errorf("failed to pack ticks: %v", err)
			}
			calldata = append(calldata, multicall.Multicall3Call{
				Target:   p.address,
				CallData: tickCalldata,
			})
		}
		if int64(len(calldata)) != 2*Default_unsiwapV3_num_ticks+1 {
			return nil, fmt.Errorf("unexpected calldata length: %d", len(calldata))
		}
		return calldata, nil
	}

	returnDataBatchHandler = func(index int, returnDatas [][]byte) error {
		if len(returnDatas) != int(2*Default_unsiwapV3_num_ticks+1) {
			return fmt.Errorf("unexpected return data length: %d expected %d", len(returnDatas), 2*Default_unsiwapV3_num_ticks+1)
		}
		for tickIdx, returnData := range returnDatas {
			res, err := UniswapV3PoolAbi.Unpack("ticks", returnData)
			if err != nil {
				return fmt.Errorf("failed to unpack ticks: %v", err)
			}
			p := pools[index]
			tick := p.tick - Default_unsiwapV3_num_ticks*p.tickSpacing + int64(tickIdx)*p.tickSpacing
			p.mu.Lock()
			p.liquidityDeltas[tick] = res[1].(*big.Int)
			p.mu.Unlock()
		}
		return nil
	}

	batchSizeForTicksMulticall := 5000 / (Default_unsiwapV3_num_ticks*2 + 1)
	if err := utils.Multicall(
		client,
		maxConcurrentRpcCalls,
		int(batchSizeForTicksMulticall),
		calldataBatchGenerator,
		returnDataBatchHandler,
		indices,
	); err != nil {
		return fmt.Errorf("failed to multicall for ticks: %v", err)
	}

	log.Info("fetched uniswapV3 pools liquidity deltas",
		"numPools", len(pools),
		"took", time.Since(startTime),
	)

	for _, pool := range pools {
		pool.stateLoaderStatus = &stateLoaderStatus{
			ethClient:                client,
			liquidityDeltasLoadBlock: 0,
			maxConcurrentRpcCalls:    maxConcurrentRpcCalls,
		}
	}
	return nil
}

func fetchUniswapV3PoolAddresses(client *ethclient.Client, factoryAddr common.Address, maxConcurrentRpcCalls int) ([]*UniswapV3Pool, error) {
	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %v", err)
	}
	poolCreatedLogs := utils.HistoricalEventFetchBatched(
		client,
		int64(0),
		int64(latestBlock),
		[]common.Address{factoryAddr},
		[]common.Hash{UniswapV3PoolCreatedHash},
		10000000,
		maxConcurrentRpcCalls,
	)

	pools := make([]*UniswapV3Pool, 0)
	for _, log := range poolCreatedLogs {
		res, err := UniswapV3FactoryAbi.Unpack("PoolCreated", log.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack pool created: %v", err)
		}
		if len(res) != 2 {
			return nil, fmt.Errorf("unexpected unpacked length: %d", len(res))
		}
		poolAddr, ok := res[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("failed to convert unpacked to common.Address")
		}
		pools = append(pools, NewUniswapV3Pool(poolAddr))
	}

	return pools, nil
}

func fillUniswapV3StaticData(
	client *ethclient.Client,
	pools []*UniswapV3Pool,
	maxConcurrentRpcCalls int,
) error {

	indices := make([]int, 0)
	for i := 0; i < len(pools); i++ {
		indices = append(indices, i)
	}

	calldataBatchGenerator := func(index int) ([]multicall.Multicall3Call, error) {
		calldataTokenA, err := UniswapV3PoolAbi.Pack("token0")
		if err != nil {
			return nil, fmt.Errorf("failed to pack token0: %v", err)
		}
		calldataTokenB, err := UniswapV3PoolAbi.Pack("token1")
		if err != nil {
			return nil, fmt.Errorf("failed to pack token1: %v", err)
		}
		calldataFee, err := UniswapV3PoolAbi.Pack("fee")
		if err != nil {
			return nil, fmt.Errorf("failed to pack fee: %v", err)
		}
		calldataTickSpacing, err := UniswapV3PoolAbi.Pack("tickSpacing")
		if err != nil {
			return nil, fmt.Errorf("failed to pack tickSpacing: %v", err)
		}
		calldataLiquidity, err := UniswapV3PoolAbi.Pack("liquidity")
		if err != nil {
			return nil, fmt.Errorf("failed to pack liquidity: %v", err)
		}
		return []multicall.Multicall3Call{
			{
				Target:   pools[index].address,
				CallData: calldataTokenA,
			}, {
				Target:   pools[index].address,
				CallData: calldataTokenB,
			}, {
				Target:   pools[index].address,
				CallData: calldataFee,
			}, {
				Target:   pools[index].address,
				CallData: calldataTickSpacing,
			}, {
				Target:   pools[index].address,
				CallData: calldataLiquidity,
			},
		}, nil
	}

	returnDataBatchHandler := func(index int, returnDatas [][]byte) error {
		if len(returnDatas) != 5 {
			return fmt.Errorf("unexpected return data length: %d", len(returnDatas))
		}

		tokenA, err := UniswapV3PoolAbi.Unpack("token0", returnDatas[0])
		if err != nil {
			return fmt.Errorf("failed to unpack token0: %v", err)
		}
		tokenAAddr, ok := tokenA[0].(common.Address)
		if !ok {
			return fmt.Errorf("failed to convert token0 to common.Address")
		}
		tokenB, err := UniswapV3PoolAbi.Unpack("token1", returnDatas[1])
		if err != nil {
			return fmt.Errorf("failed to unpack token1: %v", err)
		}
		tokenBAddr, ok := tokenB[0].(common.Address)
		if !ok {
			return fmt.Errorf("failed to convert token1 to common.Address")
		}
		fee, err := UniswapV3PoolAbi.Unpack("fee", returnDatas[2])
		if err != nil {
			return fmt.Errorf("failed to unpack fee: %v", err)
		}
		feeInt := fee[0].(*big.Int).Int64()
		tickSpacing, err := UniswapV3PoolAbi.Unpack("tickSpacing", returnDatas[3])
		if err != nil {
			return fmt.Errorf("failed to unpack tickSpacing: %v", err)
		}
		tickSpacingInt := tickSpacing[0].(*big.Int).Int64()
		liquidity, err := UniswapV3PoolAbi.Unpack("liquidity", returnDatas[4])
		if err != nil {
			return fmt.Errorf("failed to unpack liquidity: %v", err)
		}
		liquidityInt := liquidity[0].(*big.Int)
		pools[index].tokenA = tokenAAddr
		pools[index].tokenB = tokenBAddr
		pools[index].fee = feeInt
		pools[index].tickSpacing = tickSpacingInt
		pools[index].liquidity = liquidityInt
		return nil
	}

	return utils.Multicall(
		client,
		maxConcurrentRpcCalls,
		1000,
		calldataBatchGenerator,
		returnDataBatchHandler,
		indices,
	)

}
