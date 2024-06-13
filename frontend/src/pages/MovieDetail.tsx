import { useEffect, useState } from "react";
import { Location, useLocation, useNavigate } from "react-router-dom";
import Movie from "../models/movie";
import Show from "../models/show";
import MovieService from "../services/MovieService";
import ShowService from "../services/ShowService";

const MovieDetail = () => {
  const location: Location<Movie> = useLocation();
  const [movie, setMovie] = useState<Movie>();
  const [shows, setShows] = useState<Show[]>();
  useEffect(() => {
    if(location.state){
      setMovie(location.state)
    } else{
      const id = location.pathname.split("/").pop();
      MovieService
        .getById(id!)
        .then((item) => {
          setMovie(item.data)
        })
        .catch(()=>{
          window.location.href = "/"
        })        
    }
    ShowService.getAll(movie?.imdbID)
  }, [])
  const release = new Date(movie?.released!);
  return(
    <div className="w-full px-10">
      <div className="w-full bg-[#FDF7DC] flex pt-5">
        <div className="w-[20%] pl-5">
          <img src={movie?.poster} className="w-full"/>
        </div>
        <div className="w-[80%] px-5 items-center">
          <p className="text-2xl text-[#03C04A] font-bold">{movie?.title}</p>
          <p>{movie?.plot}</p>
          <p>Rating:<b className="ml-1">{movie?.rated}</b></p>
          <p>Director:<b className="ml-1">{}</b></p>
          <p>Actor:</p>
          <p>Genre:<b className="ml-1">{}</b></p>
          <p>Premiere:<b className="ml-1">{release.getDate()+"/"+(release.getMonth()+1)+"/"+release.getFullYear()}</b></p>
          <p>Duration: <b className="ml-1">{movie?.runtime} minutes.</b></p>
          <p>Language: <b className="ml-1">{}</b></p>
          <p></p>
        </div>
      </div>
    </div>
  )
};

export default MovieDetail;