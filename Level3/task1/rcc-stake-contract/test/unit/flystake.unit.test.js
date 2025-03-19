const { assert, expect } = require("chai");
const { ethers, network, ignition, upgrades } = require("hardhat");
const { developmentChains } = require("../../scripts/config");

!developmentChains.includes(network.name)
    ? describe.skip
    : describe("RCC Stake Unit Tests", function () {
        let contract;
        let owner, admin, user1, user2;
        let Erc20MockTokenFactory;
        beforeEach(async () => {
            [owner, admin, user1, user2, ...addrs] = await ethers.getSigners();

            const startBlockOffset = 10; // Start staking 10 blocks after deployment
            const endBlockOffset = 100; // End staking at block number + 1000
            const perBlock = ethers.parseEther("10"); // 10 RCC per block
            const currentBlock = await ethers.provider.getBlockNumber();
            Erc20MockTokenFactory = await ethers.getContractFactory("RccToken")
            //  部署获取到的Rcc Token 地址
            rccToken = await Erc20MockTokenFactory.deploy();

            const Stake = await hre.ethers.getContractFactory("RCCStake");
            contract = await upgrades.deployProxy(
                Stake,
                [rccToken.target, currentBlock + startBlockOffset, currentBlock + endBlockOffset, perBlock],
                { initializer: "initialize" }
            );

            // 授予 admin 权限
            const adminRole = ethers.keccak256(ethers.toUtf8Bytes("admin_role"));
            await contract.grantRole(adminRole, admin.address);

            await contract.connect(admin).addPool(
                ethers.ZeroAddress, // 使用 .ZeroAddress 确保传递 address(0)
                100, // poolWeight 权重
                ethers.parseEther("0.1"), // minDepositAmount（根据合约逻辑）
                100, // 解除质押的锁定区块数
                false //
            );
        });
        //poolLength
        describe("User can deposit and withdraw", function () {

            // it("存入ETH到合约中", async function () {
            //     const amount = ethers.parseEther("1");
            //     await expect(contract.depositETH({ value: amount })).to.not.reverted;
            // });

            it("User can claim RCC rewards", async function () {
                const userInfo = await contract.user(0, user2);
                console.log("stAmount : " + userInfo.stAmount)
                console.log("finishedRCC : " + userInfo.finishedRCC)
                console.log("pendingRCC : " + userInfo.pendingRCC)

                const depositTx = await contract.connect(user2).depositETH({ value: ethers.parseEther("10") });
                await depositTx.wait();
                //模拟区块
                for (let i = 0; i < 90; i++) {
                    await ethers.provider.send("evm_mine", []);
                }


                const poolInfo = await contract.pool(0)

                const userInfo2 = await contract.user(0, user2);


                // uint256 pendingRCC_ = user_.stAmount * pool_.accRCCPerST / (1 ether) - user_.finishedRCC + user_.pendingRCC;
                const pendingRCc = userInfo2.stAmount * poolInfo.accRCCPerST / ethers.parseEther("1") - userInfo2.finishedRCC + userInfo2.pendingRCC;
                console.log("pendingRCc : " + pendingRCc)

                // await rccStakeProxy.connect(user2).unstake(0, ethers.parseEther("10"))

                // 用户领取奖励
                const claimTx = await contract.connect(user2).claim(0);
                await expect(claimTx)
                    .to.emit(contract, 'Claim');

                // 获取用户的 RCC 代币余额
                const userRCCBalance = await rccToken.balanceOf(user2);
                console.log("poolInfo : " + poolInfo)
                console.log("pool_.accRCCPerST : " + poolInfo.accRCCPerST)
                console.log("stAmount : " + userInfo2.stAmount)
                console.log("finishedRCC : " + userInfo2.finishedRCC)
                console.log("pendingRCC : " + userInfo2.pendingRCC)
                console.log("userRCCBalance : " + userRCCBalance)

                // 断言用户的 RCC 代币余额大于 0
                expect(userRCCBalance).to.be.eq(ethers.parseEther("0"));
            });



        });
    });
