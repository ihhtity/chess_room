import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, message, Tag, Rate, Space } from 'antd'
import { EditOutlined, DeleteOutlined, PlusOutlined } from '@ant-design/icons'
import { reviewApi, roomApi, userApi, orderApi } from '@/api'
import { Review, Room, User, Order } from '@/types'
import SearchBar from '@/components/SearchBar'

export default function ReviewManage() {
  const [data, setData] = useState<Review[]>([])
  const [rooms, setRooms] = useState<Room[]>([])
  const [users, setUsers] = useState<User[]>([])
  const [orders, setOrders] = useState<Order[]>([])
  const [open, setOpen] = useState(false)
  const [addOpen, setAddOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()
  const [addForm] = Form.useForm()

  useEffect(() => {
    fetchData()
    fetchRooms()
    fetchUsers()
    fetchOrders()
  }, [])

  const fetchData = async (params?: Record<string, string>) => {
    try {
      const result = await reviewApi.getList(params)
      setData(result)
    } catch (error) {
      console.error('Failed to fetch reviews:', error)
    }
  }

  const fetchRooms = async () => {
    try {
      const result = await roomApi.getList()
      setRooms(result)
    } catch (error) {
      console.error('Failed to fetch rooms:', error)
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

  const fetchOrders = async () => {
    try {
      const result = await orderApi.getList()
      setOrders(result)
    } catch (error) {
      console.error('Failed to fetch orders:', error)
    }
  }

  const handleEdit = (record: Review) => {
    setEditingId(record.id)
    form.setFieldsValue({
      rating: record.rating,
      content: record.content,
      status: record.status
    })
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await reviewApi.delete(id)
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
      await reviewApi.create(values)
      message.success('添加成功')
      setAddOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to create review:', error)
    }
  }

  const handleSubmit = async (values: Partial<Review>) => {
    try {
      if (editingId) {
        await reviewApi.update(editingId, values)
        message.success('更新成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '隐藏', 1: '显示' }
    return map[status] || '未知'
  }

  const getStatusTag = (status: number) => {
    const colors: Record<number, string> = { 0: 'red', 1: 'green' }
    return <Tag color={colors[status]}>{getStatusText(status)}</Tag>
  }

  const getRoomName = (roomId: number) => {
    const room = rooms.find(r => r.id === roomId)
    return room ? room.name : '-'
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '订单ID', dataIndex: 'order_id', key: 'order_id' },
    { title: '用户ID', dataIndex: 'user_id', key: 'user_id' },
    { title: '用户名', dataIndex: 'user', key: 'user', render: (user: any) => user?.nickname || user?.realname || '-' },
    { title: '房间', dataIndex: 'room_id', key: 'room_id', render: (v: number) => getRoomName(v) },
    { title: '评分', dataIndex: 'rating', key: 'rating', render: (v: number) => <Rate disabled defaultValue={v} /> },
    { title: '评价内容', dataIndex: 'content', key: 'content', ellipsis: true },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusTag(s) },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
    { title: '操作', key: 'action', render: (_: any, record: Review) => (
      <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
        <Button size="small" icon={<EditOutlined />} onClick={() => handleEdit(record)} style={{ backgroundColor: '#52c41a', color: '#fff', borderColor: '#52c41a' }}>编辑</Button>
        <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  const handleSearch = (values: Record<string, string>) => {
    fetchData(values)
  }

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>评价管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'user_id', label: '用户', type: 'select', options: users.map(u => ({ label: u.nickname || u.realname, value: String(u.id) })) },
              { key: 'room_id', label: '房间', type: 'select', options: rooms.map(r => ({ label: r.name, value: String(r.id) })) },
              { key: 'rating', label: '评分', type: 'select', options: [
                { label: '1星', value: '1' },
                { label: '2星', value: '2' },
                { label: '3星', value: '3' },
                { label: '4星', value: '4' },
                { label: '5星', value: '5' }
              ]},
              { key: 'review_status', label: '状态', type: 'select', options: [
                { label: '显示', value: '1' },
                { label: '隐藏', value: '0' }
              ]}
            ]}
            onSearch={handleSearch}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加评价</Button>
        </Space>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title="编辑评价"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="rating" label="评分">
            <Rate defaultValue={5} />
          </Form.Item>
          <Form.Item name="content" label="评价内容">
            <Input.TextArea placeholder="评价内容" rows={4} />
          </Form.Item>
          <Form.Item name="status" label="状态">
            <Select>
              <Select.Option value={0}>隐藏</Select.Option>
              <Select.Option value={1}>显示</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">保存</Button>
            <Button onClick={() => setOpen(false)} style={{ marginLeft: 8 }}>取消</Button>
          </Form.Item>
        </Form>
      </Modal>

      <Modal
        title="添加评价"
        open={addOpen}
        onCancel={() => setAddOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={addForm} onFinish={handleAddSubmit} layout="vertical">
          <Form.Item name="order_id" label="订单" rules={[{ required: true, message: '请选择订单' }]}>
            <Select style={{ width: '100%' }}>
              {orders.map(order => (
                <Select.Option key={order.id} value={order.id}>{order.order_no}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="user_id" label="用户" rules={[{ required: true, message: '请选择用户' }]}>
            <Select style={{ width: '100%' }}>
              {users.map(user => (
                <Select.Option key={user.id} value={user.id}>{user.nickname || user.realname}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="room_id" label="房间" rules={[{ required: true, message: '请选择房间' }]}>
            <Select style={{ width: '100%' }}>
              {rooms.map(room => (
                <Select.Option key={room.id} value={room.id}>{room.name}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="rating" label="评分" rules={[{ required: true, message: '请选择评分' }]}>
            <Rate defaultValue={5} />
          </Form.Item>
          <Form.Item name="content" label="评价内容" rules={[{ required: true, message: '请输入评价内容' }]}>
            <Input.TextArea placeholder="评价内容" rows={4} />
          </Form.Item>
          <Form.Item name="images" label="图片">
            <Input placeholder="图片URL，多个用逗号分隔" />
          </Form.Item>
          <Form.Item name="status" label="状态">
            <Select defaultValue={1}>
              <Select.Option value={0}>隐藏</Select.Option>
              <Select.Option value={1}>显示</Select.Option>
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
