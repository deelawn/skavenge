import React, { useState } from 'react';
import { loadConfig } from '../config.js';
import { getBrowserGatewayUrl } from '../utils.js';
import { Web3 } from 'web3';

/**
 * Sign a linkage message using MetaMask
 * @param {string} address - The Ethereum address
 * @param {string} skavengerPublicKey - The Skavenge public key
 * @returns {Promise<{signature: string, message: string}>}
 */
async function signLinkageMessage(address, skavengerPublicKey) {
  const message = `Link MetaMask address ${address} to Skavenge key ${skavengerPublicKey}`;

  try {
    const signature = await window.ethereum.request({
      method: 'personal_sign',
      params: [message, address],
    });

    return { signature, message };
  } catch (error) {
    if (error.code === 4001) {
      throw new Error('Signature request rejected by user.');
    }
    throw new Error('Failed to sign message: ' + (error.message || 'Unknown error'));
  }
}

/**
 * Verify signature locally
 * @param {string} message - The message that was signed
 * @param {string} signature - The signature
 * @param {string} address - The expected address
 * @returns {boolean}
 */
function verifySignature(message, signature, address) {
  try {
    const web3 = new Web3();
    const recoveredAddress = web3.eth.accounts.recover(message, signature);
    return recoveredAddress.toLowerCase() === address.toLowerCase();
  } catch (error) {
    console.error('Signature verification failed:', error);
    return false;
  }
}

/**
 * Send linkage to gateway server
 * @param {string} ethereumAddress - The Ethereum address
 * @param {string} skavengePublicKey - The Skavenge public key
 * @param {string} message - The signed message
 * @param {string} signature - The signature
 * @returns {Promise<{success: boolean, error?: string}>}
 */
async function sendLinkageToGateway(ethereumAddress, skavengePublicKey, message, signature) {
  try {
    const config = await loadConfig();
    const gatewayUrl = getBrowserGatewayUrl(config.gatewayUrl);

    const response = await fetch(`${gatewayUrl}/link`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        ethereum_address: ethereumAddress,
        skavenge_public_key: skavengePublicKey,
        message: message,
        signature: signature,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      if (response.status === 409) {
        return {
          success: false,
          error: 'This Ethereum address is already linked to a different Skavenge key. Each address can only be linked once.',
        };
      } else if (response.status === 401) {
        return {
          success: false,
          error: 'Signature verification failed on the server.',
        };
      } else {
        return {
          success: false,
          error: data.error || 'Failed to link account on the server.',
        };
      }
    }

    return { success: true };
  } catch (error) {
    console.error('Error sending linkage to gateway:', error);
    return {
      success: false,
      error: 'Failed to connect to the gateway server. Please try again later.',
    };
  }
}

/**
 * RelinkPrompt Component
 * Displayed when both accounts are connected locally but linkage doesn't exist on server
 *
 * @param {object} props
 * @param {string} props.metamaskAddress - The MetaMask address
 * @param {string} props.skavengerPublicKey - The Skavenge public key
 * @param {function} props.onSuccess - Callback when re-linking succeeds
 * @param {function} props.onToast - Callback to show toast message
 */
function RelinkPrompt({ metamaskAddress, skavengerPublicKey, onSuccess, onToast }) {
  const [isRetrying, setIsRetrying] = useState(false);
  const [error, setError] = useState(null);

  const handleRelink = async () => {
    setIsRetrying(true);
    setError(null);

    try {
      // Request user to sign the linkage message
      onToast('Please sign the message in MetaMask to verify ownership', 'info');
      const { signature, message } = await signLinkageMessage(metamaskAddress, skavengerPublicKey);

      // Verify signature locally
      const isValid = verifySignature(message, signature, metamaskAddress);

      if (!isValid) {
        setError('Signature verification failed. Unable to confirm address ownership.');
        onToast('Signature verification failed', 'error');
        setIsRetrying(false);
        return;
      }

      // Send to gateway server
      onToast('Sending linkage to server for verification...', 'info');
      const result = await sendLinkageToGateway(metamaskAddress, skavengerPublicKey, message, signature);

      if (!result.success) {
        setError(result.error);
        onToast(result.error, 'error');
        setIsRetrying(false);
        return;
      }

      // Success!
      onToast('Account linkage verified successfully!', 'success');
      setIsRetrying(false);
      onSuccess();
    } catch (err) {
      setError(err.message);
      onToast(err.message, 'error');
      setIsRetrying(false);
    }
  };

  return (
    <div className="screen">
      <div className="step" style={{
        border: '2px solid #f59e0b',
        backgroundColor: '#fffbeb'
      }}>
        <h2 style={{ color: '#d97706' }}>⚠️ Account Linkage Required</h2>

        <div style={{
          backgroundColor: '#fef3c7',
          padding: '16px',
          borderRadius: '8px',
          marginBottom: '20px',
          border: '1px solid #fbbf24'
        }}>
          <p style={{ margin: '0 0 12px 0', fontWeight: '600', color: '#92400e' }}>
            Your accounts are connected locally, but not verified on the server.
          </p>
          <p style={{ margin: '0', color: '#92400e' }}>
            This can happen if a previous linkage request failed due to network issues.
            Please sign the linkage message again to complete the verification.
          </p>
        </div>

        <div className="account-info" style={{ marginBottom: '20px' }}>
          <div style={{ marginBottom: '12px' }}>
            <div style={{ fontWeight: '600', color: '#4b5563', marginBottom: '4px' }}>
              Ethereum Address:
            </div>
            <div style={{
              fontFamily: 'monospace',
              fontSize: '14px',
              padding: '8px',
              backgroundColor: '#f3f4f6',
              borderRadius: '4px',
              wordBreak: 'break-all'
            }}>
              {metamaskAddress}
            </div>
          </div>

          <div>
            <div style={{ fontWeight: '600', color: '#4b5563', marginBottom: '4px' }}>
              Skavenge Public Key:
            </div>
            <div style={{
              fontFamily: 'monospace',
              fontSize: '14px',
              padding: '8px',
              backgroundColor: '#f3f4f6',
              borderRadius: '4px',
              wordBreak: 'break-all'
            }}>
              {skavengerPublicKey}
            </div>
          </div>
        </div>

        {error && (
          <div style={{
            backgroundColor: '#fee2e2',
            border: '1px solid #fca5a5',
            padding: '12px',
            borderRadius: '8px',
            marginBottom: '16px',
            color: '#991b1b'
          }}>
            <strong>Error:</strong> {error}
          </div>
        )}

        <button
          className="primary"
          onClick={handleRelink}
          disabled={isRetrying}
          style={{
            width: '100%',
            padding: '14px',
            fontSize: '16px',
            fontWeight: '600'
          }}
        >
          {isRetrying ? 'Verifying...' : 'Sign & Verify Linkage'}
        </button>

        <p style={{
          marginTop: '16px',
          fontSize: '14px',
          color: '#6b7280',
          textAlign: 'center'
        }}>
          All app features are disabled until linkage is verified
        </p>
      </div>
    </div>
  );
}

export default RelinkPrompt;
