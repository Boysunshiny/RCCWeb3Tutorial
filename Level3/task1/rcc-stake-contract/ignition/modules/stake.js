
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
const stakeProxyModule = require("./stake-proxy");
const tokenModule = require("./rcc");
module.exports = buildModule("RCCStakeModule", (m) => {


    const stratBlock = 6529999;
    const endBlock = 9529999;// * 10 ** 18;
    const perBlock = "20000000000000000";// * 10 ** 18;

    const { token } = m.useModule(tokenModule);
    const { address } = m.contractAt("RccToken", token);
    const { stake, stakeProxy } = m.useModule(stakeProxyModule);

    const deployer = m.getAccount(0);
    const stakev2 = m.contractAt("RCCStake", stakeProxy);
    m.call(stake, "initialize", [address, stratBlock, endBlock, perBlock], {
        from: deployer
    })

    return { token, stake, stakeProxy };
});