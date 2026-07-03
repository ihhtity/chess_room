import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, Tag, message, Space } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { rechargeRecordApi } from '@/api'
import { RechargeRecord } from '@/types'
import SearchBar from '@/components/SearchBar'

export default function RechargeRecordManage() {
  const [data, setData] = useState<RechargeRecord[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>) => {
    try {
      const result = await rechargeRecordApi.getList(params)
      setData(result)
    } catch (error) {
      console.error('Failed to fetch recharge records:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: RechargeRecord) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await rechargeRecordApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: Partial<RechargeRecord>) => {
    try {
      if (editingId) {
        await rechargeRecordApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await rechargeRecordApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '待充值', 1: '已充值', 2: '充值失败' }
    return map[status] || '未知'
  }

  const getStatusTag = (status: number) => {
    const colors: Record<number, string> = { 0: 'default', 1: 'success', 2: 'error' }
    return <Tag color={colors[status]}>{getStatusText(status)}</Tag>
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '用户ID', dataIndex: 'user_id', key: 'user_id' },
    { title: '充值金额', dataIndex: 'amount', key: 'amount', render: (a: number) => `¥${a.toFixed(2)}` },
    { title: '赠送金额', dataIndex: 'gift_amount', key: 'gift_amount', render: (a: number) => `¥${a.toFixed(2)}` },
    { title: '支付ID', dataIndex: 'payment_id', key: 'payment_id' },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusTag(s) },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
    { title: '操作', key: 'action', render: (_: any, record: RechargeRecord) => (
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
        <h2>充值记录管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'user_id', label: '用户ID', type: 'input', placeholder: '请输入用户ID' },
              { key: 'status', label: '状态', type: 'select', options: [
                { label: '待充值', value: '0' },
                { label: '已充值', value: '1' },
                { label: '充值失败', value: '2' }
              ]}
            ]}
            onSearch={handleSearch}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加充值记录</Button>
        </Space>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title={editingId ? '编辑充值记录' : '添加充值记录'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="user_id" label="用户ID" rules={[{ required: true, message: '请输入用户ID' }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="amount" label="充值金额" rules={[{ required: true, message: '请输入充值金额' }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="gift_amount" label="赠送金额">
            <Input type="number" />
          </Form.Item>
          <Form.Item name="payment_id" label="支付ID">
            <Input type="number" />
          </Form.Item>
          <Form.Item name="status" label="状态">
            <Select defaultValue={1}>
              <Select.Option value={0}>待充值</Select.Option>
              <Select.Option value={1}>已充值</Select.Option>
              <Select.Option value={2}>充值失败</Select.Option>
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
