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
  // 获取管理员个人信息
  getProfile: async () => {
    const res = await request.get('/admin/profile')
    return cast<Admin>(res)
  },
  // 更新管理员个人信息
  updateProfile: async (data: Partial<Admin>) => {
    const res = await request.put('/admin/profile', data)
    return cast<Admin>(res)
  },
  // 更新管理员密码
  changePassword: async (data: { old_password: string; new_password: string }) => {
    await request.post('/admin/change-password', data)
  },
  // 获取管理员列表
  getList: async (params?: { username?: string; realname?: string; role_id?: string; status?: string }) => {
    const res = await request.get('/admin/admins', { params })
    return cast<Admin[]>(res)
  },
  // 创建管理员
  create: async (data: { username: string; password: string; realname: string; phone?: string; email?: string; role_id?: number }) => {
    const res = await request.post('/admin/admins', data)
    return cast<Admin>(res)
  },
  // 更新管理员
  update: async (id: number, data: Partial<Admin>) => {
    const res = await request.put(`/admin/admins/${id}`, data)
    return cast<Admin>(res)
  },
  // 删除管理员
  delete: async (id: number) => {
    await request.delete(`/admin/admins/${id}`)
  },
  // 重置管理员密码
  resetPassword: async (id: number) => {
    await request.post(`/admin/admins/${id}/reset-password`)
  }
}

// 房间接口
export const roomApi = {
  // 获取房间列表
  getList: async (params?: { type_id?: string; floor?: string; status?: string }) => {
    const res = await request.get('/room', { params })
    return cast<Room[]>(res)
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
  // 更新房间
  update: async (id: number, data: Partial<Room>) => {
    const res = await request.put(`/room/${id}`, data)
    return cast<Room>(res)
  },
  // 删除房间
  delete: async (id: number) => {
    await request.delete(`/room/${id}`)
  }
}

// 房间类型接口
export const roomTypeApi = {
  // 获取房间类型列表
  getList: async (params?: { name?: string }) => {
    const res = await request.get('/room-type', { params })
    return cast<RoomType[]>(res)
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
  // 更新房间类型
  update: async (id: number, data: Partial<RoomType>) => {
    const res = await request.put(`/room-type/${id}`, data)
    return cast<RoomType>(res)
  },
  // 删除房间类型
  delete: async (id: number) => {
    await request.delete(`/room-type/${id}`)
  }
}

// 订单接口
export const orderApi = {
  // 获取订单列表
  getList: async (params?: { order_no?: string; user_id?: string; room_id?: string; status?: string }) => {
    const res = await request.get('/order', { params })
    return cast<Order[]>(res)
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
  // 更新订单
  update: async (id: number, data: Partial<Order>) => {
    const res = await request.put(`/order/${id}`, data)
    return cast<Order>(res)
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
  }
}

// 会员接口
export const membershipApi = {
  // 获取会员列表
  getList: async (params?: { user_id?: string; level?: string; membership_status?: string }) => {
    const res = await request.get('/memberships', { params })
    return cast<Membership[]>(res)
  },
  // 获取会员等级详情
  getDetail: async (id: number) => {
    const res = await request.get(`/memberships/${id}`)
    return cast<Membership>(res)
  },
  // 创建会员
  create: async (data: Partial<Membership>) => {
    const res = await request.post('/memberships', data)
    return cast<Membership>(res)
  },
  // 更新会员等级
  update: async (id: number, data: Partial<Membership>) => {
    const res = await request.put(`/memberships/${id}`, data)
    return cast<Membership>(res)
  },
  // 删除会员等级
  delete: async (id: number) => {
    await request.delete(`/memberships/${id}`)
  }
}

// 用户接口
export const userApi = {
  // 获取用户列表
  getList: async () => {
    const res = await request.get('/user')
    return cast<User[]>(res)
  },
  // 更新用户状态
  updateStatus: async (id: number, status: number) => {
    const res = await request.put(`/user/${id}/status`, { status })
    return cast<User>(res)
  }
}

// 活动接口
export const activityApi = {
  // 获取活动列表
  getList: async (params?: { title?: string; activity_status?: string }) => {
    const res = await request.get('/activities', { params })
    return cast<Activity[]>(res)
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
  // 更新活动
  update: async (id: number, data: Partial<Activity>) => {
    const res = await request.put(`/activities/${id}`, data)
    return cast<Activity>(res)
  },
  delete: async (id: number) => {
  // 删除活动
    await request.delete(`/activities/${id}`)
  }
}

// 公告接口
export const announcementApi = {
  // 获取公告列表
  getList: async (params?: { title?: string; type?: string; status?: string }) => {
    const res = await request.get('/announcements', { params })
    return cast<Announcement[]>(res)
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
  // 更新公告
  update: async (id: number, data: Partial<Announcement>) => {
    const res = await request.put(`/announcements/${id}`, data)
    return cast<Announcement>(res)
  },
  // 删除公告
  delete: async (id: number) => {
    await request.delete(`/announcements/${id}`)
  }
}

// 充值套餐接口
export const rechargePackageApi = {
  // 获取充值套餐列表
  getList: async (params?: { name?: string; package_status?: string }) => {
    const res = await request.get('/recharge-packages', { params })
    return cast<RechargePackage[]>(res)
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
  // 更新充值套餐
  update: async (id: number, data: Partial<RechargePackage>) => {
    const res = await request.put(`/recharge-packages/${id}`, data)
    return cast<RechargePackage>(res)
  },
  // 删除充值套餐
  delete: async (id: number) => {
    await request.delete(`/recharge-packages/${id}`)
  }
}

// 支付接口
export const paymentApi = {
  getList: async (params?: { user_id?: string; payment_type?: string; status?: string }) => {
    const res = await request.get('/payments', { params })
    return cast<Payment[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/payments/${id}`)
    return cast<Payment>(res)
  },
  create: async (data: Partial<Payment>) => {
    const res = await request.post('/payments', data)
    return cast<Payment>(res)
  },
  update: async (id: number, data: Partial<Payment>) => {
    const res = await request.put(`/payments/${id}`, data)
    return cast<Payment>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/payments/${id}`)
  }
}

// 充值记录接口
export const rechargeRecordApi = {
  getList: async (params?: { user_id?: string; status?: string }) => {
    const res = await request.get('/recharge-records', { params })
    return cast<RechargeRecord[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/recharge-records/${id}`)
    return cast<RechargeRecord>(res)
  },
  create: async (data: Partial<RechargeRecord>) => {
    const res = await request.post('/recharge-records', data)
    return cast<RechargeRecord>(res)
  },
  update: async (id: number, data: Partial<RechargeRecord>) => {
    const res = await request.put(`/recharge-records/${id}`, data)
    return cast<RechargeRecord>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/recharge-records/${id}`)
  }
}

// 通知接口
export const notificationApi = {
  getList: async (params?: { user_id?: string; type?: string; read_status?: string }) => {
    const res = await request.get('/notifications', { params })
    return cast<Notification[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/notifications/${id}`)
    return cast<Notification>(res)
  },
  create: async (data: Partial<Notification>) => {
    const res = await request.post('/notifications', data)
    return cast<Notification>(res)
  },
  update: async (id: number, data: Partial<Notification>) => {
    const res = await request.put(`/notifications/${id}`, data)
    return cast<Notification>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/notifications/${id}`)
  },
  markAllRead: async (user_id: string) => {
    await request.post('/notifications/mark-all-read', { user_id })
  }
}

// 操作日志接口
export const operationLogApi = {
  getList: async (params?: { admin_id?: string; action?: string; module?: string }) => {
    const res = await request.get('/operation-logs', { params })
    return cast<OperationLog[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/operation-logs/${id}`)
    return cast<OperationLog>(res)
  },
  create: async (data: Partial<OperationLog>) => {
    const res = await request.post('/operation-logs', data)
    return cast<OperationLog>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/operation-logs/${id}`)
  }
}

// 评价接口
export const reviewApi = {
  // 获取评价列表
  getList: async (params?: { user_id?: string; room_id?: string; rating?: string; review_status?: string }) => {
    const res = await request.get('/reviews', { params })
    return cast<Review[]>(res)
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
  // 更新评价
  update: async (id: number, data: Partial<Review>) => {
    const res = await request.put(`/reviews/${id}`, data)
    return cast<Review>(res)
  },
  // 删除评价
  delete: async (id: number) => {
    await request.delete(`/reviews/${id}`)
  }
}

// 节假日接口
export const holidayApi = {
  getList: async (params?: { is_holiday?: string }) => {
    const res = await request.get('/holidays', { params })
    return cast<Holiday[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/holidays/${id}`)
    return cast<Holiday>(res)
  },
  create: async (data: Partial<Holiday>) => {
    const res = await request.post('/holidays', data)
    return cast<Holiday>(res)
  },
  update: async (id: number, data: Partial<Holiday>) => {
    const res = await request.put(`/holidays/${id}`, data)
    return cast<Holiday>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/holidays/${id}`)
  }
}

// 时间槽接口
export const timeSlotApi = {
  getList: async (params?: { type_id?: string }) => {
    const res = await request.get('/time-slots', { params })
    return cast<TimeSlot[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/time-slots/${id}`)
    return cast<TimeSlot>(res)
  },
  create: async (data: Partial<TimeSlot>) => {
    const res = await request.post('/time-slots', data)
    return cast<TimeSlot>(res)
  },
  update: async (id: number, data: Partial<TimeSlot>) => {
    const res = await request.put(`/time-slots/${id}`, data)
    return cast<TimeSlot>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/time-slots/${id}`)
  }
}

// 角色管理接口
export const roleApi = {
  getList: async () => {
    const res = await request.get('/admin/roles')
    return cast<AdminRole[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/admin/roles/${id}`)
    return cast<AdminRole>(res)
  },
  getAvailable: async () => {
    const res = await request.get('/admin/roles/available')
    return cast<AdminRole[]>(res)
  },
  create: async (data: { name: string; level: number; description?: string }) => {
    const res = await request.post('/admin/roles', data)
    return cast<AdminRole>(res)
  },
  update: async (id: number, data: Partial<AdminRole>) => {
    const res = await request.put(`/admin/roles/${id}`, data)
    return cast<AdminRole>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/admin/roles/${id}`)
  }
}

// 权限管理接口
export const permissionApi = {
  getList: async () => {
    const res = await request.get('/admin/permissions')
    return cast<Permission[]>(res)
  },
  getGrouped: async () => {
    const res = await request.get('/admin/permissions/grouped')
    return cast<Permission[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/admin/permissions/${id}`)
    return cast<Permission>(res)
  },
  getRolePermissions: async (roleId: number) => {
    const res = await request.get('/admin/permissions/role', { params: { role_id: roleId } })
    return cast<Permission[]>(res)
  },
  getMyPermissions: async () => {
    const res = await request.get('/admin/permissions/mine')
    return cast<string[]>(res)
  },
  create: async (data: Partial<Permission>) => {
    const res = await request.post('/admin/permissions', data)
    return cast<Permission>(res)
  },
  update: async (id: number, data: Partial<Permission>) => {
    const res = await request.put(`/admin/permissions/${id}`, data)
    return cast<Permission>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/admin/permissions/${id}`)
  },
  assign: async (roleId: number, permissionIds: number[]) => {
    await request.post(`/admin/permissions/role/${roleId}`, { permission_ids: permissionIds })
  }
}
