import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, InputNumber, Select, message } from 'antd'
import { EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { membershipApi } from '@/api'
import { Membership } from '@/types'
import { getUserAvatar } from '@/utils/image'

export default function MemberManage() {
  const [data, setData] = useState<Membership[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const result = await membershipApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch members:', error)
    }
  }

  const getLevelText = (level: number) => {
    const map: Record<number, string> = { 0: '普通会员', 1: '白银会员', 2: '黄金会员', 3: '钻石会员' }
    return map[level] || '普通会员'
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '禁用', 1: '正常' }
    return map[status] || '未知'
  }

  const handleEdit = (record: Membership) => {
    setEditingId(record.id)
    form.setFieldsValue({
      level: record.level,
      balance: record.balance,
      points: record.points,
      membership_status: record.membership_status
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

  const handleSubmit = async (values: any) => {
    try {
      if (editingId) {
        await membershipApi.update(editingId, values)
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
    { title: '会员等级', dataIndex: 'level', key: 'level', render: (l: number) => getLevelText(l) },
    { title: '余额', dataIndex: 'balance', key: 'balance', render: (v: number) => `¥${v.toFixed(2)}` },
    { title: '积分', dataIndex: 'points', key: 'points' },
    { title: '累计消费', dataIndex: 'total_consumed', key: 'total_consumed', render: (v: number) => `¥${v.toFixed(2)}` },
    { title: '累计储值', dataIndex: 'total_recharged', key: 'total_recharged', render: (v: number) => `¥${v.toFixed(2)}` },
    { title: '状态', dataIndex: 'membership_status', key: 'membership_status', render: (s: number) => getStatusText(s) },
    { title: '加入时间', dataIndex: 'joined_at', key: 'joined_at' },
    { title: '操作', key: 'action', render: (_: any, record: Membership) => (
      <div>
        <Button type="text" icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
        <Button type="text" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <h2>会员管理</h2>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title="编辑会员"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="level" label="会员等级">
            <Select>
              <Select.Option value={0}>普通会员</Select.Option>
              <Select.Option value={1}>白银会员</Select.Option>
              <Select.Option value={2}>黄金会员</Select.Option>
              <Select.Option value={3}>钻石会员</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="balance" label="余额">
            <InputNumber placeholder="余额" precision={2} />
          </Form.Item>
          <Form.Item name="points" label="积分">
            <InputNumber placeholder="积分" />
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
    </div>
  )
}