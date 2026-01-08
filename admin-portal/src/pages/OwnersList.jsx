import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { IndexerAPI } from '../utils/api';
import { Loading } from '../components/Loading';
import { ErrorMessage } from '../components/ErrorMessage';
import { OwnerLink } from '../components/OwnerLink';

export const OwnersList = () => {
  const [owners, setOwners] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [searchTerm, setSearchTerm] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    fetchOwners();
  }, []);

  const fetchOwners = async () => {
    try {
      setLoading(true);
      setError(null);

      const api = new IndexerAPI();
      const ownerships = await api.getOwnership();

      const ownerMap = new Map();
      (ownerships || []).forEach((ownership) => {
        const address = ownership.owner_address;
        if (!ownerMap.has(address)) {
          ownerMap.set(address, {
            address,
            clueCount: 0,
            clues: [],
          });
        }
        const owner = ownerMap.get(address);
        owner.clueCount++;
        owner.clues.push(ownership.clue_id);
      });

      setOwners(Array.from(ownerMap.values()));
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  const filteredOwners = owners.filter((owner) => {
    if (!searchTerm) return true;
    return owner.address.toLowerCase().includes(searchTerm.toLowerCase());
  });

  const handleOwnerClick = (address) => {
    navigate(`/owners/${address}`);
  };

  const handleSearch = (e) => {
    setSearchTerm(e.target.value);
  };

  if (loading) return <Loading message="Loading owners..." />;
  if (error) return <ErrorMessage error={error} onRetry={fetchOwners} />;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold text-gray-900">Clue Owners</h1>
        <div className="text-sm text-gray-600">
          Total: {owners.length} owners
        </div>
      </div>

      <div className="card">
        <div className="mb-4">
          <input
            type="text"
            placeholder="Search by Ethereum Address..."
            value={searchTerm}
            onChange={handleSearch}
            className="input-field"
          />
        </div>

        {filteredOwners.length === 0 ? (
          <div className="text-center py-8 text-gray-500">
            {searchTerm ? 'No owners found matching your search.' : 'No owners found.'}
          </div>
        ) : (
          <div className="overflow-x-auto">
            <table className="table">
              <thead className="table-header">
                <tr>
                  <th className="table-header-cell">Owner Address</th>
                  <th className="table-header-cell">Clues Owned</th>
                  <th className="table-header-cell">Actions</th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {filteredOwners.map((owner) => (
                  <tr
                    key={owner.address}
                    onClick={() => handleOwnerClick(owner.address)}
                    className="table-row"
                  >
                    <td className="table-cell">
                      <OwnerLink address={owner.address} showFull />
                    </td>
                    <td className="table-cell">
                      <span className="badge-info">{owner.clueCount} clues</span>
                    </td>
                    <td className="table-cell">
                      <button
                        onClick={(e) => {
                          e.stopPropagation();
                          handleOwnerClick(owner.address);
                        }}
                        className="text-primary-600 hover:text-primary-800 text-sm font-medium"
                      >
                        View Details →
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
};
