import { ACCESS_TOKEN_KEY } from '@/common/constant';

const isLogin = () => {
  return !!localStorage.getItem(ACCESS_TOKEN_KEY);
};

const getToken = () => {
  return localStorage.getItem(ACCESS_TOKEN_KEY);
};

const setToken = (token: string) => {
  localStorage.setItem(ACCESS_TOKEN_KEY, token);
};

const clearToken = () => {
  localStorage.removeItem(ACCESS_TOKEN_KEY);
};

export { isLogin, getToken, setToken, clearToken };
