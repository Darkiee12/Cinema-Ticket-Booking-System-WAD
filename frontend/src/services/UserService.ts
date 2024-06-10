import request from "../utils/request";
import User, { Account } from "../models/user";
import Pagination from "../utils/pagination";
import { AxiosRequestConfig } from "axios";

const login = ({ email, password }: {email: string, password: string}): Promise<Pagination<Account>> => {
  const options: AxiosRequestConfig = {
    method: "POST",
    url: `/login`,
    headers: {
      "Content-Type": "application/json",
    },
    data: {
      email,
      password
    }
  }
  return request<Account>(options);
}

const getProfile = () => {
  const token = localStorage.getItem('token');
  const options: AxiosRequestConfig = {
    method: "GET",
    url: `/profile`,
    headers: {
      "Content-Type": "application/json",
      "Authorization": `${token}`
    }
  }
  return request<User>(options);
}

const register = ({email, name, password}: {email: string, name: string, password: string}) => {
  const options: AxiosRequestConfig = {
    method: "POST",
    url: `/register`,
    headers: {
      "Content-Type": "application/json",
    },
    data: {
      email,
      name,
      password
    }
  }
  return request<User>(options);
}

const UserService = {
  login,
  getProfile,
  register
};
export default UserService;