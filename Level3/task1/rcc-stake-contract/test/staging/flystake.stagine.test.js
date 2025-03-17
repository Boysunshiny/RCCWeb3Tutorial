const { assert } = require("chai");
const { ethers, network, deployments } = require("hardhat");
const { developmentChains } = require("../../scripts/config");

developmentChains.includes(network.name)
    ? describe.skip
    : describe("Fly Stake Staging Tests", async () => {
        let contract;
        let deployer
        beforeEach(async () => {
            const accounts = await ethers.getSigners();
            deployer = accounts[0];
            const flyStakeDeployment = await deployments.get("FLYStake");
            contract = await ethers.getContractAt(
                "FLYStake",
                flyStakeDeployment.address,
                deployer
            );
        });

    });
