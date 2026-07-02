import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Input, Cell, CellGroup } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { userApi } from '@/api'

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
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="login-page">
      <div className="login-container">
        <div className="login-header">
          <span className="login-title">棋牌室预订系统</span>
          <span className="login-subtitle">欢迎回来</span>
        </div>

        <CellGroup>
          <Cell>
            <Input
              type="tel"
              placeholder="请输入手机号"
              value={phone}
              onChange={(value) => setPhone(value)}
              maxLength={11}
              prefix="📱"
            />
          </Cell>
          <Cell>
            <Input
              type="password"
              placeholder="请输入密码"
              value={password}
              onChange={(value) => setPassword(value)}
              prefix="🔑"
            />
          </Cell>
        </CellGroup>

        <Button type="primary" block loading={isLoading} onClick={handleLogin}>
          {isLoading ? '登录中...' : '登录'}
        </Button>

        <div className="login-footer">
          <span className="forgot-password">忘记密码？</span>
        </div>
      </div>
    </div>
  )
}
