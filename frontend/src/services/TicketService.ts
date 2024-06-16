import request from "../utils/request";
import Ticket from "../models/ticket";
import Pagination from "../utils/pagination";

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
  const token = localStorage.getItem("token");
  const options = {
    method: "PUT",
    url: `/tickets`,
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json",
      "Authorization": `Bearer ${token}`
    },
    data: seats
  }
  return request(options);
}

const getByUser = () => {
  const token = localStorage.getItem("token");
  const options = {
    method: "GET",
    url: `/tickets/user`,
    headers: {
      "Accept": "application/json",
      "Authorization": `Bearer ${token}`
    },
  };
  return request<Pagination<Ticket>>(options);
}

const TicketService = {
  getByShowId,
  put,
  getByUser,
};
export default TicketService;