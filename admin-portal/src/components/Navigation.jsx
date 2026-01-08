import { NavLink } from 'react-router-dom';

export const Navigation = () => {
  const navLinkClass = ({ isActive }) =>
    `px-4 py-2 rounded-lg font-medium transition-colors ${
      isActive
        ? 'bg-primary-600 text-white'
        : 'text-gray-700 hover:bg-gray-100'
    }`;

  return (
    <nav className="bg-white shadow-md mb-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center">
            <h1 className="text-xl font-bold text-gray-900">
              Skavenge Admin Portal
            </h1>
          </div>
          <div className="flex space-x-4">
            <NavLink to="/" className={navLinkClass}>
              Dashboard
            </NavLink>
            <NavLink to="/clues" className={navLinkClass}>
              Clues
            </NavLink>
            <NavLink to="/events" className={navLinkClass}>
              Events
            </NavLink>
            <NavLink to="/owners" className={navLinkClass}>
              Owners
            </NavLink>
          </div>
        </div>
      </div>
    </nav>
  );
};
