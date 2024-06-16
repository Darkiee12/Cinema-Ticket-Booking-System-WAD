import Show from '../models/show';
import Ticket, { bookingTicket } from '../models/ticket';
import TicketService from '../services/TicketService';
import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import ShowService from '../services/ShowService';
import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated';
import Button from '../components/button';

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
      <div className="w-full px-2 lg:px-10 min-w-[40rem] overflow-x-scroll">
        <div className="w-full bg-[#FDF7DC]">
          <p className="w-full text-center py-5 font-bold text-2xl font-Montserrat">
            Seat selection
          </p>
          <div className="w-full px-20">
            <div className="w-full border-2 border-solid rounded-lg border-black">
              <div className="px-5 py-1 w-full mt-5">
                <div className="border-2 border-solid border-[#03C04A] text-center text-2xl font-bold mb-10 py-5 bg-[#03C04A] text-white font-Montserrat">
                  Screen
                </div>
                {show ? <Seats show={show} /> : <div></div>}
              </div>
            </div>
          </div>
        </div>
      </div>
    );
};
const Seats: React.FC<{ show: Show }> = ({ show }) => {
  const [numSeats, setNumSeats] = useState<
    Array<{ position: number; selected: boolean }>
  >([]);
  const [booked, setBooked] = useState<number[]>([]);

  const handleSelectSeat = (seat: number, isSelected: boolean) => {
    if (isSelected) setBooked([...booked, seat].sort((a, b) => a - b));
    else setBooked(booked.filter((item) => item !== seat));
  };

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
        console.log('Booking successful' + res);
        window.location.href = '/profile';
      })
      .catch((err) => {
        console.log(err);
      });
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
    <div>
      <div className="w-full grid grid-cols-10 gap-4">
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
      <div className="w-full flex justify-center mt-2">
        <p className="font-semibold text-2xl font-Montserrat">
          Booked:{' '}
          {booked.map((seat, index) => (
            <span key={index} className="mx-">
              {seat}{' '}
            </span>
          ))}
        </p>
      </div>
      <div className="w-full flex justify-center mt-2">
        <Button text="Book" hollow={false} onClick={() => handleBooked()} />
      </div>
    </div>
  );
};

const NotAvailableSeatUnit: React.FC<{ position: number }> = ({ position }) => {
  return (
    <div className="w-3/4 inline-block border-2 border-solid border-black text-white p-2 text-center bg-red-500">
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
          className="w-3/4 inline-block border-2 border-solid border-black p-2 text-center bg-[#03C04A] text-white"
          type="button"
          onClick={() => {
            setSelected(false), updateSelection(position, false);
          }}
        >
          {num}
        </button>
      ) : (
        <button
          className="w-3/4 inline-block border-2 border-solid border-black p-2 text-center hover:bg-[#03C04A] hover:text-white hover:border-white"
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

export default SeatSelectionPage;
