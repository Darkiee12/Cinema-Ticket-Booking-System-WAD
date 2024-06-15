import request from "../utils/request";
import User, { Account, Register, Credential } from "../models/user";
import { AxiosRequestConfig } from "axios";

const login = ({ email, password }: Credential) => {
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
  return request<{data: Account}>(options);
}

const getProfile = () => {
  const token = localStorage.getItem('token');
  const options: AxiosRequestConfig = {
    method: "GET",
    url: '/profile',
    headers: {
      "Accept": "application/json",
      "Authorization": `Bearer ${token}`
    }
  }
  return request<User>(options);
}

const register = ({email, name, password}: Register) => {
  const options: AxiosRequestConfig = {
    method: "POST",
    url: `/register`,
    headers: {
      "Content-Type": "application/json"
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