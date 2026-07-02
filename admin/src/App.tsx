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

function App() {
  const [loading, setLoading] = useState(true)
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const token = localStorage.getItem('admin_token')
    setIsLoggedIn(!!token)
    setLoading(false)
  }, [])

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
          element={isLoggedIn ? <Layout /> : <Navigate to="/login" />}
        >
          <Route index element={<Dashboard />} />
          <Route path="room" element={<RoomManage />} />
          <Route path="room-type" element={<RoomTypeManage />} />
          <Route path="order" element={<OrderManage />} />
          <Route path="member" element={<MemberManage />} />
          <Route path="activity" element={<ActivityManage />} />
          <Route path="announcement" element={<AnnouncementManage />} />
          <Route path="recharge-package" element={<RechargePackageManage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App