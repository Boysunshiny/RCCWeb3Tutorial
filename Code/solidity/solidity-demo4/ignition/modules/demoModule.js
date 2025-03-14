
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
const proxyModule = require("./proxyModule");
module.exports = buildModule("DemoModule", (m) => {
    const { proxy, proxyAdmin } = m.useModule(proxyModule);

    const demo = m.contractAt("Demo", proxy);

    return { demo, proxy, proxyAdmin };
});