// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");


module.exports = buildModule("ProxyModule", (m) => {
    let args = [];
    let options = {};
    const proxy = m.contract("Proxy", args, options);
    return { proxy };
});
