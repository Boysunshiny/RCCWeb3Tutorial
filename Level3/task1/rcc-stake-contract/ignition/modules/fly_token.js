
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
module.exports = buildModule("FLYTokenModule", (m) => {
    const args = ["FLY Token", "FLY", 10000000];
    const options = {};
    const token = m.contract("FLYToken", args, options)
    return { token };
});