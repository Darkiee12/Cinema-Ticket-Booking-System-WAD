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
  useEffect(() => {
    if (location.state && (location.state as Show).auditorium) {
      console.log("Show from prev page: ",location.state)
      const cache = location.state as Show;
      setShow(cache);
    } else {
      const id = location.pathname.split('/').pop();
      if (id) {
        ShowService.getById(id)
          .then((show) => {
            console.log("Show from API: ",show)
            setShow(show);
          })
          .catch((error) => {
            console.error('Error fetching show by ID:', error);
          });
      } else {
        window.location.href = '/'; 
      }
    }
  }, []);
   if (!isAuthenticated){
    setTimeout(() => {
      window.location.href = "/login"
    }, 2000);
    return(
      <div className="w-full text-center">You need to login to view this page</div>
    )
  } else
  return (
    <div className="w-full px-10">
      <div className="w-full bg-[#FDF7DC]">
        <p className="w-full text-center py-5 font-bold text-2xl">Seat selection</p>
        <div className="w-full px-20">
          <div className="w-full border-2 border-solid rounded-lg border-black h-screen">
            <div className="px-48 py-5 w-full">
              <div className="border-2 border-solid border-black text-center font-bold mb-10">Screen</div>
              {show && <Seats show={show} />}
            </div>

          </div>

        </div>
      </div>
    </div>
  )
};
const Seats: React.FC<{ show: Show }> = ({ show }) => {
  const [numSeats, setNumSeats] = useState<number>(0);
  useEffect(() => {
    AuditoriumService
      .getById(show.auditorium.id)
      .then((res) => {
        setNumSeats(res.seats)
      })
  })
  return (
    <div className="w-full flex flex-wrap">
      {Array.from({ length: numSeats }).map((_, index) => {
        return (
          <div key={index} className="inline-block m-1 w-8 h-8 border-2 border-solid border-black">{index + 1}</div>
        )
      })}
    </div>
  )
}
export default SeatSelectionPage;