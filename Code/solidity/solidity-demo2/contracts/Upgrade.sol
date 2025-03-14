// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

event ValueChanged(uint256 value);
contract Upgrade {
    uint256 private _value; // Emitted when the stored value changesevent   ValueChanged(uint256 value);// Stores a new value in the contract

    function store(uint256 value) public {
        _value = value;
        emit ValueChanged(value);
    } // Reads the last stored valuefunction retrieve() public view returns (uint256) {return _value;
}
