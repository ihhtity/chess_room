import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, DatePicker, Select, message } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { activityApi } from '@/api'
import { Activity } from '@/types'

export default function ActivityManage() {
  const [data, setData] = useState<Activity[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const result = await activityApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch activities:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: Activity) => {
    setEditingId(record.id)
    form.setFieldsValue({
      ...record,
      valid_from: record.valid_from ? new Date(record.valid_from) : null,
      valid_to: record.valid_to ? new Date(record.valid_to) : null,
      discount: record.discount * 100
    })
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await activityApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: any) => {
    try {
      const data = {
        ...values,
        discount: (values.discount || 0) / 100,
        valid_from: values.valid_from ? values.valid_from.toISOString() : null,
        valid_to: values.valid_to ? values.valid_to.toISOString() : null
      }
      if (editingId) {
        await activityApi.update(editingId, data)
        message.success('更新成功')
      } else {
        await activityApi.create(data)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '禁用', 1: '启用' }
    return map[status] || '未知'
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '活动名称', dataIndex: 'name', key: 'name' },
    { title: '描述', dataIndex: 'description', key: 'description' },
    { title: '折扣', dataIndex: 'discount', key: 'discount', render: (v: number) => `${(v * 100).toFixed(0)}%` },
    { title: '开始时间', dataIndex: 'valid_from', key: 'valid_from' },
    { title: '结束时间', dataIndex: 'valid_to', key: 'valid_to' },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusText(s) },
    { title: '操作', key: 'action', render: (_: any, record: Activity) => (
      <div>
        <Button type="text" icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
        <Button type="text" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>活动管理</h2>
        <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加活动</Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title={editingId ? '编辑活动' : '添加活动'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="name" rules={[{ required: true, message: '请输入活动名称' }]}>
            <Input placeholder="活动名称" />
          </Form.Item>
          <Form.Item name="description">
            <Input.TextArea placeholder="活动描述" />
          </Form.Item>
          <Form.Item name="image">
            <Input placeholder="活动图片URL" />
          </Form.Item>
          <Form.Item name="discount" rules={[{ required: true, message: '请输入折扣率' }]}>
            <InputNumber placeholder="折扣率(%)" min={0} max={100} />
          </Form.Item>
          <Form.Item name="valid_from" label="开始时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" />
          </Form.Item>
          <Form.Item name="valid_to" label="结束时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" />
          </Form.Item>
          <Form.Item name="status" label="状态">
            <Select placeholder="状态">
              <Select.Option value={0}>禁用</Select.Option>
              <Select.Option value={1}>启用</Select.Option>
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