import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { IndexerAPI } from '../utils/api';
import { getClueFromContract, formatWei, formatDuration } from '../utils/web3';
import { formatEventMetadata } from '../utils/json';
import { Loading } from '../components/Loading';
import { ErrorMessage } from '../components/ErrorMessage';
import { OwnerLink } from '../components/OwnerLink';
import { EventBadge } from '../components/EventBadge';

export const ClueDetails = () => {
  const { clueId } = useParams();
  const navigate = useNavigate();
  const [clue, setClue] = useState(null);
  const [contractData, setContractData] = useState(null);
  const [events, setEvents] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [expandedEvent, setExpandedEvent] = useState(null);
  const [sortOrder, setSortOrder] = useState('desc');

  useEffect(() => {
    fetchClueData();
  }, [clueId]);

  const fetchClueData = async () => {
    try {
      setLoading(true);
      setError(null);

      const api = new IndexerAPI();

      const [clueData, eventsData, contractClue] = await Promise.all([
        api.getClue(clueId),
        api.getEventsByClueId(clueId, { order: sortOrder }),
        getClueFromContract(clueId).catch(() => null)
      ]);

      setClue(clueData);
      setEvents(eventsData || []);
      setContractData(contractClue);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  const toggleEventExpansion = (eventIndex) => {
    setExpandedEvent(expandedEvent === eventIndex ? null : eventIndex);
  };

  const toggleSortOrder = async () => {
    const newOrder = sortOrder === 'desc' ? 'asc' : 'desc';
    setSortOrder(newOrder);

    try {
      const api = new IndexerAPI();
      const eventsData = await api.getEventsByClueId(clueId, { order: newOrder });
      setEvents(eventsData || []);
    } catch (err) {
      console.error('Failed to sort events:', err);
    }
  };

  const formatTimestamp = (timestamp) => {
    if (!timestamp) return 'N/A';
    // Convert BigInt to Number before multiplying
    const timestampNum = typeof timestamp === 'bigint' ? Number(timestamp) : timestamp;
    const date = new Date(timestampNum * 1000);
    return date.toLocaleString();
  };

  if (loading) return <Loading message="Loading clue details..." />;
  if (error) return <ErrorMessage error={error} onRetry={fetchClueData} />;
  if (!clue) return <div className="text-center py-8">Clue not found</div>;

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <button
          onClick={() => navigate('/clues')}
          className="btn-secondary"
        >
          ← Back to Clues
        </button>
        <h1 className="text-3xl font-bold text-gray-900">Clue #{clueId}</h1>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="card">
          <h2 className="text-xl font-semibold mb-4">Indexer Data</h2>
          <dl className="space-y-3">
            <div>
              <dt className="text-sm font-medium text-gray-500">Clue ID</dt>
              <dd className="mt-1 text-sm text-gray-900">{clue.clue_id}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Owner</dt>
              <dd className="mt-1 text-sm text-gray-900">
                {clue.ownership?.owner_address ? (
                  <OwnerLink address={clue.ownership.owner_address} showFull />
                ) : (
                  <span className="text-gray-400">Unknown</span>
                )}
              </dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Point Value</dt>
              <dd className="mt-1 text-sm text-gray-900">{clue.point_value || 0} points</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Solve Reward</dt>
              <dd className="mt-1 text-sm text-gray-900">
                {clue.solve_reward ? `${formatWei(clue.solve_reward)} ETH` : '0 ETH'}
              </dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Solution Hash</dt>
              <dd className="mt-1 text-sm text-gray-900 font-mono break-all">
                {clue.solution_hash || 'N/A'}
              </dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Encrypted Contents</dt>
              <dd className="mt-1 text-sm text-gray-900 font-mono break-all max-h-32 overflow-y-auto bg-gray-50 p-2 rounded">
                {clue.contents || 'N/A'}
              </dd>
            </div>
          </dl>
        </div>

        {contractData && (
          <div className="card">
            <h2 className="text-xl font-semibold mb-4">Smart Contract Data</h2>
            <dl className="space-y-3">
              <div>
                <dt className="text-sm font-medium text-gray-500">Is Solved</dt>
                <dd className="mt-1">
                  <span className={`badge ${contractData.isSolved ? 'badge-success' : 'badge-warning'}`}>
                    {contractData.isSolved ? 'Solved' : 'Unsolved'}
                  </span>
                </dd>
              </div>
              <div>
                <dt className="text-sm font-medium text-gray-500">Sale Price</dt>
                <dd className="mt-1 text-sm text-gray-900">
                  {contractData.salePrice && contractData.salePrice !== '0'
                    ? `${formatWei(contractData.salePrice)} ETH`
                    : 'Not for sale'}
                </dd>
              </div>
              <div>
                <dt className="text-sm font-medium text-gray-500">R Value</dt>
                <dd className="mt-1 text-sm text-gray-900 font-mono break-all">
                  {contractData.rValue?.toString() || 'N/A'}
                </dd>
              </div>
              <div>
                <dt className="text-sm font-medium text-gray-500">Timeout</dt>
                <dd className="mt-1 text-sm text-gray-900">
                  {contractData.timeout ? formatDuration(contractData.timeout) : 'N/A'}
                </dd>
              </div>
              <div>
                <dt className="text-sm font-medium text-gray-500">Point Value (Contract)</dt>
                <dd className="mt-1 text-sm text-gray-900">{contractData.pointValue?.toString() || 0} points</dd>
              </div>
              <div>
                <dt className="text-sm font-medium text-gray-500">Solve Reward (Contract)</dt>
                <dd className="mt-1 text-sm text-gray-900">
                  {contractData.solveReward ? `${formatWei(contractData.solveReward)} ETH` : '0 ETH'}
                </dd>
              </div>
            </dl>
          </div>
        )}
      </div>

      {clue.ownership && (
        <div className="card">
          <h2 className="text-xl font-semibold mb-4">Ownership Details</h2>
          <dl className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Ownership Granted Block</dt>
              <dd className="mt-1 text-sm text-gray-900">{clue.ownership.ownership_granted_block_number}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Ownership Granted Event</dt>
              <dd className="mt-1">
                <EventBadge eventType={clue.ownership.ownership_granted_event_type} />
              </dd>
            </div>
          </dl>
        </div>
      )}

      <div className="card">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-semibold">Events ({events.length})</h2>
          <button
            onClick={toggleSortOrder}
            className="btn-secondary text-sm"
          >
            Sort: {sortOrder === 'desc' ? 'Newest First' : 'Oldest First'}
          </button>
        </div>

        {events.length === 0 ? (
          <div className="text-center py-8 text-gray-500">No events found for this clue.</div>
        ) : (
          <div className="space-y-2">
            {events.map((event, index) => (
              <div key={index} className="border border-gray-200 rounded-lg overflow-hidden">
                <div
                  onClick={() => toggleEventExpansion(index)}
                  className="flex justify-between items-center p-4 cursor-pointer hover:bg-gray-50 transition-colors"
                >
                  <div className="flex items-center gap-4">
                    <EventBadge eventType={event.event_type} />
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
                        <dt className="text-xs font-medium text-gray-500">Transaction Hash</dt>
                        <dd className="mt-1 text-xs text-gray-900 font-mono break-all">
                          {event.transaction_hash}
                        </dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Initiated By</dt>
                        <dd className="mt-1 text-xs text-gray-900">
                          <OwnerLink address={event.initiated_by} />
                        </dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Block Number</dt>
                        <dd className="mt-1 text-xs text-gray-900">{event.block_number}</dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Transaction Index</dt>
                        <dd className="mt-1 text-xs text-gray-900">{event.transaction_index}</dd>
                      </div>
                      <div>
                        <dt className="text-xs font-medium text-gray-500">Event Index</dt>
                        <dd className="mt-1 text-xs text-gray-900">{event.event_index}</dd>
                      </div>
                      <div>
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
