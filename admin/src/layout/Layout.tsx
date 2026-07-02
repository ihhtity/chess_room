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
  AuditOutlined
} from '@ant-design/icons'
import './Layout.scss'

const { Header, Sider, Content } = AntLayout

const menuItems = [
  { key: '/', icon: <DashboardOutlined />, label: '仪表盘' },
  { key: '/room-type', icon: <FolderOpenOutlined />, label: '包间类型' },
  { key: '/room', icon: <HomeOutlined />, label: '包间管理' },
  { key: '/order', icon: <ShoppingCartOutlined />, label: '订单管理' },
  { key: '/member', icon: <UserOutlined />, label: '会员管理' },
  { key: '/activity', icon: <CalendarOutlined />, label: '活动管理' },
  { key: '/announcement', icon: <BellOutlined />, label: '公告管理' },
  { key: '/recharge-package', icon: <CreditCardOutlined />, label: '充值套餐' },
  { key: '/time-slot', icon: <ClockCircleOutlined />, label: '时间槽管理' },
  { key: '/review', icon: <MessageOutlined />, label: '评价管理' },
  { key: '/holiday', icon: <CalendarOutlined />, label: '节假日管理' },
  { key: '/payment', icon: <WalletOutlined />, label: '支付管理' },
  { key: '/recharge-record', icon: <FileTextOutlined />, label: '充值记录' },
  { key: '/notification', icon: <NotificationOutlined />, label: '通知管理' },
  { key: '/operation-log', icon: <AuditOutlined />, label: '操作日志' },
  { key: '/profile', icon: <SettingOutlined />, label: '个人资料' }
]

export default function Layout() {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()

  const handleLogout = () => {
    localStorage.removeItem('admin_token')
    navigate('/login')
  }

  const handleMenuClick = ({ key }: { key: string }) => {
    navigate(key)
  }

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
          items={menuItems}
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