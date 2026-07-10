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

export interface TimeSlot {
  id: number
  type_id: number
  name: string
  start_time: string
  end_time: string
  price: number
  weekday_price: number
  weekend_price: number
  holiday_price: number
  sort_order: number
  status: number
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
  paid_at: string | null
  cancel_time: string | null
  completed_at: string | null
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
  paid_at: string | null
  refunded_at: string | null
  created_at: string
  updated_at: string
}

export interface Membership {
  id: number
  user_id: number
  user: User
  level: number
  level_name: string
  member_no: string
  points: number
  balance: number
  discount: string
  remaining_hours: number
  total_consumed: number
  total_recharged: number
  membership_status: number
  joined_at: string
  expired_at: string | null
  created_at: string
  updated_at: string
}

export interface RechargeRecord {
  id: number
  user_id: number
  amount: number
  gift_amount: number
  payment_id: number
  status: number
  created_at: string
}

export interface Notification {
  id: number
  user_id: number
  type: number
  title: string
  content: string
  read_status: number
  read: boolean
  link: string
  created_at: string
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

export interface Coupon {
  id: number
  name: string
  amount: number
  min_amount: number
  valid_from: string
  valid_to: string
  status: number
  created_at: string
  updated_at: string
}

export interface UserSetting {
  id: number
  user_id: number
  notifications: number
  sound: number
  vibrate: number
  language: string
  theme: string
  allow_push: number
  allow_marketing: number
  created_at: string
  updated_at: string
}

export interface Feedback {
  id: number
  user_id: number
  content: string
  contact: string
  type: number
  status: number
  reply: string
  replied_at: string | null
  created_at: string
  updated_at: string
}

export const OrderStatus = {
  Pending: 0,
  Active: 1,
  Completed: 2,
  Cancelled: 3,
  Refunding: 4,
  Refunded: 5
}

export const RoomStatus = {
  Maintenance: 0,
  Available: 1,
  InUse: 2,
  Reserved: 3
}

export const MemberLevel = {
  Normal: 0,
  Silver: 1,
  Gold: 2,
  Diamond: 3
}
