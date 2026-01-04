import React, { useState, useEffect } from 'react';
import Web3 from 'web3';
import { SKAVENGE_ABI } from '../contractABI';
import { sendToExtension } from '../extensionUtils';

function TokenDisplay({ metamaskAddress, config, onToast }) {
  const [tokens, setTokens] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [revealedClues, setRevealedClues] = useState({});
  const [listingToken, setListingToken] = useState(null);
  const [listingPrice, setListingPrice] = useState('');
  const [listingTimeout, setListingTimeout] = useState('');
  const [timeoutUnit, setTimeoutUnit] = useState('hours'); // 'minutes' or 'hours'
  const [minTimeout, setMinTimeout] = useState(null);
  const [maxTimeout, setMaxTimeout] = useState(null);
  const [processingTx, setProcessingTx] = useState(false);

  useEffect(() => {
    if (!metamaskAddress || !config || !config.contractAddress) {
      setLoading(false);
      return;
    }

    fetchUserTokens();
    fetchTimeoutBounds();

    // Clear revealed clues when tokens are refetched for privacy
    setRevealedClues({});
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [metamaskAddress, config]);

  // Clear revealed clues when component unmounts for privacy
  useEffect(() => {
    return () => {
      setRevealedClues({});
    };
  }, []);

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
            salePrice: clueData.salePrice.toString(),
            isForSale: isForSale,
            encryptedContents: clueData.encryptedContents,
            // Convert rValue to hex for display (only owner can see this)
            rValue: clueData.rValue.toString(),
            pointValue: Number(clueData.pointValue),
            timeout: Number(clueData.timeout),
            solveReward: clueData.solveReward.toString()
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

  const fetchTimeoutBounds = async () => {
    try {
      // Create cache key based on contract address
      const cacheKey = `timeout_bounds_${config.contractAddress}`;

      // Check if we have cached values
      const cached = localStorage.getItem(cacheKey);
      if (cached) {
        try {
          const { minTimeout: cachedMin, maxTimeout: cachedMax, timestamp } = JSON.parse(cached);
          // Cache is valid for 24 hours (86400000 ms)
          const cacheAge = Date.now() - timestamp;
          if (cacheAge < 86400000 && cachedMin && cachedMax) {
            console.log('Using cached timeout bounds:', { min: cachedMin, max: cachedMax });
            setMinTimeout(Number(cachedMin));
            setMaxTimeout(Number(cachedMax));
            return;
          }
        } catch (parseErr) {
          console.warn('Failed to parse cached timeout bounds:', parseErr);
          localStorage.removeItem(cacheKey);
        }
      }

      // Fetch from contract if no valid cache
      console.log('Fetching timeout bounds from contract...');
      const web3 = new Web3(window.ethereum);
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Fetch MIN_TIMEOUT and MAX_TIMEOUT from the contract
      const minTimeoutValue = await contract.methods.MIN_TIMEOUT().call();
      const maxTimeoutValue = await contract.methods.MAX_TIMEOUT().call();

      const min = Number(minTimeoutValue);
      const max = Number(maxTimeoutValue);

      setMinTimeout(min);
      setMaxTimeout(max);

      // Cache the values
      localStorage.setItem(cacheKey, JSON.stringify({
        minTimeout: min,
        maxTimeout: max,
        timestamp: Date.now()
      }));
      console.log('Cached timeout bounds:', { min, max });
    } catch (err) {
      console.error('Error fetching timeout bounds:', err);
      // Set default values if fetch fails (1 hour and 24 hours)
      setMinTimeout(3600);
      setMaxTimeout(86400);
    }
  };

  const formatPrice = (priceWei) => {
    if (!priceWei || priceWei === '0') return '0';
    const web3 = new Web3();
    return web3.utils.fromWei(priceWei, 'ether');
  };


  const handleRevealClue = async (tokenId, encryptedContents, rValue) => {
    try {
      console.log('=== WEBAPP: Preparing to decrypt ===');
      console.log('tokenId:', tokenId);
      console.log('encryptedContents (raw):', encryptedContents);
      console.log('rValue (raw decimal):', rValue);

      // Clean the encrypted hex string (remove '0x' prefix if present)
      const cleanEncryptedHex = encryptedContents.startsWith('0x')
        ? encryptedContents.substring(2)
        : encryptedContents;

      // Convert rValue from decimal string to hex
      // rValue comes from the contract as a decimal string representation of uint256
      // eslint-disable-next-line no-undef
      const rValueBigInt = BigInt(rValue);
      const rValueHex = rValueBigInt.toString(16).padStart(64, '0'); // Pad to 64 hex chars (32 bytes)

      console.log('cleanEncryptedHex:', cleanEncryptedHex);
      console.log('rValueHex:', rValueHex);
      console.log('encryptedHex length:', cleanEncryptedHex.length);

      // Call the extension to decrypt (with auto-unlock support)
      const response = await sendToExtension({
        action: 'decryptElGamal',
        encryptedHex: cleanEncryptedHex,
        rValueHex: rValueHex
      }, () => {
        // Callback when unlock is prompted
        if (onToast) {
          onToast('Please unlock the Skavenger extension to continue', 'info');
        }
      });

      console.log('Extension response:', response);

      if (response.success) {
        setRevealedClues(prev => ({
          ...prev,
          [tokenId]: response.plaintext
        }));
        if (onToast) {
          onToast('Clue revealed successfully', 'success');
        }
      } else {
        if (onToast) {
          onToast('Failed to reveal clue: ' + (response.error || 'Unknown error'), 'error');
        }
      }
    } catch (error) {
      console.error('Error revealing clue:', error);
      if (onToast) {
        onToast('Failed to reveal clue: ' + error.message, 'error');
      }
    }
  };

  const handleSetSalePrice = async (tokenId, priceInEth, timeoutInSeconds) => {
    try {
      setProcessingTx(true);
      const web3 = new Web3(window.ethereum);
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Convert ETH to Wei
      const priceInWei = web3.utils.toWei(priceInEth, 'ether');

      // Call setSalePrice function with timeout
      await contract.methods.setSalePrice(tokenId, priceInWei, timeoutInSeconds).send({
        from: metamaskAddress
      });

      if (onToast) {
        const timeoutDisplay = formatTimeout(timeoutInSeconds);
        onToast(`Token #${tokenId} listed for ${priceInEth} ETH with ${timeoutDisplay} timeout`, 'success');
      }

      // Refresh tokens to show updated sale status
      await fetchUserTokens();
      setListingToken(null);
      setListingPrice('');
      setListingTimeout('');
      setTimeoutUnit('hours');
    } catch (error) {
      console.error('Error setting sale price:', error);
      if (onToast) {
        onToast('Failed to list token: ' + (error.message || 'Unknown error'), 'error');
      }
    } finally {
      setProcessingTx(false);
    }
  };

  const formatTimeout = (seconds) => {
    if (seconds < 3600) {
      return `${Math.round(seconds / 60)} minutes`;
    } else if (seconds % 3600 === 0) {
      return `${seconds / 3600} hours`;
    } else {
      const hours = Math.floor(seconds / 3600);
      const minutes = Math.round((seconds % 3600) / 60);
      return `${hours}h ${minutes}m`;
    }
  };

  const handleRemoveSalePrice = async (tokenId) => {
    try {
      setProcessingTx(true);
      const web3 = new Web3(window.ethereum);
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Call removeSalePrice function
      await contract.methods.removeSalePrice(tokenId).send({
        from: metamaskAddress
      });

      if (onToast) {
        onToast(`Token #${tokenId} removed from sale`, 'success');
      }

      // Refresh tokens to show updated sale status
      await fetchUserTokens();
    } catch (error) {
      console.error('Error removing sale price:', error);
      if (onToast) {
        onToast('Failed to remove token from sale: ' + (error.message || 'Unknown error'), 'error');
      }
    } finally {
      setProcessingTx(false);
    }
  };

  const startListing = (tokenId) => {
    setListingToken(tokenId);
    setListingPrice('');
    setListingTimeout('');
    setTimeoutUnit('hours');
  };

  const cancelListing = () => {
    setListingToken(null);
    setListingPrice('');
    setListingTimeout('');
    setTimeoutUnit('hours');
  };

  const confirmListing = (tokenId) => {
    if (!listingPrice || parseFloat(listingPrice) <= 0) {
      if (onToast) {
        onToast('Please enter a valid price', 'error');
      }
      return;
    }

    if (!listingTimeout || parseFloat(listingTimeout) <= 0) {
      if (onToast) {
        onToast('Please enter a valid timeout', 'error');
      }
      return;
    }

    // Convert timeout to seconds based on unit
    const timeoutValue = parseFloat(listingTimeout);
    const timeoutInSeconds = timeoutUnit === 'hours'
      ? Math.floor(timeoutValue * 3600)
      : Math.floor(timeoutValue * 60);

    // Validate timeout bounds
    if (minTimeout !== null && timeoutInSeconds < minTimeout) {
      if (onToast) {
        const minDisplay = formatTimeout(minTimeout);
        onToast(`Timeout must be at least ${minDisplay}`, 'error');
      }
      return;
    }

    if (maxTimeout !== null && timeoutInSeconds > maxTimeout) {
      if (onToast) {
        const maxDisplay = formatTimeout(maxTimeout);
        onToast(`Timeout cannot exceed ${maxDisplay}`, 'error');
      }
      return;
    }

    handleSetSalePrice(tokenId, listingPrice, timeoutInSeconds);
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
                  <span className="detail-label">Point Value:</span>
                  <span className="detail-value" style={{ fontWeight: '600', color: '#667eea' }}>
                    {token.pointValue} {token.pointValue === 1 ? 'point' : 'points'}
                  </span>
                </div>

                {!token.isSolved && token.solveReward && token.solveReward !== '0' && (
                  <div className="token-detail-row">
                    <span className="detail-label">
                      Bonus Reward:
                      <span
                        style={{
                          display: 'inline-block',
                          marginLeft: '6px',
                          cursor: 'help',
                          fontSize: '12px',
                          color: '#667eea',
                          fontWeight: '600'
                        }}
                        title="This ETH reward is awarded to you immediately when you provide the correct solution"
                      >
                        ⓘ
                      </span>
                    </span>
                    <span className="detail-value" style={{ fontWeight: '600', color: '#48bb78' }}>
                      {formatPrice(token.solveReward)} ETH
                    </span>
                  </div>
                )}

                {token.isForSale && (
                  <>
                    <div className="token-detail-row">
                      <span className="detail-label">Sale Price:</span>
                      <span className="detail-value">{formatPrice(token.salePrice)} ETH</span>
                    </div>
                    {token.timeout > 0 && (
                      <div className="token-detail-row">
                        <span className="detail-label">Transfer Timeout:</span>
                        <span className="detail-value">{formatTimeout(token.timeout)}</span>
                      </div>
                    )}
                  </>
                )}

                {token.encryptedContents && token.rValue && token.rValue !== '0' && (
                  <div className="token-detail-row" style={{ flexDirection: 'column', alignItems: 'flex-start', gap: '8px' }}>
                    <div style={{ display: 'flex', alignItems: 'center', gap: '8px', width: '100%' }}>
                      <span className="detail-label">Clue:</span>
                      {!revealedClues[token.tokenId] && (
                        <button
                          onClick={() => handleRevealClue(token.tokenId, token.encryptedContents, token.rValue)}
                          className="btn-reveal-clue"
                          style={{
                            padding: '6px 12px',
                            fontSize: '12px',
                            backgroundColor: '#667eea',
                            color: 'white',
                            border: 'none',
                            borderRadius: '4px',
                            cursor: 'pointer',
                            fontWeight: '500'
                          }}
                        >
                          Reveal Clue
                        </button>
                      )}
                    </div>
                    {revealedClues[token.tokenId] && (
                      <div style={{
                        padding: '12px',
                        backgroundColor: '#f7fafc',
                        borderRadius: '6px',
                        width: '100%',
                        wordBreak: 'break-word'
                      }}>
                        <span className="detail-value">{revealedClues[token.tokenId]}</span>
                      </div>
                    )}
                  </div>
                )}

                {/* Listing Controls */}
                <div className="token-actions" style={{ marginTop: '16px', display: 'flex', flexDirection: 'column', gap: '8px' }}>
                  {token.isForSale ? (
                    <button
                      onClick={() => handleRemoveSalePrice(token.tokenId)}
                      disabled={processingTx}
                      className="btn-unlist"
                      style={{
                        padding: '10px',
                        fontSize: '14px',
                        backgroundColor: '#e53e3e',
                        color: 'white',
                        border: 'none',
                        borderRadius: '6px',
                        cursor: processingTx ? 'not-allowed' : 'pointer',
                        fontWeight: '500',
                        opacity: processingTx ? 0.6 : 1
                      }}
                    >
                      {processingTx ? 'Processing...' : 'Remove from Sale'}
                    </button>
                  ) : (
                    <>
                      {listingToken === token.tokenId ? (
                        <div style={{ display: 'flex', flexDirection: 'column', gap: '8px' }}>
                          <input
                            type="number"
                            step="0.001"
                            min="0"
                            placeholder="Price in ETH"
                            value={listingPrice}
                            onChange={(e) => setListingPrice(e.target.value)}
                            className="price-input"
                            style={{
                              padding: '8px 12px',
                              fontSize: '14px',
                              border: '1px solid #cbd5e0',
                              borderRadius: '6px',
                              outline: 'none'
                            }}
                          />

                          {/* Timeout Input Section */}
                          <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                            <label style={{ fontSize: '12px', color: '#718096', fontWeight: '500' }}>
                              Transfer Timeout
                            </label>
                            <div style={{ display: 'flex', gap: '8px' }}>
                              <input
                                type="number"
                                step="0.5"
                                min="0"
                                placeholder="Timeout"
                                value={listingTimeout}
                                onChange={(e) => setListingTimeout(e.target.value)}
                                className="timeout-input"
                                style={{
                                  flex: 1,
                                  padding: '8px 12px',
                                  fontSize: '14px',
                                  border: '1px solid #cbd5e0',
                                  borderRadius: '6px',
                                  outline: 'none'
                                }}
                              />
                              <select
                                value={timeoutUnit}
                                onChange={(e) => setTimeoutUnit(e.target.value)}
                                style={{
                                  padding: '8px 12px',
                                  fontSize: '14px',
                                  border: '1px solid #cbd5e0',
                                  borderRadius: '6px',
                                  backgroundColor: 'white',
                                  cursor: 'pointer',
                                  outline: 'none'
                                }}
                              >
                                <option value="minutes">Minutes</option>
                                <option value="hours">Hours</option>
                              </select>
                            </div>
                            {minTimeout !== null && maxTimeout !== null && (
                              <span style={{ fontSize: '11px', color: '#718096' }}>
                                Range: {formatTimeout(minTimeout)} - {formatTimeout(maxTimeout)}
                              </span>
                            )}
                          </div>

                          <div style={{ display: 'flex', gap: '8px' }}>
                            <button
                              onClick={() => confirmListing(token.tokenId)}
                              disabled={processingTx}
                              className="btn-confirm-list"
                              style={{
                                flex: 1,
                                padding: '10px',
                                fontSize: '14px',
                                backgroundColor: '#48bb78',
                                color: 'white',
                                border: 'none',
                                borderRadius: '6px',
                                cursor: processingTx ? 'not-allowed' : 'pointer',
                                fontWeight: '500',
                                opacity: processingTx ? 0.6 : 1
                              }}
                            >
                              {processingTx ? 'Processing...' : 'Confirm'}
                            </button>
                            <button
                              onClick={cancelListing}
                              disabled={processingTx}
                              className="btn-cancel-list"
                              style={{
                                flex: 1,
                                padding: '10px',
                                fontSize: '14px',
                                backgroundColor: '#718096',
                                color: 'white',
                                border: 'none',
                                borderRadius: '6px',
                                cursor: processingTx ? 'not-allowed' : 'pointer',
                                fontWeight: '500',
                                opacity: processingTx ? 0.6 : 1
                              }}
                            >
                              Cancel
                            </button>
                          </div>
                        </div>
                      ) : (
                        <button
                          onClick={() => startListing(token.tokenId)}
                          disabled={processingTx}
                          className="btn-list-for-sale"
                          style={{
                            padding: '10px',
                            fontSize: '14px',
                            backgroundColor: '#667eea',
                            color: 'white',
                            border: 'none',
                            borderRadius: '6px',
                            cursor: processingTx ? 'not-allowed' : 'pointer',
                            fontWeight: '500',
                            opacity: processingTx ? 0.6 : 1
                          }}
                        >
                          List for Sale
                        </button>
                      )}
                    </>
                  )}
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

export default TokenDisplay;
