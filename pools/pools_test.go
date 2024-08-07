package pools

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func genUniswapV2Pool(address, tokenA, tokenB common.Address) *UniswapV2Pool {
	return &UniswapV2Pool{
		address: address,
		tokenA:  tokenA,
		tokenB:  tokenB,
	}
}

func genUniswapV3Pool(address, tokenA, tokenB common.Address) *UniswapV3Pool {
	return &UniswapV3Pool{
		address:     address,
		tokenA:      tokenA,
		tokenB:      tokenB,
		fee:         3000,
		tickSpacing: 60,
	}
}

func TestSerializeDeserializePoolsAndPaths(t *testing.T) {
	poolFileName := "testPools.bin"
	pathFileName := "testPaths.bin"

	// Generate pools
	p1 := genUniswapV2Pool(common.HexToAddress("0xA"), common.HexToAddress("0x1"), common.HexToAddress("0x2"))
	p2 := genUniswapV2Pool(common.HexToAddress("0xB"), common.HexToAddress("0x2"), common.HexToAddress("0x3"))
	p3 := genUniswapV3Pool(common.HexToAddress("0xC"), common.HexToAddress("0x3"), common.HexToAddress("0x4"))
	p4 := genUniswapV3Pool(common.HexToAddress("0xD"), common.HexToAddress("0x4"), common.HexToAddress("0x1"))

	pools := []Pool{p1, p2, p3, p4}

	// Serialize and dump pools to file
	err := SerializePoolsToFile(poolFileName, pools)
	if err != nil {
		t.Errorf("failed to serialize pools: %v", err)
	}

	// Generate paths
	paths, err := GenerateSwapPaths(pools)
	if err != nil {
		t.Errorf("failed to generate swap paths: %v", err)
	}

	// Convert paths map to slice for serialization
	var pathList []*SwapPath
	for _, paths := range paths {
		pathList = append(pathList, paths...)
	}

	// Serialize and dump paths to file
	err = SerializeSwapPathsToFile(pathFileName, pathList)
	if err != nil {
		t.Errorf("failed to serialize swap paths: %v", err)
	}

	// Deserialize pools from file
	deserializedPools, err := DeserializePoolsFromFile(poolFileName)
	if err != nil {
		t.Errorf("failed to deserialize pools: %v", err)
	}

	// Deserialize paths from file
	deserializedPaths, err := DeserializeSwapPathsFromFile(pathFileName)
	if err != nil {
		t.Errorf("failed to deserialize swap paths: %v", err)
	}

	// Clean up files
	err = os.Remove(poolFileName)
	if err != nil {
		t.Errorf("failed to remove pool file: %v", err)
	}
	err = os.Remove(pathFileName)
	if err != nil {
		t.Errorf("failed to remove path file: %v", err)
	}

	// Verify pools
	if len(deserializedPools) != len(pools) {
		t.Errorf("unexpected number of pools: %d", len(deserializedPools))
	}
	for i, pool := range pools {
		if deserializedPools[i].Address() != pool.Address() ||
			deserializedPools[i].TokenA() != pool.TokenA() ||
			deserializedPools[i].TokenB() != pool.TokenB() {
			t.Errorf("pool mismatch at index %d", i)
		}
		if _, ok := pool.(*UniswapV3Pool); ok {
			if deserializedPools[i].(*UniswapV3Pool).fee != pool.(*UniswapV3Pool).fee ||
				deserializedPools[i].(*UniswapV3Pool).tickSpacing != pool.(*UniswapV3Pool).tickSpacing {
				t.Errorf("uniswap v3 pool mismatch at index %d", i)
			}
		}
	}

	// Verify paths
	if len(deserializedPaths) != len(pathList) {
		t.Errorf("unexpected number of paths: %d", len(deserializedPaths))
	}
	for i, path := range pathList {
		if deserializedPaths[i].FromToken != path.FromToken ||
			deserializedPaths[i].ToToken != path.ToToken ||
			len(deserializedPaths[i].IntermediateTokens) != len(path.IntermediateTokens) ||
			len(deserializedPaths[i].PoolAddrs) != len(path.PoolAddrs) {
			t.Errorf("path mismatch at index %d", i)
		}
		for j := range path.IntermediateTokens {
			if deserializedPaths[i].IntermediateTokens[j] != path.IntermediateTokens[j] {
				t.Errorf("intermediate token mismatch at path index %d, token index %d", i, j)
			}
		}
		for j := range path.PoolAddrs {
			if deserializedPaths[i].PoolAddrs[j] != path.PoolAddrs[j] {
				t.Errorf("pool address mismatch at path index %d, pool address index %d", i, j)
			}
		}
	}
}
