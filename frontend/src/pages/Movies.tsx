import { useState, useEffect } from 'react';
import MovieService from '../services/MovieService';
import Movie from '../models/movie';
import { Link } from 'react-router-dom';
import Button from '../components/button';

export const MovieUnit: React.FC<{ movie: Movie }> = ({ movie }) => {
  return (
    <div className="">
      <div className="w-full">
        <img
          src={movie.poster}
          alt={movie.title}
          className="w-full aspect-auto"
        />
      </div>
      <div>
        <p>{movie.rated}</p>
        <p className="truncate text-xl font-bold">{movie.title}</p>
        <p className="truncate text-sm">
          {movie.genres
            .map((genre) => {
              return genre.name;
            })
            .join(', ')}
        </p>
      </div>
      <Link to={{ pathname: `/movies/${movie.imdbID}` }} state={movie} className='w-full flex justify-center items-center'>
        <Button hollow={false}>Buy ticket</Button>
      </Link>
    </div>
  );
};

const MovieList: React.FC<{ movies: Movie[] }> = ({ movies }) => {
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setTimeout(() => {
      setLoading(false);
    }, 2000);
  }, []);

  return (
    <div className="grid grid-cols-3 md:grid-cols-4 gap-3 md:gap-4 md:p-5 pt-2">
      {loading ? (
        <>
          {[...Array(10)].map((_, index) => (
            <div key={index} className="skeleton-row animate-pulse w-full">
              <div className="skeleton-img bg-[#f4e8b0] h-[16rem]"></div>
              <div className="skeleton-text w-3/4 h-[1rem] mt-2 bg-[#f4e8b0]"></div>
              <div className="skeleton-text w-1/2 h-[1rem] mt-2 bg-[#f4e8b0]"></div>
              <div className="skeleton-text rounded-lg p-1 w-1/2 h-[2rem] bg-[#f4e8b0] mt-2"></div>
            </div>
          ))}
        </>
      ) : (
        movies.map((movie) => <MovieUnit key={movie.imdbID} movie={movie} />)
      )}
    </div>
  );
};

const MoviePage = () => {
  const [movies, setMovies] = useState<Movie[]>([]);
  useEffect(() => {
    MovieService.getAll().then((res) => {
      setMovies(res.data);
    });
  }, []);
  return (
    <div className="max-w-[1040px] mx-auto bg-[#FDF7DC]">
      <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">
        Currently premiere movies
      </p>
      <MovieList movies={movies} />
    </div>
  );
};

export default MoviePage;
