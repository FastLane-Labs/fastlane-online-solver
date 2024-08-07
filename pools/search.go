package pools

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Swap struct {
	FromToken common.Address
	ToToken   common.Address
	Pool      Pool
}

type BestPath struct {
	FromToken common.Address
	ToToken   common.Address
	Swaps     []Swap
	Input     *big.Int
	Output    *big.Int
}

func (ps *PoolsStore) BestPath(fromToken common.Address, toToken common.Address, amountIn *big.Int) (*BestPath, error) {
	startTime := time.Now()

	pathKey := crypto.Keccak256(fromToken[:], toToken[:])
	paths := ps.Paths[common.Hash(pathKey)]

	if len(paths) == 0 {
		return nil, fmt.Errorf("no paths found for %s -> %s", fromToken.Hex(), toToken.Hex())
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	solutions := make([]*BestPath, len(paths))

	for i, path := range paths {
		wg.Add(1)
		go func(idx int, path *SwapPath) {
			wg.Done()

			output := big.NewInt(0).Set(amountIn)
			swaps := make([]Swap, len(path.PoolAddrs))

			ft := path.FromToken
			tt := path.ToToken
			if len(path.IntermediateTokens) > 0 {
				tt = path.IntermediateTokens[0]
			}

			for j, poolAddr := range path.PoolAddrs {
				mu.Lock()
				pool := ps.Pools[poolAddr]
				mu.Unlock()

				out, err := pool.Output(ft, tt, output)
				if err != nil {
					log.Debug("search error while getting pool output", "pool", pool.Address(), "error", err, "poolType", pool.PoolType())
					return
				}

				output = out
				swaps[j] = Swap{
					FromToken: ft,
					ToToken:   tt,
					Pool:      pool,
				}

				ft = tt
				tt = path.ToToken
				if j+1 < len(path.IntermediateTokens) {
					tt = path.IntermediateTokens[j+1]
				}
			}
			solutions[idx] = &BestPath{
				FromToken: fromToken,
				ToToken:   toToken,
				Swaps:     swaps,
				Input:     amountIn,
				Output:    output,
			}

		}(i, path)
	}
	wg.Wait()

	log.Info("finished exploring paths", "from", fromToken.Hex(), "to", toToken.Hex(), "paths", len(paths), "took", time.Since(startTime))

	var bestPath *BestPath
	for _, solution := range solutions {
		if solution == nil {
			continue
		}
		if bestPath == nil || solution.Output.Cmp(bestPath.Output) == 1 {
			bestPath = solution
		}
	}
	if bestPath == nil {
		return nil, fmt.Errorf("no best path found for %s -> %s after exploring %d paths", fromToken.Hex(), toToken.Hex(), len(paths))
	}
	return bestPath, nil
}
