import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, Tag, message, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { notificationApi } from '@/api'
import { Notification } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'
import { formatDateTime } from '@/utils'

export default function NotificationManage() {
  const [data, setData] = useState<Notification[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedNotifications, setSelectedNotifications] = useState<Notification[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await notificationApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch notifications:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: Notification) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await notificationApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(n => ids.includes(String(n.id)))
    setSelectedNotifications(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await notificationApi.batchUpdate(updatedRecords)
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
      await notificationApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleSubmit = async (values: Partial<Notification>) => {
    try {
      if (editingId) {
        await notificationApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await notificationApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getTypeText = (type: number) => {
    const map: Record<number, string> = { 0: '系统通知', 1: '订单通知', 2: '活动通知', 3: '会员通知' }
    return map[type] || '未知'
  }

  const getTypeTag = (type: number) => {
    const colors: Record<number, string> = { 0: 'default', 1: 'blue', 2: 'orange', 3: 'purple' }
    return <Tag color={colors[type]}>{getTypeText(type)}</Tag>
  }

  const getReadStatusTag = (status: number) => {
    return <Tag color={status === 1 ? 'green' : 'red'}>{status === 1 ? '已读' : '未读'}</Tag>
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: Notification) => (
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
    { title: '用户ID', dataIndex: 'user_id', key: 'user_id', width: 60 },
    { title: '类型', dataIndex: 'type', key: 'type', render: (t: number) => getTypeTag(t) },
    { title: '标题', dataIndex: 'title', key: 'title' },
    { title: '内容', dataIndex: 'content', key: 'content', ellipsis: true },
    { title: '状态', dataIndex: 'read_status', key: 'read_status', render: (s: number) => getReadStatusTag(s) },
    { title: '链接', dataIndex: 'link', key: 'link', ellipsis: true },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at', render: (t: string) => formatDateTime(t) },
    { title: '操作', key: 'action', render: (_: any, record: Notification) => (
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
      {searchVisible && (
        <div style={{ marginBottom: 16, width: '100%' }}>
          <SearchBar
            fields={[
              { key: 'user_id', label: '用户ID', type: 'input', placeholder: '请输入用户ID' },
              { key: 'type', label: '类型', type: 'select', options: [
                { label: '系统通知', value: '0' },
                { label: '订单通知', value: '1' },
                { label: '活动通知', value: '2' },
                { label: '会员通知', value: '3' }
              ]},
              { key: 'read_status', label: '状态', type: 'select', options: [
                { label: '未读', value: '0' },
                { label: '已读', value: '1' }
              ]}
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>通知管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="通知管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchDelete={handleBatchDelete}
            onBatchEdit={handleBatchEdit}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加通知</Button>
        </div>
      </div>
      <Table dataSource={data} columns={columns} rowKey="id" pagination={false} />
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
        title={editingId ? '编辑通知' : '添加通知'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="user_id" label="用户ID">
            <Input type="number" placeholder="0表示全体用户" />
          </Form.Item>
          <Form.Item name="type" label="类型">
            <Select defaultValue={0}>
              <Select.Option value={0}>系统通知</Select.Option>
              <Select.Option value={1}>订单通知</Select.Option>
              <Select.Option value={2}>活动通知</Select.Option>
              <Select.Option value={3}>会员通知</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="title" label="标题" rules={[{ required: true, message: '请输入标题' }]}>
            <Input />
          </Form.Item>
          <Form.Item name="content" label="内容">
            <Input.TextArea placeholder="通知内容" rows={3} />
          </Form.Item>
          <Form.Item name="read_status" label="状态">
            <Select defaultValue={0}>
              <Select.Option value={0}>未读</Select.Option>
              <Select.Option value={1}>已读</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="link" label="链接">
            <Input placeholder="跳转链接" />
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
        records={selectedNotifications}
        fields={[
          { key: 'type', label: '类型', type: 'select', options: [{ label: '系统通知', value: 0 }, { label: '订单通知', value: 1 }, { label: '活动通知', value: 2 }, { label: '会员通知', value: 3 }] },
          { key: 'title', label: '标题', type: 'input' },
          { key: 'content', label: '内容', type: 'textarea' },
          { key: 'read_status', label: '状态', type: 'select', options: [{ label: '未读', value: 0 }, { label: '已读', value: 1 }] },
          { key: 'link', label: '链接', type: 'input' }
        ]}
        title="批量编辑通知"
      />
    </div>
  )
}
