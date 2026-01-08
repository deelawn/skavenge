import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Navigation } from './components/Navigation';
import { Dashboard } from './pages/Dashboard';
import { CluesList } from './pages/CluesList';
import { ClueDetails } from './pages/ClueDetails';
import { EventsList } from './pages/EventsList';
import { OwnersList } from './pages/OwnersList';
import { OwnerDetails } from './pages/OwnerDetails';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <Navigation />
        <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/clues" element={<CluesList />} />
            <Route path="/clues/:clueId" element={<ClueDetails />} />
            <Route path="/events" element={<EventsList />} />
            <Route path="/owners" element={<OwnersList />} />
            <Route path="/owners/:address" element={<OwnerDetails />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
