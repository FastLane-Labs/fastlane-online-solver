package events

import (
	"context"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/log"
	pb "github.com/bloXroute-Labs/gateway/v2/protobuf"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ Mempool = &BloxrouteMempoolConnection{}

type BloxrouteMempoolConnection struct {
	client         pb.GatewayClient
	pendingTxsChan chan *types.Transaction

	url       string
	authToken string
}

type blxrCredentials struct {
	authorization string
}

func (bc *blxrCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": bc.authorization,
	}, nil
}

func (bc *blxrCredentials) RequireTransportSecurity() bool {
	return false
}

func NewBloxrouteMempoolConnection(grpcUrl, authToken string) (*BloxrouteMempoolConnection, error) {
	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.Dial(
		grpcUrl,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&blxrCredentials{authorization: authToken}),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewGatewayClient(conn)

	return &BloxrouteMempoolConnection{
		client:         client,
		pendingTxsChan: make(chan *types.Transaction),
		url:            grpcUrl,
		authToken:      authToken,
	}, nil
}

func (c *BloxrouteMempoolConnection) reconnect() {
	minBackoff := 100 * time.Millisecond
	maxBackoff := time.Minute
	backoff := minBackoff

	for {
		time.Sleep(backoff)
		log.Info("bloxrouteClient conn: reconnecting...")
		newConn, err := NewBloxrouteMempoolConnection(c.url, c.authToken)
		if err == nil {
			c.client = newConn.client
			c.Start()
			return
		}

		backoff *= 3
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
		log.Error("bloxrouteClientConn:", "reconnect failed:", err, "next attempt in", backoff)
	}
}

func (c *BloxrouteMempoolConnection) Start() error {
	log.Info("bloxrouteClient conn: starting...")

	ctx, cancel := context.WithTimeout(context.Background(), 24*time.Hour)

	stream, err := c.client.PendingTxs(ctx, &pb.TxsRequest{Filters: ""})
	if err != nil {
		cancel()
		return err
	}
	log.Info("Subscribed to pending transactions")

	go func() {
		defer stream.CloseSend()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				subscriptionNotification, err := stream.Recv()
				if err != nil {
					log.Error("Error in pending transactions subscription", "error", err)
					cancel()
					return
				}
				for _, pendingTransaction := range subscriptionNotification.Tx {
					rawTx := pendingTransaction.RawTx
					tx := &types.Transaction{}
					err := tx.UnmarshalJSON(rawTx)
					if err != nil {
						log.Error("failed to unmarshal raw tx", "err", err)
						return
					}
					c.pendingTxsChan <- tx
				}
			}
		}
	}()

	go func() {
		for {
			<-ctx.Done()
			c.reconnect()
			return
		}
	}()
	return nil
}

func (c *BloxrouteMempoolConnection) MempoolTxChan() <-chan *types.Transaction {
	return c.pendingTxsChan
}
