// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
contract RccToken is ERC20 {
    constructor() ERC20("RccToken", "RCC") {
        // 初始供应量可以在这里定义，或者留空以便之后通过 mint 函数铸造
        _mint(msg.sender, 10000000 * 1_000_000_000_000_000_000);
    }
}
