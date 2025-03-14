// // contract FruitStore is BasicAuth {
// //     address _latestVersion;
// //     mapping(bytes => uint) _fruitStock;
// //     modifier onlyLatestVersion() {
// //         require(msg.sender == _latestVersion);
// //         _;
// //     }

// //     function upgradeVersion(address newVersion) public {
// //         require(msg.sender == _owner);
// //         _latestVersion = newVersion;
// //     }

// //     function setFruitStock(bytes fruit, uint stock) external onlyLatestVersion {
// //         _fruitStock[fruit] = stock;
// //     }
// // }
// contract BasicAuth {
//     address public _owner;

//     constructor() public {
//         _owner = msg.sender;
//     }

//     function setOwner(address owner) public onlyOwner {
//         _owner = owner;
//     }

//     modifier onlyOwner() {
//         require(msg.sender == _owner, "BasicAuth: only owner is authorized.");
//         _;
//     }
// }

// contract FruitStore is BasicAuth {
//     mapping(bytes => uint) _fruitStock;

//     function setFruitStock(
//         bytes fruitName,
//         uint stock
//     ) external onlyOwner validFruitName(fruitName) {
//         _fruitStock[fruitName] = stock;
//     }
// }

// contract Admin is BasicAuth {
//     function upgradeContract(
//         FruitStore fruitStore,
//         address newController
//     ) external isAuthorized {
//         fruitStore.upgradeVersion(newController);
//     }
// }

// contract FruitStoreController is BasicAuth {
//     function upgradeStock(bytes fruit, uint stock) external isAuthorized {
//         fruitStore.setFruitStock(fruit, stock);
//     }
// }
