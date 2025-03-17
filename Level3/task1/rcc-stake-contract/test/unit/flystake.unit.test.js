const { assert } = require("chai");
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
            const deployment = await ignition.deploy(FLYStakeModule);
            contract = await ethers.getContractAt(
                "FLYStake",
                deployment.stake.target,
                deployer
            );
        });

        it("test for poolLength", async function () {
            const s = await contract.poolLength();
            assert.equal(s, "0");
        });

    });
