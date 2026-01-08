import { Link } from 'react-router-dom';

export const ClueLink = ({ clueId, className = "" }) => {
  return (
    <Link to={`/clues/${clueId}`} className={`link ${className}`}>
      Clue #{clueId}
    </Link>
  );
};
