// scripts/compile.js
const path = require("path");
const fs = require("fs");
const solc = require("solc");

// 1. 读取 Solidity 合约文件
const contractPath = path.resolve(__dirname, "../contracts/Proxy.sol");
const source = fs.readFileSync(contractPath, "utf8");

// 2. 配置 solc 输入
const input = {
    language: "Solidity",
    sources: {
        "Proxy.sol": {
            content: source,
        },
    },
    settings: {
        outputSelection: {
            "*": {
                "*": ["*"],
            },
        },
    },
};

// 3. 编译合约
const output = JSON.parse(solc.compile(JSON.stringify(input)));

// // 4. 检查错误
// if (output.errors) {
//     output.errors.forEach((err) => console.error(err.formattedMessage));
//     throw new Error("Compilation failed");
// }
console.log(output);
// 5. 提取 ABI 和字节码
const contract = output.contracts["Proxy.sol"]["Proxy"];
const abi = contract.abi;
const bytecode = contract.evm.bytecode.object;

// 6. 保存 ABI 和字节码到文件
const artifactPath = path.resolve(__dirname, "../artifacts/Proxy.json");
fs.writeFileSync(
    artifactPath,
    JSON.stringify({ abi, bytecode }, null, 2)
);

console.log("Compilation successful!");
console.log("ABI and bytecode saved to artifacts/Proxy.json");