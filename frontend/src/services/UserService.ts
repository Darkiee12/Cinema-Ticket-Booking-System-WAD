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

const update = ({date_of_birth, gender, name, phone}: {date_of_birth:string, gender:string, name:string, phone:string}) => {
  const options: AxiosRequestConfig = {
    method: "PUT",
    url: '/profile',
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${getCookie("_auth")}`
    },
    data: {
      date_of_birth,
      gender,
      name,
      phone
    } 
  }
  return request<{data: boolean, filter:any, paging:any}>(options);
}

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

const UserService = {
  login,
  getProfile,
  register,
  update
};
export default UserService;

