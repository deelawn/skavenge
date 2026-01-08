import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { IndexerAPI, GatewayAPI } from '../utils/api';
import { formatEventMetadata } from '../utils/json';
import { Loading } from '../components/Loading';
import { ErrorMessage } from '../components/ErrorMessage';
import { ClueLink } from '../components/ClueLink';
import { OwnerLink } from '../components/OwnerLink';
import { EventBadge } from '../components/EventBadge';
import { formatAddress } from '../utils/web3';

export const OwnerDetails = () => {
  const { address } = useParams();
  const navigate = useNavigate();
  const [ownership, setOwnership] = useState([]);
  const [events, setEvents] = useState([]);
  const [linkedKey, setLinkedKey] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [expandedEvent, setExpandedEvent] = useState(null);
  const [sortOrder, setSortOrder] = useState('desc');

  useEffect(() => {
    fetchOwnerData();
  }, [address, sortOrder]);

  const fetchOwnerData = async () => {
    try {
      setLoading(true);
      setError(null);

      const indexerApi = new IndexerAPI();
      const gatewayApi = new GatewayAPI();

      const [ownershipData, eventsData, linkData] = await Promise.all([
        indexerApi.getOwnershipByAddress(address),
        indexerApi.getEventsByInitiator(address, { order: sortOrder }),
        gatewayApi.getLinkedPublicKey(address).catch(() => ({ success: false }))
      ]);

      setOwnership(ownershipData || []);
      setEvents(eventsData || []);
      setLinkedKey(linkData?.success ? linkData.skavenge_public_key : null);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  const toggleEventExpansion = (eventIndex) => {
    setExpandedEvent(expandedEvent === eventIndex ? null : eventIndex);
  };

  const toggleSortOrder = () => {
    setSortOrder(sortOrder === 'desc' ? 'asc' : 'desc');
  };

  const formatTimestamp = (timestamp) => {
    if (!timestamp) return 'N/A';
    // Convert BigInt to Number before multiplying
    const timestampNum = typeof timestamp === 'bigint' ? Number(timestamp) : timestamp;
    const date = new Date(timestampNum * 1000);
    return date.toLocaleString();
  };

  if (loading) return <Loading message="Loading owner details..." />;
  if (error) return <ErrorMessage error={error} onRetry={fetchOwnerData} />;

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <button
          onClick={() => navigate('/owners')}
          className="btn-secondary"
        >
          ← Back to Owners
        </button>
        <h1 className="text-2xl font-bold text-gray-900 font-mono break-all">
          {formatAddress(address)}
        </h1>
      </div>

      <div className="card">
        <h2 className="text-xl font-semibold mb-4">Owner Information</h2>
        <dl className="space-y-3">
          <div>
            <dt className="text-sm font-medium text-gray-500">Ethereum Address</dt>
            <dd className="mt-1 text-sm text-gray-900 font-mono break-all">{address}</dd>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Skavenge Public Key</dt>
            <dd className="mt-1 text-sm text-gray-900 font-mono break-all">
              {linkedKey || (
                <span className="text-gray-400 italic">Not linked</span>
              )}
            </dd>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Clues Owned</dt>
            <dd className="mt-1">
              <span className="badge-info">{ownership.length} clues</span>
            </dd>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Total Activity</dt>
            <dd className="mt-1">
              <span className="badge-info">{events.length} events</span>
            </dd>
          </div>
        </dl>
      </div>

      <div className="card">
        <h2 className="text-xl font-semibold mb-4">Owned Clues ({ownership.length})</h2>
        {ownership.length === 0 ? (
          <div className="text-center py-8 text-gray-500">This owner has no clues.</div>
        ) : (
          <div className="overflow-x-auto">
            <table className="table">
              <thead className="table-header">
                <tr>
                  <th className="table-header-cell">Clue ID</th>
                  <th className="table-header-cell">Ownership Granted</th>
                  <th className="table-header-cell">Event Type</th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {ownership.map((own) => (
                  <tr key={own.clue_id} className="hover:bg-gray-50">
                    <td className="table-cell">
                      <ClueLink clueId={own.clue_id} />
                    </td>
                    <td className="table-cell">
                      Block {own.ownership_granted_block_number}
                    </td>
                    <td className="table-cell">
                      <EventBadge eventType={own.ownership_granted_event_type} />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>

      <div className="card">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-semibold">Activity Log ({events.length})</h2>
          <button
            onClick={toggleSortOrder}
            className="btn-secondary text-sm"
          >
            Sort: {sortOrder === 'desc' ? 'Newest First' : 'Oldest First'}
          </button>
        </div>

        {events.length === 0 ? (
          <div className="text-center py-8 text-gray-500">No activity found for this owner.</div>
        ) : (
          <div className="space-y-2">
            {events.map((event, index) => (
              <div key={index} className="border border-gray-200 rounded-lg overflow-hidden">
                <div
                  onClick={() => toggleEventExpansion(index)}
                  className="flex justify-between items-center p-4 cursor-pointer hover:bg-gray-50 transition-colors"
                >
                  <div className="flex items-center gap-4 flex-wrap">
                    <EventBadge eventType={event.event_type} />
                    {event.clue_id && (
                      <ClueLink clueId={event.clue_id} />
                    )}
                    <span className="text-sm text-gray-600">
                      {formatTimestamp(event.timestamp)}
                    </span>
                    <span className="text-xs text-gray-500 font-mono">
                      Block {event.block_number}
                    </span>
                  </div>
                  <svg
                    className={`w-5 h-5 text-gray-400 transition-transform ${expandedEvent === index ? 'rotate-180' : ''}`}
                    fill="none"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path d="M19 9l-7 7-7-7"></path>
                  </svg>
                </div>

                {expandedEvent === index && (
                  <div className="border-t border-gray-200 bg-gray-50 p-4">
                    <dl className="grid grid-cols-1 md:grid-cols-2 gap-3">
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Event Type</dt>
                        <dd className="mt-1">
                          <EventBadge eventType={event.event_type} />
                        </dd>
                      </div>
                      {event.clue_id && (
                        <div>
                          <dt className="text-xs font-medium text-gray-500">Clue ID</dt>
                          <dd className="mt-1 text-xs">
                            <ClueLink clueId={event.clue_id} />
                          </dd>
                        </div>
                      )}
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Transaction Hash</dt>
                        <dd className="mt-1 text-xs text-gray-900 font-mono break-all">
                          {event.transaction_hash}
                        </dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Block Number</dt>
                        <dd className="mt-1 text-xs text-gray-900">{event.block_number}</dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Timestamp</dt>
                        <dd className="mt-1 text-xs text-gray-900">{formatTimestamp(event.timestamp)}</dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Transaction Index</dt>
                        <dd className="mt-1 text-xs text-gray-900">{event.transaction_index}</dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Event Index</dt>
                        <dd className="mt-1 text-xs text-gray-900">{event.event_index}</dd>
                      </div>
                      <div className="md:col-span-2">
                        <dt className="text-xs font-medium text-gray-500">Block Hash</dt>
                        <dd className="mt-1 text-xs text-gray-900 font-mono break-all">
                          {event.block_hash}
                        </dd>
                      </div>
                      {event.metadata && (
                        <div className="md:col-span-2">
                          <dt className="text-xs font-medium text-gray-500">Event Metadata</dt>
                          <dd className="mt-1 text-xs text-gray-900 font-mono bg-white p-2 rounded overflow-x-auto">
                            <pre>{formatEventMetadata(event.metadata)}</pre>
                          </dd>
                        </div>
                      )}
                    </dl>
                  </div>
                )}
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};
