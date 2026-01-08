import { Web3 } from 'web3';
import { SKAVENGE_ABI } from './contractABI';
import { getConfig } from './config';

let web3Instance = null;
let contractInstance = null;

export const initWeb3 = () => {
  if (web3Instance) return web3Instance;

  const config = getConfig();
  web3Instance = new Web3(config.rpcUrl);
  return web3Instance;
};

export const getWeb3 = () => {
  if (!web3Instance) {
    return initWeb3();
  }
  return web3Instance;
};

export const getContract = () => {
  if (contractInstance) return contractInstance;

  const config = getConfig();
  const web3 = getWeb3();
  contractInstance = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);
  return contractInstance;
};

export const getClueFromContract = async (tokenId) => {
  const contract = getContract();
  const clueData = await contract.methods.clues(tokenId).call();

  return {
    encryptedContents: clueData.encryptedContents,
    solutionHash: clueData.solutionHash,
    isSolved: clueData.isSolved,
    salePrice: clueData.salePrice,
    rValue: clueData.rValue,
    timeout: clueData.timeout,
    pointValue: clueData.pointValue,
    solveReward: clueData.solveReward,
  };
};

export const getClueOwner = async (tokenId) => {
  const contract = getContract();
  return await contract.methods.ownerOf(tokenId).call();
};

export const getTotalSupply = async () => {
  const contract = getContract();
  return await contract.methods.totalSupply().call();
};

export const getCurrentTokenId = async () => {
  const contract = getContract();
  return await contract.methods.getCurrentTokenId().call();
};

export const getCluesForSale = async (offset = 0, limit = 100) => {
  const contract = getContract();
  return await contract.methods.getCluesForSale(offset, limit).call();
};

export const getTotalCluesForSale = async () => {
  const contract = getContract();
  return await contract.methods.getTotalCluesForSale().call();
};

export const formatWei = (weiValue) => {
  const web3 = getWeb3();
  return web3.utils.fromWei(weiValue.toString(), 'ether');
};

export const formatAddress = (address) => {
  if (!address || address.length < 10) return address;
  return `${address.slice(0, 6)}...${address.slice(-4)}`;
};
