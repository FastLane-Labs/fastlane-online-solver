//SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.22;

interface IAtlas {
    function shortfall() external view returns (uint256);
    function reconcile(uint256 maxApprovedGasSpend) external payable returns (uint256 owed);
}
