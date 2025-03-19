const { ethers, upgrades, ignition } = require("hardhat");

const { developmentChains } = require("./config");
const FLYStakeModule = require("../ignition/modules/stake");
async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contract with the account:", deployer.address);
    console.log("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^");


    const module = await ignition.deploy(FLYStakeModule);

    console.log("Token address:", module.token.target);
    console.log("Stake address:", module.stake.target);
    console.log("StakeProxy address:", module.stakeProxy.target);



    const contract = await ethers.getContractAt(
        "FLYStake",
        module.stakeProxy.target, // 这里代用的是代理合约地址
        deployer
    );
    console.log("proxy address:", contract.target);


    const len = await contract.poolLength()  //成功
    console.log("测试调用成功", len);



    // 授予 admin 权限
    const adminRole = ethers.keccak256(ethers.toUtf8Bytes("admin_role"));
    await contract.grantRole(adminRole, deployer.address);
    console.log("测试调用成功", adminRole);


    // contract2 = await ethers.getContractAt("FLYStake", module.stake.target);
    // console.log("contract2 address:", contract2.target);
    // const len2 = await contract2.poolLength()  //成功
    // console.log("测试调用成功", len2);

    //     // 获取实现合约地址
    //     // const implementationAddress = await upgrades.erc1967.getImplementationAddress(contract1.target);
    //     // console.log("实现合约地址:", implementationAddress);
    //     // const implementationAddress2 = await upgrades.erc1967.getImplementationAddress(contract2.target);
    //     // console.log("??合约地址:", implementationAddress2);
    //     /*


    //    function addPool(
    //         address _stTokenAddress,
    //         uint256 _poolWeight,
    //         uint256 _minDepositAmount,
    //         uint256 _unstakeLockedBlocks,
    //         bool _withUpdate
    //     ) public onlyRole(ADMIN_ROLE) {
    //     */
    //     //         address _stTokenAddress = address(0x0);
    //     //         uint256 _poolWeight = 100;
    //     //         uint256 _minDepositAmount = 100;
    //     //         uint256 _withdrawLockedBlocks = 100;
    //     //         bool _withUpdate = true;


    // console.log("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^");
    // const addPoolData = await contract.connect(deployer).addPool(
    //     ethers.ZeroAddress,
    //     100,
    //     100,
    //     100,
    //     false
    // )  //成功

    // console.log("测试调用addPool成功", addPoolData);
    // console.log("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^");
}

main().then(
    () => process.exit(0)
).catch((error) => {
    console.error(error);
    process.exitCode = 1;
})