import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Space, message, InputNumber, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import { permissionApi } from '@/api'
import { Permission } from '@/types'
import { usePermission } from '@/context/PermissionContext'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'

export default function PermissionManage() {
  const [permissions, setPermissions] = useState<Permission[]>([])
  const [loading, setLoading] = useState(false)
  const [form] = Form.useForm()
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [isEdit, setIsEdit] = useState(false)
  const [currentPermission, setCurrentPermission] = useState<Permission | null>(null)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const { hasPermission } = usePermission()

  const fetchPermissions = async () => {
    setLoading(true)
    try {
      const data = await permissionApi.getList()
      setPermissions(data)
    } catch (error) {
      message.error('获取权限列表失败')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchPermissions()
  }, [])

  const handleAdd = () => {
    setIsEdit(false)
    setCurrentPermission(null)
    form.resetFields()
    setIsModalOpen(true)
  }

  const handleEdit = (permission: Permission) => {
    setIsEdit(true)
    setCurrentPermission(permission)
    form.setFieldsValue({
      code: permission.code,
      name: permission.name,
      group: permission.group,
      module: permission.module,
      description: permission.description,
      sort_order: permission.sort_order
    })
    setIsModalOpen(true)
  }

  const handleDelete = async (permission: Permission) => {
    Modal.confirm({
      title: '确认删除',
      content: `确定要删除权限 "${permission.name}" 吗？`,
      onOk: async () => {
        try {
          await permissionApi.delete(permission.id)
          message.success('删除成功')
          fetchPermissions()
        } catch (error) {
          message.error('删除失败')
        }
      }
    })
  }

  const handleBatchDelete = async (ids: string[]) => {
    try {
      for (const id of ids) {
        await permissionApi.delete(parseInt(id))
      }
      message.success('批量删除成功')
      fetchPermissions()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields()
      if (isEdit && currentPermission) {
        await permissionApi.update(currentPermission.id, values)
        message.success('更新成功')
      } else {
        await permissionApi.create(values)
        message.success('创建成功')
      }
      setIsModalOpen(false)
      fetchPermissions()
    } catch (error) {
      message.error('操作失败')
    }
  }

  const handleSearch = (values: Record<string, string>) => {
    const params: Record<string, string> = {}
    if (values.name) params.name = values.name
    if (values.group) params.group = values.group
    fetchPermissions()
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      render: (_: any, record: Permission) => (
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
    { title: '权限编码', dataIndex: 'code', key: 'code' },
    { title: '权限名称', dataIndex: 'name', key: 'name' },
    { title: '分组', dataIndex: 'group', key: 'group' },
    { title: '模块', dataIndex: 'module', key: 'module' },
    { title: '描述', dataIndex: 'description', key: 'description' },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: Permission) => (
        <Space>
          {hasPermission('permission_view') && (
            <Button icon={<EditOutlined />} onClick={() => handleEdit(record)}>编辑</Button>
          )}
          {hasPermission('permission_view') && (
            <Button danger icon={<DeleteOutlined />} onClick={() => handleDelete(record)}>删除</Button>
          )}
        </Space>
      )
    }
  ]

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>权限管理</h2>
        <Space>
          <SearchBar
            fields={[
              { key: 'name', label: '权限名称', type: 'input', placeholder: '请输入权限名称' },
              { key: 'group', label: '权限分组', type: 'input', placeholder: '请输入权限分组' }
            ]}
            onSearch={handleSearch}
          />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchDelete={handleBatchDelete}
          />
          {hasPermission('permission_view') && (
            <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加权限</Button>
          )}
        </Space>
      </div>

      <Table
        dataSource={permissions}
        columns={columns}
        rowKey="id"
        loading={loading}
        pagination={{ pageSize: 10 }}
      />

      <Modal
        title={isEdit ? '编辑权限' : '添加权限'}
        open={isModalOpen}
        onCancel={() => setIsModalOpen(false)}
        onOk={handleSubmit}
        width={500}
      >
        <Form form={form} layout="vertical">
          <Form.Item name="code" label="权限编码" rules={[{ required: true, message: '请输入权限编码' }]}>
            <Input placeholder="如: order_view" />
          </Form.Item>
          <Form.Item name="name" label="权限名称" rules={[{ required: true, message: '请输入权限名称' }]}>
            <Input placeholder="如: 订单查看" />
          </Form.Item>
          <Form.Item name="group" label="权限分组">
            <Input placeholder="如: 订单管理" />
          </Form.Item>
          <Form.Item name="module" label="所属模块">
            <Input placeholder="如: order" />
          </Form.Item>
          <Form.Item name="description" label="描述">
            <Input.TextArea placeholder="请输入权限描述" />
          </Form.Item>
          <Form.Item name="sort_order" label="排序">
            <InputNumber min={0} />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}