import axios, { AxiosInstance, AxiosRequestConfig, AxiosError } from 'axios'
import { message } from 'antd'
import { ApiResponse } from '@/types'

declare module 'axios' {
  interface InternalAxiosRequestConfig {
    __retryCount?: number
  }
}

const MAX_RETRIES = 3
const RETRY_DELAY = 1000

const request = axios.create({
  baseURL: '/api',
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json'
  }
}) as AxiosInstance & {
  get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T>
  post<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T>
  put<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T>
  delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<T>
}

request.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('admin_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

const shouldRetry = (error: AxiosError): boolean => {
  if (error.response) {
    const status = error.response.status
    return status >= 500 || status === 408
  }
  return !error.config?.cancelToken && !error.message.includes('cancel')
}

const retryRequest = async (config: AxiosRequestConfig, retryCount: number): Promise<any> => {
  try {
    const response = await axios.request(config)
    return response
  } catch (error) {
    if (retryCount < MAX_RETRIES && shouldRetry(error as AxiosError)) {
      await new Promise(resolve => setTimeout(resolve, RETRY_DELAY * Math.pow(2, retryCount)))
      return retryRequest(config, retryCount + 1)
    }
    throw error
  }
}

request.interceptors.response.use(
  (response) => {
    const res = response.data as ApiResponse
    if (res.code !== 200) {
      if (res.code === 401) {
        message.error('登录已过期，请重新登录')
        localStorage.removeItem('admin_token')
        window.location.href = '/login'
      } else {
        message.error(res.message || '请求失败')
      }
      return Promise.reject(res)
    }
    return res.data
  },
  async (error: AxiosError) => {
    const config = error.config
    if (config && shouldRetry(error)) {
      const retryCount = config.__retryCount || 0
      if (retryCount < MAX_RETRIES) {
        config.__retryCount = retryCount + 1
        try {
          const response = await retryRequest(config, retryCount)
          const res = response.data as ApiResponse
          if (res.code !== 200) {
            message.error(res.message || '请求失败')
            return Promise.reject(res)
          }
          return res.data
        } catch (retryError) {
          message.error('网络异常，请稍后重试')
          return Promise.reject(retryError)
        }
      }
    }

    if (error.response) {
      const status = error.response.status
      switch (status) {
        case 401:
          message.error('登录已过期，请重新登录')
          localStorage.removeItem('admin_token')
          window.location.href = '/login'
          break
        case 403:
          message.error('权限不足，无法操作')
          break
        case 404:
          message.error('请求的资源不存在')
          break
        case 500:
          message.error('服务器内部错误')
          break
        default:
          message.error(error.message || '请求失败')
      }
    } else if (error.request) {
      message.error('网络连接失败，请检查网络')
    } else {
      message.error('请求配置错误')
    }

    return Promise.reject(error)
  }
)

export default request