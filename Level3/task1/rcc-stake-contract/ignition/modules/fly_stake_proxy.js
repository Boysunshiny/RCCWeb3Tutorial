const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
module.exports = buildModule("RCCStakeProxyModule", (m) => {
    const stake = m.contract("FLYStake");
    const stakeProxy = m.contract("ERC1967Proxy", [
        stake,
        "0x",
    ]);
    const implementation = m.readEventArgument(
        stakeProxy,
        "Upgraded",
        "implementation",
    );
    console.log(`emit event --> IERC1967.Upgraded ${implementation.id}`);
    return { stake, stakeProxy };
});