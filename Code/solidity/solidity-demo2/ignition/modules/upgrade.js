// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");


module.exports = buildModule("UpgradeModule", (m) => {

    const proxyAdminOwner = m.getAccount(0);

    let args = [];
    let options = {};
    const upgrade = m.contract("Upgrade", args, options);

    const proxy = m.contract("TransparentUpgradeableProxy", [
        upgrade,
        proxyAdminOwner,
        "0x",
    ]);

    const proxyAdminAddress = m.readEventArgument(
        proxy,
        "AdminChanged",
        "newAdmin"
    );

    const proxyAdmin = m.contractAt("ProxyAdmin", proxyAdminAddress);

    return { proxyAdmin, proxy };


});


