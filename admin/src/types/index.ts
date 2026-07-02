export interface Admin {
  id: number
  username: string
  realname: string
  role: number
  status: number
  created_at: string
  updated_at: string
}

export interface RoomType {
  id: number
  name: string
  description: string
  base_price: number
  max_people: number
  sort_order: number
  status: number
  created_at: string
  updated_at: string
}

export interface Room {
  id: number
  name: string
  type_id: number
  type: RoomType
  floor: string
  capacity: number
  equipment: string
  images: string
  description: string
  status: number
  sort_order: number
  created_at: string
  updated_at: string
}

export interface Order {
  id: number
  order_no: string
  user_id: number
  user: User
  room_id: number
  room: Room
  start_time: string
  end_time: string
  duration: number
  status: number
  total_amount: number
  paid_amount: number
  remark: string
  paid_at: string
  cancel_time: string
  completed_at: string
  created_at: string
  updated_at: string
}

export interface Payment {
  id: number
  order_id: number
  user_id: number
  amount: number
  payment_type: number
  status: number
  transaction_no: string
  created_at: string
  paid_at: string
  refunded_at: string
}

export interface Membership {
  id: number
  user_id: number
  user: User
  level: number
  points: number
  balance: number
  total_consumed: number
  total_recharged: number
  membership_status: number
  joined_at: string
  expired_at: string
  created_at: string
  updated_at: string
}

export interface User {
  id: number
  open_id: string
  phone: string
  nickname: string
  realname: string
  avatar: string
  gender: number
  status: number
  created_at: string
  updated_at: string
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
  sort_order: number
  created_at: string
  updated_at: string
}

export interface Announcement {
  id: number
  title: string
  content: string
  type: number
  status: number
  sort_order: number
  created_at: string
  updated_at: string
}

export interface RechargePackage {
  id: number
  name: string
  amount: number
  gift_amount: number
  gift_points: number
  description: string
  sort_order: number
  status: number
  created_at: string
  updated_at: string
}

export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}