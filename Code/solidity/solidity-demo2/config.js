const networkConfig = {
    localhost: {},
    hardhat: {},
    ganache: {},
    sepolia: {
        blockConfirmations: 6,
        ethUsdPriceFeed: "0x694AA1769357215DE4FAC081bf1f309aDC325306",
    },
}

const developmentChains = ["hardhat", "localhost", "ganache"]
module.exports = {
    networkConfig,
    developmentChains,
}