import { useNavigate, useLocation } from 'react-router-dom'

export default function CustomTabBar() {
  const navigate = useNavigate()
  const location = useLocation()

  const tabs = [
    { path: '/', label: '首页', icon: '🏠' },
    { path: '/orders', label: '订单', icon: '📋' },
    { path: '/member', label: '会员', icon: '👑' },
    { path: '/profile', label: '我的', icon: '👤' }
  ]

  const isActive = (path: string) => {
    if (path === '/') {
      return location.pathname === '/'
    }
    return location.pathname.startsWith(path)
  }

  return (
    <div className="tab-bar">
      {tabs.map(tab => (
        <div
          key={tab.path}
          className={`tab-item ${isActive(tab.path) ? 'active' : ''}`}
          onClick={() => navigate(tab.path)}
        >
          <span className="tab-icon">{tab.icon}</span>
          <span className="tab-label">{tab.label}</span>
        </div>
      ))}
    </div>
  )
}