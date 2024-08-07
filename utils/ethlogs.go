package utils

import (
	"context"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// HistoricalEventFetchBatched fetches historical events in batches from an *ethclient.Client.
// This function retrieves logs from a range of blocks (`from` to `to`) for the specified addresses and topics[0]s (hashes).
//
// Arguments:
//   - client: the *ethclient.Client instance
//   - from: the starting block number to fetch events from
//   - to: the ending block number to fetch events until
//   - addressWhitelist: a list of addresses to filter the logs by
//   - topic0s: a list of event topics[0]s to filter the logs by
//   - blockBatchSize: the number of blocks to process in each RPC call
//   - maxConcurrentRpcCalls: the maximum number of concurrent RPC calls to use (mainly to manage rate limits)
//
// Returns:
//   - []types.Log: a slice of logs matching the specified criteria within the given block range

func HistoricalEventFetchBatched(
	client *ethclient.Client,
	from int64,
	to int64,
	addressWhitelist []common.Address,
	topic0s []common.Hash,
	blockBatchSize int64,
	maxConcurrentRpcCalls int,
) []types.Log {
	var logs []types.Log
	var mu sync.Mutex
	var wg sync.WaitGroup

	startTime := time.Now()
	sem := make(chan struct{}, maxConcurrentRpcCalls)

	for start := from; start <= to; start += blockBatchSize {
		end := start + blockBatchSize - 1
		if end > to {
			end = to
		}

		wg.Add(1)
		sem <- struct{}{} // acquire semaphore

		go func(start, end int64) {
			defer wg.Done()
			defer func() { <-sem }() // release semaphore

			arg := map[string]interface{}{
				"fromBlock": hexutil.EncodeBig(big.NewInt(start)),
				"toBlock":   hexutil.EncodeBig(big.NewInt(end)),
				"topics":    [][]common.Hash{topic0s},
				"address":   addressWhitelist,
			}

			var batchLogs []types.Log

			err := client.Client().CallContext(context.Background(), &batchLogs, "eth_getLogs", arg)
			if err != nil {
				log.Error("Error fetching logs", "error", err)
				return
			}

			log.Debug("found historical logs", "numLogs", len(batchLogs), "start", start, "end", end)
			mu.Lock()
			logs = append(logs, batchLogs...)
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].BlockNumber < logs[j].BlockNumber
	})

	log.Info("found historical logs", "numLogs", len(logs), "elapsed", time.Since(startTime))
	return logs
}

func HistoricalEventFetch(
	client *ethclient.Client,
	from int64,
	to int64,
	addressWhitelist []common.Address,
	topic0s []common.Hash,
) []types.Log {
	arg := map[string]interface{}{
		"fromBlock": hexutil.EncodeBig(big.NewInt(from)),
		"toBlock":   hexutil.EncodeBig(big.NewInt(to)),
		"topics":    [][]common.Hash{topic0s},
		"address":   addressWhitelist,
	}

	var logs []types.Log

	err := client.Client().CallContext(context.Background(), &logs, "eth_getLogs", arg)
	if err != nil {
		panic(err)
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].BlockNumber < (logs[j].BlockNumber)
	})

	log.Info("found historical logs", "numLogs", len(logs))

	return logs
}
