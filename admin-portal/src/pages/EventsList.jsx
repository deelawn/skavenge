import { useState, useEffect } from 'react';
import { IndexerAPI } from '../utils/api';
import { formatEventMetadata } from '../utils/json';
import { Loading } from '../components/Loading';
import { ErrorMessage } from '../components/ErrorMessage';
import { Pagination } from '../components/Pagination';
import { EventBadge } from '../components/EventBadge';
import { ClueLink } from '../components/ClueLink';
import { OwnerLink } from '../components/OwnerLink';

const ITEMS_PER_PAGE = 50;

const EVENT_TYPES = [
  'All Events',
  'ClueMinted',
  'ClueAttempted',
  'ClueSolved',
  'SalePriceSet',
  'SalePriceRemoved',
  'TransferInitiated',
  'ProofProvided',
  'ProofVerified',
  'TransferCompleted',
  'TransferCancelled',
  'AuthorizedMinterUpdated',
  'Transfer',
  'Approval',
  'ApprovalForAll',
];

export const EventsList = () => {
  const [events, setEvents] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [expandedEvent, setExpandedEvent] = useState(null);
  const [filterType, setFilterType] = useState('All Events');
  const [sortOrder, setSortOrder] = useState('desc');

  useEffect(() => {
    fetchEvents();
  }, [filterType, sortOrder]);

  const fetchEvents = async () => {
    try {
      setLoading(true);
      setError(null);

      const api = new IndexerAPI();
      const params = { order: sortOrder };

      let data;
      if (filterType === 'All Events') {
        data = await api.getEvents(params);
      } else {
        data = await api.getEventsByType(filterType, params);
      }

      setEvents(data || []);
      setCurrentPage(1);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  const totalPages = Math.ceil(events.length / ITEMS_PER_PAGE);
  const startIndex = (currentPage - 1) * ITEMS_PER_PAGE;
  const paginatedEvents = events.slice(startIndex, startIndex + ITEMS_PER_PAGE);

  const handlePageChange = (page) => {
    setCurrentPage(page);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const toggleEventExpansion = (eventIndex) => {
    setExpandedEvent(expandedEvent === eventIndex ? null : eventIndex);
  };

  const formatTimestamp = (timestamp) => {
    if (!timestamp) return 'N/A';
    // Convert BigInt to Number before multiplying
    const timestampNum = typeof timestamp === 'bigint' ? Number(timestamp) : timestamp;
    const date = new Date(timestampNum * 1000);
    return date.toLocaleString();
  };

  if (loading) return <Loading message="Loading events..." />;
  if (error) return <ErrorMessage error={error} onRetry={fetchEvents} />;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold text-gray-900">Events</h1>
        <div className="text-sm text-gray-600">
          Total: {events.length} events
        </div>
      </div>

      <div className="card">
        <div className="flex flex-col md:flex-row gap-4 mb-6">
          <div className="flex-1">
            <label className="label">Filter by Event Type</label>
            <select
              value={filterType}
              onChange={(e) => setFilterType(e.target.value)}
              className="input-field"
            >
              {EVENT_TYPES.map((type) => (
                <option key={type} value={type}>
                  {type}
                </option>
              ))}
            </select>
          </div>
          <div className="flex-1">
            <label className="label">Sort Order</label>
            <select
              value={sortOrder}
              onChange={(e) => setSortOrder(e.target.value)}
              className="input-field"
            >
              <option value="desc">Newest First</option>
              <option value="asc">Oldest First</option>
            </select>
          </div>
        </div>

        {events.length === 0 ? (
          <div className="text-center py-8 text-gray-500">
            No events found {filterType !== 'All Events' && `for ${filterType}`}.
          </div>
        ) : (
          <>
            <div className="space-y-2 mb-4">
              {paginatedEvents.map((event, index) => {
                const globalIndex = startIndex + index;
                return (
                  <div key={globalIndex} className="border border-gray-200 rounded-lg overflow-hidden">
                    <div
                      onClick={() => toggleEventExpansion(globalIndex)}
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
                        {event.initiated_by && (
                          <OwnerLink address={event.initiated_by} />
                        )}
                      </div>
                      <svg
                        className={`w-5 h-5 text-gray-400 transition-transform ${expandedEvent === globalIndex ? 'rotate-180' : ''}`}
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

                    {expandedEvent === globalIndex && (
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
                            <dt className="text-xs font-medium text-gray-500">Initiated By</dt>
                            <dd className="mt-1 text-xs">
                              <OwnerLink address={event.initiated_by} />
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
                );
              })}
            </div>

            {totalPages > 1 && (
              <Pagination
                currentPage={currentPage}
                totalPages={totalPages}
                onPageChange={handlePageChange}
                itemsPerPage={ITEMS_PER_PAGE}
                totalItems={events.length}
              />
            )}
          </>
        )}
      </div>
    </div>
  );
};
