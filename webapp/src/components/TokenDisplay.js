import React, { useState, useEffect } from 'react';
import Web3 from 'web3';
import { SKAVENGE_ABI } from '../contractABI';

function TokenDisplay({ metamaskAddress, config, onToast }) {
  const [tokens, setTokens] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!metamaskAddress || !config || !config.contractAddress) {
      setLoading(false);
      return;
    }

    fetchUserTokens();
  }, [metamaskAddress, config]);

  const fetchUserTokens = async () => {
    try {
      setLoading(true);
      setError(null);

      // Initialize Web3 with the user's MetaMask provider
      const web3 = new Web3(window.ethereum);

      // Create contract instance
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Get the number of tokens owned by the user
      const balance = await contract.methods.balanceOf(metamaskAddress).call();
      const balanceNum = Number(balance);

      if (balanceNum === 0) {
        setTokens([]);
        setLoading(false);
        return;
      }

      // Fetch all token IDs owned by the user
      const tokenPromises = [];
      for (let i = 0; i < balanceNum; i++) {
        tokenPromises.push(
          contract.methods.tokenOfOwnerByIndex(metamaskAddress, i).call()
        );
      }

      const tokenIds = await Promise.all(tokenPromises);

      // Fetch metadata for each token
      const tokenDataPromises = tokenIds.map(async (tokenId) => {
        try {
          // Fetch clue data from the contract
          const clueData = await contract.methods.clues(tokenId).call();
          const isForSale = await contract.methods.cluesForSale(tokenId).call();

          return {
            tokenId: tokenId.toString(),
            isSolved: clueData.isSolved,
            solveAttempts: clueData.solveAttempts.toString(),
            salePrice: clueData.salePrice.toString(),
            isForSale: isForSale,
            encryptedContents: clueData.encryptedContents,
            // Convert rValue to hex for display (only owner can see this)
            rValue: clueData.rValue.toString()
          };
        } catch (err) {
          console.error(`Error fetching data for token ${tokenId}:`, err);
          return {
            tokenId: tokenId.toString(),
            error: 'Failed to fetch token data'
          };
        }
      });

      const tokensData = await Promise.all(tokenDataPromises);
      setTokens(tokensData);
      setLoading(false);
    } catch (err) {
      console.error('Error fetching user tokens:', err);
      setError(err.message || 'Failed to fetch tokens');
      setLoading(false);

      if (onToast) {
        onToast('Failed to fetch tokens: ' + (err.message || 'Unknown error'), 'error');
      }
    }
  };

  const formatPrice = (priceWei) => {
    if (!priceWei || priceWei === '0') return '0';
    const web3 = new Web3();
    return web3.utils.fromWei(priceWei, 'ether');
  };

  const truncateHex = (hex, length = 10) => {
    if (!hex || hex.length <= length) return hex;
    return `${hex.substring(0, length)}...${hex.substring(hex.length - 6)}`;
  };

  if (!metamaskAddress) {
    return (
      <div className="token-display-container">
        <div className="token-display-message">
          Please connect your MetaMask wallet to view your NFTs.
        </div>
      </div>
    );
  }

  if (!config || !config.contractAddress) {
    return (
      <div className="token-display-container">
        <div className="token-display-message">
          Contract configuration not loaded. Please ensure the setup has been completed.
        </div>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="token-display-container">
        <div className="token-display-message">Loading your NFTs...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="token-display-container">
        <div className="token-display-error">
          <h3>Error</h3>
          <p>{error}</p>
          <button onClick={fetchUserTokens} className="btn-retry">Retry</button>
        </div>
      </div>
    );
  }

  if (tokens.length === 0) {
    return (
      <div className="token-display-container">
        <div className="token-display-message">
          <h3>No NFTs Found</h3>
          <p>You don't own any Skavenge NFTs yet.</p>
        </div>
      </div>
    );
  }

  return (
    <div className="token-display-container">
      <div className="token-display-header">
        <h2>Your Skavenge NFTs</h2>
        <p className="token-count">{tokens.length} {tokens.length === 1 ? 'token' : 'tokens'}</p>
        <button onClick={fetchUserTokens} className="btn-refresh">Refresh</button>
      </div>

      <div className="token-grid">
        {tokens.map((token) => (
          <div key={token.tokenId} className="token-card">
            <div className="token-card-header">
              <h3>Token #{token.tokenId}</h3>
              {token.isSolved && <span className="badge-solved">Solved</span>}
              {token.isForSale && <span className="badge-for-sale">For Sale</span>}
            </div>

            {token.error ? (
              <div className="token-error">{token.error}</div>
            ) : (
              <div className="token-details">
                <div className="token-detail-row">
                  <span className="detail-label">Status:</span>
                  <span className={`detail-value ${token.isSolved ? 'text-success' : 'text-warning'}`}>
                    {token.isSolved ? 'Solved' : 'Unsolved'}
                  </span>
                </div>

                <div className="token-detail-row">
                  <span className="detail-label">Solve Attempts:</span>
                  <span className="detail-value">{token.solveAttempts} / 3</span>
                </div>

                {token.isForSale && (
                  <div className="token-detail-row">
                    <span className="detail-label">Sale Price:</span>
                    <span className="detail-value">{formatPrice(token.salePrice)} ETH</span>
                  </div>
                )}

                {token.encryptedContents && (
                  <div className="token-detail-row">
                    <span className="detail-label">Encrypted Data:</span>
                    <span className="detail-value mono">
                      {truncateHex(token.encryptedContents)}
                    </span>
                  </div>
                )}

                {token.rValue && token.rValue !== '0' && (
                  <div className="token-detail-row">
                    <span className="detail-label">R Value:</span>
                    <span className="detail-value mono">
                      {truncateHex(token.rValue, 16)}
                    </span>
                  </div>
                )}
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

export default TokenDisplay;
