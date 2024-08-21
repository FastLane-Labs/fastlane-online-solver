// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.22;

import { IUniswapV2Pair, IUniswapV3Pool, IUniswapV3Factory} from "./interfaces/IUniswap.sol";
import { IAtlas } from "src/interfaces/IAtlas.sol";
import { SafeTransferLib, ERC20 } from "solmate/utils/SafeTransferLib.sol";
import { console } from "forge-std/console.sol";

interface IWETH9 {
    function deposit() external payable;
    function withdraw(uint256 wad) external payable;
}

enum DexType {
    Null,
    UniswapV2,
    UniswapV3
}

struct Swap {
    DexType dexType;
    address poolAddr;
    address tokenIn;
    address tokenOut;
}

// ----------------------------------------------------------------------------
// WARNING: DO NOT STORE FUNDS IN THIS CONTRACT
// WARNING: DO NOT STORE FUNDS IN THIS CONTRACT
// WARNING: DO NOT STORE FUNDS IN THIS CONTRACT
// ----------------------------------------------------------------------------

contract FastlaneOnlineSolver {
    address public immutable WETH_ADDRESS;
    address internal immutable _owner;
    address internal immutable _atlas;

    constructor(
        address weth,
        address atlas,
        address owner
    )
    {
        WETH_ADDRESS = weth;
        _owner = owner;
        _atlas = atlas;
    }

    receive() external payable {}

    function atlasSolverCall(
        address solverOpFrom,
        address executionEnvironment,
        address bidToken,
        uint256 bidAmount,
        bytes calldata solverOpData,
        bytes calldata
    )
        external
        payable
        virtual
        safetyFirst(executionEnvironment, solverOpFrom)
        payBid(executionEnvironment, bidToken, bidAmount)
    {
        (bool success,) = address(this).call{ value: msg.value }(solverOpData);

        require(success, "CALL UNSUCCESSFUL");
    }

    modifier safetyFirst(address executionEnvironment, address solverOpFrom) {
        require(msg.sender == _atlas, "INVALID ENTRY");
        require(solverOpFrom == _owner, "INVALID CALLER");

        _;

        uint256 shortfall = IAtlas(_atlas).shortfall();

        if (shortfall < msg.value) shortfall = 0;
        else shortfall -= msg.value;

        if (msg.value > address(this).balance) {
            IWETH9(WETH_ADDRESS).withdraw(msg.value - address(this).balance);
        }

        IAtlas(_atlas).reconcile{ value: msg.value }(shortfall);
    }

    modifier payBid(address executionEnvironment, address bidToken, uint256 bidAmount) {

        _;

        if (bidToken == address(0)) {
            SafeTransferLib.safeTransferETH(executionEnvironment, bidAmount);
        } else {
            SafeTransferLib.safeTransfer(ERC20(bidToken), executionEnvironment, bidAmount);
        }
    }

    function execute(
        Swap[] calldata swapPath,
        address tokenUserSells,
        address tokenUserBuys,
        address bidToken,
        uint256 bidAmount
    ) external payable {
        uint256 amount = tokenUserSells == address(0) ? address(this).balance : ERC20(tokenUserSells).balanceOf(address(this));
        if (tokenUserSells == address(0)) {
            require(swapPath[0].tokenIn == WETH_ADDRESS, "swapPath[0].tokenIn != WETH_ADDRESS");
            IWETH9(WETH_ADDRESS).deposit{value: amount}();
        }

        for (uint256 i = 0; i < swapPath.length; i++) {
            amount = executeSwap(swapPath[i], amount);
        }

        require(amount >= bidAmount, "amountOut < bidAmount");

        if (bidToken == address(0)) {
            require (tokenUserBuys == address(0) || tokenUserBuys == WETH_ADDRESS, "Invalid tokenUserBuys for ETH bid");
            require (swapPath[swapPath.length - 1].tokenOut == WETH_ADDRESS, "swapPath[swapPath.length - 1].tokenOut != WETH_ADDRESS");
            IWETH9(WETH_ADDRESS).withdraw(amount);
        }

        uint256 excessTokenUserBuys = amount - bidAmount;

        if (bidToken == address(0)) {
            SafeTransferLib.safeTransferETH(_owner, excessTokenUserBuys);
        } else {
            SafeTransferLib.safeTransfer(ERC20(tokenUserBuys), _owner, excessTokenUserBuys);
        }
    }

    function executeSwap(Swap memory swap, uint256 amountIn) internal returns (uint256) {
        if (swap.dexType == DexType.UniswapV2) {
            IUniswapV2Pair pair = IUniswapV2Pair(swap.poolAddr);
            (uint112 reserve0, uint112 reserve1, ) = pair.getReserves();
            require(pair.token0() == swap.tokenIn || pair.token1() == swap.tokenIn, "Invalid tokenIn");
            require(pair.token0() == swap.tokenOut || pair.token1() == swap.tokenOut, "Invalid tokenOut");

            SafeTransferLib.safeTransfer(ERC20(swap.tokenIn), swap.poolAddr, amountIn);
            if (swap.tokenIn == pair.token0()) {
                uint amtOut = getUniswapV2AmountOut(amountIn, reserve0, reserve1);
                pair.swap(0, amtOut, address(this), bytes(""));
                return uint256(amtOut);
            } else {
                uint amtOut = getUniswapV2AmountOut(amountIn, reserve1, reserve0);
                pair.swap(amtOut, 0, address(this), bytes(""));
                return uint256(amtOut);
            }
        } else if (swap.dexType == DexType.UniswapV3) {
            uint160 MIN_SQRT_RATIO = 4295128739;
            uint160 MAX_SQRT_RATIO = 1461446703485210103287273052203988822378723970342;
            IUniswapV3Pool pool = IUniswapV3Pool(swap.poolAddr);
            require(pool.token0() == swap.tokenIn || pool.token1() == swap.tokenIn, "Invalid tokenIn");
            require(pool.token0() == swap.tokenOut || pool.token1() == swap.tokenOut, "Invalid tokenOut");

            if (swap.tokenIn == pool.token0()) {
                (,int256 amountOut) = pool.swap(
                    address(this), 
                    true, 
                    int256(amountIn),
                    MIN_SQRT_RATIO + 1,
                    abi.encode(swap)
                );
                return uint256(-amountOut);
            } else {
                (int256 amountOut,) = pool.swap(
                    address(this), 
                    false, 
                    int256(amountIn), 
                    MAX_SQRT_RATIO - 1, 
                    abi.encode(swap)
                );
                return uint256(-amountOut);
            }
        } else {
            revert("Invalid DexType");
        }
    }

    // vulnurable to attack: don't store funds in contract
    function uniswapV3SwapCallback(
      int256 amount0Delta,
      int256 amount1Delta,
      bytes calldata _data
    ) external {
        Swap memory swap = abi.decode(_data, (Swap));
        require(msg.sender == swap.poolAddr, "Invalid sender");
        IUniswapV3Pool pool = IUniswapV3Pool(swap.poolAddr);
        require(IUniswapV3Factory(pool.factory()).getPool(swap.tokenIn, swap.tokenOut, pool.fee()) == swap.poolAddr, "Invalid pool");
        require(amount0Delta > 0 || amount1Delta > 0, "Invalid amountDeltas");
        if (amount0Delta > 0) {
            SafeTransferLib.safeTransfer(ERC20(swap.tokenIn), swap.poolAddr, uint256(amount0Delta));
        } else {
            SafeTransferLib.safeTransfer(ERC20(swap.tokenIn), swap.poolAddr, uint256(amount1Delta));
        }
    }

    function getUniswapV2AmountOut(
        uint amountIn,
        uint reserveIn,
        uint reserveOut
    ) internal pure returns (uint256 amountOut) {
        uint amountInWithFee = amountIn * 997;
        uint numerator = amountInWithFee * reserveOut;
        uint denominator = (reserveIn * 1000) + amountInWithFee;
        amountOut = numerator / denominator;
    }
}
