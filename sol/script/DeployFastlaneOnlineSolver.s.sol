// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import "../src/FastlaneOnlineSolver.sol";

contract DeployFastlaneOnlineSolverScript is Script {
    function run() external {
        console.log("\n=== DEPLOYING FastlaneOnlineSolver ===\n");

        uint256 deployerPrivateKey = vm.envUint("SOLVER1_PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        address atlasAddress = vm.envAddress("ATLAS_ADDRESS"); 
        address wethAddress = vm.envAddress("WETH_ADDRESS");

        console.log("===============================");
        console.log("Deploying FastlaneOnlineSolver...");
        console.log("Deployer address:", deployer);
        console.log("Atlas address:", atlasAddress);
        console.log("WETH address:", wethAddress);
        console.log("===============================");

        vm.startBroadcast(deployerPrivateKey);

        FastlaneOnlineSolver solver = new FastlaneOnlineSolver(
            wethAddress,
            atlasAddress,
            deployer
        );

        console.log("FastlaneOnlineSolver deployed to:", address(solver));

        vm.stopBroadcast();
    }
}
