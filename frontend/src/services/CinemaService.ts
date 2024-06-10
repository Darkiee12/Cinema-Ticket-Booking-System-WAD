import request from "../utils/request";
import Cinema from "../models/cinema";

const getAll = (page?: number, limit?: number, cursor?: string, owner_id?: string) => {
  const filter = Object.entries({ page, limit, cursor, owner_id })
    .filter(([_, value]) => value !== undefined && value !== '')
    .map(([key, value]) => `${key}=${encodeURIComponent(String(value))}`)
    .join('&');
  const options = {
    method: "GET",
    url: `/cinemas?${filter}`,
    headers: {
      "Content-Type": "application/json",
    },
  }
  return request<Cinema>(options);
}

const getByName = (name: string) => {
  const options = {
    method: "GET",
    url: `/cinemas/name/${name}`,
    headers: {
      "Content-Type": "application/json",
    },
  }
  return request<Cinema>(options);
}

const getById = (id: string) => {
  const options = {
    method: "GET",
    url: `/cinemas/${id}`,
    headers: {
      "Content-Type": "application/json",
    },
  }
  return request<Cinema>(options);
}

const CinemaService = {
  getAll,
  getByName,
  getById
};
export default CinemaService;