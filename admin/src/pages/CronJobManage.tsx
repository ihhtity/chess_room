import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, message, Tag, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { cronJobApi } from '@/api'
import { CronJob } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'
import { formatDateTime } from '@/utils'

export default function CronJobManage() {
  const [data, setData] = useState<CronJob[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedCronJobs, setSelectedCronJobs] = useState<CronJob[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await cronJobApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch cron jobs:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: CronJob) => {
    setEditingId(record.id)
    form.setFieldsValue({
      ...record
    })
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await cronJobApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchDelete = async (ids: string[]) => {
    try {
      await cronJobApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(c => ids.includes(String(c.id)))
    setSelectedCronJobs(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await cronJobApi.batchUpdate(updatedRecords)
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
      if (editingId) {
        await cronJobApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await cronJobApi.create(values)
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
      render: (_: any, record: CronJob) => (
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
    { title: '任务名称', dataIndex: 'name', key: 'name' },
    { title: 'Cron表达式', dataIndex: 'cron_expression', key: 'cron_expression', ellipsis: true },
    { title: '处理函数', dataIndex: 'handler', key: 'handler', ellipsis: true },
    { title: '运行次数', dataIndex: 'run_count', key: 'run_count', width: 80 },
    { title: '最后运行', dataIndex: 'last_run_at', key: 'last_run_at', render: (t: string) => formatDateTime(t) },
    { title: '下次运行', dataIndex: 'next_run_at', key: 'next_run_at', render: (t: string) => formatDateTime(t) },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusTag(s) },
    { title: '操作', key: 'action', render: (_: any, record: CronJob) => (
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
              { key: 'name', label: '任务名称', type: 'input', placeholder: '请输入任务名称' },
              { key: 'status', label: '状态', type: 'select', options: [
                { label: '启用', value: '1' },
                { label: '禁用', value: '0' }
              ]}
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>定时任务管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="定时任务数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchEdit={handleBatchEdit}
            onBatchDelete={handleBatchDelete}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加任务</Button>
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
        title={editingId ? '编辑定时任务' : '添加定时任务'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="name" label="任务名称" rules={[{ required: true, message: '请输入任务名称' }]}>
            <Input placeholder="任务名称" />
          </Form.Item>
          <Form.Item name="cron_expression" label="Cron表达式" rules={[{ required: true, message: '请输入Cron表达式' }]}>
            <Input placeholder="如: 0 0 * * * (每天凌晨0点)" />
          </Form.Item>
          <Form.Item name="handler" label="处理函数标识" rules={[{ required: true, message: '请输入处理函数标识' }]}>
            <Input placeholder="任务处理函数标识" />
          </Form.Item>
          <Form.Item name="params" label="参数(JSON)">
            <Input.TextArea placeholder="任务参数，JSON格式" rows={3} />
          </Form.Item>
          <Form.Item name="description" label="任务描述">
            <Input.TextArea placeholder="任务描述" rows={3} />
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
        records={selectedCronJobs}
        fields={[
          { key: 'name', label: '任务名称', type: 'input' },
          { key: 'handler', label: '处理函数标识', type: 'input' },
          { key: 'description', label: '任务描述', type: 'textarea' },
          { key: 'status', label: '状态', type: 'select', options: [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }] }
        ]}
        title="批量编辑定时任务"
      />
    </div>
  )
}