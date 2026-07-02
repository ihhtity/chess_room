import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, message } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { roomTypeApi } from '@/api'
import { RoomType } from '@/types'

export default function RoomTypeManage() {
  const [data, setData] = useState<RoomType[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const result = await roomTypeApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch room types:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: RoomType) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await roomTypeApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: Partial<RoomType>) => {
    try {
      if (editingId) {
        await roomTypeApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await roomTypeApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id' },
    { title: '名称', dataIndex: 'name', key: 'name' },
    { title: '描述', dataIndex: 'description', key: 'description' },
    { title: '基础价格', dataIndex: 'base_price', key: 'base_price', render: (v: number) => `¥${v}/小时` },
    { title: '最大人数', dataIndex: 'max_people', key: 'max_people' },
    { title: '操作', key: 'action', render: (_: any, record: RoomType) => (
      <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
        <Button size="small" icon={<EditOutlined />} onClick={() => handleEdit(record)} style={{ backgroundColor: '#52c41a', color: '#fff', borderColor: '#52c41a' }}>编辑</Button>
        <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>包间类型管理</h2>
        <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加类型</Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title={editingId ? '编辑包间类型' : '添加包间类型'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="name" rules={[{ required: true, message: '请输入名称' }]}>
            <Input placeholder="名称" />
          </Form.Item>
          <Form.Item name="description">
            <Input placeholder="描述" />
          </Form.Item>
          <Form.Item name="base_price" rules={[{ required: true, message: '请输入基础价格' }]}>
            <InputNumber placeholder="基础价格" />
          </Form.Item>
          <Form.Item name="max_people" rules={[{ required: true, message: '请输入最大人数' }]}>
            <InputNumber placeholder="最大人数" />
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