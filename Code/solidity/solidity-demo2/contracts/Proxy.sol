// SPDX-Licence-Identifier: MIT

pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

interface ProxyInterface {
    function add() external;

    function sub() external;
}

contract Proxy {

    address public implementation;
    uint256 public x;

    function setImplementation(address _newImplementation) public {
        implementation = _newImplementation;

    }

    function call() internal {
        (bool success, bytes memory data) = implementation.delegatecall(
            msg.data
        );
        if (!success) {
            revert("proxy call failed");
        }
    }

    receive() external payable {
        call();
    }

    fallback() external payable {
        call();
    }
}

contract A {
    address public implementation;
    uint256 public x;

    function add() external returns (uint256) {
        return x = x + 1;
    }
}

contract B {
    address public implementation;
    uint256 public x;

    function add() external returns (uint256) {
        return x = x + 1;
    }

    function sub() external returns (uint256) {
        return x = x - 1;
    }
}
