import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, Select, message } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { roomApi, roomTypeApi } from '@/api'
import { Room, RoomType } from '@/types'
import { getRoomImage } from '@/utils/image'

export default function RoomManage() {
  const [data, setData] = useState<Room[]>([])
  const [types, setTypes] = useState<RoomType[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
    fetchTypes()
  }, [])

  const fetchData = async () => {
    try {
      const result = await roomApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch rooms:', error)
    }
  }

  const fetchTypes = async () => {
    try {
      const result = await roomTypeApi.getList()
      setTypes(result)
    } catch (error) {
      console.error('Failed to fetch types:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: Room) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await roomApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: Partial<Room>) => {
    try {
      if (editingId) {
        await roomApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await roomApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '维护中', 1: '空闲', 2: '使用中', 3: '已预约' }
    return map[status] || '未知'
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '图片', dataIndex: 'images', key: 'images', width: 80, render: (images: string) => (
      <img 
        src={getRoomImage(images, 120, 90)} 
        alt="房间图片" 
        style={{ width: '60px', height: '45px', objectFit: 'cover', borderRadius: '4px' }}
      />
    )},
    { title: '名称', dataIndex: 'name', key: 'name' },
    { title: '类型', dataIndex: 'type', key: 'type', render: (t: RoomType) => t.name },
    { title: '楼层', dataIndex: 'floor', key: 'floor' },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusText(s) },
    { title: '操作', key: 'action', render: (_: any, record: Room) => (
      <div>
        <Button type="text" icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
        <Button type="text" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>包间管理</h2>
        <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加包间</Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title={editingId ? '编辑包间' : '添加包间'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="name" rules={[{ required: true, message: '请输入名称' }]}>
            <Input placeholder="名称" />
          </Form.Item>
          <Form.Item name="type_id" rules={[{ required: true, message: '请选择类型' }]}>
            <Select placeholder="选择类型">
              {types.map(t => (
                <Select.Option key={t.id} value={t.id}>{t.name}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="floor" rules={[{ required: true, message: '请输入楼层' }]}>
            <Input placeholder="楼层（如：1F、2F）" />
          </Form.Item>
          <Form.Item name="capacity" rules={[{ required: true, message: '请输入容量' }]}>
            <InputNumber placeholder="容量" />
          </Form.Item>
          <Form.Item name="images">
            <Input placeholder="图片URL" />
          </Form.Item>
          <Form.Item name="status">
            <Select placeholder="状态">
              <Select.Option value={0}>维护中</Select.Option>
              <Select.Option value={1}>空闲</Select.Option>
              <Select.Option value={2}>使用中</Select.Option>
              <Select.Option value={3}>已预约</Select.Option>
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