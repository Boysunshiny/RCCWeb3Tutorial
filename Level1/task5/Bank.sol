// SPDX-License-Identifier:  MIT
pragma solidity ^0.8.7;

 

event Bank__Store(address indexed _from, uint256 _value);

contract Bank {
    /*State Variable */
    address public immutable owner;


    /*Functions */
    constructor() {
        owner = msg.sender;
    } 


    receive() external payable {
       emit  Bank__Store( msg.sender, msg.value);
    }
 
    function withDraw()  public  {
        require(msg.sender == owner, "You are not the owner");
        payable(msg.sender).transfer(address(this).balance);
    }

    function getBalance() public view   returns(uint256){
        return address(this).balance;
 
    }

}
