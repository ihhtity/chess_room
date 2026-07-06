import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, DatePicker, Select, message, Tag, Checkbox } from 'antd'
import type { Dayjs } from 'dayjs'
import dayjs from 'dayjs'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { activityApi } from '@/api'
import { Activity } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'
import { formatDateTime } from '@/utils'

export default function ActivityManage() {
  const [data, setData] = useState<Activity[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedActivities, setSelectedActivities] = useState<Activity[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await activityApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch activities:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: Activity) => {
    setEditingId(record.id)
    form.setFieldsValue({
      ...record,
      valid_from: record.valid_from ? dayjs(record.valid_from) : null,
      valid_to: record.valid_to ? dayjs(record.valid_to) : null,
      discount: record.discount * 100,
      sort_order: record.sort_order
    })
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await activityApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchDelete = async (ids: string[]) => {
    try {
      await activityApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(a => ids.includes(String(a.id)))
    setSelectedActivities(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await activityApi.batchUpdate(updatedRecords)
      message.success('批量编辑成功')
      fetchData()
      setSelectedRowKeys([])
      setBatchEditVisible(false)
    } catch (error) {
      message.error('批量编辑失败')
    }
  }

  const handleSubmit = async (values: any) => {
    try {
      const data = {
        ...values,
        discount: (values.discount || 0) / 100,
        valid_from: values.valid_from ? (values.valid_from as Dayjs).format('YYYY-MM-DD HH:mm:ss') : null,
        valid_to: values.valid_to ? (values.valid_to as Dayjs).format('YYYY-MM-DD HH:mm:ss') : null
      }
      if (editingId) {
        await activityApi.update(editingId, data)
        message.success('更新成功')
      } else {
        await activityApi.create(data)
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

  const getStatusTag = (status: number) => {
    const colors: Record<number, string> = { 0: 'red', 1: 'green' }
    return <Tag color={colors[status]}>{getStatusText(status)}</Tag>
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: Activity) => (
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
    { title: '活动名称', dataIndex: 'name', key: 'name' },
    { title: '描述', dataIndex: 'description', key: 'description', ellipsis: true },
    { title: '折扣', dataIndex: 'discount', key: 'discount', width: 90, render: (v: number) => <span style={{ color: '#ff4d4f', fontWeight: 'bold' }}>{(v * 100).toFixed(0)}%</span> },
    { title: '开始时间', dataIndex: 'valid_from', key: 'valid_from', render: (t: string) => formatDateTime(t) },
    { title: '结束时间', dataIndex: 'valid_to', key: 'valid_to', render: (t: string) => formatDateTime(t) },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusTag(s) },
    { title: '操作', key: 'action', render: (_: any, record: Activity) => (
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
              { key: 'name', label: '活动名称', type: 'input', placeholder: '请输入活动名称' },
              { key: 'activity_status', label: '状态', type: 'select', options: [
                { label: '启用', value: '1' },
                { label: '禁用', value: '0' }
              ]}
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>活动管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="活动管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchEdit={handleBatchEdit}
            onBatchDelete={handleBatchDelete}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加活动</Button>
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
        title={editingId ? '编辑活动' : '添加活动'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="name" label="活动名称" rules={[{ required: true, message: '请输入活动名称' }]}>
            <Input placeholder="活动名称" />
          </Form.Item>
          <Form.Item name="description" label="活动描述">
            <Input.TextArea placeholder="活动描述" rows={3} />
          </Form.Item>
          <Form.Item name="image" label="活动图片URL">
            <Input placeholder="活动图片URL" />
          </Form.Item>
          <Form.Item name="discount" label="折扣率(%)" rules={[{ required: true, message: '请输入折扣率' }]}>
            <InputNumber placeholder="折扣率(%)" min={0} max={100} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="valid_from" label="开始时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="valid_to" label="结束时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="sort_order" label="排序">
            <InputNumber placeholder="排序" min={0} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="status" label="状态">
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
      <BatchEditModal
        visible={batchEditVisible}
        onCancel={() => setBatchEditVisible(false)}
        onOk={handleBatchEditSubmit}
        records={selectedActivities}
        fields={[
          { key: 'name', label: '活动名称', type: 'input' },
          { key: 'description', label: '活动描述', type: 'textarea' },
          { key: 'start_time', label: '开始时间', type: 'input' },
          { key: 'end_time', label: '结束时间', type: 'input' },
          { key: 'sort_order', label: '排序', type: 'number' },
          { key: 'status', label: '状态', type: 'select', options: [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }] }
        ]}
        title="批量编辑活动"
      />
    </div>
  )
}