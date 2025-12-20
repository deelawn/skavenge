import React, { useState, useEffect } from 'react';
import Web3 from 'web3';
import { SKAVENGE_ABI } from '../contractABI';
import { getPublicKeyByEthereumAddress, storeTransferCiphertext } from '../utils';

// Extension ID - hardcoded to match manifest.json
const SKAVENGER_EXTENSION_ID = "hnbligdjmpihmhhgajlfjmckcnmnofbn";

/**
 * Send message to Skavenger extension
 */
async function sendToExtension(message) {
  return new Promise((resolve, reject) => {
    if (typeof window.chrome === 'undefined' || !window.chrome.runtime) {
      reject(new Error('Chrome extension API not available'));
      return;
    }

    window.chrome.runtime.sendMessage(SKAVENGER_EXTENSION_ID, message, (response) => {
      if (window.chrome.runtime.lastError) {
        reject(new Error(window.chrome.runtime.lastError.message));
      } else if (response && response.success !== false) {
        resolve(response);
      } else {
        reject(new Error(response?.error || 'Unknown error'));
      }
    });
  });
}

function Transfers({ metamaskAddress, config, onToast }) {
  const [transfers, setTransfers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [processingTransfer, setProcessingTransfer] = useState(null); // transferId being processed

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

  /**
   * Handle providing proof for a transfer
   * This is the seller's action after a buyer initiates a purchase
   */
  const handleProvideProof = async (transfer) => {
    setProcessingTransfer(transfer.transferId);

    try {
      const web3 = new Web3(window.ethereum);
      const contract = new web3.eth.Contract(SKAVENGE_ABI, config.contractAddress);

      // Step 1: Get the buyer's Skavenge public key from the gateway
      onToast('Getting buyer public key...', 'info');
      const buyerPubKeyResult = await getPublicKeyByEthereumAddress(transfer.buyer);
      if (!buyerPubKeyResult.success) {
        throw new Error(buyerPubKeyResult.error || 'Failed to get buyer public key');
      }
      const buyerPublicKey = buyerPubKeyResult.publicKey;

      // Step 2: Get the clue contents and r value from the contract
      onToast('Retrieving clue data...', 'info');
      const encryptedContents = await contract.methods.getClueContents(transfer.tokenId).call({
        from: metamaskAddress
      });
      const rValue = await contract.methods.getRValue(transfer.tokenId).call({
        from: metamaskAddress
      });

      // Remove '0x' prefix if present for the encrypted contents
      let encryptedHex = encryptedContents;
      if (encryptedHex.startsWith('0x')) {
        encryptedHex = encryptedHex.slice(2);
      }

      // Convert rValue to hex string
      let rValueHex = window.BigInt(rValue).toString(16);
      if (rValueHex.length % 2 !== 0) rValueHex = '0' + rValueHex;

      // Step 3: Decrypt the clue content using the extension
      onToast('Decrypting clue content...', 'info');
      const decryptResponse = await sendToExtension({
        action: 'decryptElGamal',
        encryptedHex: encryptedHex,
        rValueHex: rValueHex
      });

      if (!decryptResponse.success) {
        throw new Error(decryptResponse.error || 'Failed to decrypt clue content');
      }
      const plaintext = decryptResponse.plaintext;

      // Step 4: Generate the transfer proof using the extension
      onToast('Generating transfer proof...', 'info');
      const proofResponse = await sendToExtension({
        action: 'generateTransferProof',
        plaintext: plaintext,
        buyerPublicKey: buyerPublicKey
      });

      if (!proofResponse.success) {
        throw new Error(proofResponse.error || 'Failed to generate transfer proof');
      }

      // Step 5: Compute the buyer ciphertext hash
      const buyerCiphertextHash = web3.utils.keccak256(proofResponse.buyerCiphertext);

      // Debug logging
      console.log('=== PROVIDE PROOF DEBUG ===');
      console.log('Transfer ID:', transfer.transferId);
      console.log('Proof (hex):', proofResponse.proof);
      console.log('Proof length (bytes):', (proofResponse.proof.length - 2) / 2);
      console.log('Buyer ciphertext hash:', buyerCiphertextHash);
      console.log('Sending from:', metamaskAddress);
      console.log('Token ID:', transfer.tokenId);
      console.log('=== END DEBUG ===');

      // Step 6: Send the provideProof transaction via MetaMask
      onToast('Please confirm the transaction in MetaMask...', 'info');

      // Try to estimate gas first to catch any revert reasons
      try {
        const gasEstimate = await contract.methods.provideProof(
          transfer.transferId,
          proofResponse.proof,
          buyerCiphertextHash
        ).estimateGas({ from: metamaskAddress });
        console.log('Gas estimate:', gasEstimate);
      } catch (gasError) {
        console.error('Gas estimation failed:', gasError);
        // Try to extract revert reason
        if (gasError.message) {
          throw new Error('Transaction would fail: ' + gasError.message);
        }
        throw gasError;
      }

      const tx = await contract.methods.provideProof(
        transfer.transferId,
        proofResponse.proof,
        buyerCiphertextHash
      ).send({
        from: metamaskAddress
      });

      console.log('ProvideProof transaction:', tx);
      onToast('Proof provided successfully! Storing ciphertexts...', 'success');

      // Step 7: Store the ciphertexts on the gateway
      const storeResult = await storeTransferCiphertext(
        transfer.transferId,
        proofResponse.buyerCiphertext,
        proofResponse.sellerCiphertext,
        SKAVENGER_EXTENSION_ID
      );

      if (!storeResult.success) {
        console.warn('Failed to store ciphertexts on gateway:', storeResult.error);
        onToast('Proof provided but failed to store ciphertexts: ' + storeResult.error, 'warning');
      } else {
        onToast('Transfer proof submitted successfully!', 'success');
      }

      // Reload transfers to update the UI
      await loadTransfers();

    } catch (err) {
      console.error('Error providing proof:', err);

      // Handle user rejection
      if (err.code === 4001 || err.message?.includes('User denied')) {
        onToast('Transaction was cancelled', 'info');
      } else {
        onToast('Failed to provide proof: ' + err.message, 'error');
      }
    } finally {
      setProcessingTransfer(null);
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

              {/* Action Buttons */}
              {transfer.userRole === 'seller' && transfer.status === 'initiated' && (
                <div style={{ marginTop: '16px', paddingTop: '16px', borderTop: '1px solid #e2e8f0' }}>
                  <button
                    onClick={() => handleProvideProof(transfer)}
                    disabled={processingTransfer === transfer.transferId}
                    style={{
                      width: '100%',
                      padding: '12px 16px',
                      backgroundColor: processingTransfer === transfer.transferId ? '#adb5bd' : '#667eea',
                      color: 'white',
                      border: 'none',
                      borderRadius: '6px',
                      cursor: processingTransfer === transfer.transferId ? 'not-allowed' : 'pointer',
                      fontSize: '14px',
                      fontWeight: '600',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      gap: '8px'
                    }}
                  >
                    {processingTransfer === transfer.transferId ? 'Processing...' : 'Provide Proof'}
                  </button>
                  <p style={{
                    fontSize: '12px',
                    color: '#718096',
                    marginTop: '8px',
                    textAlign: 'center'
                  }}>
                    Generate and submit the transfer proof for the buyer
                  </p>
                </div>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default Transfers;
