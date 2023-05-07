import axios from 'axios';
import qs from 'qs';
enum types {
  POST = 'POST',
  GET = 'GET'
}

enum State {
  SUCCESS = 'POST',
  ERROR = 'GET'
}

export interface ResponetFrom {
  data?: any;
  list?: any;
  msg: State;
}

const JsonParse = (res:any):Object=>{
  try {
    return JSON.parse(res)
  } catch (error) {
    return {}
  }
}

//拦截器部分
axios.interceptors.request.use(
  function (config) {
    config.headers['Authorization'] = localStorage.getItem('token');
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

/**
 * http请求
 * @param options
 * @returns
 */
const request = (options: any): Promise<ResponetFrom> => {
  const axiosOptions = Object.assign(
    {
      transformResponse: [data => data],
      headers: {
        Accept: 'application/json',
        ContentType: 'application/json;charset=UTF-8'
      },
      withCredentials: false,
      timeout: 400000,
      paramsSerializer: params => qs.stringify(params),
      validateStatus: status => status >= 200 && status < 300,
      baseURL: process.env.REACT_APP_BASE_URL
    },
    options
  );

  return new Promise((resolve:Function,reject:Function) => {
    axios(axiosOptions)
      .then((res: any) => {
        resolve({
          ...JsonParse(res.data),
          msg: State.SUCCESS
        });
      })
      .catch(error => {
        reject({
          data: error,
          msg: State.ERROR
        });
      });
  });
};

/**
 * http post请求方式
 * @param {*} url
 * @param {*} data
 */
export const httpPost = (
  url: string,
  data: any = {}
): Promise<ResponetFrom> => {
  return request({
    url,
    method: types.POST,
    data
  });
};

/**
 * http get请求方式
 * @param url
 * @param params
 * @returns
 */
export const httpGet = (
  url: string,
  params: any = {}
): Promise<ResponetFrom> => {
  return request({
    url,
    method: types.GET,
    params
  });
};

/**
 * http post formData 请求方式
 */
export const httpFormData = (
  url: string,
  params: any = {}
): Promise<ResponetFrom> => {
  const headers:any = {
    'Content-Type': 'multipart/form-data'
  };
  const formData = new FormData();
  for (const field in params) {
    if (params[field]) {
      formData.append(field, params[field]);
    }
  }
  return request({
    url,
    method: types.POST,
    data: formData,
    headers
  });
};