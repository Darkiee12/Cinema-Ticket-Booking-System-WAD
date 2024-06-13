import { useEffect, useState } from "react";
import { Location, useLocation } from "react-router-dom";
import Movie from "../models/movie";
import Show from "../models/show";
import MovieService from "../services/MovieService";
import ShowService from "../services/ShowService";

const MovieDetail = () => {
  const location: Location<Movie> = useLocation();
  const [movie, setMovie] = useState<Movie>();
  const [shows, setShows] = useState<Show[]>();
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());
  const [loading, setLoading] = useState(true);
  useEffect(() => {
    if (location.state) {
      setMovie(location.state)
    } else {
      const id = location.pathname.split("/").pop();
      MovieService
        .getById(id!)
        .then((item) => {
          setMovie(item.data)
        })
        .catch(() => {
          window.location.href = "/"
        })
    }
    ShowService.getAll(movie?.imdbID)
  }, [])
  const release = new Date(movie?.released!);
  const handleDateChange = (date: Date) => {
    setSelectedDate(date);
  };
  return (
    <div className="w-full px-10">
      <div className="w-full bg-[#FDF7DC] flex pt-5">
        <div className="w-[20%] pl-5">
          <img src={movie?.poster} className="w-full" />
        </div>
        <div className="w-[80%] px-5 items-center md:text-base text-xs">
          <p className="md:text-2xl text-[#03C04A] font-bold">{movie?.title}</p>
          <p>{movie?.plot}</p>
          <p>Rating:<b className="ml-1">{movie?.rated}</b></p>
          <p>Director:<b className="ml-1">{ }</b></p>
          <p>Actor:</p>
          <p>Genre:<b className="ml-1">{ }</b></p>
          <p>Premiere:<b className="ml-1">{release.getDate() + "/" + (release.getMonth() + 1) + "/" + release.getFullYear()}</b></p>
          <p>Duration: <b className="ml-1">{movie?.runtime} minutes.</b></p>
          <p>Language: <b className="ml-1">{ }</b></p>
          <p>You are choosing: {selectedDate.getDate()}</p>
        </div>
      </div>
      <div className="w-full mt-5">
        <DatePicker onDateChange={handleDateChange}/>
      </div>
    </div>
  )
};

const DatePicker: React.FC<{onDateChange: (date: Date) => void}> = ({ onDateChange }) => {
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());
  const dates = Array.from({ length: 7 }, (_, i) => {
    const date = new Date();
    date.setDate(date.getDate() + i);
    return date;
  });
  const handleDateClick = (date: Date) => {
    setSelectedDate(date);
    onDateChange(date);
  };
  return (
    <div className="w-full flex flex-wrap justify-between gap-2 px-5">
      {dates.map((date) => (
        <button
          key={date.toDateString()}
          aria-label={`Select date ${date.toLocaleDateString()}`}
          className={`border-2 px-5 py-2 text-lg font-semibold rounded-[10px]  ${
            date.toDateString() === selectedDate.toDateString() ? 'bg-green-500 text-white' : 'border-[#03C04A] hover:bg-green-300'
          }`}
          onClick={() => handleDateClick(date) }  
        >
          <p>{date.getDate() + "/" + ((date.getMonth() < 9) ? "0" + (date.getMonth() + 1) : (date.getMonth() + 1)) + "/" + date.getFullYear()}</p>
        </button>
      ))}
    </div>
  )
}

const ShowComponent: React.FC<{ shows: Show[] }> = ({ shows }) => {
  return(
    <div></div>
  )
};

const CinemaComponent: React.FC<{}> = () => {
  return(
    <div></div>
  )
}
export default MovieDetail;