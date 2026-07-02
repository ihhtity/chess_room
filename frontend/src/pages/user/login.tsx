import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Input, Cell, CellGroup } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { userApi } from '@/api'
import './login.scss'

export default function LoginPage() {
  const navigate = useNavigate()
  const [phone, setPhone] = useState('')
  const [password, setPassword] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  const handleLogin = async () => {
    if (!phone || !password) {
      showToast({ message: '请输入手机号和密码', type: 'warning' })
      return
    }

    if (!/^1[3-9]\d{9}$/.test(phone)) {
      showToast({ message: '请输入正确的手机号', type: 'warning' })
      return
    }

    if (password.length < 6) {
      showToast({ message: '密码长度不能少于6位', type: 'warning' })
      return
    }

    try {
      setIsLoading(true)
      const data = await userApi.login({ phone, password })
      localStorage.setItem('token', data.token)
      localStorage.setItem('user', JSON.stringify(data.user))
      showToast({ message: '登录成功', type: 'success' })
      setTimeout(() => {
        navigate('/')
      }, 1000)
    } catch (error) {
      console.error('Login failed:', error)
      showToast({ message: '登录失败，请检查账号密码', type: 'error' })
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="login-page">
      <div className="login-bg"></div>
      
      <div className="login-container">
        <div className="login-header">
          <div className="logo-wrapper">
            <span className="logo-icon">🀄</span>
          </div>
          <span className="login-title">棋牌室预订系统</span>
          <span className="login-subtitle">尊享棋牌时光</span>
        </div>

        <CellGroup className="login-form">
          <Cell className="form-cell">
            <Input
              type="tel"
              placeholder="请输入手机号"
              value={phone}
              onChange={(value) => setPhone(value)}
              maxLength={11}
            />
          </Cell>
          <Cell className="form-cell">
            <Input
              type="password"
              placeholder="请输入密码"
              value={password}
              onChange={(value) => setPassword(value)}
            />
          </Cell>
        </CellGroup>

        <Button type="primary" block className="login-btn" loading={isLoading} onClick={handleLogin}>
          {isLoading ? '登录中...' : '登录'}
        </Button>

        <div className="login-footer">
          <span className="forgot-password">忘记密码？</span>
          <span className="register-link">注册账号</span>
        </div>
      </div>
    </div>
  )
}