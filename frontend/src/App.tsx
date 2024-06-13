import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NavBar from './components/navbar'
import { ImageSlider } from './components/imageslider'

import MoviePage from './pages/Movies'

import img1 from './assets/oppenheimer.jpg'
import img2 from './assets/spiderman.jpg'
import img3 from './assets/avenger.jpg'
import img4 from './assets/titannic.jpg'
import img5 from './assets/kungfupanda.jpeg'

const image = [img1 , img2, img3, img4, img5]

function App() {
  return (
    <div className="w-full h-max bg-[#FDFCF0]">
      <NavBar />
      <div className='max-w-[1040px] h-screen mx-auto bg-[#FDF7DC]'>
        <div className='w-[1040px] h-[585px]'>
          <ImageSlider imageUrls={image}/>
        </div>
      </div>
      
      
      <Routes>
          <Route path="/movies" element={<MoviePage />} />
      </Routes>

    </div>
  )
}

export default App
