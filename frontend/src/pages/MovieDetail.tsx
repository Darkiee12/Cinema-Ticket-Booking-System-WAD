import { useEffect, useState } from 'react';
import { Link, Location, useLocation } from 'react-router-dom';
import Movie from '../models/movie';
import Show, { CinemaAuditorium, MovieShow } from '../models/show';
import MovieService from '../services/MovieService';
import ShowService from '../services/ShowService';

const MovieDetail = () => {
  const location: Location<Movie> = useLocation();
  const [movie, setMovie] = useState<Movie>();
  const [shows, setShows] = useState<Show[]>();
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());
  const [, setLoading] = useState(true);
  useEffect(() => {
    if (location.state) {
      setMovie(location.state);
    } else {
      const id = location.pathname.split('/').pop();
      if (id) {
        MovieService.getById(id)
          .then((item) => {
            setMovie(item.data);
          })
          .catch(() => {
            window.location.href = '/';
          });
      } else {
        window.location.href = '/';
      }
    }
  }, []);
  useEffect(() => {
    ShowService.getAll(
      movie?.imdbID,
      `"${selectedDate.toISOString().substring(0, 10)}"`,
    )
      .then((res) => {
        setShows(res.data);
        setLoading(false);
      })
      .catch(() => {
        setShows(undefined);
      });
  }, [selectedDate]);

  const handleDateChange = (date: Date) => {
    setSelectedDate(date);
  };
  return (
    <div className="w-full md:px-10">
      <div className="w-full bg-[#FDF7DC] pt-5">
        <div className="w-full">
          <MovieSection movie={movie!} />
        </div>
        <div className="w-full mt-5">
          <DatePicker onDateChange={handleDateChange} />
        </div>
        <div className="w-full px-5">
          {!shows ? (
            <div>Loading....</div>
          ) : (
            <div>
              {shows.length === 0 ? (
                <div className="w-full text-center italic pt-2">
                  No shows are available on this date
                </div>
              ) : (
                <div className="">
                  <ShowComponent shows={shows!} selectedDate={selectedDate} />
                </div>
              )}
            </div>
          )}
        </div>
        <div className="w-full py-5 text-center italic">The end</div>
      </div>
    </div>
  );
};

const MovieSection: React.FC<{ movie: Movie }> = ({ movie }) => {
  const release = new Date(movie?.released!);
  return (
    <div className="flex h-full">
      <div className="w-[30%] md:w-[20%] pl-2 lg:pl-5 flex items-center justify-center">
        <img src={movie?.poster} className="w-full h-full" />
      </div>
      <div className="w-[70%] md:w-[80%] px-5 items-center md:text-base text-xs max-h-[15rem] overflow-y-scroll lg:overflow-y-visible">
        <p className="md:text-2xl text-[#03C04A] font-bold">{movie?.title}</p>
        <p>{movie?.plot}</p>
        <p>
          Rating:<b className="ml-1">{movie?.rated}</b>
        </p>
        <p>
          Director:<b className="ml-1">{}</b>
        </p>
        <p>Actor:</p>
        <p>
          Genre:
          <b className="ml-1">
            {movie?.genres
              ? movie.genres.map((genre) => genre.name).join(', ')
              : ''}
          </b>
        </p>
        <p>
          Premiere:
          <b className="ml-1">
            {release.getDate() +
              '/' +
              (release.getMonth() + 1) +
              '/' +
              release.getFullYear()}
          </b>
        </p>
        <p>
          Duration: <b className="ml-1">{movie?.runtime} minutes.</b>
        </p>
        <p>
          Language: <b className="ml-1">{}</b>
        </p>
      </div>
    </div>
  );
};

const DatePicker: React.FC<{ onDateChange: (date: Date) => void }> = ({
  onDateChange,
}) => {
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
    <div className="w-full flex flex-wrap justify-around gap-2 px-5">
      {dates.map((date) => (
        <button
          key={date.toDateString()}
          aria-label={`Select date ${date.toLocaleDateString()}`}
          className={`border-2 px-2 py-1 md:px-5 md:py-2 text-base md:text-lg font-semibold rounded-[10px]  ${
            date.toDateString() === selectedDate.toDateString()
              ? 'bg-green-500 text-white'
              : 'border-[#03C04A] hover:bg-green-300'
          }`}
          onClick={() => handleDateClick(date)}
        >
          <p>
            {date.getDate() +
              '/' +
              (date.getMonth() < 9
                ? '0' + (date.getMonth() + 1)
                : date.getMonth() + 1) +
              '/' +
              date.getFullYear()}
          </p>
        </button>
      ))}
    </div>
  );
};

const ShowComponent: React.FC<{ shows: Show[]; selectedDate: Date }> = ({
  shows,
  selectedDate,
}) => {
  const cinemas = new MovieShow(shows);
  const date = selectedDate.toISOString().substring(0, 10);
  return (
    <div className="w-full flex flex-wrap gap-y-3 mt-5">
      {cinemas.getCinemas(date).map((cinema) => {
        return <CinemaComponent key={cinema.id} cinema={cinema} />;
      })}
    </div>
  );
};

const CinemaComponent: React.FC<{ cinema: CinemaAuditorium }> = ({
  cinema,
}) => {
  const shows = cinema.auditoriums
    .flatMap((auditorium) => auditorium.shows)
    .sort((a, b) => a.startTime.localeCompare(b.startTime));
  return (
    <div className="w-full border-2 border-black rounded-lg p-2">
      <b>{cinema.name}</b>
      <div className="w-full flex flex-wrap gap-5">
        {shows.map((show, index) => {
          return <ShowUnit key={index} show={show} />;
        })}
      </div>
    </div>
  );
};

const ShowUnit: React.FC<{ show: Show }> = ({ show }) => {
  const display = show.startTime.substring(0, 5);
  return (
    <Link
      to={'/show/' + show.id}
      state={{ show }}
      className="border-2 border-[#03C04A] rounded-lg p-2 hover:bg-[#03C04A] hover:text-white"
    >
      <b>{display}</b>
    </Link>
  );
};

export default MovieDetail;
