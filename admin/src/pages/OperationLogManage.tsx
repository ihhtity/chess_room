import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, Tag, message } from 'antd'
import { PlusOutlined, DeleteOutlined } from '@ant-design/icons'
import { operationLogApi } from '@/api'
import { OperationLog } from '@/types'

export default function OperationLogManage() {
  const [data, setData] = useState<OperationLog[]>([])
  const [open, setOpen] = useState(false)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const result = await operationLogApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch operation logs:', error)
    }
  }

  const handleAdd = () => {
    form.resetFields()
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await operationLogApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleSubmit = async (values: Partial<OperationLog>) => {
    try {
      await operationLogApi.create(values)
      message.success('创建成功')
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getActionText = (action: string) => {
    const map: Record<string, string> = {
      'create': '创建',
      'update': '更新',
      'delete': '删除',
      'login': '登录',
      'logout': '退出',
      'view': '查看'
    }
    return map[action] || action
  }

  const getActionTag = (action: string) => {
    const colors: Record<string, string> = {
      'create': 'green',
      'update': 'blue',
      'delete': 'red',
      'login': 'orange',
      'logout': 'default',
      'view': 'purple'
    }
    return <Tag color={colors[action] || 'default'}>{getActionText(action)}</Tag>
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '管理员ID', dataIndex: 'admin_id', key: 'admin_id' },
    { title: '操作类型', dataIndex: 'action', key: 'action', render: (a: string) => getActionTag(a) },
    { title: '模块', dataIndex: 'module', key: 'module' },
    { title: '目标ID', dataIndex: 'target_id', key: 'target_id' },
    { title: '操作内容', dataIndex: 'content', key: 'content', ellipsis: true },
    { title: 'IP', dataIndex: 'ip', key: 'ip' },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
    { title: '操作', key: 'action', render: (_: any, record: OperationLog) => (
      <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
        <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>操作日志管理</h2>
        <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加日志</Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title="添加操作日志"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="admin_id" label="管理员ID" rules={[{ required: true, message: '请输入管理员ID' }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="action" label="操作类型" rules={[{ required: true, message: '请选择操作类型' }]}>
            <Select>
              <Select.Option value="create">创建</Select.Option>
              <Select.Option value="update">更新</Select.Option>
              <Select.Option value="delete">删除</Select.Option>
              <Select.Option value="login">登录</Select.Option>
              <Select.Option value="logout">退出</Select.Option>
              <Select.Option value="view">查看</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="module" label="模块">
            <Input placeholder="如：房间管理" />
          </Form.Item>
          <Form.Item name="target_id" label="目标ID">
            <Input type="number" />
          </Form.Item>
          <Form.Item name="content" label="操作内容">
            <Input.TextArea placeholder="操作详情" rows={3} />
          </Form.Item>
          <Form.Item name="ip" label="IP地址">
            <Input />
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
