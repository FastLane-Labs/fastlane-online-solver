package pools

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/fastlane-online-solver/contract/uniswapV2Factory"
	"github.com/FastLane-Labs/fastlane-online-solver/contract/uniswapV2Pair"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ Pool = &UniswapV2Pool{}

var (
	UniswapV2SyncEventHash = common.HexToHash("0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1")
)

var (
	UniswapV2PairAbi, _    = uniswapV2Pair.UniswapV2PairMetaData.GetAbi()
	UniswapV2FactoryAbi, _ = uniswapV2Factory.UniswapV2FactoryMetaData.GetAbi()
)

type UniswapV2Pool struct {
	address  common.Address
	tokenA   common.Address
	tokenB   common.Address
	reserveA *big.Int
	reserveB *big.Int
}

func NewUniswapV2Pool(address common.Address) *UniswapV2Pool {
	return &UniswapV2Pool{
		address: address,
	}
}

func (p *UniswapV2Pool) PoolType() PoolType {
	return UniswapV2PoolType
}

func (p *UniswapV2Pool) Address() common.Address {
	return p.address
}

func (p *UniswapV2Pool) TokenA() common.Address {
	return p.tokenA
}

func (p *UniswapV2Pool) TokenB() common.Address {
	return p.tokenB
}

func (p *UniswapV2Pool) HandleLog(aLog *types.Log) error {
	if aLog.Address != p.address {
		return nil
	}
	if aLog.Topics[0] != UniswapV2SyncEventHash {
		return nil
	}

	p.reserveA = new(big.Int).SetBytes(aLog.Data[0:32])
	p.reserveB = new(big.Int).SetBytes(aLog.Data[32:64])

	log.Debug("uniswapV2 log update", "event", "sync", "pool", p.address.Hex(), "reserveA", p.reserveA, "reserveB", p.reserveB)
	return nil
}

func (p *UniswapV2Pool) Output(inputToken common.Address, outputToken common.Address, input *big.Int) (*big.Int, error) {
	if input.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("input amount must be greater than 0")
	}
	if input.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), nil
	}
	var reserveIn, reserveOut *big.Int
	if inputToken == p.tokenA && outputToken == p.tokenB {
		reserveIn, reserveOut = p.reserveA, p.reserveB
	} else if inputToken == p.tokenB && outputToken == p.tokenA {
		reserveIn, reserveOut = p.reserveB, p.reserveA
	} else {
		return nil, fmt.Errorf("invalid input/output tokens - got %s,%s, expected %s,%s", inputToken.Hex(), outputToken.Hex(), p.tokenA.Hex(), p.tokenB.Hex())
	}
	amountInWithFee := new(big.Int).Mul(input, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, big.NewInt(1000)), amountInWithFee)
	return new(big.Int).Div(numerator, denominator), nil
}

func (p *UniswapV2Pool) SerializedPoolLength() int {
	return 60
}

func (p *UniswapV2Pool) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.Write(p.address.Bytes()); err != nil {
		return nil, err
	}
	if _, err := buf.Write(p.tokenA.Bytes()); err != nil {
		return nil, err
	}
	if _, err := buf.Write(p.tokenB.Bytes()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *UniswapV2Pool) Deserialize(data []byte) error {
	if len(data) != p.SerializedPoolLength() {
		return fmt.Errorf("cannot deserialize UniswapV2Pool from data of length %d, expect 60", len(data))
	}

	p.address = common.BytesToAddress(data[:20])
	p.tokenA = common.BytesToAddress(data[20:40])
	p.tokenB = common.BytesToAddress(data[40:60])

	return nil
}
