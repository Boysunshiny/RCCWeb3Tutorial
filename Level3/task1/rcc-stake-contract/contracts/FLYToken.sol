// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract FLYToken is ERC20 {
    constructor(
        string memory name,
        string memory symbol,
        uint256 initialSupply
    ) ERC20(name, symbol) {
        // 初始供应量可以在这里定义，或者留空以便之后通过 mint 函数铸造
        _mint(msg.sender, initialSupply * 1_000_000_000_000_000_000);
    }
}
