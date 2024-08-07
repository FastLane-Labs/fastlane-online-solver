package pools

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type PoolType byte

var (
	UniswapV2PoolType = PoolType(1)
	UniswapV3PoolType = PoolType(2)
)

type Pool interface {
	PoolType() PoolType      //type of the pool
	Address() common.Address //address of the pool
	TokenA() common.Address  //address of token A
	TokenB() common.Address  //address of token B

	HandleLog(*types.Log) error // handle a log event

	Output(inputToken common.Address,
		outputToken common.Address,
		amountIn *big.Int) (*big.Int, error) //output amount of outputToken for amountIn of inputToken

	SerializedPoolLength() int  //length of the serialized pool
	Serialize() ([]byte, error) //serialize the pool
	Deserialize([]byte) error   //deserialize the pool
}
