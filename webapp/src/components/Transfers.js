import React, { useState, useEffect } from 'react';
import Web3 from 'web3';
import { SKAVENGE_ABI } from '../contractABI';

function Transfers({ metamaskAddress, config, onToast }) {
  const [transfers, setTransfers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (metamaskAddress && config) {
      loadTransfers();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [metamaskAddress, config]);

  const loadTransfers = async () => {
    setLoading(true);
    setError(null);

    try {
      const web3 = new Web3(window.ethereum);
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Get TransferInitiated events from the beginning of the contract
      const events = await contract.getPastEvents('TransferInitiated', {
        fromBlock: 0,
        toBlock: 'latest'
      });

      // Filter and fetch transfer details
      const transferPromises = events.map(async (event) => {
        const { transferId, buyer, tokenId } = event.returnValues;

        // Get the current transfer state
        const transfer = await contract.methods.transfers(transferId).call();

        // Skip if transfer is completed (buyer would be address(0))
        if (transfer.buyer === '0x0000000000000000000000000000000000000000') {
          return null;
        }

        // Get token owner to determine if user is seller
        const owner = await contract.methods.ownerOf(tokenId).call();

        // Only include if user is buyer or seller
        const isBuyer = buyer.toLowerCase() === metamaskAddress.toLowerCase();
        const isSeller = owner.toLowerCase() === metamaskAddress.toLowerCase();

        if (!isBuyer && !isSeller) {
          return null;
        }

        // Determine status based on proof and proofVerified
        let status;
        if (transfer.proofVerified) {
          status = 'proof verified';
        } else if (transfer.proof && transfer.proof !== '0x') {
          status = 'proof provided';
        } else {
          status = 'initiated';
        }

        return {
          transferId,
          tokenId: tokenId.toString(),
          buyer,
          seller: owner,
          value: transfer.value.toString(),
          status,
          initiatedAt: new Date(Number(transfer.initiatedAt) * 1000).toLocaleString(),
          proofProvidedAt: transfer.proofProvidedAt > 0
            ? new Date(Number(transfer.proofProvidedAt) * 1000).toLocaleString()
            : null,
          verifiedAt: transfer.verifiedAt > 0
            ? new Date(Number(transfer.verifiedAt) * 1000).toLocaleString()
            : null,
          userRole: isBuyer ? 'buyer' : 'seller'
        };
      });

      const transfersData = (await Promise.all(transferPromises)).filter(t => t !== null);

      // Sort by initiation time (most recent first)
      transfersData.sort((a, b) => {
        const timeA = new Date(a.initiatedAt).getTime();
        const timeB = new Date(b.initiatedAt).getTime();
        return timeB - timeA;
      });

      setTransfers(transfersData);
    } catch (err) {
      console.error('Error loading transfers:', err);
      setError('Failed to load transfers. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const getStatusColor = (status) => {
    switch (status) {
      case 'initiated':
        return '#e67700';
      case 'proof provided':
        return '#2b8a3e';
      case 'proof verified':
        return '#51cf66';
      default:
        return '#718096';
    }
  };

  const formatAddress = (address) => {
    return `${address.substring(0, 6)}...${address.substring(38)}`;
  };

  const formatValue = (valueWei) => {
    const web3 = new Web3(window.ethereum);
    return web3.utils.fromWei(valueWei, 'ether');
  };

  if (loading) {
    return (
      <div className="token-display-container">
        <div style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          minHeight: '400px',
          color: '#718096'
        }}>
          Loading transfers...
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="token-display-container">
        <div style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          minHeight: '400px',
          gap: '16px'
        }}>
          <div style={{ color: '#e03131', fontSize: '16px' }}>{error}</div>
          <button
            onClick={loadTransfers}
            style={{
              padding: '8px 16px',
              backgroundColor: '#667eea',
              color: 'white',
              border: 'none',
              borderRadius: '6px',
              cursor: 'pointer'
            }}
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="token-display-container">
      <div className="token-display-header">
        <div style={{ display: 'flex', alignItems: 'center', gap: '12px' }}>
          <h2 style={{ margin: 0, fontSize: '24px', fontWeight: '600' }}>
            Transfers
          </h2>
          <span style={{
            padding: '4px 12px',
            backgroundColor: '#e9ecef',
            borderRadius: '12px',
            fontSize: '14px',
            fontWeight: '500',
            color: '#495057'
          }}>
            {transfers.length}
          </span>
        </div>
        <button
          onClick={loadTransfers}
          disabled={loading}
          style={{
            padding: '8px 16px',
            backgroundColor: loading ? '#adb5bd' : '#667eea',
            color: 'white',
            border: 'none',
            borderRadius: '6px',
            cursor: loading ? 'not-allowed' : 'pointer',
            fontSize: '14px',
            fontWeight: '500'
          }}
        >
          Refresh
        </button>
      </div>

      {transfers.length === 0 ? (
        <div style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          minHeight: '300px',
          color: '#718096'
        }}>
          No active transfers found
        </div>
      ) : (
        <div className="token-grid">
          {transfers.map((transfer) => (
            <div key={transfer.transferId} className="token-card">
              <div className="token-card-header">
                <h3 style={{ margin: 0, fontSize: '18px', fontWeight: '600' }}>
                  Token #{transfer.tokenId}
                </h3>
                <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
                  <span style={{
                    padding: '4px 8px',
                    backgroundColor: getStatusColor(transfer.status),
                    color: 'white',
                    borderRadius: '4px',
                    fontSize: '12px',
                    fontWeight: '500',
                    textTransform: 'capitalize'
                  }}>
                    {transfer.status}
                  </span>
                  <span style={{
                    padding: '4px 8px',
                    backgroundColor: transfer.userRole === 'buyer' ? '#667eea' : '#2b8a3e',
                    color: 'white',
                    borderRadius: '4px',
                    fontSize: '12px',
                    fontWeight: '500',
                    textTransform: 'capitalize'
                  }}>
                    {transfer.userRole}
                  </span>
                </div>
              </div>

              <div className="token-details">
                <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                  <span className="detail-label">Buyer</span>
                  <span className="detail-value" style={{ fontFamily: 'monospace', fontSize: '13px' }}>
                    {formatAddress(transfer.buyer)}
                  </span>
                </div>

                <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                  <span className="detail-label">Seller</span>
                  <span className="detail-value" style={{ fontFamily: 'monospace', fontSize: '13px' }}>
                    {formatAddress(transfer.seller)}
                  </span>
                </div>

                <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                  <span className="detail-label">Price</span>
                  <span className="detail-value">{formatValue(transfer.value)} ETH</span>
                </div>

                <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                  <span className="detail-label">Initiated</span>
                  <span className="detail-value" style={{ fontSize: '13px' }}>
                    {transfer.initiatedAt}
                  </span>
                </div>

                {transfer.proofProvidedAt && (
                  <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                    <span className="detail-label">Proof Provided</span>
                    <span className="detail-value" style={{ fontSize: '13px' }}>
                      {transfer.proofProvidedAt}
                    </span>
                  </div>
                )}

                {transfer.verifiedAt && (
                  <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                    <span className="detail-label">Verified</span>
                    <span className="detail-value" style={{ fontSize: '13px' }}>
                      {transfer.verifiedAt}
                    </span>
                  </div>
                )}
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default Transfers;
