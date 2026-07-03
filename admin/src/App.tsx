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
import RoleManage from '@/pages/RoleManage'
import PermissionManage from '@/pages/PermissionManage'
import AdminManage from '@/pages/AdminManage'
import { PermissionProvider } from '@/context/PermissionContext'

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
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login onLogin={() => setIsLoggedIn(true)} />} />
        <Route
          path="/*"
          element={isLoggedIn ? (
            <PermissionProvider>
              <Layout />
            </PermissionProvider>
          ) : (
            <Navigate to="/login" />
          )}
        >
          <Route index element={<Dashboard />} />
          <Route path="room" element={<RoomManage />} />
          <Route path="room-type" element={<RoomTypeManage />} />
          <Route path="order" element={<OrderManage />} />
          <Route path="member" element={<MemberManage />} />
          <Route path="activity" element={<ActivityManage />} />
          <Route path="announcement" element={<AnnouncementManage />} />
          <Route path="recharge-package" element={<RechargePackageManage />} />
          <Route path="time-slot" element={<TimeSlotManage />} />
          <Route path="review" element={<ReviewManage />} />
          <Route path="holiday" element={<HolidayManage />} />
          <Route path="payment" element={<PaymentManage />} />
          <Route path="recharge-record" element={<RechargeRecordManage />} />
          <Route path="notification" element={<NotificationManage />} />
          <Route path="operation-log" element={<OperationLogManage />} />
          <Route path="role" element={<RoleManage />} />
          <Route path="permission" element={<PermissionManage />} />
          <Route path="admin-manage" element={<AdminManage />} />
          <Route path="profile" element={<Profile />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App