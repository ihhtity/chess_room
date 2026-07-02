import { useState, useEffect } from 'react'
import { Card, Form, Input, Button, message, Avatar } from 'antd'
import { UserOutlined } from '@ant-design/icons'
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
        realname: result.realname
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
    <div>
      <h2>个人资料</h2>
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '24px' }}>
        <Card title="基本信息">
          <div style={{ textAlign: 'center', marginBottom: 24 }}>
            <Avatar size={128} icon={<UserOutlined />} />
            <div style={{ marginTop: 16, fontSize: '20px', fontWeight: 'bold' }}>
              {profile?.realname || profile?.username}
            </div>
            <div style={{ color: '#666' }}>管理员</div>
          </div>
          <Form form={form} onFinish={handleSubmit} layout="vertical">
            <Form.Item name="username" label="用户名" rules={[{ required: true, message: '请输入用户名' }]}>
              <Input placeholder="用户名" />
            </Form.Item>
            <Form.Item name="realname" label="真实姓名" rules={[{ required: true, message: '请输入真实姓名' }]}>
              <Input placeholder="真实姓名" />
            </Form.Item>
            <Form.Item>
              <Button type="primary" htmlType="submit" loading={loading}>保存修改</Button>
            </Form.Item>
          </Form>
        </Card>

        <Card title="修改密码">
          <Form form={passwordForm} onFinish={handlePasswordSubmit} layout="vertical">
            <Form.Item name="old_password" label="原密码" rules={[{ required: true, message: '请输入原密码' }]}>
              <Input.Password placeholder="原密码" />
            </Form.Item>
            <Form.Item name="new_password" label="新密码" rules={[{ required: true, message: '请输入新密码' }]}>
              <Input.Password placeholder="新密码" />
            </Form.Item>
            <Form.Item name="confirm_password" label="确认密码" rules={[{ required: true, message: '请确认密码' }]}>
              <Input.Password placeholder="确认密码" />
            </Form.Item>
            <Form.Item>
              <Button type="primary" htmlType="submit" loading={loading}>修改密码</Button>
            </Form.Item>
          </Form>
        </Card>
      </div>
    </div>
  )
}