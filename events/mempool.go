package events

import "github.com/ethereum/go-ethereum/core/types"

type Mempool interface {
	Start() error
	MempoolTxChan() <-chan *types.Transaction
}
