# fastlane-online-solver

This is a reference solver (searcher) implementation for FastLane Online, which is an `Atlas` (https://github.com/FastLane-Labs/atlas) powered RFQ dApp. It fulfills `SwapIntent`s generated as `UserOperation` transactions by the FastLane online front-end by generating `SolverOperation`s which in addition to fulfilling intents, aim to maximise profits for the solver.

## overview

* The solver connects to the blockchain mempool and listens for any targeted `UserOperation` transactions targeted for FastlaneOnline. These transactions have the method `fastOnlineSwap` and the `to` address is the `FastLaneOnline` contract.
* It tries to "solve" the `SwapIntent` exploring various swap paths which are made availble to it beforehand. It also keeps track of the state of the DEX pools in the swap paths.
* Once it finds an appropriate fulfillment of the `SwapIntent` (one where `amountTokenOut > minAmountTokenOut`), it generates a `SolverOperation` in which the `bid` amount is set based on certain configured parameters.
* It broadcasts this `SolverOperation` transaction with a gas price higher than the `UserOperation` transaction thereby frontrunning it.
* The FastlaneOnline `DAppControl` cache gets updated as a result of this `SolverOperation` transaction which is then used in the `metacall` of Atlas.

## configuration

The configuration parameters can be passed in the `config.yaml` file in the project root, as cli-flags or as environment variables. The precedence is as follows - environment variables > cli-flags > config.yaml. Some configuration parameters such as `SOLVER_PK` can only be passed in as environment variables. Other more detailed configurations such as `pools_config` and `interesting_tokens` are passed in the yaml config file.

## requirements

* Some bonded atlETH in the Atlas contract by the solver EOA.
* A mempool connection to listen to `SwapIntent` user operations.
* An ETH RPC connection to listen to chain events and broadcast transactions.
* Some native ETH for gas.

## modules

* **pools** - 
Includes all funcitonality related to DEX pools. Following DEXs are currently supported
    * Uniswap V2
    * Uniswap V3
It includes functionality to serialize, deserialize static and dynamic pool data.

* **events** -
Includes all functionality related to chain events such as RPC calls and mempool connections.

* **bot** -
Instantiate all components, connections and run the main event loop with provided configurations.

* **sol** - 
Smart contract implementation of the `FastlaneOnlineSolver` contract.

## run

Deploy the `FastlaneOnlineSolver` smart contract using - 
```
cd sol/
SOLVER1_PRIVATE_KEY=<SOLVER_PK> forge script script/DeployFastlaneOnlineSolver.s.sol:DeployFastlaneOnlineSolverScript --rpc-url <RPC_URL> --broadcast
```

Then generate pools and paths static data using
```
MODE=data_gen go run main.go
```
This will create pools and paths binary files as specified by the config parameters `pools_bin_file` and `swap_paths_bin_file`.

Run the bot using - 
```
SOLVER_PK=<SOLVER_PK> go run main.go
```

## deployment

After updating the configurations in `docker-compose.yml`, deploy using 

```
SOLVER_PK=<SOLVER_PK> MODE=run_bot CTX=<DEPLOY_CONTEXT> make deploy
```