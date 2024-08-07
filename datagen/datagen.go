package datagen

import (
	"fmt"
	"os"

	"github.com/FastLane-Labs/fastlane-online-solver/config"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/FastLane-Labs/fastlane-online-solver/pools"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DataGen generates data files (binary files) for pools and swap paths.
// It performs the following steps:
// 1. Checks write permissions for the specified file paths.
// 2. Deletes the existing files.
// 3. Generates static data for dex pools specified in the `config.PoolsConfigFile`.
// 4. Generates swap paths for the pools initialized in step 3.
// 5. Serializes the pools and swap paths to the files specified in the `config.PoolsBinFile` and `config.SwapPathsBinFile`.

// Arguments:
//   - conf: the *config.Config instance
//
// Returns:
//   - error: if any error occurs during the data generation
func DataGen(conf *config.Config) error {
	if err := checkWritePermissions(conf); err != nil {
		return err
	}

	log.Info("generating data files", "poolsFile", conf.PoolsBinFile, "pathsFile", conf.SwapPathsBinFile)

	client, err := ethclient.Dial(conf.EthRpcUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client - %v", err)
	}

	poolsStore, err := pools.InitializePoolsAndPaths(client, conf)
	if err != nil {
		return fmt.Errorf("failed to initialize pools and paths - %v", err)
	}

	poolsFile := conf.PoolsBinFile
	pathsFile := conf.SwapPathsBinFile

	swapPools := make([]pools.Pool, 0)
	swapPaths := make([]*pools.SwapPath, 0)

	for _, pool := range poolsStore.Pools {
		swapPools = append(swapPools, pool)
	}
	for _, paths := range poolsStore.Paths {
		swapPaths = append(swapPaths, paths...)
	}

	log.Info("serializing pools to file", "poolsFile", poolsFile, "numPools", len(swapPools))
	err = pools.SerializePoolsToFile(poolsFile, swapPools)
	if err != nil {
		return fmt.Errorf("failed to serialize pools to file - %v", err)
	}

	log.Info("serializing swap paths to file", "pathsFile", pathsFile, "numPaths", len(swapPaths))
	err = pools.SerializeSwapPathsToFile(pathsFile, swapPaths)
	if err != nil {
		return fmt.Errorf("failed to serialize swap paths to file - %v", err)
	}

	log.Info("data files generation complete", "poolsFile", poolsFile, "pathsFile", pathsFile)
	return nil
}

func checkWritePermissions(conf *config.Config) error {
	if conf.PoolsBinFile == "" {
		return fmt.Errorf("pools file path not set")
	}
	if conf.SwapPathsBinFile == "" {
		return fmt.Errorf("swap paths file path not set")
	}

	err := checkAndDeleteFile(conf.PoolsBinFile)
	if err != nil {
		return err
	}

	err = checkAndDeleteFile(conf.SwapPathsBinFile)
	if err != nil {
		return err
	}

	return nil
}

func checkAndDeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("cannot delete file: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("cannot write to file: %v", err)
	}
	file.Close()
	return nil
}
