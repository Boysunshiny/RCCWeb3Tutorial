
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
module.exports = buildModule("RCCStakeModule", (m) => {
    const deployer = m.getAccount(0);
    const args = [];
    const options = {
        from: deployer,
    };
    const stake = m.contract("RCCStake", args, options)
    return { stake };
});