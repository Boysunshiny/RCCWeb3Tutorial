NFT Marketplace 后端代码 👈
NFT Marketplace 合约代码 👈
NFT Marketplace 前端代码 👈

📺 视频地址 :
⁠NFT marketplace 项目架构分享
⁠NFT market 合约拆解

背景介绍
2020 年是 DeFi 元年

2021 年是 NFT 元年

...

数字资产的兴起与需求
去中心化的市场需求
版权保护和二级市场
全球化市场的潜力
技术与金融的融合

项目意义
当前的 NFT 交易市场不仅是一个基于区块链的应用，也是链上技术与链下服务高度结合的典型范例。通过项目的设计和开发，可以探索如何将区块链的去中心化、透明性、不可篡改等特点与传统的链下业务流程进行有机融合，创建一个灵活、可扩展的系统架构，不仅服务于 NFT 交易市场，还能支持未来其他潜在的链上应用，如 Bitcoin 上的铭文、符文等新兴数字资产。以下从几个关键角度说明项目的深远意义：

技术架构的通用性和可扩展性
链上技术原理与链下服务的结合
去中心化应用的场景扩展

NFT 基本概念
图片
NFT 的核心操作
详见：[https://eips.ethereum.org/EIPS/eip-721 ] (https://eips.ethereum.org/EIPS/eip-721)
transfer
safeTransferFrom(address \_from, address \_to, uint256 \_tokenId)
transferFrom(address \_from, address \_to, uint256 \_tokenId)

approve
approve(address \_approved, uint256 \_tokenId)
setApprovalForAll(address \_operator, bool \_approved)

RCC 教学助理-Alex — 2024/11/12 19:08
NFT 数据模型
Collection ———— NFT 集合的实体

Item ———— 代表交易系统中代表 NFT 的实体

Ownership ———— 代表 NFT 的所有权，也就是 Item 的 Owner， 即 Item 和 Wallet 的关联关系

Order ———— 代表出售或购买 NFT 意愿的实体。

Activity ———— 代表 NFT 状态下发生的事件：mint, transfer, list, buy 等
图片
NFT 交易模式
NFT 订单在链下: 非 dex
NFT 订单在链上:dex

订单簿 OrderBook: Maker, Taker: 用户; 价格确定于订单

做市商 AMM: ERC721——AMM: Maker, Taker: 一方是池子, 一方是用户; 价格是随池子变化的;

项目描述
架构图
API 服务

交易同步服务

NFT 导入服务

订单中继服务

定时任务模块
图片
微服务功能描述
API 服务

涉及 Collection，Item， Order， Activity 等实体的相关接口请求
create table ob_collection_sepolia
(
id bigint auto_increment comment '主键'
primary key,
chain_id tinyint default 1 not null comment '链类型(1:以太坊)',
symbol varchar(128) not null comment '项目标识',
展开
message.txt
6 KB
图片
图片
图片
交易合约核心功能
基于智能合约实现基于订单簿模型的 NFT 交易系统, 即能够支持以下写入操作和查询操作;

写入操作:

链上订单簿(OrderBook DEX)支持 create limit sell/buy, market sell/buy order, edit(cancel&create)/cancel order 功能;

limit sell order:
limit buy order:
market sell order:
market buy order:
edit/cancel order:

查询操作:

支持从链上查询订单（包括已经过期订单）;
合约构成及组件
OrderBookExchange: 实现完整的订单簿交易逻辑
OrderStorage: 用于存储订单信息的模块
OrderValidator: 用于处理订单逻辑验证的模块
ProtocolManager: 用于管理协议费的模块
...
OrderVault: 独立存储订单相关资产的模块;

RCC 教学助理-Alex — 2025/3/19 14:08
问题模版
https://ncnlo95gua0y.feishu.cn/wiki/MbtOw9RJrim9PKkrOk3coNLjnfc?from=from_copylink
图片
图片
