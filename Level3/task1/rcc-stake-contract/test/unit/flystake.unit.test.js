const { assert, expect } = require("chai");
const { ethers, network, ignition } = require("hardhat");
const { developmentChains } = require("../../scripts/config");
const FLYStakeModule = require("../../ignition/modules/fly_stake");

!developmentChains.includes(network.name)
    ? describe.skip
    : describe("Fly Stake Unit Tests", function () {
        let contract;
        let deployer
        beforeEach(async () => {
            const accounts = await ethers.getSigners();
            deployer = accounts[0];
            const module = await ignition.deploy(FLYStakeModule);
            contract = await ethers.getContractAt(
                "FLYStake",
                module.stakeProxy.target,
                deployer
            );
        });
        //poolLength
        describe("depositETH", function () {
            it("直接调用depositETH会失败", async function () {
                await expect(contract.depositETH()).to.be.reverted;
            });
            it("存入ETH到合约中", async function () {
                const amount = ethers.parseEther("1");
                await expect(contract.depositETH({ value: amount })).to.not.reverted;
            });
        });
    });
