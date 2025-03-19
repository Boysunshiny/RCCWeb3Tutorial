const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules")
const { ethers } = require('ethers');
const { ignition } = require("hardhat")
const RccTokenModule = require("./rcc");
const RCCStakeModule = require("./stake");
 
async function main() {
    let token = await ignition.deploy(RccTokenModule);
    let stake = await ignition.deploy(RCCStakeModule); 
    Promise.all([token, stake]).then((results) => {
        // 所有合约部署完成
        console.log("All contracts deployed successfully");
        console.log("RccToken address:", results[0].address);
        console.log("RCCStake address:", results[1].address);
    })  
}
   

    // const deployer = m.getAccount(0);
    // const stratBlock = 100;
    // const endBlock = 100;// * 10 ** 18;
    // const perBlock = 10;// * 10 ** 18;

    // const { token } = m.useModule(tokenModule);
    // // 获取合约地址

    // const { address } = m.contractAt("RccToken", token)

    // console.log("address----->", address);

    // const args = [address, stratBlock, endBlock, perBlock];

    // console.log("args----->", args);


    // // 假设 RccToken 有一个函数叫 "transfer"
    // const initializer = "initialize(address,uint256,uint256,uint256)"; // 函数签名 
    // const fragment = ethers.id(initializer).slice(0, 10);

    // const encodedArgs = ethers.AbiCoder.defaultAbiCoder().encode(
    //     ["address", "uint256", "uint256", "uint256"], // 参数类型
    //     [address, stratBlock, endBlock, perBlock] // 参数值
    // );
    // // console.log("fragment", fragment)


    // // // 创建 Interface 对象
    // // const abi = ["initialize(address,uint256,uint256,uint256)"];
    // // const iface = new ethers.Interface(abi);
    // // console.log("calldata", iface);
    // // // 生成 calldata
    // const calldata = fragment + encodedArgs.slice(2);





    // const stake = m.contract("RCCStake", [], {
    //     from: deployer
    // });


    // // stake.Interface.encodeFunctionData(fragment, args);
    // const stakeProxy = m.contract("ERC1967Proxy", [
    //     stake,
    //     "0x",
    // ], {
    //     from: deployer
    // });


    // const implementation = m.readEventArgument(
    //     stakeProxy,
    //     "Upgraded",
    //     "implementation",
    // );

 
 