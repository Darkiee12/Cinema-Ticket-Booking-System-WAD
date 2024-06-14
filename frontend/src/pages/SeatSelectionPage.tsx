import Show from "../models/show";
import Ticket from "../models/ticket";
import Auditorium from "../models/auditorium";
import TicketService from "../services/TicketService";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import ShowService from "../services/ShowService";

const SeatSelectionPage = () => {
  const location = useLocation();
  const [show, setShow] = useState<Show>();
  const [selectedSeat, setSelectedSeat] = useState<number[]>([]);
  useEffect(() => {
    if(location.state){
      setShow(location.state);
    } else{
      const id = location.pathname.split('/').pop()
      if(id){
        ShowService.getById(id)
      } else{
        window.location.href = '/'
      }
      
    }
  },[]);
  
  return(
    <div>
      Landing page of seat selection
    </div>
  )
};

export default SeatSelectionPage;