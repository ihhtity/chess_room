import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Cell, CellGroup, Input, Loading, Dialog, Radio, RadioGroup } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { userApi, User } from '@/api'
import CustomTabBar from '@/components/CustomTabBar'
import { getUserAvatar } from '@/utils/image'

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
    { icon: '🎫', label: '我的优惠券', onClick: () => {} },
    { icon: '📞', label: '联系客服', onClick: () => {} },
    { icon: '⚙️', label: '设置', onClick: () => {} },
    { icon: '📄', label: '关于我们', onClick: () => {} }
  ]

  if (!user) {
    return (
      <Loading type="circular" color="#667eea" style={{ marginTop: '100rpx' }} />
    )
  }

  return (
    <div className="page">
      <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 120px)' }}>
        <div className="profile-header">
          <div className="profile-avatar">
            <img src={getUserAvatar(user.avatar)} alt="头像" className="avatar-img" />
          </div>
          <div className="profile-info">
            <span className="profile-name">{user.realname || user.nickname || '用户'}</span>
            <span className="profile-phone">{user.phone}</span>
          </div>
          {!isEditing && (
            <Button type="primary" size="small" onClick={() => setIsEditing(true)}>
              编辑
            </Button>
          )}
        </div>

        {isEditing && (
          <CellGroup>
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
          </CellGroup>
        )}

        {isEditing && (
          <div className="edit-actions">
            <Button type="default" onClick={handleCancelEdit}>取消</Button>
            <Button type="primary" onClick={handleSaveProfile}>保存</Button>
          </div>
        )}

        <CellGroup>
          <Cell title="修改密码" onClick={() => setShowChangePassword(true)} />
          <Cell title="我的订单" onClick={() => navigate('/orders')} />
          <Cell title="会员中心" onClick={() => navigate('/member')} />
        </CellGroup>

        <CellGroup>
          {menuItems.map((item, index) => (
            <Cell key={index} title={item.label} onClick={item.onClick} />
          ))}
        </CellGroup>

        <Button type="default" block onClick={handleLogout} style={{ marginTop: '20px' }}>
          退出登录
        </Button>
      </div>
      <CustomTabBar />

      <Dialog
        visible={showChangePassword}
        title="修改密码"
        onCancel={() => setShowChangePassword(false)}
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
        <div style={{ display: 'flex', gap: '10px', marginTop: '20px' }}>
          <Button type="default" style={{ flex: 1 }} onClick={() => setShowChangePassword(false)}>取消</Button>
          <Button type="primary" style={{ flex: 1 }} onClick={handleChangePassword}>确认修改</Button>
        </div>
      </Dialog>
    </div>
  )
}
