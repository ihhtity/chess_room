import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, message, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { announcementApi } from '@/api'
import { Announcement } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'
import { formatDateTime } from '@/utils'

export default function AnnouncementManage() {
  const [data, setData] = useState<Announcement[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedAnnouncements, setSelectedAnnouncements] = useState<Announcement[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await announcementApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch announcements:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: Announcement) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await announcementApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchDelete = async (ids: string[]) => {
    try {
      await announcementApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(a => ids.includes(String(a.id)))
    setSelectedAnnouncements(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await announcementApi.batchUpdate(updatedRecords)
      message.success('批量编辑成功')
      fetchData()
      setSelectedRowKeys([])
      setBatchEditVisible(false)
    } catch (error) {
      message.error('批量编辑失败')
    }
  }

  const handleSubmit = async (values: Partial<Announcement>) => {
    try {
      if (editingId) {
        await announcementApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await announcementApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '隐藏', 1: '显示' }
    return map[status] || '未知'
  }

  const getTypeText = (type: number) => {
    const map: Record<number, string> = { 0: '普通', 1: '重要', 2: '紧急' }
    return map[type] || '普通'
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: Announcement) => (
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
    { title: '标题', dataIndex: 'title', key: 'title' },
    { title: '内容', dataIndex: 'content', key: 'content', ellipsis: true },
    { title: '类型', dataIndex: 'type', key: 'type', render: (t: number) => getTypeText(t) },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusText(s) },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at', render: (t: string) => formatDateTime(t) },
    { title: '操作', key: 'action', render: (_: any, record: Announcement) => (
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
              { key: 'title', label: '标题', type: 'input', placeholder: '请输入标题' },
              { key: 'type', label: '类型', type: 'select', options: [
                { label: '普通', value: '0' },
                { label: '重要', value: '1' },
                { label: '紧急', value: '2' }
              ]},
              { key: 'status', label: '状态', type: 'select', options: [
                { label: '显示', value: '1' },
                { label: '隐藏', value: '0' }
              ]}
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>公告管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="公告管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchDelete={handleBatchDelete}
            onBatchEdit={handleBatchEdit}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加公告</Button>
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
        title={editingId ? '编辑公告' : '添加公告'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="title" rules={[{ required: true, message: '请输入标题' }]}>
            <Input placeholder="标题" />
          </Form.Item>
          <Form.Item name="content" rules={[{ required: true, message: '请输入内容' }]}>
            <Input.TextArea placeholder="内容" rows={4} />
          </Form.Item>
          <Form.Item name="type">
            <Select placeholder="类型">
              <Select.Option value={0}>普通</Select.Option>
              <Select.Option value={1}>重要</Select.Option>
              <Select.Option value={2}>紧急</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="status">
            <Select placeholder="状态">
              <Select.Option value={0}>隐藏</Select.Option>
              <Select.Option value={1}>显示</Select.Option>
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
        records={selectedAnnouncements}
        fields={[
          { key: 'title', label: '标题', type: 'input' },
          { key: 'content', label: '内容', type: 'textarea' },
          { key: 'type', label: '类型', type: 'select', options: [{ label: '普通', value: 0 }, { label: '重要', value: 1 }, { label: '紧急', value: 2 }] },
          { key: 'status', label: '状态', type: 'select', options: [{ label: '显示', value: 1 }, { label: '隐藏', value: 0 }] }
        ]}
        title="批量编辑公告"
      />
    </div>
  )
}