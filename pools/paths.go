package pools

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// FromToken --> IntermediateToken[0] --> IntermediateToken[1] --> ToToken
// [-----------Pool1-----------------]    [-------------Pool2-------------]
type SwapPath struct {
	FromToken          common.Address
	ToToken            common.Address
	IntermediateTokens []common.Address
	PoolAddrs          []common.Address
}

func (path *SwapPath) Key() common.Hash {
	return crypto.Keccak256Hash(path.FromToken[:], path.ToToken[:])
}

func (path *SwapPath) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if _, err := buf.Write(path.FromToken.Bytes()); err != nil {
		return nil, err
	}

	if _, err := buf.Write(path.ToToken.Bytes()); err != nil {
		return nil, err
	}

	intermediateLen := uint16(len(path.IntermediateTokens))
	if err := binary.Write(buf, binary.LittleEndian, intermediateLen); err != nil {
		return nil, err
	}
	for _, token := range path.IntermediateTokens {
		if _, err := buf.Write(token.Bytes()); err != nil {
			return nil, err
		}
	}

	poolLen := uint16(len(path.PoolAddrs))
	if err := binary.Write(buf, binary.LittleEndian, poolLen); err != nil {
		return nil, err
	}
	for _, pool := range path.PoolAddrs {
		if _, err := buf.Write(pool.Bytes()); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (path *SwapPath) Deserialize(data []byte) error {
	buf := bytes.NewReader(data)

	fromTokenBytes := make([]byte, common.AddressLength)
	if _, err := buf.Read(fromTokenBytes); err != nil {
		return err
	}
	path.FromToken = common.BytesToAddress(fromTokenBytes)

	toTokenBytes := make([]byte, common.AddressLength)
	if _, err := buf.Read(toTokenBytes); err != nil {
		return err
	}
	path.ToToken = common.BytesToAddress(toTokenBytes)

	var intermediateLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &intermediateLen); err != nil {
		return err
	}
	path.IntermediateTokens = make([]common.Address, intermediateLen)
	for i := uint16(0); i < intermediateLen; i++ {
		tokenBytes := make([]byte, common.AddressLength)
		if _, err := buf.Read(tokenBytes); err != nil {
			return err
		}
		path.IntermediateTokens[i] = common.BytesToAddress(tokenBytes)
	}

	var poolLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &poolLen); err != nil {
		return err
	}
	path.PoolAddrs = make([]common.Address, poolLen)
	for i := uint16(0); i < poolLen; i++ {
		poolBytes := make([]byte, common.AddressLength)
		if _, err := buf.Read(poolBytes); err != nil {
			return err
		}
		path.PoolAddrs[i] = common.BytesToAddress(poolBytes)
	}

	return nil
}

func GenerateSwapPaths(pools []Pool) (map[common.Hash][]*SwapPath, error) {
	paths := make(map[common.Hash][]*SwapPath)

	var mu sync.Mutex
	var wg sync.WaitGroup
	completedGoRoutines := make(chan struct{})
	doneChan := make(chan struct{})
	totalGoRoutines := len(pools) + len(pools)*(len(pools)-1)/2

	go func() {
		completed := 0
		lastRecord := 0.0
		for {
			select {
			case <-doneChan:
				return
			case <-completedGoRoutines:
				completed++
				progressPerc := float64(completed) * 100 / float64(totalGoRoutines)
				if progressPerc > lastRecord {
					log.Debug("Generating swap paths", "progress", fmt.Sprintf("%.2f%%", progressPerc))
					lastRecord += 10.0
				}
			}
		}
	}()

	addPath := func(path *SwapPath) {
		mu.Lock()
		defer mu.Unlock()

		key := path.Key()
		if _, ok := paths[key]; !ok {
			paths[key] = make([]*SwapPath, 0)
		}
		paths[key] = append(paths[key], path)
	}

	for _, p := range pools {
		wg.Add(1)
		go func(pool Pool) {
			defer func() {
				completedGoRoutines <- struct{}{}
				wg.Done()
			}()
			path1 := &SwapPath{
				FromToken:          pool.TokenA(),
				ToToken:            pool.TokenB(),
				IntermediateTokens: []common.Address{},
				PoolAddrs:          []common.Address{pool.Address()},
			}
			addPath(path1)
			path2 := &SwapPath{
				FromToken:          pool.TokenB(),
				ToToken:            pool.TokenA(),
				IntermediateTokens: []common.Address{},
				PoolAddrs:          []common.Address{pool.Address()},
			}
			addPath(path2)
		}(p)
	}

	for i, p1 := range pools {
		for j, p2 := range pools {
			if i >= j {
				continue
			}

			wg.Add(1)
			go func(pool1, pool2 Pool) {
				defer func() {
					completedGoRoutines <- struct{}{}
					wg.Done()
				}()

				t1A := pool1.TokenA()
				t1B := pool1.TokenB()
				t2A := pool2.TokenA()
				t2B := pool2.TokenB()
				if (t1A == t2A && t1B == t2B) || (t1A == t2B && t1B == t2A) {
					return
				}
				if (t1A != t2A && t1A != t2B) && (t1B != t2A && t1B != t2B) {
					return
				}
				if t1A == t2A {
					path1 := &SwapPath{
						FromToken:          t1B,
						ToToken:            t2B,
						IntermediateTokens: []common.Address{t1A},
						PoolAddrs:          []common.Address{pool1.Address(), pool2.Address()},
					}
					addPath(path1)
					path2 := &SwapPath{
						FromToken:          t2B,
						ToToken:            t1B,
						IntermediateTokens: []common.Address{t1A},
						PoolAddrs:          []common.Address{pool2.Address(), pool1.Address()},
					}
					addPath(path2)
				} else if t1A == t2B {
					path1 := &SwapPath{
						FromToken:          t1B,
						ToToken:            t2A,
						IntermediateTokens: []common.Address{t1A},
						PoolAddrs:          []common.Address{pool1.Address(), pool2.Address()},
					}
					addPath(path1)
					path2 := &SwapPath{
						FromToken:          t2A,
						ToToken:            t1B,
						IntermediateTokens: []common.Address{t1A},
						PoolAddrs:          []common.Address{pool2.Address(), pool1.Address()},
					}
					addPath(path2)
				} else if t1B == t2A {
					path1 := &SwapPath{
						FromToken:          t1A,
						ToToken:            t2B,
						IntermediateTokens: []common.Address{t1B},
						PoolAddrs:          []common.Address{pool1.Address(), pool2.Address()},
					}
					addPath(path1)
					path2 := &SwapPath{
						FromToken:          t2B,
						ToToken:            t1A,
						IntermediateTokens: []common.Address{t1B},
						PoolAddrs:          []common.Address{pool2.Address(), pool1.Address()},
					}
					addPath(path2)
				} else if t1B == t2B {
					path1 := &SwapPath{
						FromToken:          t1A,
						ToToken:            t2A,
						IntermediateTokens: []common.Address{t1B},
						PoolAddrs:          []common.Address{pool1.Address(), pool2.Address()},
					}
					addPath(path1)
					path2 := &SwapPath{
						FromToken:          t2A,
						ToToken:            t1A,
						IntermediateTokens: []common.Address{t1B},
						PoolAddrs:          []common.Address{pool2.Address(), pool1.Address()},
					}
					addPath(path2)
				}
			}(p1, p2)
		}
	}

	wg.Wait()
	close(completedGoRoutines)
	close(doneChan)
	return paths, nil
}
