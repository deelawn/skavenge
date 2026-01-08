import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { IndexerAPI } from '../utils/api';
import { getTotalSupply } from '../utils/web3';
import { Loading } from '../components/Loading';
import { ErrorMessage } from '../components/ErrorMessage';
import { getConfig } from '../utils/config';

export const Dashboard = () => {
  const [stats, setStats] = useState({
    totalClues: 0,
    totalOwners: 0,
    totalEvents: 0,
    recentEvents: [],
  });
  const [config, setConfig] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetchDashboardData();
  }, []);

  const fetchDashboardData = async () => {
    try {
      setLoading(true);
      setError(null);

      const api = new IndexerAPI();
      const appConfig = getConfig();

      const [clues, events, ownership, totalSupply] = await Promise.all([
        api.getAllClues(),
        api.getEvents({ limit: 10, order: 'desc' }),
        api.getOwnership(),
        getTotalSupply().catch(() => 0)
      ]);

      const ownerSet = new Set();
      (ownership || []).forEach((own) => {
        ownerSet.add(own.owner_address);
      });

      setStats({
        totalClues: clues?.length || 0,
        totalOwners: ownerSet.size,
        totalEvents: events?.length || 0,
        recentEvents: events?.slice(0, 5) || [],
        contractTotalSupply: totalSupply,
      });

      setConfig(appConfig);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  const formatTimestamp = (timestamp) => {
    if (!timestamp) return 'N/A';
    // Convert BigInt to Number before multiplying
    const timestampNum = typeof timestamp === 'bigint' ? Number(timestamp) : timestamp;
    const date = new Date(timestampNum * 1000);
    return date.toLocaleString();
  };

  if (loading) return <Loading message="Loading dashboard..." />;
  if (error) return <ErrorMessage error={error} onRetry={fetchDashboardData} />;

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div className="card cursor-pointer hover:shadow-lg transition-shadow" onClick={() => navigate('/clues')}>
          <div className="flex items-center">
            <div className="flex-shrink-0 bg-primary-100 rounded-md p-3">
              <svg className="h-6 w-6 text-primary-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Total Clues</dt>
                <dd className="text-3xl font-semibold text-gray-900">{stats.totalClues}</dd>
              </dl>
            </div>
          </div>
        </div>

        <div className="card cursor-pointer hover:shadow-lg transition-shadow" onClick={() => navigate('/owners')}>
          <div className="flex items-center">
            <div className="flex-shrink-0 bg-green-100 rounded-md p-3">
              <svg className="h-6 w-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Total Owners</dt>
                <dd className="text-3xl font-semibold text-gray-900">{stats.totalOwners}</dd>
              </dl>
            </div>
          </div>
        </div>

        <div className="card cursor-pointer hover:shadow-lg transition-shadow" onClick={() => navigate('/events')}>
          <div className="flex items-center">
            <div className="flex-shrink-0 bg-yellow-100 rounded-md p-3">
              <svg className="h-6 w-6 text-yellow-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Total Events</dt>
                <dd className="text-3xl font-semibold text-gray-900">{stats.totalEvents}</dd>
              </dl>
            </div>
          </div>
        </div>

        <div className="card">
          <div className="flex items-center">
            <div className="flex-shrink-0 bg-purple-100 rounded-md p-3">
              <svg className="h-6 w-6 text-purple-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Contract Supply</dt>
                <dd className="text-3xl font-semibold text-gray-900">{stats.contractTotalSupply?.toString() || '0'}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      {config && (
        <div className="card">
          <h2 className="text-xl font-semibold mb-4">Configuration</h2>
          <dl className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Indexer API</dt>
              <dd className="mt-1 text-sm text-gray-900 font-mono">{config.indexerApiUrl}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Gateway API</dt>
              <dd className="mt-1 text-sm text-gray-900 font-mono">{config.gatewayApiUrl}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">RPC URL</dt>
              <dd className="mt-1 text-sm text-gray-900 font-mono">{config.rpcUrl}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Contract Address</dt>
              <dd className="mt-1 text-sm text-gray-900 font-mono break-all">{config.contractAddress}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Chain ID</dt>
              <dd className="mt-1 text-sm text-gray-900">{config.chainId}</dd>
            </div>
          </dl>
        </div>
      )}

      <div className="card">
        <h2 className="text-xl font-semibold mb-4">Recent Events</h2>
        {stats.recentEvents.length === 0 ? (
          <div className="text-center py-8 text-gray-500">No recent events</div>
        ) : (
          <div className="space-y-3">
            {stats.recentEvents.map((event, index) => (
              <div key={index} className="flex justify-between items-center p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors">
                <div className="flex items-center gap-3">
                  <span className="badge-info">{event.event_type}</span>
                  {event.clue_id && (
                    <span className="text-sm text-gray-600">Clue #{event.clue_id}</span>
                  )}
                </div>
                <span className="text-sm text-gray-500">{formatTimestamp(event.timestamp)}</span>
              </div>
            ))}
          </div>
        )}
        {stats.recentEvents.length > 0 && (
          <div className="mt-4 text-center">
            <button
              onClick={() => navigate('/events')}
              className="btn-primary"
            >
              View All Events
            </button>
          </div>
        )}
      </div>
    </div>
  );
};
