
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
const stakeProxyModule = require("./fly_stake_proxy");
const tokenModule = require("./fly_token");
module.exports = buildModule("FLYStakeModule", (m) => {


    const stratBlock = 0;
    const endBlock = 10000000;// * 10 ** 18;
    const perBlock = 1;// * 10 ** 18;

    const { token } = m.useModule(tokenModule);
    const { address } = m.contractAt("FLYToken", token);
    const { stake, stakeProxy } = m.useModule(stakeProxyModule);


    m.call(stake, "initialize", [address, stratBlock, endBlock, perBlock])

    return { token, stake, stakeProxy };
});