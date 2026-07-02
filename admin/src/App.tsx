import { useState, useEffect } from 'react'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { Spin } from 'antd'
import Login from '@/pages/Login'
import Layout from '@/layout/Layout'
import Dashboard from '@/pages/Dashboard'
import RoomManage from '@/pages/RoomManage'
import RoomTypeManage from '@/pages/RoomTypeManage'
import OrderManage from '@/pages/OrderManage'
import MemberManage from '@/pages/MemberManage'
import ActivityManage from '@/pages/ActivityManage'
import AnnouncementManage from '@/pages/AnnouncementManage'
import RechargePackageManage from '@/pages/RechargePackageManage'
import TimeSlotManage from '@/pages/TimeSlotManage'
import ReviewManage from '@/pages/ReviewManage'
import HolidayManage from '@/pages/HolidayManage'
import PaymentManage from '@/pages/PaymentManage'
import RechargeRecordManage from '@/pages/RechargeRecordManage'
import NotificationManage from '@/pages/NotificationManage'
import OperationLogManage from '@/pages/OperationLogManage'
import Profile from '@/pages/Profile'

function App() {
  // 加载状态
  const [loading, setLoading] = useState(true)
  // 登录状态
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  // 检查登录状态
  useEffect(() => {
    // 检查本地存储中的 token
    const token = localStorage.getItem('admin_token')
    // 如果 token 存在，设置登录状态为 true
    setIsLoggedIn(!!token)
    // 加载完成后，设置加载状态为 false
    setLoading(false)
  }, [])

  // 加载中，显示加载动画
  if (loading) {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <Spin size="large" />
      </div>
    )
  }

  return (
    // 路由配置
    <BrowserRouter>
      {/* 路由配置 */}
      <Routes>
        {/* 登录路由 */}
        <Route path="/login" element={<Login onLogin={() => setIsLoggedIn(true)} />} />
        {/* 布局路由 */}
        <Route
          path="/*"
          element={isLoggedIn ? <Layout /> : <Navigate to="/login" />}
        >
          {/* 首页路由 */}
          <Route index element={<Dashboard />} />
          {/* 房间管理路由 */}
          <Route path="room" element={<RoomManage />} />
          {/* 房间类型管理路由 */}
          <Route path="room-type" element={<RoomTypeManage />} />
          {/* 订单管理路由 */}
          <Route path="order" element={<OrderManage />} />
          {/* 会员管理路由 */}
          <Route path="member" element={<MemberManage />} />
          {/* 活动管理路由 */}
          <Route path="activity" element={<ActivityManage />} />
          {/* 公告管理路由 */}
          <Route path="announcement" element={<AnnouncementManage />} />
          {/* 充值套餐管理路由 */}
          <Route path="recharge-package" element={<RechargePackageManage />} />
          {/* 时间槽管理路由 */}
          <Route path="time-slot" element={<TimeSlotManage />} />
          {/* 评价管理路由 */}
          <Route path="review" element={<ReviewManage />} />
          {/* 节假日管理路由 */}
          <Route path="holiday" element={<HolidayManage />} />
          {/* 支付管理路由 */}
          <Route path="payment" element={<PaymentManage />} />
          {/* 充值记录管理路由 */}
          <Route path="recharge-record" element={<RechargeRecordManage />} />
          {/* 通知管理路由 */}
          <Route path="notification" element={<NotificationManage />} />
          {/* 操作日志管理路由 */}
          <Route path="operation-log" element={<OperationLogManage />} />
          {/* 个人信息路由 */}
          <Route path="profile" element={<Profile />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App