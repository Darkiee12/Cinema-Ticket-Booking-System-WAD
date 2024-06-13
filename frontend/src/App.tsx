
import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NavBar from './components/navbar'
import MoviePage from './pages/Movies'

function App() {
  return (
    <div className="w-full h-screen">
      <NavBar />
      <Router>
        <Routes>
          <Route path="/movies" element={<MoviePage />} />
        </Routes>
      </Router>

    </div>
  )
}

export default App
