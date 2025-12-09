import React, { useState, useEffect } from 'react';
import Web3 from 'web3';
import { SKAVENGE_ABI } from '../contractABI';

function ClueMarketplace({ metamaskAddress, config, onToast }) {
  const [cluesForSale, setCluesForSale] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [processingPurchase, setProcessingPurchase] = useState(null);

  useEffect(() => {
    if (!config || !config.contractAddress) {
      setLoading(false);
      return;
    }

    fetchCluesForSale();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [config]);

  const fetchCluesForSale = async () => {
    try {
      setLoading(true);
      setError(null);

      // Initialize Web3
      const web3 = new Web3(window.ethereum);
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Get total number of clues for sale
      const total = await contract.methods.getTotalCluesForSale().call();
      const totalNum = Number(total);

      if (totalNum === 0) {
        setCluesForSale([]);
        setLoading(false);
        return;
      }

      // Fetch all clues for sale (offset=0, limit=totalNum)
      const result = await contract.methods.getCluesForSale(0, totalNum).call();

      // Parse the result into a more usable format
      const clues = result.tokenIds.map((tokenId, index) => ({
        tokenId: tokenId.toString(),
        owner: result.owners[index],
        price: result.prices[index].toString(),
        isSolved: result.solvedStatus[index]
      }));

      setCluesForSale(clues);
      setLoading(false);
    } catch (err) {
      console.error('Error fetching clues for sale:', err);
      setError(err.message || 'Failed to fetch marketplace listings');
      setLoading(false);

      if (onToast) {
        onToast('Failed to fetch marketplace: ' + (err.message || 'Unknown error'), 'error');
      }
    }
  };

  const formatPrice = (priceWei) => {
    if (!priceWei || priceWei === '0') return '0';
    const web3 = new Web3();
    return web3.utils.fromWei(priceWei, 'ether');
  };

  const truncateAddress = (address) => {
    if (!address) return '';
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
  };

  const handlePurchase = async (tokenId, price) => {
    try {
      setProcessingPurchase(tokenId);

      if (!metamaskAddress) {
        if (onToast) {
          onToast('Please connect your MetaMask wallet to purchase tokens', 'error');
        }
        return;
      }

      const web3 = new Web3(window.ethereum);
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Call initiatePurchase with the token price as value
      await contract.methods.initiatePurchase(tokenId).send({
        from: metamaskAddress,
        value: price
      });

      if (onToast) {
        onToast(`Purchase initiated for token #${tokenId}`, 'success');
      }

      // Refresh the marketplace
      await fetchCluesForSale();
    } catch (error) {
      console.error('Error purchasing token:', error);
      if (onToast) {
        onToast('Failed to purchase token: ' + (error.message || 'Unknown error'), 'error');
      }
    } finally {
      setProcessingPurchase(null);
    }
  };

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
        <div className="token-display-message">Loading marketplace...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="token-display-container">
        <div className="token-display-error">
          <h3>Error</h3>
          <p>{error}</p>
          <button onClick={fetchCluesForSale} className="btn-retry">Retry</button>
        </div>
      </div>
    );
  }

  if (cluesForSale.length === 0) {
    return (
      <div className="token-display-container">
        <div className="token-display-message">
          <h3>No Clues For Sale</h3>
          <p>There are currently no clues listed in the marketplace.</p>
        </div>
      </div>
    );
  }

  return (
    <div className="token-display-container">
      <div className="token-display-header">
        <h2>Clue Marketplace</h2>
        <p className="token-count">{cluesForSale.length} {cluesForSale.length === 1 ? 'clue' : 'clues'} for sale</p>
        <button onClick={fetchCluesForSale} className="btn-refresh">Refresh</button>
      </div>

      <div className="token-grid">
        {cluesForSale.map((clue) => (
          <div key={clue.tokenId} className="token-card">
            <div className="token-card-header">
              <h3>Token #{clue.tokenId}</h3>
              {clue.isSolved && <span className="badge-solved">Solved</span>}
              <span className="badge-for-sale">For Sale</span>
            </div>

            <div className="token-details">
              <div className="token-detail-row">
                <span className="detail-label">Price:</span>
                <span className="detail-value" style={{ color: '#667eea', fontWeight: '600' }}>
                  {formatPrice(clue.price)} ETH
                </span>
              </div>

              <div className="token-detail-row">
                <span className="detail-label">Owner:</span>
                <span className="detail-value" style={{ fontFamily: 'Monaco, monospace', fontSize: '12px' }}>
                  {truncateAddress(clue.owner)}
                </span>
              </div>

              <div className="token-detail-row">
                <span className="detail-label">Status:</span>
                <span className={`detail-value ${clue.isSolved ? 'text-success' : 'text-warning'}`}>
                  {clue.isSolved ? 'Solved' : 'Unsolved'}
                </span>
              </div>

              {/* Purchase Button */}
              {metamaskAddress && metamaskAddress.toLowerCase() !== clue.owner.toLowerCase() && (
                <div className="token-actions" style={{ marginTop: '16px' }}>
                  <button
                    onClick={() => handlePurchase(clue.tokenId, clue.price)}
                    disabled={processingPurchase === clue.tokenId}
                    className="btn-purchase"
                    style={{
                      width: '100%',
                      padding: '10px',
                      fontSize: '14px',
                      backgroundColor: '#48bb78',
                      color: 'white',
                      border: 'none',
                      borderRadius: '6px',
                      cursor: processingPurchase === clue.tokenId ? 'not-allowed' : 'pointer',
                      fontWeight: '500',
                      opacity: processingPurchase === clue.tokenId ? 0.6 : 1
                    }}
                  >
                    {processingPurchase === clue.tokenId ? 'Processing...' : `Buy for ${formatPrice(clue.price)} ETH`}
                  </button>
                </div>
              )}

              {metamaskAddress && metamaskAddress.toLowerCase() === clue.owner.toLowerCase() && (
                <div style={{ marginTop: '16px', padding: '8px', backgroundColor: '#f7fafc', borderRadius: '6px', textAlign: 'center' }}>
                  <span style={{ fontSize: '12px', color: '#718096' }}>You own this token</span>
                </div>
              )}

              {!metamaskAddress && (
                <div style={{ marginTop: '16px', padding: '8px', backgroundColor: '#fff5f5', borderRadius: '6px', textAlign: 'center' }}>
                  <span style={{ fontSize: '12px', color: '#e53e3e' }}>Connect wallet to purchase</span>
                </div>
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default ClueMarketplace;
