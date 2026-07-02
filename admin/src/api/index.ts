import request from '@/utils/request'
import { Admin, Room, RoomType, Order, Membership, User, Activity, Announcement, RechargePackage } from '@/types'

const cast = <T>(res: any): T => res as unknown as T

export const adminApi = {
  login: async (data: { username: string; password: string }) => {
    const res = await request.post('/admin/login', data)
    return cast<{ admin: Admin; token: string }>(res)
  },
  getProfile: async () => {
    const res = await request.get('/admin/profile')
    return cast<Admin>(res)
  }
}

export const roomApi = {
  getList: async (params?: { type_id?: string; floor?: string; status?: string }) => {
    const res = await request.get('/room', { params })
    return cast<Room[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/room/${id}`)
    return cast<Room>(res)
  },
  create: async (data: Partial<Room>) => {
    const res = await request.post('/room', data)
    return cast<Room>(res)
  },
  update: async (id: number, data: Partial<Room>) => {
    const res = await request.put(`/room/${id}`, data)
    return cast<Room>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/room/${id}`)
  }
}

export const roomTypeApi = {
  getList: async () => {
    const res = await request.get('/room-type')
    return cast<RoomType[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/room-type/${id}`)
    return cast<RoomType>(res)
  },
  create: async (data: Partial<RoomType>) => {
    const res = await request.post('/room-type', data)
    return cast<RoomType>(res)
  },
  update: async (id: number, data: Partial<RoomType>) => {
    const res = await request.put(`/room-type/${id}`, data)
    return cast<RoomType>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/room-type/${id}`)
  }
}

export const orderApi = {
  getList: async (params?: { room_id?: string; status?: string }) => {
    const res = await request.get('/order', { params })
    return cast<Order[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/order/${id}`)
    return cast<Order>(res)
  },
  confirm: async (id: number) => {
    const res = await request.put(`/order/${id}/confirm`)
    return cast<Order>(res)
  },
  complete: async (id: number) => {
    const res = await request.put(`/order/${id}/complete`)
    return cast<Order>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/order/${id}`)
  }
}

export const membershipApi = {
  getList: async (params?: { level?: string; status?: string }) => {
    const res = await request.get('/memberships', { params })
    return cast<Membership[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/memberships/${id}`)
    return cast<Membership>(res)
  },
  update: async (id: number, data: Partial<Membership>) => {
    const res = await request.put(`/memberships/${id}`, data)
    return cast<Membership>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/memberships/${id}`)
  }
}

export const userApi = {
  getList: async () => {
    const res = await request.get('/user')
    return cast<User[]>(res)
  },
  updateStatus: async (id: number, status: number) => {
    const res = await request.put(`/user/${id}/status`, { status })
    return cast<User>(res)
  }
}

export const activityApi = {
  getList: async () => {
    const res = await request.get('/activities')
    return cast<Activity[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/activities/${id}`)
    return cast<Activity>(res)
  },
  create: async (data: Partial<Activity>) => {
    const res = await request.post('/activities', data)
    return cast<Activity>(res)
  },
  update: async (id: number, data: Partial<Activity>) => {
    const res = await request.put(`/activities/${id}`, data)
    return cast<Activity>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/activities/${id}`)
  }
}

export const announcementApi = {
  getList: async () => {
    const res = await request.get('/announcements')
    return cast<Announcement[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/announcements/${id}`)
    return cast<Announcement>(res)
  },
  create: async (data: Partial<Announcement>) => {
    const res = await request.post('/announcements', data)
    return cast<Announcement>(res)
  },
  update: async (id: number, data: Partial<Announcement>) => {
    const res = await request.put(`/announcements/${id}`, data)
    return cast<Announcement>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/announcements/${id}`)
  }
}

export const rechargePackageApi = {
  getList: async () => {
    const res = await request.get('/recharge-packages')
    return cast<RechargePackage[]>(res)
  },
  getDetail: async (id: number) => {
    const res = await request.get(`/recharge-packages/${id}`)
    return cast<RechargePackage>(res)
  },
  create: async (data: Partial<RechargePackage>) => {
    const res = await request.post('/recharge-packages', data)
    return cast<RechargePackage>(res)
  },
  update: async (id: number, data: Partial<RechargePackage>) => {
    const res = await request.put(`/recharge-packages/${id}`, data)
    return cast<RechargePackage>(res)
  },
  delete: async (id: number) => {
    await request.delete(`/recharge-packages/${id}`)
  }
}