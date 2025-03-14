
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
const upgradeModule = require("./upgradeModule");
module.exports = buildModule("DemoV2Module", (m) => {
    const { proxy } = m.useModule(upgradeModule);

    const demo = m.contractAt("DemoV2", proxy);

    return { demo };
});