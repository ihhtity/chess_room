import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, Space, message, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, UnlockOutlined } from '@ant-design/icons'
import { roleApi, permissionApi } from '@/api'
import { AdminRole, Permission } from '@/types'
import { usePermission } from '@/context/PermissionContext'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'

export default function RoleManage() {
  const [roles, setRoles] = useState<AdminRole[]>([])
  const [permissions, setPermissions] = useState<Permission[]>([])
  const [loading, setLoading] = useState(false)
  const [form] = Form.useForm()
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [isAssignModalOpen, setIsAssignModalOpen] = useState(false)
  const [isEdit, setIsEdit] = useState(false)
  const [currentRole, setCurrentRole] = useState<AdminRole | null>(null)
  const [selectedRole, setSelectedRole] = useState<number>(0)
  const [selectedPermissions, setSelectedPermissions] = useState<number[]>([])
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const { hasPermission } = usePermission()

  const fetchRoles = async () => {
    setLoading(true)
    try {
      const data = await roleApi.getList()
      setRoles(data)
    } catch (error) {
      message.error('获取角色列表失败')
    } finally {
      setLoading(false)
    }
  }

  const fetchPermissions = async () => {
    try {
      const data = await permissionApi.getList()
      setPermissions(data)
    } catch (error) {
      message.error('获取权限列表失败')
    }
  }

  useEffect(() => {
    fetchRoles()
    fetchPermissions()
  }, [])

  const handleAdd = () => {
    setIsEdit(false)
    setCurrentRole(null)
    form.resetFields()
    setIsModalOpen(true)
  }

  const handleEdit = (role: AdminRole) => {
    setIsEdit(true)
    setCurrentRole(role)
    form.setFieldsValue({
      name: role.name,
      level: role.level,
      description: role.description,
      status: role.status
    })
    setIsModalOpen(true)
  }

  const handleDelete = async (role: AdminRole) => {
    Modal.confirm({
      title: '确认删除',
      content: `确定要删除角色 "${role.name}" 吗？`,
      onOk: async () => {
        try {
          await roleApi.delete(role.id)
          message.success('删除成功')
          fetchRoles()
        } catch (error) {
          message.error('删除失败')
        }
      }
    })
  }

  const handleBatchDelete = async (ids: string[]) => {
    try {
      for (const id of ids) {
        await roleApi.delete(parseInt(id))
      }
      message.success('批量删除成功')
      fetchRoles()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields()
      if (isEdit && currentRole) {
        await roleApi.update(currentRole.id, values)
        message.success('更新成功')
      } else {
        await roleApi.create(values)
        message.success('创建成功')
      }
      setIsModalOpen(false)
      fetchRoles()
    } catch (error) {
      message.error('操作失败')
    }
  }

  const handleAssign = (roleId: number) => {
    setSelectedRole(roleId)
    permissionApi.getRolePermissions(roleId).then(data => {
      setSelectedPermissions(data.map(p => p.id))
    })
    setIsAssignModalOpen(true)
  }

  const handleAssignSubmit = async () => {
    try {
      await permissionApi.assign(selectedRole, selectedPermissions)
      message.success('权限分配成功')
      setIsAssignModalOpen(false)
    } catch (error) {
      message.error('权限分配失败')
    }
  }

  const handleSearch = (values: Record<string, string>) => {
    const params: Record<string, string> = {}
    if (values.name) params.name = values.name
    if (values.status) params.status = values.status
    fetchRoles()
  }

  const groupedPermissions = permissions.reduce((acc, perm) => {
    const group = perm.group || '未分组'
    if (!acc[group]) {
      acc[group] = []
    }
    acc[group].push(perm)
    return acc
  }, {} as Record<string, Permission[]>)

  const columns = [
    {
      title: '选择',
      key: 'selection',
      render: (_: any, record: AdminRole) => (
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
    { title: '角色名称', dataIndex: 'name', key: 'name' },
    { title: '层级', dataIndex: 'level', key: 'level', render: (level: number) => `第${level}级` },
    { title: '描述', dataIndex: 'description', key: 'description' },
    { title: '状态', dataIndex: 'status', key: 'status', render: (status: number) => status === 1 ? '启用' : '禁用' },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at' },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: AdminRole) => (
        <Space>
          {hasPermission('permission_assign') && (
            <Button icon={<UnlockOutlined />} onClick={() => handleAssign(record.id)}>分配权限</Button>
          )}
          {hasPermission('role_edit') && (
            <Button icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
          )}
          {hasPermission('role_delete') && (
            <Button danger icon={<DeleteOutlined />} onClick={() => handleDelete(record)}>删除</Button>
          )}
        </Space>
      )
    }
  ]

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>角色管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'name', label: '角色名称', type: 'input', placeholder: '请输入角色名称' },
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
          {hasPermission('role_create') && (
            <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加角色</Button>
          )}
        </Space>
      </div>
      <Table
        dataSource={roles}
        columns={columns}
        rowKey="id"
        loading={loading}
        pagination={{ pageSize: 10 }}
      />
      <Modal
        title={isEdit ? '编辑角色' : '添加角色'}
        open={isModalOpen}
        onCancel={() => setIsModalOpen(false)}
        onOk={handleSubmit}
        width={500}
      >
        <Form form={form} layout="vertical">
          <Form.Item name="name" label="角色名称" rules={[{ required: true, message: '请输入角色名称' }]}>
            <Input placeholder="请输入角色名称" />
          </Form.Item>
          <Form.Item name="level" label="角色层级" rules={[{ required: true, message: '请输入角色层级' }]}>
            <InputNumber min={1} max={10} placeholder="数字越小层级越高" />
          </Form.Item>
          <Form.Item name="description" label="角色描述">
            <Input.TextArea placeholder="请输入角色描述" />
          </Form.Item>
          {isEdit && (
            <Form.Item name="status" label="状态">
              <InputNumber min={0} max={1} />
            </Form.Item>
          )}
        </Form>
      </Modal>
      <Modal
        title="分配权限"
        open={isAssignModalOpen}
        onCancel={() => setIsAssignModalOpen(false)}
        onOk={handleAssignSubmit}
        width={600}
      >
        <div style={{ maxHeight: 400, overflowY: 'auto' }}>
          {Object.entries(groupedPermissions).map(([group, perms]) => (
            <div key={group} style={{ marginBottom: 16 }}>
              <h4>{group}</h4>
              <div style={{ display: 'flex', flexWrap: 'wrap', gap: '8px 16px' }}>
                {perms.map(perm => (
                  <Checkbox
                    key={perm.id}
                    checked={selectedPermissions.includes(perm.id)}
                    onChange={(e) => {
                      if (e.target.checked) {
                        setSelectedPermissions([...selectedPermissions, perm.id])
                      } else {
                        setSelectedPermissions(selectedPermissions.filter(id => id !== perm.id))
                      }
                    }}
                  >
                    {perm.name}
                  </Checkbox>
                ))}
              </div>
            </div>
          ))}
        </div>
      </Modal>
    </div>
  )
}