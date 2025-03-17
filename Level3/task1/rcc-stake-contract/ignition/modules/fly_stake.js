
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
const stakeProxyModule = require("./fly_stake_proxy");
const tokenModule = require("./fly_token");
module.exports = buildModule("FLYStakeModule", (m) => {


    const defaultAdminRole = m.getAccount(0);
    const upgradedRole = m.getAccount(0);
    const adminRole = m.getAccount(0);

    const { token } = m.useModule(tokenModule);
    const { address } = m.contractAt("FLYToken", token);
    const { stake, stakeProxy } = m.useModule(stakeProxyModule);
    m.call(stake, "initialize", [address, defaultAdminRole, upgradedRole, adminRole])

    return { token, stake, stakeProxy };
});