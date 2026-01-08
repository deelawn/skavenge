const eventColors = {
  ClueMinted: 'bg-green-100 text-green-800',
  ClueAttempted: 'bg-yellow-100 text-yellow-800',
  ClueSolved: 'bg-purple-100 text-purple-800',
  SalePriceSet: 'bg-blue-100 text-blue-800',
  SalePriceRemoved: 'bg-gray-100 text-gray-800',
  TransferInitiated: 'bg-indigo-100 text-indigo-800',
  ProofProvided: 'bg-cyan-100 text-cyan-800',
  ProofVerified: 'bg-teal-100 text-teal-800',
  TransferCompleted: 'bg-green-100 text-green-800',
  TransferCancelled: 'bg-red-100 text-red-800',
  AuthorizedMinterUpdated: 'bg-orange-100 text-orange-800',
  Transfer: 'bg-blue-100 text-blue-800',
  Approval: 'bg-gray-100 text-gray-800',
  ApprovalForAll: 'bg-gray-100 text-gray-800',
};

export const EventBadge = ({ eventType }) => {
  const colorClass = eventColors[eventType] || 'bg-gray-100 text-gray-800';

  return (
    <span className={`badge ${colorClass}`}>
      {eventType}
    </span>
  );
};
