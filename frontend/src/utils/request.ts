import axios from 'axios';
import { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios';
import Pagination from './pagination';

export interface Error {
  status_code: number;
  message: string;
  log: string;
  error_key: string;
}

const client = axios.create({
  baseURL: import.meta.env.VITE_API_URL + '/v1',
});

const request = function <T = any>(options: AxiosRequestConfig) {
  const onSuccess = function (response: AxiosResponse<T>) {
    console.debug('Request Successful!', response);
    return response.data;
  };

  const onError = function (error: AxiosError<Error>) {
    console.error('Request Failed:', error.config);

    if (error.response) {
      console.error('Status:', error.response.status);
      console.error('Data:', error.response.data);
      console.error('Headers:', error.response.headers);
    } else {
      console.error('Error Message:', error.message);
    }

    return Promise.reject(error.response || error.message);
  };

  return client(options).then(onSuccess).catch(onError);
};

export default request;
