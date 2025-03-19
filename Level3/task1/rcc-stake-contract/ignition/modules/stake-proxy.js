const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
module.exports = buildModule("RCCStakeProxyModule", (m) => {
    const deployer = m.getAccount(0);
    const stake = m.contract("RCCStake", [], {
        from: deployer
    });

    const stakeProxy = m.contract("ERC1967Proxy", [
        stake,
        "0x",
    ], {
        from: deployer
    });
    const implementation = m.readEventArgument(
        stakeProxy,
        "Upgraded",
        "implementation",
    );

    return { stake, stakeProxy };
});