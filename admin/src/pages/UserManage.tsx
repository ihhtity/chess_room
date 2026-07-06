import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, message, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { userApi } from '@/api'
import { User } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'
import { formatDateTime } from '@/utils'

export default function UserManage() {
  const [data, setData] = useState<User[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedUsers, setSelectedUsers] = useState<User[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await userApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch users:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: User) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await userApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(u => ids.includes(String(u.id)))
    setSelectedUsers(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await userApi.batchUpdate(updatedRecords)
      message.success('批量编辑成功')
      fetchData()
      setSelectedRowKeys([])
      setBatchEditVisible(false)
    } catch (error) {
      message.error('批量编辑失败')
    }
  }

  const handleBatchDelete = async (ids: string[]) => {
    try {
      await userApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleSubmit = async (values: Partial<User>) => {
    try {
      if (editingId) {
        await userApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await userApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '禁用', 1: '正常' }
    return map[status] || '未知'
  }

  const getGenderText = (gender: number) => {
    const map: Record<number, string> = { 0: '未知', 1: '男', 2: '女' }
    return map[gender] || '未知'
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: User) => (
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
    { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
    {
      title: '头像', dataIndex: 'avatar', key: 'avatar', width: 80, render: (avatar: string) => (
        <img
          src={avatar || 'https://api.dicebear.com/7.x/avataaars/svg?seed=default'}
          alt="用户头像"
          style={{ width: '40px', height: '40px', borderRadius: '50%', objectFit: 'cover' }}
        />
      )
    },
    { title: '昵称', dataIndex: 'nickname', key: 'nickname' },
    { title: '真实姓名', dataIndex: 'realname', key: 'realname' },
    { title: '手机号', dataIndex: 'phone', key: 'phone' },
    { title: '性别', dataIndex: 'gender', key: 'gender', render: (g: number) => getGenderText(g) },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusText(s) },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at', width: 160, render: (t: string) => formatDateTime(t) },
    {
      title: '操作', key: 'action', render: (_: any, record: User) => (
        <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
          <Button size="small" icon={<EditOutlined />} onClick={() => handleEdit(record)} style={{ backgroundColor: '#52c41a', color: '#fff', borderColor: '#52c41a' }}>编辑</Button>
          <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
        </div>
      )
    }
  ]

  const handleSearch = (values: Record<string, string>) => {
    fetchData(values)
  }

  return (
    <div>
      {searchVisible && (
        <div style={{ marginBottom: 16, width: '100%' }}>
          <SearchBar
            fields={[
              { key: 'nickname', label: '昵称', type: 'input', placeholder: '请输入昵称' },
              { key: 'phone', label: '手机号', type: 'input', placeholder: '请输入手机号' },
              {
                key: 'status', label: '状态', type: 'select', options: [
                  { label: '禁用', value: '0' },
                  { label: '正常', value: '1' }
                ]
              }
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>用户管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="用户管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchEdit={handleBatchEdit}
            onBatchDelete={handleBatchDelete}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加用户</Button>
        </div>
      </div>
      <Table 
        dataSource={data} 
        columns={columns} 
        rowKey="id"
        pagination={false}
      />
      <CustomPagination
        current={currentPage}
        pageSize={pageSize}
        total={total}
        onChange={(page, size) => {
          setCurrentPage(page)
          setPageSize(size)
          fetchData(undefined, page, size)
        }}
      />

      <Modal
        title={editingId ? '编辑用户' : '添加用户'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="nickname" label="昵称" rules={[{ required: true, message: '请输入昵称' }]}>
            <Input placeholder="昵称" />
          </Form.Item>
          <Form.Item name="realname" label="真实姓名">
            <Input placeholder="真实姓名" />
          </Form.Item>
          <Form.Item name="phone" label="手机号">
            <Input placeholder="手机号" />
          </Form.Item>
          <Form.Item name="avatar" label="头像URL">
            <Input placeholder="头像URL" />
          </Form.Item>
          <Form.Item name="gender" label="性别">
            <Select placeholder="选择性别">
              <Select.Option value={0}>未知</Select.Option>
              <Select.Option value={1}>男</Select.Option>
              <Select.Option value={2}>女</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="status" label="状态">
            <Select placeholder="状态">
              <Select.Option value={0}>禁用</Select.Option>
              <Select.Option value={1}>正常</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">保存</Button>
            <Button onClick={() => setOpen(false)} style={{ marginLeft: 8 }}>取消</Button>
          </Form.Item>
        </Form>
      </Modal>
      <BatchEditModal
        visible={batchEditVisible}
        onCancel={() => setBatchEditVisible(false)}
        onOk={handleBatchEditSubmit}
        records={selectedUsers}
        fields={[
          { key: 'nickname', label: '昵称', type: 'input' },
          { key: 'realname', label: '真实姓名', type: 'input' },
          { key: 'phone', label: '手机号', type: 'input' },
          { key: 'gender', label: '性别', type: 'select', options: [{ label: '未知', value: 0 }, { label: '男', value: 1 }, { label: '女', value: 2 }] },
          { key: 'status', label: '状态', type: 'select', options: [{ label: '禁用', value: 0 }, { label: '正常', value: 1 }] }
        ]}
        title="批量编辑用户"
      />
    </div>
  )
}
