package utils

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/contract/multicall"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	multicallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11") // mutlticall3 common for all chains
)

// Multicall makes multiple batched RPC calls to an *ethclient.Client
// based on generating []multicall.Multicall3Call from indices and handling [][]byte return data.
// Example use case: `pools.fillUniswapV2StaticData`.
// WARNING: Parallel execution of `returnDataBatchHandlerFunc` and `callDataBatchGeneratorFunc` is used,
// so take care of any contention in these functions.
//
// Arguments:
//   - client: *ethclient.Client
//   - maxConcurrentRpcCalls: number of concurrent goroutines to use while making multicalls (mainly to manage rate limits enforced by the RPC provider)
//   - numBatchesPerMulticall: number of batches to club together in a single multicall
//     (there's a limit on the multicall calldata size, so use a smaller `numBatchesPerMulticall` if batch size is bigger)
//   - callDataBatchGeneratorFunc: function that takes an index and returns []multicall.Multicall3Call
//   - returnDataBatchHandlerFunc: function that takes an index and [][]byte return data
//   - indices: indices for which to call `callDataBatchGeneratorFunc` and then `returnDataBatchHandlerFunc`
//
// Returns:
//   - error: if any error occurs during the multicall
func Multicall(
	client *ethclient.Client,
	maxConcurrentRpcCalls int,
	numBatchesPerMulticall int,
	callDataBatchGeneratorFunc func(index int) ([]multicall.Multicall3Call, error),
	returnDataBatchHandlerFunc func(index int, returnDataAtIndex [][]byte) error,
	indices []int,
) error {

	startTime := time.Now()

	batchAtIndex := make(map[int][]multicall.Multicall3Call)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, idx := range indices {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			batch, err := callDataBatchGeneratorFunc(index)
			if err != nil {
				log.Error("failed to generate call data batch", "index", index, "err", err)
				return
			}

			mu.Lock()
			if _, ok := batchAtIndex[index]; ok {
				log.Error("duplicate indices received in multicall")
				mu.Unlock()
				return
			}
			batchAtIndex[index] = batch
			mu.Unlock()
		}(idx)
	}
	wg.Wait()

	indexReturnDataLocation := make(map[int]struct {
		chunkIndex       int
		fromIndexInChunk int
		toIndexInChunk   int
	})

	// Make chunks of calldata batches such that max(len(chunk)) == numBatchesPerMulticall
	calldataChunks := make([][]multicall.Multicall3Call, 0)
	from := 0
	for {
		if from == len(indices) {
			break
		}
		to := from + numBatchesPerMulticall
		if to > len(indices) {
			to = len(indices)
		}
		chunk := make([]multicall.Multicall3Call, 0)
		for i := from; i < to; i++ {
			if batch, ok := batchAtIndex[indices[i]]; ok {
				fromIndexInChunk := len(chunk)
				chunk = append(chunk, batch...)
				toIndexInChunk := len(chunk)
				indexReturnDataLocation[indices[i]] = struct {
					chunkIndex       int
					fromIndexInChunk int
					toIndexInChunk   int
				}{
					chunkIndex:       len(calldataChunks),
					fromIndexInChunk: fromIndexInChunk,
					toIndexInChunk:   toIndexInChunk,
				}
			} else {
				return fmt.Errorf("failed to generate call data batch for index %d", indices[i])
			}
		}
		calldataChunks = append(calldataChunks, chunk)
		from = to
	}

	returnDataChunks := make([][][]byte, len(calldataChunks))

	sem := make(chan struct{}, maxConcurrentRpcCalls)
	completed := 0
	for j, ch := range calldataChunks {
		wg.Add(1)
		go func(chunkIdx int, calldataChunk []multicall.Multicall3Call) {
			defer wg.Done()
			defer func() {
				<-sem
				completed++
				log.Debug("multicall progress",
					"done", fmt.Sprintf("%f%%", float64(completed*100)/float64(len(calldataChunks))),
					"doneChunks", completed,
					"totalChunks", len(calldataChunks),
				)
			}()
			sem <- struct{}{}

			returnDataBatch, err := multicall_inner(client, calldataChunk)
			if err != nil {
				returnDataChunks[chunkIdx] = nil
				log.Error("failed to multicall", "err", err)
				return
			}

			returnDataChunks[chunkIdx] = returnDataBatch
		}(j, ch)
	}
	wg.Wait()

	for chIdx, chunkReturnData := range returnDataChunks {
		if chunkReturnData == nil {
			return fmt.Errorf("failed to multicall at chunk index %d", chIdx)
		}
	}

	for _, idx := range indices {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			loc := indexReturnDataLocation[index]
			if err := returnDataBatchHandlerFunc(index, returnDataChunks[loc.chunkIndex][loc.fromIndexInChunk:loc.toIndexInChunk]); err != nil {
				log.Error("failed to handle return data", "err", err)
			}
		}(idx)
	}
	wg.Wait()

	log.Debug("multicall done",
		"numChunks", len(calldataChunks),
		"numBatchesPerChunk", numBatchesPerMulticall,
		"took", time.Since(startTime),
	)

	return nil
}

func multicall_inner(client *ethclient.Client, multicallData []multicall.Multicall3Call) ([][]byte, error) {
	multicallAbi, err := multicall.MulticallMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get multicall abi: %v", err)
	}

	ethCalldata, err := multicallAbi.Pack("aggregate", multicallData)
	if err != nil {
		return nil, fmt.Errorf("failed to pack multicall data: %v", err)
	}

	ethMsg := ethereum.CallMsg{
		To:   &multicallAddress,
		Data: ethCalldata,
	}

	resp, err := client.CallContract(context.Background(), ethMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call multicall contract: %v", err)
	}

	res, err := multicallAbi.Unpack("aggregate", resp)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack multicall response: %v", err)
	}

	if len(res) != 2 {
		return nil, fmt.Errorf("unexpected response length: %d", len(res))
	}

	returnData, ok := res[1].([][]byte)
	if !ok {
		return nil, fmt.Errorf("failed to convert return data to [][]byte")
	}

	return returnData, nil
}
