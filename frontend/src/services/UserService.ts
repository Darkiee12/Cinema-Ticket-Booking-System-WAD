import request from "../utils/request";
import User, { Account, Register, Credential } from "../models/User";
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

const getProfile = (token?: string) => {
  const options: AxiosRequestConfig = {
    method: "GET",
    url: '/profile',
    headers: {
      "Accept": "application/json",
      "Authorization": `Bearer ${token || getCookie("_auth")}`
    }
  }
  return request<{data: User}>(options);
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
  return request<string>(options);
}

const UserService = {
  login,
  getProfile,
  register
};
export default UserService;

function getCookie(name: string): string | undefined {
  const cookies = document.cookie.split(';');
  for (let cookie of cookies) {
      const [cookieName, cookieValue] = cookie.split('=').map(c => c.trim());
      if (cookieName === name) {
          return decodeURIComponent(cookieValue);
      }
  }
  return undefined;
}