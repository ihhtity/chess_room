import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Space, message, Checkbox, Select, InputNumber } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, RestOutlined } from '@ant-design/icons'
import { adminApi, roleApi } from '@/api'
import { Admin, AdminRole } from '@/types'
import { usePermission } from '@/context/PermissionContext'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'

export default function AdminManage() {
  const [admins, setAdmins] = useState<Admin[]>([])
  const [roles, setRoles] = useState<AdminRole[]>([])
  const [loading, setLoading] = useState(false)
  const [form] = Form.useForm()
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [isEdit, setIsEdit] = useState(false)
  const [currentAdmin, setCurrentAdmin] = useState<Admin | null>(null)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const { hasPermission } = usePermission()

  const fetchAdmins = async (params: Record<string, string> = {}) => {
    setLoading(true)
    try {
      const data = await adminApi.getList(params)
      setAdmins(data)
    } catch (error) {
      message.error('获取管理员列表失败')
    } finally {
      setLoading(false)
    }
  }

  const fetchRoles = async () => {
    try {
      const data = await roleApi.getAvailable()
      setRoles(data)
    } catch (error) {
      message.error('获取角色列表失败')
    }
  }

  useEffect(() => {
    fetchAdmins()
    fetchRoles()
  }, [])

  const handleAdd = () => {
    setIsEdit(false)
    setCurrentAdmin(null)
    form.resetFields()
    setIsModalOpen(true)
  }

  const handleEdit = (admin: Admin) => {
    setIsEdit(true)
    setCurrentAdmin(admin)
    form.setFieldsValue({
      username: admin.username,
      realname: admin.realname,
      phone: admin.phone,
      email: admin.email,
      role_id: admin.role_id,
      status: admin.status
    })
    setIsModalOpen(true)
  }

  const handleDelete = async (admin: Admin) => {
    if (admin.id === parseInt(localStorage.getItem('admin_id') || '0')) {
      message.error('不能删除自己')
      return
    }
    Modal.confirm({
      title: '确认删除',
      content: `确定要删除管理员 "${admin.username}" 吗？`,
      onOk: async () => {
        try {
          await adminApi.delete(admin.id)
          message.success('删除成功')
          fetchAdmins()
        } catch (error) {
          message.error('删除失败')
        }
      }
    })
  }

  const handleBatchDelete = async (ids: string[]) => {
    const currentId = parseInt(localStorage.getItem('admin_id') || '0')
    if (ids.includes(String(currentId))) {
      message.error('不能删除自己')
      return
    }
    try {
      for (const id of ids) {
        await adminApi.delete(parseInt(id))
      }
      message.success('批量删除成功')
      fetchAdmins()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleResetPassword = (admin: Admin) => {
    Modal.confirm({
      title: '重置密码',
      content: `确定要重置管理员 "${admin.username}" 的密码吗？重置后密码为 123456`,
      onOk: async () => {
        try {
          await adminApi.resetPassword(admin.id)
          message.success('密码重置成功')
        } catch (error) {
          message.error('密码重置失败')
        }
      }
    })
  }

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields()
      if (isEdit && currentAdmin) {
        await adminApi.update(currentAdmin.id, values)
        message.success('更新成功')
      } else {
        await adminApi.create(values)
        message.success('创建成功')
      }
      setIsModalOpen(false)
      fetchAdmins()
    } catch (error) {
      message.error('操作失败')
    }
  }

  const handleSearch = (values: Record<string, string>) => {
    fetchAdmins(values)
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      render: (_: any, record: Admin) => (
        <Checkbox
          checked={selectedRowKeys.includes(String(record.id))}
          onChange={(e) => {
            if (e.target.checked) {
              setSelectedRowKeys([...selectedRowKeys, String(record.id)])
            } else {
              setSelectedRowKeys(selectedRowKeys.filter(key => key !== String(record.id)))
            }
          }}
        />
      )
    },
    { title: '用户名', dataIndex: 'username', key: 'username' },
    { title: '真实姓名', dataIndex: 'realname', key: 'realname' },
    { title: '手机号', dataIndex: 'phone', key: 'phone' },
    { title: '邮箱', dataIndex: 'email', key: 'email' },
    { 
      title: '角色', 
      dataIndex: 'role', 
      key: 'role', 
      render: (role: AdminRole) => role ? role.name : '无角色' 
    },
    { title: '状态', dataIndex: 'status', key: 'status', render: (status: number) => status === 1 ? '启用' : '禁用' },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: Admin) => (
        <Space>
          {hasPermission('admin_edit') && (
            <Button icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
          )}
          {hasPermission('admin_delete') && (
            <Button danger icon={<DeleteOutlined />} onClick={() => handleDelete(record)}>删除</Button>
          )}
          {hasPermission('admin_edit') && (
            <Button icon={<RestOutlined />} onClick={() => handleResetPassword(record)}>重置密码</Button>
          )}
        </Space>
      )
    }
  ]

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>管理者管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'username', label: '用户名', type: 'input', placeholder: '请输入用户名' },
              { key: 'realname', label: '真实姓名', type: 'input', placeholder: '请输入真实姓名' },
              { key: 'role_id', label: '角色', type: 'select', options: roles.map(r => ({ label: r.name, value: String(r.id) })) },
              { key: 'status', label: '状态', type: 'select', options: [
                { label: '启用', value: '1' },
                { label: '禁用', value: '0' }
              ]}
            ]}
            onSearch={handleSearch}
          />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchDelete={handleBatchDelete}
          />
          {hasPermission('admin_create') && (
            <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加管理者</Button>
          )}
        </Space>
      </div>
      <Table
        dataSource={admins}
        columns={columns}
        rowKey="id"
        loading={loading}
        pagination={{ pageSize: 10 }}
      />
      <Modal
        title={isEdit ? '编辑管理者' : '添加管理者'}
        open={isModalOpen}
        onCancel={() => setIsModalOpen(false)}
        onOk={handleSubmit}
        width={500}
      >
        <Form form={form} layout="vertical">
          {!isEdit && (
            <Form.Item name="username" label="用户名" rules={[{ required: true, message: '请输入用户名' }]}>
              <Input placeholder="请输入用户名" />
            </Form.Item>
          )}
          {!isEdit && (
            <Form.Item name="password" label="密码" rules={[{ required: true, message: '请输入密码' }]}>
              <Input.Password placeholder="请输入密码" />
            </Form.Item>
          )}
          <Form.Item name="realname" label="真实姓名" rules={[{ required: true, message: '请输入真实姓名' }]}>
            <Input placeholder="请输入真实姓名" />
          </Form.Item>
          <Form.Item name="phone" label="手机号">
            <Input placeholder="请输入手机号" />
          </Form.Item>
          <Form.Item name="email" label="邮箱">
            <Input placeholder="请输入邮箱" />
          </Form.Item>
          <Form.Item name="role_id" label="角色">
            <Select
              placeholder="请选择角色"
              options={roles.map(r => ({ label: r.name, value: r.id }))}
            />
          </Form.Item>
          {isEdit && (
            <Form.Item name="status" label="状态">
              <InputNumber min={0} max={1} />
            </Form.Item>
          )}
        </Form>
      </Modal>
    </div>
  )
}