package pools

import (
	"fmt"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/config"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PoolsStore struct {
	Pools map[common.Address]Pool
	Paths map[common.Hash][]*SwapPath
}

func NewPoolsStoreFromBinFiles(conf *config.Config) (*PoolsStore, error) {
	poolsFileName := conf.PoolsBinFile
	pathsFileName := conf.SwapPathsBinFile

	poolsSlice, err := DeserializePoolsFromFile(poolsFileName)
	if err != nil {
		return nil, err
	}
	pathsSlice, err := DeserializeSwapPathsFromFile(pathsFileName)
	if err != nil {
		return nil, err
	}

	pools := make(map[common.Address]Pool)
	paths := make(map[common.Hash][]*SwapPath)

	for _, pool := range poolsSlice {
		pools[pool.Address()] = pool
	}
	log.Info("loaded pools", "numPools", len(pools), "poolsFile", poolsFileName)

	numPaths := 0
	for _, path := range pathsSlice {
		if _, ok := paths[path.Key()]; !ok {
			paths[path.Key()] = make([]*SwapPath, 0)
		}
		paths[path.Key()] = append(paths[path.Key()], path)
		numPaths++
	}
	log.Info("loaded swap paths", "numPaths", numPaths, "pathsFile", pathsFileName)

	return &PoolsStore{
		Pools: pools,
		Paths: paths,
	}, nil
}

func NewPoolsStoreFromPoolsAndPaths(pools []Pool, paths map[common.Hash][]*SwapPath) (*PoolsStore, error) {
	poolsMap := make(map[common.Address]Pool)
	for _, pool := range pools {
		poolsMap[pool.Address()] = pool
	}
	return &PoolsStore{
		Pools: poolsMap,
		Paths: paths,
	}, nil
}

func (ps *PoolsStore) InitializeDynamicState(client *ethclient.Client, maxConcurrentRpcCalls int) {
	uniswapV2Pools := make([]*UniswapV2Pool, 0)
	uniswapV3Pools := make([]*UniswapV3Pool, 0)
	for _, pool := range ps.Pools {
		switch pool.PoolType() {
		case UniswapV2PoolType:
			uniswapV2Pool := pool.(*UniswapV2Pool)
			uniswapV2Pools = append(uniswapV2Pools, uniswapV2Pool)

		case UniswapV3PoolType:
			uniswapV3Pool := pool.(*UniswapV3Pool)
			uniswapV3Pools = append(uniswapV3Pools, uniswapV3Pool)
		}
	}

	GeneratePoolsDynamicData_UniswapV2(client, uniswapV2Pools, maxConcurrentRpcCalls)
	GeneratePoolsDynamicData_UniswapV3(client, uniswapV3Pools, maxConcurrentRpcCalls)
}

func InitializePoolsAndPaths(client *ethclient.Client, conf *config.Config) (*PoolsStore, error) {
	poolsConfig, err := conf.PoolsConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get pools config: %v", err)
	}

	log.Info("generating pools...")
	pools := make([]Pool, 0)
	for _, poolConfig := range poolsConfig {
		switch poolConfig.DexType {
		case "UniswapV2":
			v2Pools, err := GeneratePoolsStaticData_UniswapV2(
				client,
				common.HexToAddress(poolConfig.FactoryAddress),
				conf.GetInterestingTokens(),
				conf.MaxConcurrentRpcCalls,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to generate uniswap v2 pools: %v", err)
			}
			for _, v2Pool := range v2Pools {
				pools = append(pools, v2Pool)
			}
		case "UniswapV3":
			v3Pools, err := GeneratePoolsStaticData_UniswapV3(
				client,
				common.HexToAddress(poolConfig.FactoryAddress),
				conf.GetInterestingTokens(),
				conf.MaxConcurrentRpcCalls,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to generate uniswap v3 pools: %v", err)
			}
			for _, v3Pool := range v3Pools {
				pools = append(pools, v3Pool)
			}
		default:
			log.Error("unsupported dex type: %s", poolConfig.DexType)
			continue
		}
	}

	log.Info("generated pools", "numPools", len(pools))

	log.Info("generating swap paths...")
	startTime := time.Now()
	paths, err := GenerateSwapPaths(pools)
	if err != nil {
		return nil, fmt.Errorf("failed to generate swap paths: %v", err)
	}
	log.Info("generated swap paths", "elapsed", time.Since(startTime))
	return NewPoolsStoreFromPoolsAndPaths(pools, paths)
}

func (ps *PoolsStore) HandleLog(aLog *types.Log) {
	pool, ok := ps.Pools[aLog.Address]
	if !ok {
		return
	}
	pool.HandleLog(aLog)
}
