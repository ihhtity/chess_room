import request from '@/utils/request'

export interface User {
  id: number
  openid: string
  phone: string
  nickname: string
  realname?: string
  avatar: string
  gender: number
}

export interface RoomType {
  id: number
  name: string
  description: string
  base_price: number
  max_people: number
}

export interface Room {
  id: number
  name: string
  type_id: number
  type: RoomType
  status: number
  equipment: string
  images: string
  description: string
  floor: string
}

export interface Order {
  id: number
  order_no: string
  user_id: number
  room_id: number
  room: Room
  start_time: string
  end_time: string
  duration: number
  status: number
  total_amount: number
  paid_amount: number
  created_at: string
  remark?: string
}

export interface Membership {
  id: number
  user_id: number
  user?: User
  level: number
  points: number
  balance: number
  total_consumed: number
  total_recharged: number
  discount?: number
}

export interface Activity {
  id: number
  name: string
  description: string
  image: string
  discount: number
  valid_from: string
  valid_to: string
  status: number
}

export interface Announcement {
  id: number
  title: string
  content: string
  type: number
  created_at: string
}

export interface RechargePackage {
  id: number
  name: string
  amount: number
  gift_amount: number
  gift_points: number
  description: string
  sort_order: number
}

export interface Review {
  id: number
  order_id: number
  user_id: number
  user?: User
  room_id: number
  rating: number
  content: string
  images: string
  created_at: string
}

export interface LoginResult {
  token: string
  user: User
}

export const userApi = {
  login: (data: { code?: string; openid?: string; phone?: string; password?: string; nickname?: string; avatar?: string; gender?: number }) =>
    request.post<LoginResult>('/user/login', data),
  
  getUserInfo: () =>
    request.get<User>('/user/info'),
  
  updateUserInfo: (data: { nickname?: string; avatar?: string; gender?: number; realname?: string }) =>
    request.put<User>('/user/info', data),
  
  changePassword: (data: { old_password: string; new_password: string }) =>
    request.post<any>('/user/change-password', data)
}

export const roomApi = {
  getRoomList: (params?: { status?: number; type_id?: number }) =>
    request.get<Room[]>('/rooms', { params }),
  
  getRoomDetail: (id: number) =>
    request.get<Room>(`/rooms/${id}`),
  
  getRoomTypeList: () =>
    request.get<RoomType[]>('/room-type')
}

export const orderApi = {
  createOrder: (data: { room_id: number; start_time: string; end_time: string; remark?: string }) =>
    request.post<Order>('/orders', data),
  
  getOrderList: (params?: { room_id?: number; status?: number; start_time?: string; end_time?: string }) =>
    request.get<Order[]>('/orders', { params }),
  
  getOrderDetail: (id: number) =>
    request.get<Order>(`/orders/${id}`),
  
  cancelOrder: (id: number) =>
    request.put<any>(`/orders/${id}/cancel`)
}

export const paymentApi = {
  createPayment: (data: { order_no: string; payment_type: number }) =>
    request.post<any>('/payments', data),
  
  getPaymentByOrderNo: (orderNo: string) =>
    request.get<any>(`/payments/${orderNo}`)
}

export const membershipApi = {
  getMembership: () =>
    request.get<Membership>('/membership'),
  
  recharge: (data: { amount: number }) =>
    request.post<Membership>('/recharge', data),
  
  getRechargeRecords: () =>
    request.get<any[]>('/membership/recharges'),
  
  getRechargePackages: () =>
    request.get<RechargePackage[]>('/recharge-packages')
}

export const activityApi = {
  getActivityList: () =>
    request.get<Activity[]>('/activities'),
  
  getActivityDetail: (id: number) =>
    request.get<Activity>(`/activities/${id}`)
}

export const announcementApi = {
  getAnnouncementList: () =>
    request.get<Announcement[]>('/announcements'),
  
  getAnnouncementDetail: (id: number) =>
    request.get<Announcement>(`/announcements/${id}`)
}

export const reviewApi = {
  getReviewList: (params?: { room_id?: number; rating?: number }) =>
    request.get<Review[]>('/reviews', { params }),
  
  createReview: (data: { order_id: number; rating: number; content: string; images?: string }) =>
    request.post<Review>('/reviews', data)
}