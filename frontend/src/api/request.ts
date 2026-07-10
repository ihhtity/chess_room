interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

interface RequestOptions {
  url: string
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE'
  data?: any
  header?: Record<string, string>
}

const BASE_URL = '/api'

async function request<T = any>(options: RequestOptions): Promise<T> {
  const token = uni.getStorageSync('user_token')
  const header: Record<string, string> = {
    'Content-Type': 'application/json',
    ...options.header
  }
  
  if (token) {
    header.Authorization = `Bearer ${token}`
  }

  return new Promise((resolve, reject) => {
    uni.request({
      url: BASE_URL + options.url,
      method: options.method || 'GET',
      data: options.data,
      header,
      timeout: 15000,
      success: (res) => {
        const response = res.data as ApiResponse<T>
        if (response.code === 200) {
          resolve(response.data)
        } else {
          if (response.code === 401) {
            uni.removeStorageSync('user_token')
            uni.navigateTo({ url: '/pages/user/login' })
          }
          uni.showToast({
            title: response.message || '请求失败',
            icon: 'none'
          })
          reject(response)
        }
      },
      fail: (err) => {
        uni.showToast({
          title: '网络请求失败',
          icon: 'none'
        })
        reject(err)
      }
    })
  })
}

export const get = <T = any>(url: string, params?: any) => {
  return request<T>({ url, method: 'GET', data: params })
}

export const post = <T = any>(url: string, data?: any) => {
  return request<T>({ url, method: 'POST', data })
}

export const put = <T = any>(url: string, data?: any) => {
  return request<T>({ url, method: 'PUT', data })
}

export const del = <T = any>(url: string, data?: any) => {
  return request<T>({ url, method: 'DELETE', data })
}

export default {
  get,
  post,
  put,
  delete: del
}
