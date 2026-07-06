import request from '@/utils/request'
import { Admin, Room, RoomType, Order, Membership, User, Activity, Announcement, RechargePackage, Payment, Review, TimeSlot, Holiday, RechargeRecord, Notification, OperationLog, AdminRole, Permission } from '@/types'

// 定义类型转换函数
const cast = <T>(res: any): T => res as unknown as T

// 管理员接口
export const adminApi = {
  // 登录
  login: async (data: { username: string; password: string }) => {
    const res = await request.post('/admin/login', data)
    return cast<{ admin: Admin; token: string }>(res)
  },
  // 获取管理员信息
  getProfile: async () => {
    const res = await request.get('/admin/profile')
    return cast<Admin>(res)
  },
  // 更新管理员信息
  updateProfile: async (data: Partial<Admin>) => {
    const res = await request.put('/admin/profile', data)
    return cast<Admin>(res)
  },
  // 更新管理员密码
  changePassword: async (data: { old_password: string; new_password: string }) => {
    await request.post('/admin/change-password', data)
  },
  // 获取管理员列表
  getList: async (params?: { username?: string; realname?: string; role_id?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/admin/admins', { params })
    return cast<Admin[] | { data: Admin[]; total: number }>(res)
  },
  // 创建管理员
  create: async (data: { username: string; password: string; realname: string; phone?: string; email?: string; role_id?: number }) => {
    const res = await request.post('/admin/admins', data)
    return cast<Admin>(res)
  },
  // 更新管理员信息
  update: async (id: number, data: Partial<Admin>) => {
    const res = await request.put(`/admin/admins/${id}`, data)
    return cast<Admin>(res)
  },
  // 批量更新管理员信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/admin/admins/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除管理员
  delete: async (id: number) => {
    await request.delete(`/admin/admins/${id}`)
  },
  // 批量删除管理员
  batchDelete: async (ids: number[]) => {
    await request.delete('/admin/admins/batch', { data: { ids } })
  },
  // 重置管理员密码
  resetPassword: async (id: number) => {
    await request.post(`/admin/admins/${id}/reset-password`)
  }
}

// 房间接口
export const roomApi = {
  // 获取房间列表
  getList: async (params?: { type_id?: string; floor?: string; status?: string; name?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/room', { params })
    return cast<Room[] | { data: Room[]; total: number }>(res)
  },
  // 获取房间详情
  getDetail: async (id: number) => {
    const res = await request.get(`/room/${id}`)
    return cast<Room>(res)
  },
  // 创建房间
  create: async (data: Partial<Room>) => {
    const res = await request.post('/room', data)
    return cast<Room>(res)
  },
  // 更新房间信息
  update: async (id: number, data: Partial<Room>) => {
    const res = await request.put(`/room/${id}`, data)
    return cast<Room>(res)
  },
  // 批量更新房间信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/room/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除房间
  delete: async (id: number) => {
    await request.delete(`/room/${id}`)
  },
  // 批量删除房间
  batchDelete: async (ids: number[]) => {
    await request.delete('/room/batch', { data: { ids } })
  }
}

// 房间类型接口
export const roomTypeApi = {
  // 获取房间类型列表
  getList: async (params?: { name?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/room-type', { params })
    return cast<RoomType[] | { data: RoomType[]; total: number }>(res)
  },
  // 获取房间类型详情
  getDetail: async (id: number) => {
    const res = await request.get(`/room-type/${id}`)
    return cast<RoomType>(res)
  },
  // 创建房间类型
  create: async (data: Partial<RoomType>) => {
    const res = await request.post('/room-type', data)
    return cast<RoomType>(res)
  },
  // 更新房间类型信息
  update: async (id: number, data: Partial<RoomType>) => {
    const res = await request.put(`/room-type/${id}`, data)
    return cast<RoomType>(res)
  },
  // 批量更新房间类型信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/room-type/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除房间类型
  delete: async (id: number) => {
    await request.delete(`/room-type/${id}`)
  },
  // 批量删除房间类型
  batchDelete: async (ids: number[]) => {
    await request.delete('/room-type/batch', { data: { ids } })
  }
}

// 订单接口
export const orderApi = {
  // 获取订单列表
  getList: async (params?: { order_no?: string; user_id?: string; room_id?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/order', { params })
    return cast<Order[] | { data: Order[]; total: number }>(res)
  },
  // 获取订单详情
  getDetail: async (id: number) => {
    const res = await request.get(`/order/${id}`)
    return cast<Order>(res)
  },
  // 创建订单
  create: async (data: Partial<Order>) => {
    const res = await request.post('/order', data)
    return cast<Order>(res)
  },
  // 更新订单信息
  update: async (id: number, data: Partial<Order>) => {
    const res = await request.put(`/order/${id}`, data)
    return cast<Order>(res)
  },
  // 批量更新订单信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/order/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 确认订单
  confirm: async (id: number) => {
    const res = await request.put(`/order/${id}/confirm`)
    return cast<Order>(res)
  },
  // 完成订单
  complete: async (id: number) => {
    const res = await request.put(`/order/${id}/complete`)
    return cast<Order>(res)
  },
  // 删除订单
  delete: async (id: number) => {
    await request.delete(`/order/${id}`)
  },
  // 批量删除订单
  batchDelete: async (ids: number[]) => {
    await request.delete('/order/batch', { data: { ids } })
  }
}

// 会员接口
export const membershipApi = {
  // 获取会员列表
  getList: async (params?: { user_id?: string; level?: string; membership_status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/memberships', { params })
    return cast<Membership[] | { data: Membership[]; total: number }>(res)
  },
  // 获取会员详情
  getDetail: async (id: number) => {
    const res = await request.get(`/memberships/${id}`)
    return cast<Membership>(res)
  },
  // 创建会员
  create: async (data: Partial<Membership>) => {
    const res = await request.post('/memberships', data)
    return cast<Membership>(res)
  },
  // 更新会员信息
  update: async (id: number, data: Partial<Membership>) => {
    const res = await request.put(`/memberships/${id}`, data)
    return cast<Membership>(res)
  },
  // 批量更新会员信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/memberships/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除会员
  delete: async (id: number) => {
    await request.delete(`/memberships/${id}`)
  },
  // 批量删除会员
  batchDelete: async (ids: number[]) => {
    await request.delete('/memberships/batch', { data: { ids } })
  }
}

// 用户接口
export const userApi = {
  // 获取用户列表
  getList: async (params?: { nickname?: string; phone?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/user', { params })
    return cast<User[] | { data: User[]; total: number }>(res)
  },
  // 获取用户详情
  getDetail: async (id: number) => {
    const res = await request.get(`/user/${id}`)
    return cast<User>(res)
  },
  // 创建用户
  create: async (data: Partial<User>) => {
    const res = await request.post('/user', data)
    return cast<User>(res)
  },
  // 更新用户信息
  update: async (id: number, data: Partial<User>) => {
    const res = await request.put(`/user/${id}`, data)
    return cast<User>(res)
  },
  // 更新用户状态
  updateStatus: async (id: number, status: number) => {
    const res = await request.put(`/user/${id}/status`, { status })
    return cast<User>(res)
  },
  // 批量更新用户信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/user/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除用户
  delete: async (id: number) => {
    await request.delete(`/user/${id}`)
  },
  // 批量删除用户
  batchDelete: async (ids: number[]) => {
    await request.delete('/user/batch', { data: { ids } })
  }
}

// 活动接口
export const activityApi = {
  // 获取活动列表
  getList: async (params?: { title?: string; activity_status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/activities', { params })
    return cast<Activity[] | { data: Activity[]; total: number }>(res)
  },
  // 获取活动详情
  getDetail: async (id: number) => {
    const res = await request.get(`/activities/${id}`)
    return cast<Activity>(res)
  },
  // 创建活动
  create: async (data: Partial<Activity>) => {
    const res = await request.post('/activities', data)
    return cast<Activity>(res)
  },
  // 更新活动信息
  update: async (id: number, data: Partial<Activity>) => {
    const res = await request.put(`/activities/${id}`, data)
    return cast<Activity>(res)
  },
  // 批量更新活动信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/activities/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除活动
  delete: async (id: number) => {
    await request.delete(`/activities/${id}`)
  },
  // 批量删除活动
  batchDelete: async (ids: number[]) => {
    await request.delete('/activities/batch', { data: { ids } })
  }
}

// 公告接口
export const announcementApi = {
  // 获取公告列表
  getList: async (params?: { title?: string; type?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/announcements', { params })
    return cast<Announcement[] | { data: Announcement[]; total: number }>(res)
  },
  // 获取公告详情
  getDetail: async (id: number) => {
    const res = await request.get(`/announcements/${id}`)
    return cast<Announcement>(res)
  },
  // 创建公告
  create: async (data: Partial<Announcement>) => {
    const res = await request.post('/announcements', data)
    return cast<Announcement>(res)
  },
  // 更新公告信息
  update: async (id: number, data: Partial<Announcement>) => {
    const res = await request.put(`/announcements/${id}`, data)
    return cast<Announcement>(res)
  },
  // 批量更新公告信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/announcements/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除公告
  delete: async (id: number) => {
    await request.delete(`/announcements/${id}`)
  },
  // 批量删除公告
  batchDelete: async (ids: number[]) => {
    await request.delete('/announcements/batch', { data: { ids } })
  }
}

// 充值套餐接口
export const rechargePackageApi = {
  // 获取充值套餐列表
  getList: async (params?: { name?: string; package_status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/recharge-packages', { params })
    return cast<RechargePackage[] | { data: RechargePackage[]; total: number }>(res)
  },
  // 获取充值套餐详情
  getDetail: async (id: number) => {
    const res = await request.get(`/recharge-packages/${id}`)
    return cast<RechargePackage>(res)
  },
  // 创建充值套餐
  create: async (data: Partial<RechargePackage>) => {
    const res = await request.post('/recharge-packages', data)
    return cast<RechargePackage>(res)
  },
  // 更新充值套餐信息
  update: async (id: number, data: Partial<RechargePackage>) => {
    const res = await request.put(`/recharge-packages/${id}`, data)
    return cast<RechargePackage>(res)
  },
  // 批量更新充值套餐信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/recharge-packages/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除充值套餐
  delete: async (id: number) => {
    await request.delete(`/recharge-packages/${id}`)
  },
  // 批量删除充值套餐
  batchDelete: async (ids: number[]) => {
    await request.delete('/recharge-packages/batch', { data: { ids } })
  }
}

// 支付接口
export const paymentApi = {
  // 获取支付列表
  getList: async (params?: { user_id?: string; payment_type?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/payments', { params })
    return cast<Payment[] | { data: Payment[]; total: number }>(res)
  },
  // 获取支付详情
  getDetail: async (id: number) => {
    const res = await request.get(`/payments/${id}`)
    return cast<Payment>(res)
  },
  // 创建支付
  create: async (data: Partial<Payment>) => {
    const res = await request.post('/payments', data)
    return cast<Payment>(res)
  },
  // 更新支付信息
  update: async (id: number, data: Partial<Payment>) => {
    const res = await request.put(`/payments/${id}`, data)
    return cast<Payment>(res)
  },
  // 批量更新支付信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/payments/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除支付
  delete: async (id: number) => {
    await request.delete(`/payments/${id}`)
  },
  // 批量删除支付
  batchDelete: async (ids: number[]) => {
    await request.delete('/payments/batch', { data: { ids } })
  }
}

// 充值记录接口
export const rechargeRecordApi = {
  // 获取充值记录列表
  getList: async (params?: { user_id?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/recharge-records', { params })
    return cast<RechargeRecord[] | { data: RechargeRecord[]; total: number }>(res)
  },
  // 获取充值记录详情
  getDetail: async (id: number) => {
    const res = await request.get(`/recharge-records/${id}`)
    return cast<RechargeRecord>(res)
  },
  // 创建充值记录
  create: async (data: Partial<RechargeRecord>) => {
    const res = await request.post('/recharge-records', data)
    return cast<RechargeRecord>(res)
  },
  // 更新充值记录信息
  update: async (id: number, data: Partial<RechargeRecord>) => {
    const res = await request.put(`/recharge-records/${id}`, data)
    return cast<RechargeRecord>(res)
  },
  // 批量更新充值记录信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/recharge-records/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除充值记录
  delete: async (id: number) => {
    await request.delete(`/recharge-records/${id}`)
  },
  // 批量删除充值记录
  batchDelete: async (ids: number[]) => {
    await request.delete('/recharge-records/batch', { data: { ids } })
  }
}

// 通知接口
export const notificationApi = {
  // 获取通知列表
  getList: async (params?: { user_id?: string; type?: string; read_status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/notifications', { params })
    return cast<Notification[] | { data: Notification[]; total: number }>(res)
  },
  // 获取通知详情
  getDetail: async (id: number) => {
    const res = await request.get(`/notifications/${id}`)
    return cast<Notification>(res)
  },
  // 创建通知
  create: async (data: Partial<Notification>) => {
    const res = await request.post('/notifications', data)
    return cast<Notification>(res)
  },
  // 更新通知信息
  update: async (id: number, data: Partial<Notification>) => {
    const res = await request.put(`/notifications/${id}`, data)
    return cast<Notification>(res)
  },
  // 批量更新通知信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/notifications/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除通知
  delete: async (id: number) => {
    await request.delete(`/notifications/${id}`)
  },
  // 批量删除通知
  batchDelete: async (ids: number[]) => {
    await request.delete('/notifications/batch', { data: { ids } })
  },
  // 标记所有通知为已读
  markAllRead: async (user_id: string) => {
    await request.post('/notifications/mark-all-read', { user_id })
  }
}

// 操作日志接口
export const operationLogApi = {
  // 获取操作日志列表
  getList: async (params?: { admin_id?: string; action?: string; module?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/operation-logs', { params })
    return cast<OperationLog[] | { data: OperationLog[]; total: number }>(res)
  },
  // 获取操作日志详情
  getDetail: async (id: number) => {
    const res = await request.get(`/operation-logs/${id}`)
    return cast<OperationLog>(res)
  },
  // 创建操作日志
  create: async (data: Partial<OperationLog>) => {
    const res = await request.post('/operation-logs', data)
    return cast<OperationLog>(res)
  },
  // 删除操作日志
  delete: async (id: number) => {
    await request.delete(`/operation-logs/${id}`)
  },
  // 批量删除操作日志
  batchDelete: async (ids: number[]) => {
    await request.delete('/operation-logs/batch', { data: { ids } })
  },
  // 批量更新操作日志信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/operation-logs/batch', items)
    return cast<{ success: boolean }>(res)
  }
}

// 评价接口
export const reviewApi = {
  // 获取评价列表
  getList: async (params?: { user_id?: string; room_id?: string; rating?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/reviews', { params })
    return cast<Review[] | { data: Review[]; total: number }>(res)
  },
  // 获取评价详情
  getDetail: async (id: number) => {
    const res = await request.get(`/reviews/${id}`)
    return cast<Review>(res)
  },
  // 创建评价
  create: async (data: Partial<Review>) => {
    const res = await request.post('/reviews', data)
    return cast<Review>(res)
  },
  // 更新评价信息
  update: async (id: number, data: Partial<Review>) => {
    const res = await request.put(`/reviews/${id}`, data)
    return cast<Review>(res)
  },
  // 批量更新评价信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/reviews/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除评价
  delete: async (id: number) => {
    await request.delete(`/reviews/${id}`)
  },
  // 批量删除评价
  batchDelete: async (ids: number[]) => {
    await request.delete('/reviews/batch', { data: { ids } })
  }
}

// 节假日接口
export const holidayApi = {
  // 获取假日列表
  getList: async (params?: { is_holiday?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/holidays', { params })
    return cast<Holiday[] | { data: Holiday[]; total: number }>(res)
  },
  // 获取假日详情
  getDetail: async (id: number) => {
    const res = await request.get(`/holidays/${id}`)
    return cast<Holiday>(res)
  },
  // 创建假日
  create: async (data: Partial<Holiday>) => {
    const res = await request.post('/holidays', data)
    return cast<Holiday>(res)
  },
  // 更新假日信息
  update: async (id: number, data: Partial<Holiday>) => {
    const res = await request.put(`/holidays/${id}`, data)
    return cast<Holiday>(res)
  },
  // 批量更新假日信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/holidays/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除假日
  delete: async (id: number) => {
    await request.delete(`/holidays/${id}`)
  },
  // 批量删除假日 
  batchDelete: async (ids: number[]) => {
    await request.delete('/holidays/batch', { data: { ids } })
  }
}

// 时间槽接口
export const timeSlotApi = {
  // 获取时间槽列表
  getList: async (params?: { type_id?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/time-slots', { params })
    return cast<TimeSlot[] | { data: TimeSlot[]; total: number }>(res)
  },
  // 获取时间槽详情
  getDetail: async (id: number) => {
    const res = await request.get(`/time-slots/${id}`)
    return cast<TimeSlot>(res)
  },
  // 创建时间槽
  create: async (data: Partial<TimeSlot>) => {
    const res = await request.post('/time-slots', data)
    return cast<TimeSlot>(res)
  },
  // 更新时间槽信息
  update: async (id: number, data: Partial<TimeSlot>) => {
    const res = await request.put(`/time-slots/${id}`, data)
    return cast<TimeSlot>(res)
  },
  // 批量更新时间槽信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/time-slots/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 删除时间槽
  delete: async (id: number) => {
    await request.delete(`/time-slots/${id}`)
  },
  // 批量删除时间槽 
  batchDelete: async (ids: number[]) => {
    await request.delete('/time-slots/batch', { data: { ids } })
  }
}

// 角色管理接口
export const roleApi = {
  // 获取角色列表
  getList: async (params?: { name?: string; status?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/admin/roles', { params })
    return cast<AdminRole[] | { data: AdminRole[]; total: number }>(res)
  },
  // 获取角色详情
  getDetail: async (id: number) => {
    const res = await request.get(`/admin/roles/${id}`)
    return cast<AdminRole>(res)
  },
  // 获取可用角色列表
  getAvailable: async () => {
    const res = await request.get('/admin/roles/available')
    return cast<AdminRole[]>(res)
  },
  // 创建角色
  create: async (data: { name: string; level: number; description?: string }) => {
    const res = await request.post('/admin/roles', data)
    return cast<AdminRole>(res)
  },
  // 更新角色信息
  update: async (id: number, data: Partial<AdminRole>) => {
    const res = await request.put(`/admin/roles/${id}`, data)
    return cast<AdminRole>(res)
  },
  // 删除角色
  delete: async (id: number) => {
    await request.delete(`/admin/roles/${id}`)
  },
  // 批量更新角色信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/admin/roles/batch', items)
    return cast<{ success: boolean }>(res)
  }
}

// 权限管理接口
export const permissionApi = {
  // 获取权限列表
  getList: async (params?: { name?: string; group?: string; page?: string; page_size?: string }) => {
    const res = await request.get('/admin/permissions', { params })
    return cast<Permission[] | { data: Permission[]; total: number }>(res)
  },
  // 获取分组权限列表
  getGrouped: async () => {
    const res = await request.get('/admin/permissions/grouped')
    return cast<Permission[]>(res)
  },
  // 获取权限详情
  getDetail: async (id: number) => {
    const res = await request.get(`/admin/permissions/${id}`)
    return cast<Permission>(res)
  },
  // 获取角色权限列表
  getRolePermissions: async (roleId: number) => {
    const res = await request.get('/admin/permissions/role', { params: { role_id: roleId } })
    return cast<Permission[]>(res)
  },
  // 获取我的权限列表
  getMyPermissions: async () => {
    const res = await request.get('/admin/permissions/mine')
    return cast<string[]>(res)
  },
  // 创建权限
  create: async (data: Partial<Permission>) => {
    const res = await request.post('/admin/permissions', data)
    return cast<Permission>(res)
  },
  // 更新权限信息
  update: async (id: number, data: Partial<Permission>) => {
    const res = await request.put(`/admin/permissions/${id}`, data)
    return cast<Permission>(res)
  },
  // 删除权限
  delete: async (id: number) => {
    await request.delete(`/admin/permissions/${id}`)
  },
  // 批量更新权限信息
  batchUpdate: async (items: any[]) => {
    const res = await request.put('/admin/permissions/batch', items)
    return cast<{ success: boolean }>(res)
  },
  // 为角色分配权限
  assign: async (roleId: number, permissionIds: number[]) => {
    await request.post(`/admin/permissions/role/${roleId}`, { permission_ids: permissionIds })
  }
}
