// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Transfer {
    function transfer(uint _amount) public payable {
        payable(msg.sender).transfer(_amount);
    }

    function send(uint _amount) public payable {
        bool success = payable(msg.sender).send(_amount);
        require(success, "Send Failed");
    }

    function call(uint _amount) public payable {
        (bool success, ) = msg.sender.call{value: _amount}("");
        if (!success) revert("Unable to send eth");
    }
}
