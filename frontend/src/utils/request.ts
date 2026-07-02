import axios, { AxiosInstance, AxiosRequestConfig } from 'axios'
import { showToast } from '@/components/Toast'

const baseURL = '/api'

const instance = axios.create({
  baseURL,
  timeout: 10000
})

instance.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  (response) => {
    const res = response.data
    if (res.code === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      setTimeout(() => {
        window.location.href = '/login'
      }, 1000)
      showToast({ message: '登录已失效，请重新登录', type: 'error' })
      return Promise.reject(new Error('登录失效'))
    }
    if (res.code !== 200) {
      showToast({ message: res.message || '请求失败', type: 'error' })
      return Promise.reject(new Error(res.message))
    }
    return res.data
  },
  (error) => {
    if (error.response && error.response.data) {
      showToast({ message: error.response.data.message || '请求失败', type: 'error' })
    } else {
      showToast({ message: '网络错误', type: 'error' })
    }
    return Promise.reject(error)
  }
)

interface CustomAxiosInstance extends AxiosInstance {
  get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T>
  post<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T>
  put<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T>
  delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<T>
}

const request = instance as CustomAxiosInstance

export default request
