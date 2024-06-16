import { ImageSlider } from '../components/imageslider';
import img1 from '../assets/oppenheimer.jpg';
import img2 from '../assets/spiderman.jpg';
import img3 from '../assets/avenger.jpg';
import img4 from '../assets/titannic.jpg';
import img5 from '../assets/kungfupanda.jpeg';
import Movie from '../models/movie';
import MovieService from '../services/MovieService';
import { useEffect, useState } from 'react';
import { MovieUnit } from './Movies';

const image = [img1, img2, img3, img4, img5];

const Homepage = () => {
  const [movie, setMovie] = useState<Movie[]>([]);

  const today = new Date();

  useEffect(() => {
    MovieService.getAll()
      .then((res) => {
        setMovie(res.data);
        //console.log(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const currentMovies = () => {
    return (
      <div className="grid grid-cols-3 md:grid-cols-4 gap-3 md:gap-4 md:p-5 pt-2">
        {movie
          .filter((movie) => new Date(movie.released) <= today)
          .slice(0, 10)
          .map((movie) => (
            <MovieUnit key={movie.imdbID} movie={movie} />
          ))}
      </div>
    );
  };

  const upcomingMovies = () => {
    return (
      <div className="grid grid-cols-5 gap-4 p-5 pt-2">
        {movie
          .filter((movie) => new Date(movie.released) > today)
          .slice(0, 10)
          .map((movie) => (
            <MovieUnit key={movie.imdbID} movie={movie} />
          ))}
      </div>
    );
  };

  return (
    <div className="max-w-[1040px] h-full mx-auto bg-[#FDF7DC]">
      <div className="w-full h-full flex justify-center items-center">
        <div className="w-full h-full md:w-[1040px] md:h-[585px]">
          <ImageSlider imageUrls={image} />
        </div>
      </div>
      <div>
        <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">
          Currently premiere movies
        </p>
        {currentMovies()}
      </div>
      <div>
        <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">
          Upcoming movies
        </p>
        {upcomingMovies()}
      </div>
    </div>
  );
};
export default Homepage;
