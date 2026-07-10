import { get, post, put, del } from './request'
import type { User, Room, RoomType, Order, Membership, RechargeRecord, Activity, Announcement, TimeSlot, RechargePackage, Coupon, Notification, UserSetting, Feedback } from '@/types'

export const userApi = {
  login: (data: { phone?: string; password?: string; open_id?: string; nickname?: string; avatar?: string; gender?: number }) =>
    post<{ user: User; token: string }>('/user/login', data),
  
  register: (data: { phone: string; password: string; nickname?: string }) =>
    post<{ user: User; token: string }>('/user/register', data),
  
  sendSms: (phone: string) =>
    post('/user/send-sms', { phone }),
  
  getInfo: () =>
    get<User>('/user/info'),
  
  updateInfo: (data: { nickname?: string; avatar?: string; gender?: number; realname?: string }) =>
    put<User>('/user/info', data),
  
  changePassword: (data: { old_password: string; new_password: string }) =>
    post('/user/change-password', data)
}

export const roomApi = {
  getList: (params?: { type_id?: number; floor?: string; status?: number; name?: string; page?: number; page_size?: number }) =>
    get<{ data: Room[]; total: number }>('/rooms', params),
  
  getDetail: (id: number) =>
    get<Room>(`/rooms/${id}`),
  
  getTypes: () =>
    get<{ data: RoomType[]; total: number }>('/room-type'),
  
  checkAvailability: (params: { room_id: number; start_time: string; end_time: string }) =>
    get<{ available: boolean; message: string }>('/rooms/availability', params),
  
  getDateType: (date: string) =>
    get<{ date_type: string; date_type_text: string }>('/rooms/date-type', { date })
}

export const orderApi = {
  getList: (params?: { status?: number; page?: number; page_size?: number }) =>
    get<{ data: Order[]; total: number }>('/orders', params),
  
  getDetail: (id: number) =>
    get<Order>(`/orders/${id}`),
  
  create: (data: { room_id: number; start_time: string; end_time: string; remark?: string }) =>
    post<Order>('/orders', data),
  
  cancel: (id: number) =>
    put(`/orders/${id}/cancel`),
  
  confirm: (id: number) =>
    put<Order>(`/orders/${id}/confirm`),
  
  complete: (id: number) =>
    put<Order>(`/orders/${id}/complete`)
}

export const membershipApi = {
  getInfo: () =>
    get<Membership>('/membership'),
  
  recharge: (data: { amount: number }) =>
    post<Membership>('/membership/recharge', data),
  
  getRechargeRecords: () =>
    get<RechargeRecord[]>('/membership/recharges')
}

export const notificationApi = {
  getList: (params?: { page?: number; page_size?: number }) =>
    get<{ data: Notification[]; total: number }>('/notifications', params),
  
  markAllRead: () =>
    post('/notifications/mark-all-read')
}

export const activityApi = {
  getList: () =>
    get<{ data: Activity[]; total: number }>('/activities'),
  
  getDetail: (id: number) =>
    get<Activity>(`/activities/${id}`)
}

export const announcementApi = {
  getList: () =>
    get<{ data: Announcement[]; total: number }>('/announcements'),
  
  getDetail: (id: number) =>
    get<Announcement>(`/announcements/${id}`)
}

export const timeSlotApi = {
  getList: () =>
    get<{ data: TimeSlot[]; total: number }>('/time-slots')
}

export const wechatApi = {
  login: (data: { code: string; nickname?: string; avatar?: string; gender?: number }) =>
    post<{ user: User; token: string }>('/wechat/login', data),
  
  pay: (data: { order_id: number; amount: number; open_id: string }) =>
    post('/wechat/pay', data)
}

export const rechargePackageApi = {
  getList: () =>
    get<{ data: RechargePackage[]; total: number }>('/recharge-packages')
}

export const couponApi = {
  getList: () =>
    get<{ data: Coupon[]; total: number }>('/coupons'),
  
  claim: (id: number) =>
    post(`/coupons/${id}/claim`)
}

export const settingApi = {
  get: () =>
    get<UserSetting>('/setting'),
  
  update: (data: Partial<UserSetting>) =>
    put<UserSetting>('/setting', data),
  
  toggle: (key: string) =>
    post<UserSetting>('/setting/toggle', { key })
}

export const feedbackApi = {
  create: (data: { content: string; contact?: string; type?: number }) =>
    post<Feedback>('/feedback', data),
  
  getList: (params?: { page?: number; page_size?: number }) =>
    get<{ data: Feedback[]; total: number }>('/feedback', params),
  
  getDetail: (id: number) =>
    get<Feedback>(`/feedback/${id}`)
}
