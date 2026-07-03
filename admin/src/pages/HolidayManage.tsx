import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, DatePicker, Select, message, Tag, Space } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { holidayApi } from '@/api'
import { Holiday } from '@/types'
import SearchBar from '@/components/SearchBar'

export default function HolidayManage() {
  const [data, setData] = useState<Holiday[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>) => {
    try {
      const result = await holidayApi.getList(params)
      setData(result)
    } catch (error) {
      console.error('Failed to fetch holidays:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: Holiday) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await holidayApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: Partial<Holiday>) => {
    try {
      if (editingId) {
        await holidayApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await holidayApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getIsHolidayText = (isHoliday: number) => {
    const map: Record<number, string> = { 0: '工作日', 1: '节假日' }
    return map[isHoliday] || '未知'
  }

  const getIsHolidayTag = (isHoliday: number) => {
    const colors: Record<number, string> = { 0: 'blue', 1: 'red' }
    return <Tag color={colors[isHoliday]}>{getIsHolidayText(isHoliday)}</Tag>
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '节假日名称', dataIndex: 'name', key: 'name' },
    { title: '日期', dataIndex: 'date', key: 'date' },
    { title: '类型', dataIndex: 'is_holiday', key: 'is_holiday', render: (s: number) => getIsHolidayTag(s) },
    { title: '描述', dataIndex: 'description', key: 'description', ellipsis: true },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
    { title: '操作', key: 'action', render: (_: any, record: Holiday) => (
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
        <h2>节假日管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'name', label: '节假日名称', type: 'input', placeholder: '请输入节假日名称' },
              { key: 'is_holiday', label: '类型', type: 'select', options: [
                { label: '节假日', value: '1' },
                { label: '工作日', value: '0' }
              ]}
            ]}
            onSearch={handleSearch}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加节假日</Button>
        </Space>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title={editingId ? '编辑节假日' : '添加节假日'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="name" label="节假日名称" rules={[{ required: true, message: '请输入节假日名称' }]}>
            <Input placeholder="如：春节" />
          </Form.Item>
          <Form.Item name="date" label="日期" rules={[{ required: true, message: '请选择日期' }]}>
            <DatePicker format="YYYY-MM-DD" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="is_holiday" label="类型">
            <Select defaultValue={1}>
              <Select.Option value={0}>工作日</Select.Option>
              <Select.Option value={1}>节假日</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="description" label="描述">
            <Input.TextArea placeholder="描述" rows={3} />
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
