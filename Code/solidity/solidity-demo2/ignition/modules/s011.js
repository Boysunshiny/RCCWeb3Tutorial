// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");


module.exports = buildModule("S011Module", (m) => {
    let args = [];
    let options = {};
    const s011 = m.contract("S011", args, options);
    return { s011 };
});
