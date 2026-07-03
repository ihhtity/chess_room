import { useState } from 'react'
import { Outlet, useNavigate } from 'react-router-dom'
import { Layout as AntLayout, Menu, Button } from 'antd'
import {
  DashboardOutlined,
  HomeOutlined,
  FolderOpenOutlined,
  ShoppingCartOutlined,
  UserOutlined,
  LogoutOutlined,
  CalendarOutlined,
  BellOutlined,
  CreditCardOutlined,
  SettingOutlined,
  ClockCircleOutlined,
  MessageOutlined,
  WalletOutlined,
  FileTextOutlined,
  NotificationOutlined,
  AuditOutlined,
  LockOutlined,
  TeamOutlined
} from '@ant-design/icons'
import { usePermission } from '@/context/PermissionContext'
import './Layout.scss'

const { Header, Sider, Content } = AntLayout

const menuItems = [
  { key: '/', icon: <DashboardOutlined />, label: '仪表盘', permission: 'statistics_view' },
  { key: '/room-type', icon: <FolderOpenOutlined />, label: '包间类型', permission: 'room_view' },
  { key: '/room', icon: <HomeOutlined />, label: '包间管理', permission: 'room_view' },
  { key: '/order', icon: <ShoppingCartOutlined />, label: '订单管理', permission: 'order_view' },
  { key: '/member', icon: <UserOutlined />, label: '会员管理', permission: 'member_view' },
  { key: '/activity', icon: <CalendarOutlined />, label: '活动管理', permission: 'activity_view' },
  { key: '/announcement', icon: <BellOutlined />, label: '公告管理', permission: 'announcement_view' },
  { key: '/recharge-package', icon: <CreditCardOutlined />, label: '充值套餐', permission: 'member_view' },
  { key: '/time-slot', icon: <ClockCircleOutlined />, label: '时间槽管理', permission: 'room_view' },
  { key: '/review', icon: <MessageOutlined />, label: '评价管理', permission: 'review_view' },
  { key: '/holiday', icon: <CalendarOutlined />, label: '节假日管理', permission: 'room_view' },
  { key: '/payment', icon: <WalletOutlined />, label: '支付管理', permission: 'order_view' },
  { key: '/recharge-record', icon: <FileTextOutlined />, label: '充值记录', permission: 'member_view' },
  { key: '/notification', icon: <NotificationOutlined />, label: '通知管理', permission: 'announcement_view' },
  { key: '/operation-log', icon: <AuditOutlined />, label: '操作日志', permission: 'admin_view' },
  { key: '/role', icon: <TeamOutlined />, label: '角色管理', permission: 'role_view' },
  { key: '/permission', icon: <LockOutlined />, label: '权限管理', permission: 'permission_view' },
  { key: '/admin-manage', icon: <UserOutlined />, label: '管理者管理', permission: 'admin_view' },
  { key: '/profile', icon: <SettingOutlined />, label: '个人资料' }
]

export default function Layout() {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()
  const { hasPermission } = usePermission()

  const handleLogout = () => {
    localStorage.removeItem('admin_token')
    navigate('/login')
  }

  const handleMenuClick = ({ key }: { key: string }) => {
    navigate(key)
  }

  const filteredMenuItems = menuItems.filter(item => {
    if (!item.permission) return true
    return hasPermission(item.permission)
  }).map(item => ({
    key: item.key,
    icon: item.icon,
    label: item.label
  }))

  return (
    <AntLayout className="admin-layout">
      <Sider
        collapsible
        collapsed={collapsed}
        onCollapse={(value) => setCollapsed(value)}
        className="sider"
      >
        <div className="logo">
          <span>棋牌室管理</span>
        </div>
        <Menu
          theme="dark"
          mode="inline"
          selectedKeys={[window.location.pathname]}
          items={filteredMenuItems}
          onClick={handleMenuClick}
        />
      </Sider>
      <AntLayout>
        <Header className="header">
          <div className="header-left">
            <Button
              type="text"
              icon={collapsed ? <DashboardOutlined /> : <DashboardOutlined />}
              onClick={() => setCollapsed(!collapsed)}
            />
          </div>
          <div className="header-right">
            <Button type="text" icon={<LogoutOutlined />} onClick={handleLogout}>
              退出登录
            </Button>
          </div>
        </Header>
        <Content className="content">
          <Outlet />
        </Content>
      </AntLayout>
    </AntLayout>
  )
}