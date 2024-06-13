import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NavBar from './components/navbar'
import Homepage from './pages/Homepage'
import MoviePage from './pages/Movies'

function App() {
  return (
    <div className="w-full h-max bg-[#FDFCF0]">
      <NavBar />    
      <Routes>
          <Route path="/movies" element={<MoviePage />} />
          <Route path="/" element={<Homepage />} />
      </Routes>

    </div>
  )
}

export default App
