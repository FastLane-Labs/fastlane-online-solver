// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.22;

import "forge-std/Test.sol";
import "../src/FastlaneOnlineSolver.sol";
import { IUniswapV2Pair, IUniswapV3Pool, IUniswapV3Factory } from "../src/interfaces/IUniswap.sol";
import { SafeTransferLib, ERC20 } from "solmate/utils/SafeTransferLib.sol";

//anvil --fork-url <POLYGON_RPC_URL> --fork-block-number 60144000 --port 8546
//forge test -vv
contract FastlaneOnlineSolverTest is Test {
    address constant atlasAddress = address(1);
    address constant wethAddress = address(2);

    address constant tokenA = 0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6; //WBTC //8 decimals
    address constant tokenB = 0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619; //WETH // 18 decimals
    address constant tokenC = 0x53E0bca35eC356BD5ddDFebbD1Fc0fD03FaBad39; //LINK // 18 decimals

    address constant poolAB = 0x50eaEDB835021E4A108B7290636d62E9765cc6d7; // uniswap v3 pool
    address constant poolBC = 0x5cA6CA6c3709E1E6CFe74a50Cf6B2B6BA2Dadd67; // uniswap v2 pair
    address constant poolBC2 = 0x3e31AB7f37c048FC6574189135D108df80F0ea26; // uniswap v3 pair

    FastlaneOnlineSolver public solver;

    function setUp() public {
        vm.createSelectFork("http://localhost:8546");

        solver = new FastlaneOnlineSolver(wethAddress, atlasAddress, address(this));
    }

    function test_execute_forward() public {
        deal(tokenA, address(solver), 10e8);

        Swap[] memory swapPath = new Swap[](2);
        swapPath[0] = Swap(DexType.UniswapV3, poolAB, tokenA, tokenB);
        swapPath[1] = Swap(DexType.UniswapV2, poolBC, tokenB, tokenC);


        uint256 bidAmount = 1e17;
        solver.execute(
            swapPath,
            tokenA,
            tokenC,
            tokenC,
            bidAmount
        );

        uint256 amountOnContract = ERC20(tokenC).balanceOf(address(solver));

        assertEq(amountOnContract, bidAmount);
    }

    function test_execute_reverse() public {
        deal(tokenC, address(solver), 1e22);

        Swap[] memory swapPath = new Swap[](2);
        swapPath[0] = Swap(DexType.UniswapV2, poolBC, tokenC, tokenB);
        swapPath[1] = Swap(DexType.UniswapV3, poolAB, tokenB, tokenA);

        uint256 bidAmount = 1e5;
        solver.execute(
            swapPath,
            tokenC,
            tokenA,
            tokenA,
            bidAmount
        );

        uint256 amountOnContract = ERC20(tokenA).balanceOf(address(solver));

        assertEq(amountOnContract, bidAmount);
    }

    function test_execute_v3_v3() public {
        deal(tokenC, address(solver), 10000e18);

        Swap[] memory swapPath = new Swap[](2);
        swapPath[0] = Swap(DexType.UniswapV3, poolBC2, tokenC, tokenB);
        swapPath[1] = Swap(DexType.UniswapV3, poolAB, tokenB, tokenA);


        uint256 bidAmount = 1e5;

        solver.execute(
            swapPath,
            tokenC,
            tokenA,
            tokenA,
            bidAmount
        );

        uint256 amountOnContract = ERC20(tokenA).balanceOf(address(solver));
        uint256 amountOnEOA = ERC20(tokenA).balanceOf(address(this));

        console.log("total swap output", amountOnContract + amountOnEOA);
        assertEq(amountOnContract, bidAmount);
    }
}
