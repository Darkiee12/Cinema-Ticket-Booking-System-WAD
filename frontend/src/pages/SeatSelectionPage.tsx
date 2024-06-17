import Show from '../models/show';
import Ticket, { bookingTicket } from '../models/ticket';
import TicketService from '../services/TicketService';
import { createContext, useContext, useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import ShowService from '../services/ShowService';
import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated';
import Button from '../components/button';
import MovieService from '../services/MovieService';
import Movie from '../models/movie';
import CustomDate from '../utils/date';
interface SeatContextType {
  booked: number[];
  setBooked: React.Dispatch<React.SetStateAction<number[]>>;
}
const SeatContext = createContext<SeatContextType | undefined>(undefined);

const SeatProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [booked, setBooked] = useState<number[]>([]);

  return (
    <SeatContext.Provider value={{ booked, setBooked }}>
      {children}
    </SeatContext.Provider>
  );
};
const SeatSelectionPage = () => {
  const location = useLocation();
  const isAuthenticated = useIsAuthenticated();
  const [show, setShow] = useState<Show>();
  const [info, setInfo] = useState<string>();
  useEffect(() => {
    const id = location.pathname.split('/').pop();
    if (id) {
      ShowService.getById(parseInt(id)).then((res) => {
        setShow(res.data);
      });
    } else {
      setInfo('Show not found');
    }
  }, []);
  if (!isAuthenticated) {
    setTimeout(() => {
      window.location.href = '/login';
    }, 2000);
    return (
      <div className="w-full text-center">
        You need to login to view this page
      </div>
    );
  } else
    return (
      <SeatProvider>
        <div className="w-full px-2 lg:px-10 overflow-x-scroll">
          <div className="w-full bg-[#FDF7DC]">
            <p className="w-full text-center py-5 font-bold text-2xl font-Montserrat">
              Seat selection
            </p>
            <div className="w-full lg:px-20">
              <div className="w-full border-2 border-solid rounded-lg border-black">
                <div className="px-5 py-1 w-full mt-5">
                  <div className="border-2 border-solid border-[#03C04A] text-center text-2xl font-bold mb-10 py-5 bg-[#03C04A] text-white font-Montserrat">
                    Screen
                  </div>
                  <div className="">
                    {show ? <Seats show={show} /> : <div></div>}
                  </div>
                </div>
              </div>
            </div>
            <div className="w-full mt-2 md:mt-5">
              {show && <ShowComponent show={show} />}
            </div>
          </div>
        </div>
      </SeatProvider>
    );
};
const useSeatContext = () => {
  const context = useContext(SeatContext);
  if (!context) {
    throw new Error('useSeatContext must be used within a SeatProvider');
  }
  return context;
};
const Seats: React.FC<{ show: Show }> = ({ show }) => {
  const [numSeats, setNumSeats] = useState<
    Array<{ position: number; selected: boolean }>
  >([]);
  const { booked, setBooked } = useSeatContext();

  const handleSelectSeat = (seat: number, isSelected: boolean) => {
    if (isSelected) setBooked([...booked, seat].sort((a, b) => a - b));
    else setBooked(booked.filter((item) => item !== seat));
  };

  useEffect(() => {
    TicketService.getByShowId(show.id).then((res) => {
      const seats = res.data.map((ticket: Ticket) => {
        return {
          position: ticket.seat_number,
          selected: ticket.status === 0 ? true : false,
        };
      });
      setNumSeats(seats.sort((a, b) => a.position - b.position));
    });
  }, [show.id]);

  return (
    <div className="overflow-x-auto">
      <div className="w-full grid grid-cols-10 gap-x-4 gap-y-2">
        {numSeats &&
          numSeats.map((seat) =>
            seat.selected ? (
              <NotAvailableSeatUnit
                key={seat.position}
                position={seat.position}
              />
            ) : (
              <AvailableSeatUnit
                key={seat.position}
                position={seat.position}
                updateSelection={handleSelectSeat}
              />
            ),
          )}
      </div>
    </div>
  );
};

const NotAvailableSeatUnit: React.FC<{ position: number }> = ({ position }) => {
  return (
    <div className="w-full min-w-10 inline-block border-2 border-solid border-black text-white text-center bg-red-500">
      {position < 10 ? `0${position}` : position}
    </div>
  );
};

const AvailableSeatUnit: React.FC<{
  position: number;
  updateSelection: (seat: number, isSelected: boolean) => void;
}> = ({ position, updateSelection }) => {
  const [selected, setSelected] = useState<boolean>(false);

  const num = position < 10 ? `0${position}` : position;
  return (
    <div>
      {selected ? (
        <button
          className="w-full min-w-10 inline-block border-2 border-solid border-black p-2 text-center bg-[#03C04A] text-white"
          type="button"
          onClick={() => {
            setSelected(false), updateSelection(position, false);
          }}
        >
          {num}
        </button>
      ) : (
        <button
          className="w-full min-w-10 inline-block border-2 border-solid border-black p-2 text-center hover:bg-[#03C04A] hover:text-white hover:border-white"
          type="button"
          onClick={() => {
            setSelected(true), updateSelection(position, true);
          }}
        >
          {num}
        </button>
      )}
    </div>
  );
};

const ShowComponent: React.FC<{ show: Show }> = ({ show }) => {
  const [movie, setMovie] = useState<Movie>();
  const [price, setPrice] = useState<number>(0);
  const navigate = useNavigate();
  const { booked } = useSeatContext();
  useEffect(() => {
    MovieService.getById(show.imdbID).then((res) => {
      setMovie(res.data);
    });
  }, []);
  useEffect(() => {
    setPrice(50_000 * booked.length);
  }, [booked]);
  const handleBooked = () => {
    const bookingData: bookingTicket[] = [];
    booked.forEach((seat) => {
      bookingData.push({
        seat_number: seat,
        show_id: parseInt(show.id),
      });
    });

    console.log(bookingData);
    TicketService.put(bookingData)
      .then((res) => {
        console.log(res);
        const state = {
          movie: movie,
          show: show,
          booked: booked,
          price: price
        }
        navigate("/thankyou", {state: state})
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    movie && (
      <div className="w-full flex px-2 lg:px-20 py-3">
        <div className="w-1/4">
          <img src={movie.poster} alt={movie.poster} className="w-full lg:px-10 aspect-auto" />
        </div>
        <div className="w-3/4">
          <div className="w-full px-2 flex flex-wrap justify-between gap-y-2">
            <div className='flex justify-center items-center'>
              <div className='w-10 h-10 bg-red-500'></div>
              <p className='pl-2'>Unavailable seats</p>
            </div>
            <div className='flex justify-center items-center'>
              <div className='w-10 h-10 bg-[#03C04A]'></div>
              <p className='pl-2'>Selecting seats</p>
            </div>
            <div className='flex justify-center items-center'>
              <div className='w-10 h-10 border-2 border-black'></div>
              <p className='pl-2'>Available seats</p>
            </div>
          </div>
          <div className="w-full flex flex-wrap justify-between">
            <div className="max-w-[33.33%] px-2 py-5">
              <p className="text-xl font-bold">{movie.title}</p>
              <p>{movie.rated}</p>
              <p>{movie.runtime} minutes</p>
            </div>
            <div className="px-2 py-5">
              <p className="text-xl font-bold">{show.auditorium.cinema.name}</p>
              <p>{show.auditorium.name}</p>
              <p>{CustomDate.format(show.date)}</p>
              <p>
                {show.startTime.substring(0, 5)} ~{' '}
                {show.endTime.substring(0, 5)}
              </p>
            </div>
            <div className="px-2 py-5">
            <div className='max-h-20 overflow-y-auto'>
            <p><span className='text-xl font-bold mr-1  w-full'>Seats:</span>{booked.join(", ")}</p>
            </div>
              
              <p><span className='text-xl font-bold mr-1 overflow-x-auto w-full'>Price:</span>{price.toLocaleString('vi-VN', {
                style: 'currency',
                currency: 'VND'
              })}</p>
            </div>
          </div>
          <div className="w-full flex justify-center mt-2">
            <Button hollow={false} onClick={() => handleBooked()}>Book</Button>
          </div>
        </div>
      </div>
    )
  );
};

export default SeatSelectionPage;
