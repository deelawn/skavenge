import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { IndexerAPI } from '../utils/api';
import { formatWei } from '../utils/web3';
import { Loading } from '../components/Loading';
import { ErrorMessage } from '../components/ErrorMessage';
import { Pagination } from '../components/Pagination';
import { ClueLink } from '../components/ClueLink';
import { OwnerLink } from '../components/OwnerLink';

const ITEMS_PER_PAGE = 20;

export const CluesList = () => {
  const [clues, setClues] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [searchTerm, setSearchTerm] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    fetchClues();
  }, []);

  const fetchClues = async () => {
    try {
      setLoading(true);
      setError(null);
      const api = new IndexerAPI();
      const data = await api.getAllClues();
      setClues(data || []);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  const filteredClues = clues.filter(clue => {
    if (!searchTerm) return true;
    const term = searchTerm.toLowerCase();
    return (
      clue.clue_id?.toString().includes(term) ||
      clue.ownership?.owner_address?.toLowerCase().includes(term)
    );
  });

  const totalPages = Math.ceil(filteredClues.length / ITEMS_PER_PAGE);
  const startIndex = (currentPage - 1) * ITEMS_PER_PAGE;
  const paginatedClues = filteredClues.slice(startIndex, startIndex + ITEMS_PER_PAGE);

  const handlePageChange = (page) => {
    setCurrentPage(page);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const handleSearch = (e) => {
    setSearchTerm(e.target.value);
    setCurrentPage(1);
  };

  const handleClueClick = (clueId) => {
    navigate(`/clues/${clueId}`);
  };

  if (loading) return <Loading message="Loading clues..." />;
  if (error) return <ErrorMessage error={error} onRetry={fetchClues} />;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold text-gray-900">Clues</h1>
        <div className="text-sm text-gray-600">
          Total: {clues.length} clues
        </div>
      </div>

      <div className="card">
        <div className="mb-4">
          <input
            type="text"
            placeholder="Search by Clue ID or Owner Address..."
            value={searchTerm}
            onChange={handleSearch}
            className="input-field"
          />
        </div>

        {filteredClues.length === 0 ? (
          <div className="text-center py-8 text-gray-500">
            {searchTerm ? 'No clues found matching your search.' : 'No clues available.'}
          </div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="table">
                <thead className="table-header">
                  <tr>
                    <th className="table-header-cell">Clue ID</th>
                    <th className="table-header-cell">Owner</th>
                    <th className="table-header-cell">Point Value</th>
                    <th className="table-header-cell">Solve Reward</th>
                    <th className="table-header-cell">Status</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {paginatedClues.map((clue) => (
                    <tr
                      key={clue.clue_id}
                      onClick={() => handleClueClick(clue.clue_id)}
                      className="table-row"
                    >
                      <td className="table-cell">
                        <ClueLink clueId={clue.clue_id} />
                      </td>
                      <td className="table-cell">
                        {clue.ownership?.owner_address ? (
                          <OwnerLink address={clue.ownership.owner_address} />
                        ) : (
                          <span className="text-gray-400">Unknown</span>
                        )}
                      </td>
                      <td className="table-cell">
                        <span className="badge-info">
                          {clue.point_value || 0} pts
                        </span>
                      </td>
                      <td className="table-cell">
                        {clue.solve_reward ? `${formatWei(clue.solve_reward)} ETH` : '-'}
                      </td>
                      <td className="table-cell">
                        <div className="flex gap-2">
                          {clue.contents && (
                            <span className="badge-info">Has Content</span>
                          )}
                        </div>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>

            {totalPages > 1 && (
              <Pagination
                currentPage={currentPage}
                totalPages={totalPages}
                onPageChange={handlePageChange}
                itemsPerPage={ITEMS_PER_PAGE}
                totalItems={filteredClues.length}
              />
            )}
          </>
        )}
      </div>
    </div>
  );
};
