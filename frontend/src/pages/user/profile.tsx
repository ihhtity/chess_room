import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Cell, CellGroup, Input, Loading, Dialog, Radio, RadioGroup, Grid, GridItem } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { userApi, User } from '@/api'
import { getUserAvatar } from '@/utils/image'
import './profile.scss'

export default function ProfilePage() {
  const navigate = useNavigate()
  const [user, setUser] = useState<User | null>(null)
  const [isEditing, setIsEditing] = useState(false)
  const [editForm, setEditForm] = useState({
    nickname: '',
    realname: '',
    gender: 0
  })
  const [showChangePassword, setShowChangePassword] = useState(false)
  const [passwordForm, setPasswordForm] = useState({
    old_password: '',
    new_password: '',
    confirm_password: ''
  })

  useEffect(() => {
    fetchUserInfo()
  }, [])

  const fetchUserInfo = async () => {
    try {
      const data = await userApi.getUserInfo()
      setUser(data)
      setEditForm({
        nickname: data.nickname,
        realname: data.realname || '',
        gender: data.gender
      })
    } catch (error) {
      console.error('Failed to fetch user info:', error)
      setUser({
        id: 1, openid: '', phone: '138****8888', nickname: '用户A', realname: '张三', avatar: '', gender: 1
      })
      setEditForm({ nickname: '用户A', realname: '张三', gender: 1 })
    }
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    showToast({ message: '退出成功', type: 'success' })
    setTimeout(() => {
      navigate('/login')
    }, 1000)
  }

  const handleSaveProfile = async () => {
    try {
      const data = await userApi.updateUserInfo(editForm)
      setUser(data)
      setIsEditing(false)
      showToast({ message: '信息修改成功', type: 'success' })
    } catch (error) {
      console.error('Failed to update user info:', error)
    }
  }

  const handleCancelEdit = () => {
    setIsEditing(false)
    if (user) {
      setEditForm({
        nickname: user.nickname,
        realname: user.realname || '',
        gender: user.gender
      })
    }
  }

  const handleChangePassword = async () => {
    if (!passwordForm.old_password || !passwordForm.new_password || !passwordForm.confirm_password) {
      showToast({ message: '请填写完整密码信息', type: 'warning' })
      return
    }

    if (passwordForm.new_password !== passwordForm.confirm_password) {
      showToast({ message: '两次输入的新密码不一致', type: 'warning' })
      return
    }

    if (passwordForm.new_password.length < 6) {
      showToast({ message: '新密码长度不能少于6位', type: 'warning' })
      return
    }

    try {
      await userApi.changePassword({
        old_password: passwordForm.old_password,
        new_password: passwordForm.new_password
      })
      showToast({ message: '密码修改成功', type: 'success' })
      setShowChangePassword(false)
      setPasswordForm({
        old_password: '',
        new_password: '',
        confirm_password: ''
      })
    } catch (error) {
      console.error('Failed to change password:', error)
    }
  }

  const menuItems = [
    { icon: '🎫', label: '我的优惠券', path: '' },
    { icon: '📞', label: '联系客服', path: '' },
    { icon: '⚙️', label: '设置', path: '' },
    { icon: '📄', label: '关于我们', path: '' }
  ]

  if (!user) {
    return (
      <div className="loading-wrapper">
        <Loading type="circular" color="#667eea" />
      </div>
    )
  }

  return (
    <div className="container">
      <div className="profile-header" style={{ background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)' }}>
        <div className="profile-avatar">
          <img src={getUserAvatar(user.avatar)} alt="头像" className="avatar-img" />
        </div>
        <div className="profile-info">
          <span className="profile-name">{user.realname || user.nickname || '用户'}</span>
          <span className="profile-phone">{user.phone}</span>
        </div>
        {!isEditing && (
          <Button type="primary" size="small" className="edit-btn" onClick={() => setIsEditing(true)}>
            编辑
          </Button>
        )}
      </div>

      {isEditing && (
        <CellGroup className="edit-group">
          <Cell title="昵称">
            <Input
              value={editForm.nickname}
              onChange={(value) => setEditForm({ ...editForm, nickname: value })}
            />
          </Cell>
          <Cell title="真实姓名">
            <Input
              value={editForm.realname}
              onChange={(value) => setEditForm({ ...editForm, realname: value })}
            />
          </Cell>
          <Cell title="性别">
            <RadioGroup
              value={editForm.gender}
              onChange={(value) => setEditForm({ ...editForm, gender: Number(value) })}
            >
              <Radio value={1}>男</Radio>
              <Radio value={2}>女</Radio>
              <Radio value={0}>保密</Radio>
            </RadioGroup>
          </Cell>
          <div className="edit-actions">
            <Button type="default" onClick={handleCancelEdit}>取消</Button>
            <Button type="primary" onClick={handleSaveProfile}>保存</Button>
          </div>
        </CellGroup>
      )}

      <CellGroup className="menu-group">
        <Cell title="我的订单" onClick={() => navigate('/orders')} extra={<span className="cell-arrow">›</span>} />
        <Cell title="会员中心" onClick={() => navigate('/member')} extra={<span className="cell-arrow">›</span>} />
        <Cell title="修改密码" onClick={() => setShowChangePassword(true)} extra={<span className="cell-arrow">›</span>} />
      </CellGroup>

      <Grid columns={4} className="grid-menu">
        {menuItems.map((item, index) => (
          <GridItem key={index}>
            <div className="grid-icon-wrapper">
              <span className="grid-icon">{item.icon}</span>
            </div>
            <span className="grid-label">{item.label}</span>
          </GridItem>
        ))}
      </Grid>

      <div className="action-section">
        <Button type="default" block onClick={handleLogout}>
          退出登录
        </Button>
      </div>

      <Dialog
        visible={showChangePassword}
        title="修改密码"
        onCancel={() => setShowChangePassword(false)}
        closeOnOverlayClick={false}
      >
        <CellGroup>
          <Cell title="原密码">
            <Input
              type="password"
              placeholder="请输入原密码"
              value={passwordForm.old_password}
              onChange={(value) => setPasswordForm({ ...passwordForm, old_password: value })}
            />
          </Cell>
          <Cell title="新密码">
            <Input
              type="password"
              placeholder="请输入新密码"
              value={passwordForm.new_password}
              onChange={(value) => setPasswordForm({ ...passwordForm, new_password: value })}
            />
          </Cell>
          <Cell title="确认新密码">
            <Input
              type="password"
              placeholder="请再次输入新密码"
              value={passwordForm.confirm_password}
              onChange={(value) => setPasswordForm({ ...passwordForm, confirm_password: value })}
            />
          </Cell>
        </CellGroup>
        <div className="dialog-actions">
          <Button type="default" onClick={() => setShowChangePassword(false)}>取消</Button>
          <Button type="primary" onClick={handleChangePassword}>确认修改</Button>
        </div>
      </Dialog>
    </div>
  )
}