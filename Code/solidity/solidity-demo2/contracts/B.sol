// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ProxyInterface.sol";

contract B is ProxyInterface {
    address public implementation;
    uint256 public x;

    function add() external returns (uint256) {
        return x = x + 1;
    }

    function sub() external returns (uint256) {
        return x = x - 1;
    }
}
