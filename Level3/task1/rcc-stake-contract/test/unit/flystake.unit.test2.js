// const { expect } = require("chai");
// const { ethers } = require("hardhat");

// describe("rccStake test", function () {
//     let rccStakeProxy;
//     let rccToken;
//     let Erc20MockTokenFactory;
//     let owner, admin, user1, user2;
//     let ustToken;
//     beforeEach(async function () {
//         // 获取多个以太坊账户，用于模拟不同的用户和管理员角色
//         [owner, admin, user1, user2, ...addrs] = await ethers.getSigners();
//         //用于发布质押代币
//         Erc20MockTokenFactory = await ethers.getContractFactory("RccToken")
//         const startBlockOffset = 10; // Start staking 10 blocks after deployment
//         const endBlockOffset = 100; // End staking at block number + 1000
//         const RCCPerBlock = ethers.parseEther("10"); // 10 RCC per block
//         //  部署获取到的Rcc Token 地址
//         rccToken = await Erc20MockTokenFactory.deploy();
//         await rccToken.waitForDeployment();
//         const currentBlock = await ethers.provider.getBlockNumber();
//         const Stake = await hre.ethers.getContractFactory("RCCStake");
//         rccStakeProxy = await upgrades.deployProxy(
//             Stake,
//             [rccToken.target, currentBlock + startBlockOffset, currentBlock + endBlockOffset, RCCPerBlock],
//             { initializer: "initialize" }
//         );

//         await rccStakeProxy.waitForDeployment()

//         // 授予 admin 权限
//         const adminRole = ethers.keccak256(ethers.toUtf8Bytes("admin_role"));
//         await rccStakeProxy.grantRole(adminRole, admin.address);

//         await rccStakeProxy.connect(admin).addPool(
//             ethers.ZeroAddress, // 使用 .ZeroAddress 确保传递 address(0)
//             100, // poolWeight 权重
//             ethers.parseEther("0.1"), // minDepositAmount（根据合约逻辑）
//             100, // 解除质押的锁定区块数
//             false //
//         );
//     });

//     // });
//     it("User can claim RCC rewards", async function () {
//         const userInfo = await rccStakeProxy.user(0, user2);
//         console.log("stAmount : " + userInfo.stAmount)
//         console.log("finishedRCC : " + userInfo.finishedRCC)
//         console.log("pendingRCC : " + userInfo.pendingRCC)

//         const depositTx = await rccStakeProxy.connect(user2).depositETH({ value: ethers.parseEther("10") });
//         await depositTx.wait();
//         //模拟区块
//         for (let i = 0; i < 90; i++) {
//             await ethers.provider.send("evm_mine", []);
//         }


//         const poolInfo = await rccStakeProxy.pool(0)

//         const userInfo2 = await rccStakeProxy.user(0, user2);


//         // uint256 pendingRCC_ = user_.stAmount * pool_.accRCCPerST / (1 ether) - user_.finishedRCC + user_.pendingRCC;
//         const pendingRCc = userInfo2.stAmount * poolInfo.accRCCPerST / ethers.parseEther("1") - userInfo2.finishedRCC + userInfo2.pendingRCC;
//         console.log("pendingRCc : " + pendingRCc)

//         // await rccStakeProxy.connect(user2).unstake(0, ethers.parseEther("10"))

//         // 用户领取奖励
//         const claimTx = await rccStakeProxy.connect(user2).claim(0);
//         await expect(claimTx)
//             .to.emit(rccStakeProxy, 'Claim');

//         // 获取用户的 RCC 代币余额
//         const userRCCBalance = await rccToken.balanceOf(user2);
//         console.log("poolInfo : " + poolInfo)
//         console.log("pool_.accRCCPerST : " + poolInfo.accRCCPerST)
//         console.log("stAmount : " + userInfo2.stAmount)
//         console.log("finishedRCC : " + userInfo2.finishedRCC)
//         console.log("pendingRCC : " + userInfo2.pendingRCC)
//         console.log("userRCCBalance : " + userRCCBalance)

//         // 断言用户的 RCC 代币余额大于 0
//         expect(userRCCBalance).to.be.equals(ethers.parseEther("0"));
//     });


// });