// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { ethers } = require("hardhat");


module.exports = buildModule("MyTokenModule", (m) => {
    const args = [
        "My Token",
        "MT",
    ];
    const amount = ethers.parseEther("0");

    const myToken = m.contract("MyToken", args, {
        value: amount,
    });

    return { myToken };
});
