import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, InputNumber, Select, DatePicker, message, Tag } from 'antd'
import { EditOutlined, DeleteOutlined, PlusOutlined } from '@ant-design/icons'
import dayjs from 'dayjs'
import { membershipApi, userApi } from '@/api'
import { Membership, User } from '@/types'
import { getUserAvatar } from '@/utils/image'

export default function MemberManage() {
  const [data, setData] = useState<Membership[]>([])
  const [users, setUsers] = useState<User[]>([])
  const [open, setOpen] = useState(false)
  const [addOpen, setAddOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()
  const [addForm] = Form.useForm()

  useEffect(() => {
    fetchData()
    fetchUsers()
  }, [])

  const fetchData = async () => {
    try {
      const result = await membershipApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch members:', error)
    }
  }

  const fetchUsers = async () => {
    try {
      const result = await userApi.getList()
      setUsers(result)
    } catch (error) {
      console.error('Failed to fetch users:', error)
    }
  }

  const getLevelText = (level: number) => {
    const map: Record<number, string> = { 0: '普通会员', 1: '白银会员', 2: '黄金会员', 3: '钻石会员' }
    return map[level] || '普通会员'
  }

  const getLevelTag = (level: number) => {
    const colors: Record<number, string> = { 0: 'default', 1: 'silver', 2: 'gold', 3: 'purple' }
    return <Tag color={colors[level]}>{getLevelText(level)}</Tag>
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '禁用', 1: '正常' }
    return map[status] || '未知'
  }

  const getStatusTag = (status: number) => {
    const colors: Record<number, string> = { 0: 'red', 1: 'green' }
    return <Tag color={colors[status]}>{getStatusText(status)}</Tag>
  }

  const handleEdit = (record: Membership) => {
    setEditingId(record.id)
    form.setFieldsValue({
      level: record.level,
      balance: record.balance,
      points: record.points,
      membership_status: record.membership_status,
      total_consumed: record.total_consumed,
      total_recharged: record.total_recharged,
      expired_at: record.expired_at ? dayjs(record.expired_at) : null
    })
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await membershipApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleAdd = () => {
    addForm.resetFields()
    setAddOpen(true)
  }

  const handleAddSubmit = async (values: any) => {
    try {
      const data = {
        ...values,
        expired_at: values.expired_at ? values.expired_at.format('YYYY-MM-DD HH:mm:ss') : null
      }
      await membershipApi.create(data)
      message.success('添加成功')
      setAddOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to create membership:', error)
    }
  }

  const handleSubmit = async (values: any) => {
    try {
      if (editingId) {
        const data = {
          ...values,
          expired_at: values.expired_at ? values.expired_at.format('YYYY-MM-DD HH:mm:ss') : null
        }
        await membershipApi.update(editingId, data)
        message.success('更新成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '头像', dataIndex: 'user', key: 'user', width: 80, render: (user: any) => (
      <img 
        src={getUserAvatar(user?.avatar)} 
        alt="用户头像" 
        style={{ width: '40px', height: '40px', objectFit: 'cover', borderRadius: '50%' }}
      />
    )},
    { title: '用户ID', dataIndex: 'user_id', key: 'user_id' },
    { title: '用户名', dataIndex: 'user', key: 'nickname', render: (user: any) => user?.nickname || user?.realname || '-' },
    { title: '会员等级', dataIndex: 'level', key: 'level', render: (l: number) => getLevelTag(l) },
    { title: '余额', dataIndex: 'balance', key: 'balance', render: (v: number) => <span style={{ color: '#52c41a', fontWeight: 'bold' }}>¥{v.toFixed(2)}</span> },
    { title: '积分', dataIndex: 'points', key: 'points', render: (v: number) => <span style={{ color: '#faad14', fontWeight: 'bold' }}>{v}</span> },
    { title: '累计消费', dataIndex: 'total_consumed', key: 'total_consumed', render: (v: number) => `¥${v.toFixed(2)}` },
    { title: '累计储值', dataIndex: 'total_recharged', key: 'total_recharged', render: (v: number) => <span style={{ color: '#1890ff' }}>¥{v.toFixed(2)}</span> },
    { title: '状态', dataIndex: 'membership_status', key: 'membership_status', render: (s: number) => getStatusTag(s) },
    { title: '加入时间', dataIndex: 'joined_at', key: 'joined_at' },
    { title: '操作', key: 'action', render: (_: any, record: Membership) => (
      <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
        <Button size="small" icon={<EditOutlined />} onClick={() => handleEdit(record)} style={{ backgroundColor: '#52c41a', color: '#fff', borderColor: '#52c41a' }}>编辑</Button>
        <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>会员管理</h2>
        <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加会员</Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title="编辑会员"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="level" label="会员等级">
            <Select>
              <Select.Option value={0}>普通会员</Select.Option>
              <Select.Option value={1}>白银会员</Select.Option>
              <Select.Option value={2}>黄金会员</Select.Option>
              <Select.Option value={3}>钻石会员</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="balance" label="余额">
            <InputNumber placeholder="余额" precision={2} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="points" label="积分">
            <InputNumber placeholder="积分" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="total_consumed" label="累计消费">
            <InputNumber placeholder="累计消费" precision={2} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="total_recharged" label="累计储值">
            <InputNumber placeholder="累计储值" precision={2} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="expired_at" label="过期时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="membership_status" label="状态">
            <Select>
              <Select.Option value={0}>禁用</Select.Option>
              <Select.Option value={1}>正常</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">保存</Button>
            <Button onClick={() => setOpen(false)} style={{ marginLeft: 8 }}>取消</Button>
          </Form.Item>
        </Form>
      </Modal>

      <Modal
        title="添加会员"
        open={addOpen}
        onCancel={() => setAddOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={addForm} onFinish={handleAddSubmit} layout="vertical">
          <Form.Item name="user_id" label="用户" rules={[{ required: true, message: '请选择用户' }]}>
            <Select style={{ width: '100%' }}>
              {users.map(user => (
                <Select.Option key={user.id} value={user.id}>{user.nickname || user.realname}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="level" label="会员等级">
            <Select defaultValue={0}>
              <Select.Option value={0}>普通会员</Select.Option>
              <Select.Option value={1}>白银会员</Select.Option>
              <Select.Option value={2}>黄金会员</Select.Option>
              <Select.Option value={3}>钻石会员</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="balance" label="余额">
            <InputNumber placeholder="余额" precision={2} style={{ width: '100%' }} defaultValue={0} />
          </Form.Item>
          <Form.Item name="points" label="积分">
            <InputNumber placeholder="积分" style={{ width: '100%' }} defaultValue={0} />
          </Form.Item>
          <Form.Item name="total_consumed" label="累计消费">
            <InputNumber placeholder="累计消费" precision={2} style={{ width: '100%' }} defaultValue={0} />
          </Form.Item>
          <Form.Item name="total_recharged" label="累计储值">
            <InputNumber placeholder="累计储值" precision={2} style={{ width: '100%' }} defaultValue={0} />
          </Form.Item>
          <Form.Item name="expired_at" label="过期时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="membership_status" label="状态">
            <Select defaultValue={1}>
              <Select.Option value={0}>禁用</Select.Option>
              <Select.Option value={1}>正常</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">保存</Button>
            <Button onClick={() => setAddOpen(false)} style={{ marginLeft: 8 }}>取消</Button>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}