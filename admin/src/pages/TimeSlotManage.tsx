import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, Select, message, Tag, Space } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { timeSlotApi, roomTypeApi } from '@/api'
import { TimeSlot, RoomType } from '@/types'
import SearchBar from '@/components/SearchBar'

export default function TimeSlotManage() {
  const [data, setData] = useState<TimeSlot[]>([])
  const [roomTypes, setRoomTypes] = useState<RoomType[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
    fetchRoomTypes()
  }, [])

  const fetchData = async (params?: Record<string, string>) => {
    try {
      const result = await timeSlotApi.getList(params)
      setData(result)
    } catch (error) {
      console.error('Failed to fetch time slots:', error)
    }
  }

  const fetchRoomTypes = async () => {
    try {
      const result = await roomTypeApi.getList()
      setRoomTypes(result)
    } catch (error) {
      console.error('Failed to fetch room types:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: TimeSlot) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await timeSlotApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: Partial<TimeSlot>) => {
    try {
      if (editingId) {
        await timeSlotApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await timeSlotApi.create(values)
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

  const getStatusTag = (status: number) => {
    const colors: Record<number, string> = { 0: 'red', 1: 'green' }
    return <Tag color={colors[status]}>{getStatusText(status)}</Tag>
  }

  const getRoomTypeName = (typeId: number) => {
    const type = roomTypes.find(t => t.id === typeId)
    return type ? type.name : '-'
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '房间类型', dataIndex: 'type_id', key: 'type_id', render: (v: number) => getRoomTypeName(v) },
    { title: '时间槽名称', dataIndex: 'name', key: 'name' },
    { title: '开始时间', dataIndex: 'start_time', key: 'start_time' },
    { title: '结束时间', dataIndex: 'end_time', key: 'end_time' },
    { title: '价格', dataIndex: 'price', key: 'price', render: (v: number) => <span style={{ color: '#ff4d4f', fontWeight: 'bold' }}>¥{v.toFixed(2)}</span> },
    { title: '工作日价', dataIndex: 'weekday_price', key: 'weekday_price', render: (v: number) => `¥${v.toFixed(2)}` },
    { title: '周末价', dataIndex: 'weekend_price', key: 'weekend_price', render: (v: number) => `¥${v.toFixed(2)}` },
    { title: '节假日价', dataIndex: 'holiday_price', key: 'holiday_price', render: (v: number) => `¥${v.toFixed(2)}` },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusTag(s) },
    { title: '操作', key: 'action', render: (_: any, record: TimeSlot) => (
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
        <h2>时间槽管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'type_id', label: '房间类型', type: 'select', options: roomTypes.map(t => ({ label: t.name, value: String(t.id) })) },
              { key: 'name', label: '时间槽名称', type: 'input', placeholder: '请输入时间槽名称' }
            ]}
            onSearch={handleSearch}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加时间槽</Button>
        </Space>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title={editingId ? '编辑时间槽' : '添加时间槽'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="type_id" label="房间类型" rules={[{ required: true, message: '请选择房间类型' }]}>
            <Select placeholder="选择房间类型">
              {roomTypes.map(type => (
                <Select.Option key={type.id} value={type.id}>{type.name}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="name" label="时间槽名称" rules={[{ required: true, message: '请输入时间槽名称' }]}>
            <Input placeholder="如：上午场" />
          </Form.Item>
          <Form.Item name="start_time" label="开始时间" rules={[{ required: true, message: '请输入开始时间' }]}>
            <Input placeholder="如：09:00" />
          </Form.Item>
          <Form.Item name="end_time" label="结束时间" rules={[{ required: true, message: '请输入结束时间' }]}>
            <Input placeholder="如：12:00" />
          </Form.Item>
          <Form.Item name="price" label="价格">
            <InputNumber placeholder="价格" prefix="¥" min={0} precision={2} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="weekday_price" label="工作日价格">
            <InputNumber placeholder="工作日价格" prefix="¥" min={0} precision={2} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="weekend_price" label="周末价格">
            <InputNumber placeholder="周末价格" prefix="¥" min={0} precision={2} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="holiday_price" label="节假日价格">
            <InputNumber placeholder="节假日价格" prefix="¥" min={0} precision={2} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="sort_order" label="排序">
            <InputNumber placeholder="排序" min={0} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="status" label="状态">
            <Select>
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
