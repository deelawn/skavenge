// hardhat.config.js
module.exports = {
	solidity: "0.8.19",
	networks: {
		hardhat: {
			chainId: 1337,
			allowUnlimitedContractSize: true,
			mining: {
				auto: true,
				interval: 1000
			},
			accounts: [
				{
					privateKey: "0xa7a505c6a83e4521a6532f09c375d150257a0ec4652bdf32231c1f7ee3590af4",
					balance: "10000000000000000000000" // 10,000 ETH
				},
				{
					privateKey: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
					balance: "10000000000000000000000" // 10,000 ETH
				},
				{
					privateKey: "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
					balance: "10000000000000000000000" // 10,000 ETH
				},
				{
					privateKey: "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
					balance: "10000000000000000000000" // 10,000 ETH
				},
				{
					privateKey: "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
					balance: "10000000000000000000000" // 10,000 ETH
				}
			]
		}
	}
};