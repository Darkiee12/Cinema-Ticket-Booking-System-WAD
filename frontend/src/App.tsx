import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NavBar from './components/navbar'
import Homepage from './pages/Homepage'
import MoviePage from './pages/Movies'

import img1 from './assets/oppenheimer.jpg'
import img2 from './assets/spiderman.jpg'
import img3 from './assets/avenger.jpg'
import img4 from './assets/titannic.jpg'
import img5 from './assets/kungfupanda.jpeg'
import MovieDetail from './pages/MovieDetail';

const image = [img1, img2, img3, img4, img5]

function App() {
  return (
    <div className="w-full h-max bg-[#FDFCF0]">
      <NavBar />
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/movies" element={<MoviePage />} />
          <Route path="/movies/:imdbId" element={<MovieDetail/>}/>
        </Routes>
      </Router>

    </div>
  )
}

const Home = () => {
  return (
    <div className='max-w-[1040px] h-screen mx-auto bg-[#FDF7DC]'>
      <div className='w-[1040px] h-[585px]'>
        <ImageSlider imageUrls={image} />
      </div>
    </div>
  )
}


export default App
