import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, Tag, message, Space } from 'antd'
import { PlusOutlined, DeleteOutlined } from '@ant-design/icons'
import { operationLogApi } from '@/api'
import { OperationLog } from '@/types'
import SearchBar from '@/components/SearchBar'

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
      'view': '查看',
      'view_profile': '查看个人信息',
      'view_permissions': '查看权限',
      'reset_password': '重置密码',
      'update_status': '更新状态',
      'confirm': '确认',
      'complete': '完成',
      'cancel': '取消',
      'change_password': '修改密码',
      'assign': '分配权限',
      'recharge': '充值',
      'mark_read': '标记已读'
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
      'view': 'purple',
      'view_profile': 'purple',
      'view_permissions': 'purple',
      'reset_password': 'gold',
      'update_status': 'cyan',
      'confirm': 'green',
      'complete': 'green',
      'cancel': 'red',
      'change_password': 'gold',
      'assign': 'blue',
      'recharge': 'green',
      'mark_read': 'blue'
    }
    return <Tag color={colors[action] || 'default'}>{getActionText(action)}</Tag>
  }

  const getModuleText = (module: string) => {
    const map: Record<string, string> = {
      'admin': '管理员',
      'admins': '管理员管理',
      'roles': '角色管理',
      'permissions': '权限管理',
      'room': '包间管理',
      'room-type': '包间类型',
      'order': '订单管理',
      'orders': '订单',
      'activity': '活动管理',
      'activities': '活动',
      'announcement': '公告管理',
      'announcements': '公告',
      'review': '评价管理',
      'reviews': '评价',
      'membership': '会员管理',
      'memberships': '会员',
      'recharge-packages': '充值套餐',
      'recharge-records': '充值记录',
      'time-slots': '时间槽',
      'holidays': '节假日',
      'payments': '支付记录',
      'notifications': '通知管理',
      'operation-logs': '操作日志',
      'dashboard': '数据统计'
    }
    return map[module] || module
  }

  const columns = [
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    { title: '管理员ID', dataIndex: 'admin_id', key: 'admin_id' },
    { title: '操作类型', dataIndex: 'action', key: 'action', render: (a: string) => getActionTag(a) },
    { title: '模块', dataIndex: 'module', key: 'module', render: (m: string) => getModuleText(m) },
    { title: '目标ID', dataIndex: 'target_id', key: 'target_id' },
    { title: '操作内容', dataIndex: 'content', key: 'content', ellipsis: true },
    { title: 'IP', dataIndex: 'ip', key: 'ip' },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
    { title: '操作', key: 'operation', render: (_: any, record: OperationLog) => (
      <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
        <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  const handleSearch = (values: Record<string, string>) => {
    fetchData()
  }

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>操作日志管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'admin_id', label: '管理员ID', type: 'number', placeholder: '请输入管理员ID' },
              { key: 'action', label: '操作类型', type: 'select', options: [
                { label: '创建', value: 'create' },
                { label: '更新', value: 'update' },
                { label: '删除', value: 'delete' },
                { label: '登录', value: 'login' },
                { label: '退出', value: 'logout' },
                { label: '重置密码', value: 'reset_password' },
                { label: '确认', value: 'confirm' },
                { label: '完成', value: 'complete' },
                { label: '取消', value: 'cancel' },
                { label: '分配权限', value: 'assign' }
              ]},
              { key: 'module', label: '模块', type: 'select', options: [
                { label: '管理员', value: 'admin' },
                { label: '角色管理', value: 'roles' },
                { label: '权限管理', value: 'permissions' },
                { label: '包间管理', value: 'room' },
                { label: '订单管理', value: 'order' },
                { label: '活动管理', value: 'activity' },
                { label: '公告管理', value: 'announcement' },
                { label: '评价管理', value: 'review' },
                { label: '会员管理', value: 'membership' }
              ]}
            ]}
            onSearch={handleSearch}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加日志</Button>
        </Space>
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
