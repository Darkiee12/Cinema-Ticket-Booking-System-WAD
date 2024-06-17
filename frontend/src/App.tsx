import './App.css';
import { Routes, Route } from 'react-router-dom';
import NavBar, { NavBarMobile } from './components/navbar';
import Homepage from './pages/Homepage';
import MoviePage from './pages/Movies';
import MovieDetail from './pages/MovieDetail';
import CredentialPage from './pages/CredentialPage';
import SeatSelectionPage from './pages/SeatSelectionPage';
import Profile from './pages/Profile';
import { useState } from 'react';
import Cinemas from './pages/Cinemas';
import CinemaDetail from './pages/CinemaDetail';
import {BrowserView, MobileView} from 'react-device-detect';
import PostTicketBookingPage from './pages/PostTicketBookingPage';
function App() {
  const [username, setUsername] = useState<string>('');
  const handleUsernameUpdate = (newUsername: string) => {
    setUsername(newUsername);
  };
  return (
    <div className="w-full h-full bg-[#FDFCF0]">
      <BrowserView>
        <NavBar username={username} />
      </BrowserView>
      <MobileView>
        <NavBarMobile username={username}/>
      </MobileView>
      <Routes>
        <Route path="/login" element={<CredentialPage />} />
        <Route path="/movies" element={<MoviePage />} />
        <Route path="/movies/:imdbId" element={<MovieDetail />} />
        <Route path="/show/:showId" element={<SeatSelectionPage />} />
        <Route
          path="/profile"
          element={<Profile onUsernameUpdate={handleUsernameUpdate} />}
        />
        <Route path="/cinemas" element={<Cinemas />} />
        <Route path="/cinemas/:cinemaId" element={<CinemaDetail />} />
        <Route path="/" element={<Homepage />} />
        <Route path="/thankyou" element={<PostTicketBookingPage />} />
      </Routes>
    </div>
  );
}
export default App;
