import axios from 'axios'

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})

// http 请求拦截器
service.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return error
  }
)

// http 响应拦截器
service.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    return error
  }
)

export default service
