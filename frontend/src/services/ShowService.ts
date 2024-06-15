import request from '../utils/request'
import Show from '../models/show'
import { AxiosRequestConfig } from 'axios'
import Pagination from '../utils/pagination'

const getAll = (imdbID?: string, date?: string) => {
  const filter = Object.entries({ imdbID, date })
    .filter(([_, value]) => value !== undefined && value !== '')
    .map(([key, value]) => `${key}=${encodeURIComponent(String(value))}`)
    .join('&')
  const options: AxiosRequestConfig = {
    method: 'GET',
    url: `/shows?${filter}`,
    headers: {
      'Content-Type': 'application/json',
    },
  }
  return request<Pagination<Show>>(options)
}

const getById = (id: string) => {
  const options: AxiosRequestConfig = {
    method: 'GET',
    url: `/shows/${id}`,
    headers: {
      'Content-Type': 'application/json',
    },
  }
  return request<{data: Show}>(options)
}

const ShowService = {
  getAll,
  getById,
}
export default ShowService