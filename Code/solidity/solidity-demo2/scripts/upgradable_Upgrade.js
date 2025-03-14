


const { ethers, upgrades } = require('hardhat');

// const { verify } = require("../utils/verify")
async function main() {
    const contractFactory = await ethers.getContractFactory('Upgrade');
    let address;
    if (false) {
        console.log('Depoly...');
        const contract = await upgrades.deployProxy(contractFactory, [42], { initializer: 'store' });
        console.log('Depoly  to:', await contract.getAddress());
        const result = await contract.deploymentTransaction().wait(1);
        console.log('Depoly   to:', result.contractAddress);
        address = result.contractAddress;
    } else {
        console.log('Upgrade...');
        const contract = await upgrades.upgradeProxy("0x9f001C404c3a23244F7d14Bf94565a01725147a2", contractFactory, {});
        console.log('Upgrade deployed to:', await contract.getAddress());
        address = await contract.getAddress();
    }
    // await verify(address, []);
}
main().then(() =>
    process.exit(0)
).catch((error) => {
    console.error(error);
    process.exit(1);
});





// const contract = await upgrades.deployProxy(contractFactory, [42], { initializer: 'store' });
// const result = await contract.deploymentTransaction().wait(1);
// console.log('Upgrade deployed to:', result.contractAddress);

// // 升级合约的函数
// // scripts/deploy_upgradeable_box.js
// const { ethers, upgrades } = require('hardhat');
// async function upgradeContract() {
//     const MyContractV2 = await ethers.getContractFactory("MyContractV2");
//     const upgraded = await upgrades.upgradeProxy("0x...", MyContractV2); // 替换为你的代理合约地址
//     console.log("MyContract upgraded to V2 at:", upgraded.address);
// }
