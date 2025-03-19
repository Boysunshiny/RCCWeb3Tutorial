
const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
module.exports = buildModule("RccTokenModule", (m) => {
    const deployer = m.getAccount(0);
    const args = [];
    const options = {
        from: deployer,
    };
    const token = m.contract("RccToken", args, options)
    return { token };
}); 