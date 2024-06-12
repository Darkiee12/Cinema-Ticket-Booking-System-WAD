import request from "../utils/request";
import Ticket from "../models/ticket";

const getByShowId = (showId: string) => {
  const options = {
    method: "GET",
    url: `/tickets?show_id=${showId}`,
    headers: {
      "Accept": "application/json"
    },
  };
  return request<Ticket>(options);
};

const put = ({seat_number, show_id}: {seat_number: number, show_id: number}) => {
  const token = localStorage.getItem("token");
  const options = {
    method: "PUT",
    url: `/tickets`,
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json",
      "Authorization": `${token}`
    },
    data: {
      seat_number,
      show_id,
    },
  };
  return request(options);
}

const getByUser = () => {
  const token = localStorage.getItem("token");
  const options = {
    method: "GET",
    url: `/tickets/user`,
    headers: {
      "Accept": "application/json",
      "Authorization": `${token}`
    },
  };
  return request<Ticket>(options);
}

const TicketService = {
  getByShowId,
  put,
  getByUser,
};
export default TicketService;