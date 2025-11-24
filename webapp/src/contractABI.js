// ABI for Skavenge contract - includes ERC721Enumerable and custom Clue functions
export const SKAVENGE_ABI = [
  // ERC721Enumerable functions
  {
    "inputs": [{"internalType": "address", "name": "owner", "type": "address"}],
    "name": "balanceOf",
    "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "owner", "type": "address"},
      {"internalType": "uint256", "name": "index", "type": "uint256"}
    ],
    "name": "tokenOfOwnerByIndex",
    "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [{"internalType": "uint256", "name": "tokenId", "type": "uint256"}],
    "name": "ownerOf",
    "outputs": [{"internalType": "address", "name": "", "type": "address"}],
    "stateMutability": "view",
    "type": "function"
  },
  // Custom Skavenge functions
  {
    "inputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "name": "clues",
    "outputs": [
      {"internalType": "bytes", "name": "encryptedContents", "type": "bytes"},
      {"internalType": "bytes32", "name": "solutionHash", "type": "bytes32"},
      {"internalType": "bool", "name": "isSolved", "type": "bool"},
      {"internalType": "uint256", "name": "solveAttempts", "type": "uint256"},
      {"internalType": "uint256", "name": "salePrice", "type": "uint256"},
      {"internalType": "uint256", "name": "rValue", "type": "uint256"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "name": "cluesForSale",
    "outputs": [{"internalType": "bool", "name": "", "type": "bool"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalSupply",
    "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  }
];
