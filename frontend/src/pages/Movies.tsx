import {useState, useEffect} from "react";
import Movie from "../models/movie";
import MovieService from "../services/MovieService";

const MovieUnit: React.FC<{movie: Movie}> = ({movie}) => {
  return(
    <div>
      <div>
        <img src={movie.poster} alt={movie.title} />
      </div>
      <div>
        <p>{movie.rated}</p>
        <p className="truncate text-2xl font-bold">{movie.title}</p>
        <p>Genre: {movie.type}</p>
      </div>
    </div>
  )
};

const MovieList: React.FC<{movies: Movie[]}> = ({movies}) => {
  return(
    <div className="grid grid-rows-4 grid-flow-col gap">
      {movies.map((movie) => (
        <MovieUnit key={movie.imdbID} movie={movie} />
      ))}
    </div>
  )
};

const MoviePage = () => {
  const [movies, setMovies] = useState<Movie[]>([]);
  useEffect(() => {
    MovieService
    .getAll()
    .then((res) => {
      setMovies(res.data);
    })
  }, []);
  return(
    <div className="w-full px-10">
      <p className="w-full text-center">Currently premier movie</p>
      <MovieList movies={movies} />
    </div>
  )
}

export default MoviePage;