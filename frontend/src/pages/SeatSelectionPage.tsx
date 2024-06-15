import Show from "../models/show";
import Ticket from "../models/ticket";
import Auditorium from "../models/auditorium";
import TicketService from "../services/TicketService";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import ShowService from "../services/ShowService";
import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated'
import AuditoriumService from "../services/AuditoriumService";

const SeatSelectionPage = () => {
  const location = useLocation();
  const isAuthenticated = useIsAuthenticated()
  const [show, setShow] = useState<Show>();
  const [selectedSeat, setSelectedSeat] = useState<number[]>([]);
  const [info, setInfo] = useState<string>();
  useEffect(() => {
    const id = location.pathname.split("/").pop();
    if (id) {
      ShowService.getById(id)
        .then((res) => {
          setShow(res.data);
        })
    } else {
      setInfo("Show not found")
    }
  }, [])
  if (!isAuthenticated) {
    setTimeout(() => {
      window.location.href = '/login'
    }, 2000)
    return (
      <div className="w-full text-center">
        You need to login to view this page
      </div>
    )
  } else
    return (
      <div className="w-full px-2 lg:px-10 min-w-[40rem] overflow-x-scroll">
        <div className="w-full bg-[#FDF7DC]">
          <p className="w-full text-center py-5 font-bold text-2xl">Seat selection</p>
          <div className="w-full px-20">
            <div className="w-full border-2 border-solid rounded-lg border-black">
              <div className="px-5 py-1 w-full mt-5">
                <div className="border-2 border-solid border-black text-center text-2xl font-bold mb-10 py-5">Screen</div>
                {show ?
                  <Seats show={show} />
                  : (<div></div>)}
              </div>
            </div>
          </div>
        </div>
      </div>
    )
};
const Seats: React.FC<{ show: Show }> = ({ show }) => {
  const [numSeats, setNumSeats] = useState<number>(0);
  const [booked, setBooked] = useState<number[]>([]);

  useEffect(() => {
    console.log("Retrieving auditorium with id: ", show.auditorium.id);
    AuditoriumService.getById(show.auditorium.id)
      .then((res) => {
        setNumSeats(res.data.seats);
      })
  }, [])

  return (
    <div className="w-full grid grid-cols-10 gap-2">
      {Array.from({ length: numSeats }).map((_, index) => (
        <div key={index} className="w-full inline-block border-2 border-solid border-black p-2 text-center">
          {index + 1 < 10 ? `0${index + 1}` : index + 1}
        </div>
      ))}
    </div>
  );
};

export default SeatSelectionPage;