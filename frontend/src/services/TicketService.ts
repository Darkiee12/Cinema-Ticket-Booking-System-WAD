import request from "../utils/request";
import Ticket, { bookingTicket } from "../models/ticket";
import Pagination from "../utils/pagination";
import { getCookie } from "./UserService";

const getByShowId = (showId: string) => {
  const options = {
    method: "GET",
    url: `/tickets?show_id=${showId}`,
    headers: {
      "Accept": "application/json"
    },
  };
  return request<Pagination<Ticket>>(options);
};

const put = (seats: Array<{ seat_number: number, show_id: number }>) => {
  const options = {
    method: "PUT",
    url: "/tickets",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json",
      "Authorization": `Bearer ${getCookie("_auth")}`
    },
    data: seats
  }
  return request(options);
}

const getByUser = () => {
  const options = {
    method: "GET",
    url: `/tickets/user`,
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${getCookie("_auth")}`
    },
  };
  return request(options);
}

const TicketService = {
  getByShowId,
  put,
  getByUser,
};
export default TicketService;