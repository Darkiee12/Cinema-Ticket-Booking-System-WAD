import request from "../utils/request";
import Movie from "../models/movie";

const getAll = (page?: number, limit?: number, cursor?: string) => {
  const filter = Object.entries({ page, limit, cursor })
    .filter(([_, value]) => value !== undefined && value !== '')
    .map(([key, value]) => `${key}=${encodeURIComponent(String(value))}`)
    .join('&');
  const options = {
    method: "GET",
    url: `/movies?${filter}`,
    headers: {
      "Content-Type": "application/json",
    },
  }
  return request<Movie>(options);
}

const getById = (id: string) => {
  const options = {
    method: "GET",
    url: `/movies/${id}`,
    headers: {
      "Content-Type": "application/json",
    },
  }
  return request<Movie>(options);
}

const MovieService = {
  getAll,
  getById
};

export default MovieService;