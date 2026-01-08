import { Link } from 'react-router-dom';
import { formatAddress } from '../utils/web3';

export const OwnerLink = ({ address, showFull = false, className = "" }) => {
  const displayAddress = showFull ? address : formatAddress(address);

  return (
    <Link
      to={`/owners/${address}`}
      className={`link font-mono ${className}`}
      title={address}
    >
      {displayAddress}
    </Link>
  );
};
