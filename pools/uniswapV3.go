package pools

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/FastLane-Labs/fastlane-online-solver/contract/multicall"
	"github.com/FastLane-Labs/fastlane-online-solver/contract/uniswapV3Factory"
	uniswapV3PoolContract "github.com/FastLane-Labs/fastlane-online-solver/contract/uniswapV3Pool"
	"github.com/FastLane-Labs/fastlane-online-solver/log"
	"github.com/FastLane-Labs/fastlane-online-solver/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ Pool = &UniswapV3Pool{}

var (
	UniswapV3MintEventHash = common.HexToHash("0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde")
	UniswapV3BurnEventHash = common.HexToHash("0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c")
	UniswapV3SwapEventHash = common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67")
)

var (
	Default_unsiwapV3_num_ticks = int64(5)
)

var (
	UniswapV3PoolAbi, _    = uniswapV3PoolContract.UniswapV3PoolMetaData.GetAbi()
	UniswapV3FactoryAbi, _ = uniswapV3Factory.UniswapV3FactoryMetaData.GetAbi()
)

type UniswapV3Pool struct {
	address     common.Address
	tokenA      common.Address
	tokenB      common.Address
	tickSpacing int64
	fee         int64

	tick        int64
	liquidity   *big.Int
	sqrtPrice96 *big.Int

	liquidityDeltas map[int64]*big.Int
	mu              sync.Mutex

	stateLoaderStatus *stateLoaderStatus
}

type stateLoaderStatus struct {
	ethClient                *ethclient.Client
	maxConcurrentRpcCalls    int
	liquidityDeltasLoadBlock uint64
}

func NewUniswapV3Pool(address common.Address) *UniswapV3Pool {
	return &UniswapV3Pool{
		address:         address,
		liquidityDeltas: make(map[int64]*big.Int),
	}
}

func (p *UniswapV3Pool) PoolType() PoolType {
	return UniswapV3PoolType
}
func (p *UniswapV3Pool) Address() common.Address {
	return p.address
}
func (p *UniswapV3Pool) TokenA() common.Address {
	return p.tokenA
}
func (p *UniswapV3Pool) TokenB() common.Address {
	return p.tokenB
}

func (p *UniswapV3Pool) HandleLog(aLog *types.Log) error {
	if aLog.Address != p.address {
		return nil
	}
	if aLog.Topics[0] == UniswapV3MintEventHash {

		res, err := UniswapV3PoolAbi.Unpack("Mint", aLog.Data)
		if err != nil {
			return fmt.Errorf("failed to unpack Mint: %v", err)
		}
		tickLower := new(big.Int).SetBytes(aLog.Topics[2].Bytes()).Int64()
		tickUpper := new(big.Int).SetBytes(aLog.Topics[3].Bytes()).Int64()
		liquidityAmount := res[1].(*big.Int)

		p.mu.Lock()
		if _, ok := p.liquidityDeltas[tickLower]; ok {
			p.liquidityDeltas[tickLower].Add(p.liquidityDeltas[tickLower], liquidityAmount)
		}
		if _, ok := p.liquidityDeltas[tickUpper]; ok {
			p.liquidityDeltas[tickUpper].Sub(p.liquidityDeltas[tickUpper], liquidityAmount)
		}
		if p.tick >= tickLower && p.tick < tickUpper {
			p.liquidity.Add(p.liquidity, liquidityAmount)
		}
		p.mu.Unlock()
		log.Debug("uniswapV3 log update",
			"event", "mint",
			"pool", p.address.Hex(),
			"tickLower", tickLower,
			"tickUpper", tickUpper,
			"liquidityAmount", liquidityAmount,
		)
	} else if aLog.Topics[0] == UniswapV3BurnEventHash {

		res, err := UniswapV3PoolAbi.Unpack("Burn", aLog.Data)
		if err != nil {
			return fmt.Errorf("failed to unpack Burn: %v", err)
		}
		tickLower := new(big.Int).SetBytes(aLog.Topics[2].Bytes()).Int64()
		tickUpper := new(big.Int).SetBytes(aLog.Topics[3].Bytes()).Int64()
		liquidityAmount := res[1].(*big.Int)

		p.mu.Lock()
		if _, ok := p.liquidityDeltas[tickLower]; ok {
			p.liquidityDeltas[tickLower].Sub(p.liquidityDeltas[tickLower], liquidityAmount)
		}
		if _, ok := p.liquidityDeltas[tickUpper]; ok {
			p.liquidityDeltas[tickUpper].Add(p.liquidityDeltas[tickUpper], liquidityAmount)
		}
		if p.tick >= tickLower && p.tick < tickUpper {
			p.liquidity.Sub(p.liquidity, liquidityAmount)
		}
		p.mu.Unlock()
		log.Debug("uniswapV3 log update",
			"event", "burn",
			"pool", p.address.Hex(),
			"tickLower", tickLower,
			"tickUpper", tickUpper,
			"liquidityAmount", liquidityAmount,
		)
	} else if aLog.Topics[0] == UniswapV3SwapEventHash {
		res, err := UniswapV3PoolAbi.Unpack("Swap", aLog.Data)
		if err != nil {
			return fmt.Errorf("failed to unpack Swap: %v", err)
		}
		p.sqrtPrice96 = res[2].(*big.Int)
		oldLiquidity := big.NewInt(0).Set(p.liquidity)
		p.liquidity = res[3].(*big.Int)
		unspacedTick := (res[4].(*big.Int)).Int64()
		spacedTick := p.spaceTick(unspacedTick)
		prevTick := p.tick
		p.tick = spacedTick
		if prevTick == p.tick {
			if p.liquidity.Cmp(oldLiquidity) != 0 {
				log.Error("liquidity changed but tick stayed the same", "pool", p.address.Hex(), "liquidity", oldLiquidity, "newLiquidity", p.liquidity)
			}
		} else {
			if p.stateLoaderStatus == nil {
				panic("stateLoaderStatus is nil while handling log")
			}
			if aLog.BlockNumber > p.stateLoaderStatus.liquidityDeltasLoadBlock {
				go p.loadLiquidityDeltas(p.stateLoaderStatus.ethClient, p.stateLoaderStatus.maxConcurrentRpcCalls)
				p.stateLoaderStatus.liquidityDeltasLoadBlock = aLog.BlockNumber
			}
		}
	}
	log.Debug("uniswapV3 log update",
		"event", "swap",
		"pool", p.address.Hex(),
		"tick", p.tick,
		"liquidity", p.liquidity,
		"sqrtPrice96", p.sqrtPrice96,
	)
	return nil
}

func (p *UniswapV3Pool) loadLiquidityDeltas(ethClient *ethclient.Client, maxConcurrentRpcCalls int) error {
	log.Debug("loading uniswap v3 liquidity deltas", "pool", p.address.Hex())
	startTime := time.Now()

	totalTicks := 2*Default_unsiwapV3_num_ticks + 1

	indices := make([]int, 0)
	for i := 0; i < int(totalTicks); i++ {
		indices = append(indices, i)
	}

	tickAtIndex := func(index int) int64 {
		return p.tick - p.tickSpacing*int64(Default_unsiwapV3_num_ticks) + p.tickSpacing*int64(index)
	}

	calldataBatchGenerator := func(index int) ([]multicall.Multicall3Call, error) {
		tick := tickAtIndex(index)
		calldata, err := UniswapV3PoolAbi.Pack("ticks", big.NewInt(tick))
		if err != nil {
			return nil, fmt.Errorf("failed to pack ticks: %v", err)
		}
		return []multicall.Multicall3Call{
			{
				Target:   p.address,
				CallData: calldata,
			},
		}, nil
	}

	returnDataBatchHandler := func(index int, returnDatas [][]byte) error {
		if len(returnDatas) != 1 {
			return fmt.Errorf("unexpected return data length: %d, expected 1", len(returnDatas))
		}
		res, err := UniswapV3PoolAbi.Unpack("ticks", returnDatas[0])
		if err != nil {
			return fmt.Errorf("failed to unpack ticks: %v", err)
		}
		liquidityDelta := res[1].(*big.Int)
		tick := tickAtIndex(index)
		p.mu.Lock()
		p.liquidityDeltas[tick] = liquidityDelta
		p.mu.Unlock()
		return nil
	}

	if err := utils.Multicall(ethClient,
		maxConcurrentRpcCalls,
		1000,
		calldataBatchGenerator,
		returnDataBatchHandler,
		indices,
	); err != nil {
		return fmt.Errorf("failed to multicall for liquidity deltas: %v", err)
	}

	log.Debug("loaded uniswap v3 liquidity deltas",
		"pool", p.address.Hex(),
		"numTicksLoaded", totalTicks,
		"took", time.Since(startTime),
	)
	return nil
}

func (p *UniswapV3Pool) spaceTick(tick int64) int64 {
	if tick%p.tickSpacing == 0 {
		return tick
	} else {
		tmp := tick / p.tickSpacing
		if tick < 0 {
			tmp--
		}
		return tmp * p.tickSpacing
	}
}

func (p *UniswapV3Pool) SerializedPoolLength() int {
	return 76
}

func (p *UniswapV3Pool) Serialize() ([]byte, error) {
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
	tickSpacingBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(tickSpacingBytes, uint64(p.tickSpacing))
	if _, err := buf.Write(tickSpacingBytes); err != nil {
		return nil, err
	}
	feeBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(feeBytes, uint64(p.fee))
	if _, err := buf.Write(feeBytes); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *UniswapV3Pool) Deserialize(data []byte) error {
	if len(data) != p.SerializedPoolLength() {
		return fmt.Errorf("cannot deserialize UniswapV3Pool from data of length %d, expect 76", len(data))
	}
	p.address = common.BytesToAddress(data[0:20])
	p.tokenA = common.BytesToAddress(data[20:40])
	p.tokenB = common.BytesToAddress(data[40:60])
	p.tickSpacing = int64(binary.LittleEndian.Uint64(data[60:68]))
	p.fee = int64(binary.LittleEndian.Uint64(data[68:76]))
	p.liquidityDeltas = make(map[int64]*big.Int)
	return nil
}

func (p *UniswapV3Pool) Output(inputToken common.Address, outputToken common.Address, input *big.Int) (*big.Int, error) {
	tokenAtotokenB := p.tokenA == inputToken && p.tokenB == outputToken
	tokenBtotokenA := p.tokenA == outputToken && p.tokenB == inputToken
	if !tokenAtotokenB && !tokenBtotokenA {
		return nil, fmt.Errorf("invalid input/output tokens - got %s,%s, expected %s,%s", inputToken.Hex(), outputToken.Hex(), p.tokenA.Hex(), p.tokenB.Hex())
	}

	inputRemaining := big.NewInt(0).Set(input)
	output := big.NewInt(0)
	currPrice := big.NewInt(0).Set(p.sqrtPrice96)
	currTick := p.tick
	liq := big.NewInt(0).Set(p.liquidity)

	if tokenBtotokenA {
		for {
			inputRemainingLessFee := mulDiv(inputRemaining, big.NewInt(1000_000-p.fee), big.NewInt(1000_000))
			nextTick := currTick + p.tickSpacing
			nextTickPrice := priceAtTick(nextTick)
			dy := deltaYFromPriceDiff(liq, currPrice, nextTickPrice, true)
			if dy.Cmp(inputRemainingLessFee) >= 0 {
				finalPrice := finalPriceFromInput(inputRemainingLessFee, liq, currPrice, false)
				if finalPrice.Cmp(currPrice) < 0 {
					panic("final price less than curr price")
				}
				dx := deltaXFromPriceDiff(liq, currPrice, finalPrice, false)
				output.Add(output, dx)
				break
			} else {
				dx := deltaXFromPriceDiff(liq, currPrice, nextTickPrice, false)
				output.Add(output, dx)
				feeAmount := mulDivRoundingUp(dy, big.NewInt(p.fee), big.NewInt(1000_000-p.fee))
				inputRemaining.Sub(inputRemaining, dy.Add(dy, feeAmount))

				p.mu.Lock()
				if nextLiqDelta, ok := p.liquidityDeltas[nextTick]; ok {
					liq.Add(liq, nextLiqDelta)
				} else {
					swappedAmtPercent := new(big.Int).Mul(inputRemaining, big.NewInt(1000))
					swappedAmtPercent.Div(swappedAmtPercent, input)
					p.mu.Unlock()
					return nil, fmt.Errorf("swap too big - not enough liquidity stored: missing liquidity delta for tick %d, amount swapped: %f%%", nextTick, float64(swappedAmtPercent.Int64())/10)
				}
				p.mu.Unlock()

				currPrice = nextTickPrice
				currTick = nextTick
			}
		}
	} else {
		for {
			inputRemainingLessFee := mulDiv(inputRemaining, big.NewInt(1000_000-p.fee), big.NewInt(1000_000))
			prevTick := currTick
			prevTickPrice := priceAtTick(prevTick)
			dx := deltaXFromPriceDiff(liq, currPrice, prevTickPrice, true)
			if dx.Cmp(inputRemainingLessFee) >= 0 {
				finalPrice := finalPriceFromInput(inputRemainingLessFee, liq, currPrice, true)
				if finalPrice.Cmp(currPrice) > 0 {
					panic("final price greater than curr price")
				}
				dy := deltaYFromPriceDiff(liq, currPrice, finalPrice, false)
				output.Add(output, dy)
				break
			} else {
				dy := deltaYFromPriceDiff(liq, currPrice, prevTickPrice, false)
				output.Add(output, dy)
				feeAmount := mulDivRoundingUp(dx, big.NewInt(p.fee), big.NewInt(1000_000-p.fee))
				inputRemaining.Sub(inputRemaining, dx.Add(dx, feeAmount))

				p.mu.Lock()
				if prevLiqDelta, ok := p.liquidityDeltas[prevTick]; ok {
					liq.Sub(liq, prevLiqDelta)
				} else {
					swappedAmtPercent := new(big.Int).Mul(inputRemaining, big.NewInt(1000))
					swappedAmtPercent.Div(swappedAmtPercent, input)
					p.mu.Unlock()
					return nil, fmt.Errorf("swap too big - not enough liquidity stored: missing liquidity delta for tick %d amount swapped: %f%%", prevTick, float64(swappedAmtPercent.Int64())/10)
				}
				p.mu.Unlock()

				currPrice = prevTickPrice
				currTick = prevTick - p.tickSpacing
			}
		}
	}
	return output, nil
}

func deltaYFromPriceDiff(liquidity, currPrice, finalPrice *big.Int, roundUp bool) *big.Int {
	diff := big.NewInt(0).Sub(finalPrice, currPrice)
	diff.Abs(diff)
	if roundUp {
		return mulDivRoundingUp(liquidity, diff, new(big.Int).Lsh(big.NewInt(1), 96))
	} else {
		return mulDiv(liquidity, diff, new(big.Int).Lsh(big.NewInt(1), 96))
	}
}

func deltaXFromPriceDiff(liquidity, currPrice, finalPrice *big.Int, roundUp bool) *big.Int {
	var bigPrice, smallPrice *big.Int
	if finalPrice.Cmp(currPrice) > 0 {
		bigPrice = finalPrice
		smallPrice = currPrice
	} else {
		bigPrice = currPrice
		smallPrice = finalPrice
	}
	numer1 := new(big.Int).Lsh(liquidity, 96)
	numer2 := new(big.Int).Sub(bigPrice, smallPrice)
	if roundUp {
		r1 := mulDivRoundingUp(numer1, numer2, bigPrice)
		r2 := new(big.Int).Div(r1, smallPrice)
		if new(big.Int).Mod(r1, smallPrice).Cmp(big.NewInt(0)) > 0 {
			r2.Add(r2, big.NewInt(1))
		}
		return r2
	} else {
		return new(big.Int).Div(mulDiv(numer1, numer2, bigPrice), smallPrice)
	}
}

func finalPriceFromDeltaY(dy, liquidity, initialPrice *big.Int, add bool) *big.Int {
	if liquidity.Cmp(big.NewInt(0)) == 0 {
		return initialPrice
	} else {
		uint160Max := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 161), big.NewInt(1))
		quotient := new(big.Int)
		if dy.Cmp(uint160Max) <= 0 {
			quotient = new(big.Int).Div(new(big.Int).Lsh(dy, 96), liquidity)
		} else {
			quotient = mulDiv(dy, new(big.Int).Lsh(big.NewInt(1), 96), liquidity)
		}
		if add {
			return new(big.Int).Add(quotient, initialPrice)
		} else {
			quotientRounded := new(big.Int)
			if dy.Cmp(uint160Max) <= 0 {
				r1 := new(big.Int).Div(new(big.Int).Lsh(dy, 96), liquidity)
				if new(big.Int).Mod(new(big.Int).Lsh(dy, 96), liquidity).Cmp(big.NewInt(0)) > 0 {
					quotientRounded = new(big.Int).Add(r1, big.NewInt(1))
				} else {
					quotientRounded = r1
				}
			} else {
				quotientRounded = mulDivRoundingUp(dy, new(big.Int).Lsh(big.NewInt(1), 96), liquidity)
			}
			return new(big.Int).Sub(initialPrice, quotientRounded)
		}
	}
}

func finalPriceFromDeltaX(dx, liquidity, initialPrice *big.Int, add bool) *big.Int {
	if liquidity.Cmp(big.NewInt(0)) == 0 {
		return initialPrice
	} else {
		numer1 := new(big.Int).Lsh(liquidity, 96)
		if add {
			product := new(big.Int).Mul(dx, initialPrice)
			if new(big.Int).Div(product, dx).Cmp(initialPrice) == 0 {
				denom := new(big.Int).Add(numer1, product)
				if denom.Cmp(numer1) >= 0 {
					return mulDivRoundingUp(numer1, initialPrice, denom)
				}
			}
			d1 := new(big.Int).Add(new(big.Int).Div(numer1, initialPrice), dx)
			r := new(big.Int).Div(numer1, d1)
			if new(big.Int).Mod(numer1, d1).Cmp(big.NewInt(0)) > 0 {
				r.Add(r, big.NewInt(1))
			}
			return r
		} else {
			product := new(big.Int).Mul(dx, initialPrice)
			denom := new(big.Int).Sub(numer1, product)
			return mulDivRoundingUp(numer1, initialPrice, denom)
		}
	}
}

func finalPriceFromInput(input, liquidity, initialPrice *big.Int, zeroForOne bool) *big.Int {
	if zeroForOne {
		return finalPriceFromDeltaX(input, liquidity, initialPrice, true)
	} else {
		return finalPriceFromDeltaY(input, liquidity, initialPrice, true)
	}
}

func mulDiv(a, b, c *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(a, b), c)
}

func mulDivRoundingUp(a, b, c *big.Int) *big.Int {
	numer := new(big.Int).Mul(a, b)
	quo := new(big.Int).Div(numer, c)
	rem := new(big.Int).Mod(numer, c)
	if rem.Cmp(big.NewInt(0)) > 0 {
		quo.Add(quo, big.NewInt(1))
	}
	return quo
}

var (
	uv3_c_00, _ = big.NewInt(0).SetString("fffcb933bd6fad37aa2d162d1a594001", 16)
	uv3_c_01, _ = big.NewInt(0).SetString("100000000000000000000000000000000", 16)
	uv3_c_1, _  = big.NewInt(0).SetString("fff97272373d413259a46990580e213a", 16)
	uv3_c_2, _  = big.NewInt(0).SetString("fff2e50f5f656932ef12357cf3c7fdcc", 16)
	uv3_c_3, _  = big.NewInt(0).SetString("ffe5caca7e10e4e61c3624eaa0941cd0", 16)
	uv3_c_4, _  = big.NewInt(0).SetString("ffcb9843d60f6159c9db58835c926644", 16)
	uv3_c_5, _  = big.NewInt(0).SetString("ff973b41fa98c081472e6896dfb254c0", 16)
	uv3_c_6, _  = big.NewInt(0).SetString("ff2ea16466c96a3843ec78b326b52861", 16)
	uv3_c_7, _  = big.NewInt(0).SetString("fe5dee046a99a2a811c461f1969c3053", 16)
	uv3_c_8, _  = big.NewInt(0).SetString("fcbe86c7900a88aedcffc83b479aa3a4", 16)
	uv3_c_9, _  = big.NewInt(0).SetString("f987a7253ac413176f2b074cf7815e54", 16)
	uv3_c_10, _ = big.NewInt(0).SetString("f3392b0822b70005940c7a398e4b70f3", 16)
	uv3_c_11, _ = big.NewInt(0).SetString("e7159475a2c29b7443b29c7fa6e889d9", 16)
	uv3_c_12, _ = big.NewInt(0).SetString("d097f3bdfd2022b8845ad8f792aa5825", 16)
	uv3_c_13, _ = big.NewInt(0).SetString("a9f746462d870fdf8a65dc1f90e061e5", 16)
	uv3_c_14, _ = big.NewInt(0).SetString("70d869a156d2a1b890bb3df62baf32f7", 16)
	uv3_c_15, _ = big.NewInt(0).SetString("31be135f97d08fd981231505542fcfa6", 16)
	uv3_c_16, _ = big.NewInt(0).SetString("9aa508b5b7a84e1c677de54f3e99bc9", 16)
	uv3_c_17, _ = big.NewInt(0).SetString("5d6af8dedb81196699c329225ee604", 16)
	uv3_c_18, _ = big.NewInt(0).SetString("2216e584f5fa1ea926041bedfe98", 16)
	uv3_c_19, _ = big.NewInt(0).SetString("48a170391f7dc42444e8fa2", 16)
)

func priceAtTick(tick int64) *big.Int {
	absTick := big.NewInt(tick)
	if tick < 0 {
		absTick = big.NewInt(-tick)
	}

	comparer := big.NewInt(1)
	ratio := new(big.Int)

	if new(big.Int).And(absTick, comparer).Cmp(big.NewInt(0)) != 0 {
		ratio.Set(uv3_c_00)
	} else {
		ratio.Set(uv3_c_01)
	}

	constants := []*big.Int{
		uv3_c_1, uv3_c_2, uv3_c_3, uv3_c_4, uv3_c_5,
		uv3_c_6, uv3_c_7, uv3_c_8, uv3_c_9, uv3_c_10,
		uv3_c_11, uv3_c_12, uv3_c_13, uv3_c_14, uv3_c_15,
		uv3_c_16, uv3_c_17, uv3_c_18, uv3_c_19,
	}

	for _, c := range constants {
		comparer.Lsh(comparer, 1)
		if new(big.Int).And(absTick, comparer).Cmp(big.NewInt(0)) != 0 {
			ratio.Mul(ratio, c)
			ratio.Rsh(ratio, 128)
		}
	}

	if tick > 0 {
		maxValue := new(big.Int).Lsh(big.NewInt(1), 256)
		ratio = new(big.Int).Div(new(big.Int).Sub(maxValue, big.NewInt(1)), ratio)
	}

	sqrtPriceX96 := new(big.Int).Rsh(ratio, 32)
	if new(big.Int).Mod(ratio, new(big.Int).Lsh(big.NewInt(1), 32)).Cmp(big.NewInt(0)) != 0 {
		sqrtPriceX96.Add(sqrtPriceX96, big.NewInt(1))
	}

	return sqrtPriceX96
}
