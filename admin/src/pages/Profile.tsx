import { useState, useEffect } from 'react'
import { Card, Form, Input, Button, message, Avatar, Tag } from 'antd'
import { UserOutlined, LockOutlined, UserSwitchOutlined, MailOutlined, PhoneOutlined } from '@ant-design/icons'
import { adminApi } from '@/api'
import { Admin } from '@/types'

export default function Profile() {
  const [profile, setProfile] = useState<Admin | null>(null)
  const [loading, setLoading] = useState(false)
  const [form] = Form.useForm()
  const [passwordForm] = Form.useForm()

  useEffect(() => {
    fetchProfile()
  }, [])

  const fetchProfile = async () => {
    try {
      const result = await adminApi.getProfile()
      setProfile(result)
      form.setFieldsValue({
        username: result.username,
        realname: result.realname,
        email: result.email || '',
        phone: result.phone || ''
      })
    } catch (error) {
      console.error('Failed to fetch profile:', error)
    }
  }

  const handleSubmit = async (values: any) => {
    try {
      setLoading(true)
      await adminApi.updateProfile(values)
      message.success('更新成功')
      fetchProfile()
    } catch (error) {
      console.error('Failed to update profile:', error)
    } finally {
      setLoading(false)
    }
  }

  const handlePasswordSubmit = async (values: any) => {
    if (values.new_password !== values.confirm_password) {
      message.error('两次输入的密码不一致')
      return
    }
    try {
      setLoading(true)
      await adminApi.changePassword(values)
      message.success('密码修改成功')
      passwordForm.resetFields()
    } catch (error) {
      console.error('Failed to change password:', error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={{ minHeight: '100%', padding: '40px 24px', background: 'linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%)' }}>
      <div style={{ maxWidth: '800px', margin: '0 auto' }}>
        <div style={{ 
          textAlign: 'center', 
          marginBottom: '40px',
          padding: '40px 24px',
          background: '#fff',
          borderRadius: '16px',
          boxShadow: '0 4px 20px rgba(0,0,0,0.08)',
          position: 'relative',
          overflow: 'hidden'
        }}>
          <div style={{ 
            position: 'absolute', 
            top: '-50px', 
            right: '-50px', 
            width: '150px', 
            height: '150px', 
            background: 'rgba(24, 144, 255, 0.08)',
            borderRadius: '50%'
          }} />
          <div style={{ 
            position: 'absolute', 
            bottom: '-30px', 
            left: '-30px', 
            width: '100px', 
            height: '100px', 
            background: 'rgba(99, 102, 241, 0.06)',
            borderRadius: '50%'
          }} />
          
          <div style={{ position: 'relative', zIndex: 1 }}>
            <Avatar 
              size={120} 
              icon={<UserOutlined />} 
              style={{ 
                backgroundColor: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                marginBottom: '20px',
                border: '4px solid #fff',
                boxShadow: '0 4px 16px rgba(102, 126, 234, 0.3)'
              }} 
            />
            <h1 style={{ fontSize: '28px', fontWeight: '700', marginBottom: '8px', color: '#1f2937' }}>
              {profile?.realname || profile?.username}
            </h1>
            <div style={{ display: 'flex', justifyContent: 'center', gap: '12px', marginBottom: '16px' }}>
              <Tag color="blue" style={{ fontSize: '12px', padding: '4px 12px' }}>管理员</Tag>
              {profile?.status === 1 && (
                <Tag color="green" style={{ fontSize: '12px', padding: '4px 12px' }}>在线</Tag>
              )}
            </div>
            <p style={{ color: '#6b7280', fontSize: '14px' }}>
              用户名：{profile?.username}
            </p>
          </div>
        </div>

        <div style={{ display: 'grid', gridTemplateColumns: '1fr', gap: '24px' }}>
          <Card 
            title={
              <span style={{ display: 'flex', alignItems: 'center', gap: '8px', fontSize: '16px', fontWeight: '600', color: '#1f2937' }}>
                <UserSwitchOutlined style={{ color: '#667eea' }} /> 基本信息
              </span>
            }
            variant="outlined"
            style={{ 
              borderRadius: '12px', 
              boxShadow: '0 2px 12px rgba(0,0,0,0.06)',
              border: '1px solid #f0f0f0'
            }}
          >
            <Form form={form} onFinish={handleSubmit} layout="vertical" size="large">
              <Form.Item name="username" label="用户名" rules={[{ required: true, message: '请输入用户名' }]}>
                <Input 
                  placeholder="用户名" 
                  prefix={<UserOutlined style={{ color: '#9ca3af' }} />}
                  style={{ borderRadius: '8px', border: '1px solid #e5e7eb', transition: 'all 0.3s' }}
                />
              </Form.Item>
              <Form.Item name="realname" label="真实姓名" rules={[{ required: true, message: '请输入真实姓名' }]}>
                <Input 
                  placeholder="真实姓名" 
                  prefix={<UserOutlined style={{ color: '#9ca3af' }} />}
                  style={{ borderRadius: '8px', border: '1px solid #e5e7eb' }}
                />
              </Form.Item>
              <Form.Item name="email" label="邮箱">
                <Input 
                  placeholder="邮箱" 
                  prefix={<MailOutlined style={{ color: '#9ca3af' }} />}
                  style={{ borderRadius: '8px', border: '1px solid #e5e7eb' }}
                />
              </Form.Item>
              <Form.Item name="phone" label="手机号">
                <Input 
                  placeholder="手机号" 
                  prefix={<PhoneOutlined style={{ color: '#9ca3af' }} />}
                  style={{ borderRadius: '8px', border: '1px solid #e5e7eb' }}
                />
              </Form.Item>
              <Form.Item style={{ marginBottom: 0 }}>
                <Button 
                  type="primary" 
                  htmlType="submit" 
                  loading={loading} 
                  block
                  style={{ 
                    height: '44px',
                    borderRadius: '8px',
                    background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
                    border: 'none',
                    fontSize: '15px',
                    fontWeight: '500'
                  }}
                >
                  保存修改
                </Button>
              </Form.Item>
            </Form>
          </Card>

          <Card 
            title={
              <span style={{ display: 'flex', alignItems: 'center', gap: '8px', fontSize: '16px', fontWeight: '600', color: '#1f2937' }}>
                <LockOutlined style={{ color: '#f59e0b' }} /> 修改密码
              </span>
            }
            variant="outlined"
            style={{ 
              borderRadius: '12px', 
              boxShadow: '0 2px 12px rgba(0,0,0,0.06)',
              border: '1px solid #f0f0f0'
            }}
          >
            <Form form={passwordForm} onFinish={handlePasswordSubmit} layout="vertical" size="large">
              <Form.Item name="old_password" label="原密码" rules={[{ required: true, message: '请输入原密码' }]}>
                <Input.Password 
                  placeholder="原密码" 
                  prefix={<LockOutlined style={{ color: '#9ca3af' }} />}
                  style={{ borderRadius: '8px', border: '1px solid #e5e7eb' }}
                />
              </Form.Item>
              <Form.Item name="new_password" label="新密码" rules={[{ required: true, message: '请输入新密码' }]}>
                <Input.Password 
                  placeholder="新密码" 
                  prefix={<LockOutlined style={{ color: '#9ca3af' }} />}
                  style={{ borderRadius: '8px', border: '1px solid #e5e7eb' }}
                />
              </Form.Item>
              <Form.Item name="confirm_password" label="确认密码" rules={[{ required: true, message: '请确认密码' }]}>
                <Input.Password 
                  placeholder="确认密码" 
                  prefix={<LockOutlined style={{ color: '#9ca3af' }} />}
                  style={{ borderRadius: '8px', border: '1px solid #e5e7eb' }}
                />
              </Form.Item>
              <Form.Item style={{ marginBottom: 0 }}>
                <Button 
                  type="primary" 
                  htmlType="submit" 
                  loading={loading} 
                  block
                  style={{ 
                    height: '44px',
                    borderRadius: '8px',
                    background: 'linear-gradient(135deg, #f59e0b 0%, #d97706 100%)',
                    border: 'none',
                    fontSize: '15px',
                    fontWeight: '500'
                  }}
                >
                  修改密码
                </Button>
              </Form.Item>
            </Form>
          </Card>
        </div>
      </div>
    </div>
  )
}