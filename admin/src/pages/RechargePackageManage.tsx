import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, Select, message } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { rechargePackageApi } from '@/api'
import { RechargePackage } from '@/types'

export default function RechargePackageManage() {
  const [data, setData] = useState<RechargePackage[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const result = await rechargePackageApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch recharge packages:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: RechargePackage) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await rechargePackageApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: Partial<RechargePackage>) => {
    try {
      if (editingId) {
        await rechargePackageApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await rechargePackageApi.create(values)
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
    { title: '套餐名称', dataIndex: 'name', key: 'name' },
    { title: '充值金额', dataIndex: 'amount', key: 'amount', render: (v: number) => `¥${v}` },
    { title: '赠送金额', dataIndex: 'gift_amount', key: 'gift_amount', render: (v: number) => `¥${v}` },
    { title: '赠送积分', dataIndex: 'gift_points', key: 'gift_points' },
    { title: '描述', dataIndex: 'description', key: 'description' },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusText(s) },
    { title: '操作', key: 'action', render: (_: any, record: RechargePackage) => (
      <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
        <Button size="small" icon={<EditOutlined />} onClick={() => handleEdit(record)} style={{ backgroundColor: '#52c41a', color: '#fff', borderColor: '#52c41a' }}>编辑</Button>
        <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>充值套餐管理</h2>
        <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加套餐</Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title={editingId ? '编辑充值套餐' : '添加充值套餐'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="name" rules={[{ required: true, message: '请输入套餐名称' }]}>
            <Input placeholder="套餐名称" />
          </Form.Item>
          <Form.Item name="amount" rules={[{ required: true, message: '请输入充值金额' }]}>
            <InputNumber placeholder="充值金额" prefix="¥" min={0} />
          </Form.Item>
          <Form.Item name="gift_amount">
            <InputNumber placeholder="赠送金额" prefix="¥" min={0} />
          </Form.Item>
          <Form.Item name="gift_points">
            <InputNumber placeholder="赠送积分" min={0} />
          </Form.Item>
          <Form.Item name="description">
            <Input.TextArea placeholder="套餐描述" />
          </Form.Item>
          <Form.Item name="status">
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