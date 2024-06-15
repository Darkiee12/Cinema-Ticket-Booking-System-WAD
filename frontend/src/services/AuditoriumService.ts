import { AxiosRequestConfig } from "axios"
import request from "../utils/request";
import Auditorium from "../models/auditorium";

const getById = (id: string, page?: number, limit?: number, cursor?: string) => {
  const filter = Object.entries({ page, limit, cursor })
    .filter(([_, value]) => value !== undefined && value !== '')
    .map(([key, value]) => `${key}=${encodeURIComponent(String(value))}`)
    .join('&');
  const options: AxiosRequestConfig = {
    method: "GET",
    url: `/auditoriums/${id}${filter}`,
    headers: {
      "Content-Type": "application/json",
    },
  }
  return request<{data:Auditorium}>(options);
}

const getByName = (name: string, page?: number, limit?: number, cursor?: string) => {
  const filter = Object.entries({ page, limit, cursor })
    .filter(([_, value]) => value !== undefined && value !== '')
    .map(([key, value]) => `${key}=${encodeURIComponent(String(value))}`)
    .join('&');
  const options: AxiosRequestConfig = {
    method: "GET",
    url: `/cinemas/${name}/auditoriums?${filter}`,
    headers: {
      "Content-Type": "application/json",
    },
  }
  return request<Auditorium>(options);
}



const AuditoriumService = {
  getById,
  getByName
};
export default AuditoriumService;