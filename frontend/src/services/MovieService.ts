import request from '../utils/request';
import Pagination from '../utils/pagination';
import Movie from '../models/movie';

const getAll = (page?: number, limit?: number, cursor?: string) => {
  const filter = Object.entries({ page, limit, cursor })
    .filter(([_, value]) => value !== undefined && value !== '')
    .map(([key, value]) => `${key}=${encodeURIComponent(String(value))}`)
    .join('&');
  const options = {
    method: 'GET',
    url: `${filter ? `/movies?${filter}` : '/movies'}`,
    headers: {
      accept: 'application/json',
      'Content-Type': 'application/json',
    },
  };
  return request<Pagination<Movie>>(options);
};

const getById = (id: string) => {
  const options = {
    method: 'GET',
    url: `/movies/${id}`,
    headers: {
      'Content-Type': 'application/json',
    },
  };
  return request<{ data: Movie }>(options);
};

const MovieService = {
  getAll,
  getById,
};

export default MovieService;
