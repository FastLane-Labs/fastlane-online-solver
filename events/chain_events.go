package events

import (
	"context"
	"fmt"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/config"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/FastLane-Labs/fastlane-online-solver/pools"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ Mempool = &EthClientConnection{}

// EthClientConnection is a connection to an Eth node that listens for new blocks, logs, and pending transactions.
// It has reconnect logic built in with exponential backoff factor of 3.
// Implements the `Mempoôl` interface.
type EthClientConnection struct {
	Client         *ethclient.Client
	logsChan       chan *types.Log
	blocksChan     chan *types.Header
	pendingTxsChan chan *types.Transaction

	conf      *config.Config
	isStarted bool
}

func NewEthClientConnection(conf *config.Config) (*EthClientConnection, error) {
	client, err := ethclient.Dial(conf.EthRpcUrl)
	if err != nil {
		return nil, err
	}

	return &EthClientConnection{
		Client:         client,
		logsChan:       make(chan *types.Log),
		blocksChan:     make(chan *types.Header),
		pendingTxsChan: make(chan *types.Transaction),
		conf:           conf,
		isStarted:      false,
	}, nil
}

func (c *EthClientConnection) reconnect() {
	minBackoff := 100 * time.Millisecond
	maxBackoff := time.Minute
	backoff := minBackoff

	for {
		time.Sleep(backoff)
		log.Info("ethClient conn: reconnecting...")
		newConn, err := NewEthClientConnection(c.conf)
		if err == nil {
			c.Client = newConn.Client
			c.Start()
			return
		}

		backoff *= 3
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
		log.Error("ethClientConn:", "reconnect failed:", err, "next attempt in", backoff)
	}
}

func (c *EthClientConnection) Start() error {
	if c.isStarted {
		return nil
	}

	log.Info("ethClient conn: starting...")

	// Subscribe to new heads
	headers := make(chan *types.Header)
	headsSub, err := c.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %v", err)
	}
	log.Info("Subscribed to new heads")

	logs := make(chan types.Log)
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics: [][]common.Hash{{
			pools.UniswapV2SyncEventHash,
			pools.UniswapV3MintEventHash,
			pools.UniswapV3BurnEventHash,
			pools.UniswapV3SwapEventHash,
		}},
	}
	logsSub, err := c.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new logs: %v", err)
	}
	log.Info("Subscribed to new logs")

	pendingTxsChan := make(chan *types.Transaction)
	params := map[string]interface{}{
		"toAddress": []string{
			c.conf.FastlaneOnlineAddress,
		},
		"hashesOnly": false,
	}
	pendingTxsSub, err := c.Client.Client().EthSubscribe(context.Background(), pendingTxsChan, c.conf.MempoolConnectionEthMethod, params)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new pending transactions: %v", err)
	}
	log.Info("Subscribed to new pending transactions", "toAddress", c.conf.FastlaneOnlineAddress)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-headsSub.Err():
				log.Error("Error in new head subscription", "error", err)
				cancel()
				return
			case header := <-headers:
				c.blocksChan <- header
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-logsSub.Err():
				log.Error("Error in new logs subscription", "error", err)
				cancel()
				return
			case log := <-logs:
				c.logsChan <- &log
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-pendingTxsSub.Err():
				log.Error("Error in new pending transactions subscription", "error", err)
				cancel()
				return
			case tx := <-pendingTxsChan:
				c.pendingTxsChan <- tx
			}
		}
	}()

	go func() {
		for {
			<-ctx.Done()
			c.Client.Close()
			c.isStarted = false
			c.reconnect()
			return
		}
	}()

	c.isStarted = true
	return nil
}

func (c *EthClientConnection) BlocksChan() <-chan *types.Header {
	return c.blocksChan
}

func (c *EthClientConnection) LogsChan() <-chan *types.Log {
	return c.logsChan
}

func (c *EthClientConnection) MempoolTxChan() <-chan *types.Transaction {
	return c.pendingTxsChan
}
