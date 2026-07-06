import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, Select, Tag, message, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { paymentApi } from '@/api'
import { Payment } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'
import { formatDateTime } from '@/utils'

export default function PaymentManage() {
  const [data, setData] = useState<Payment[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedPayments, setSelectedPayments] = useState<Payment[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await paymentApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch payments:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: Payment) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await paymentApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(p => ids.includes(String(p.id)))
    setSelectedPayments(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await paymentApi.batchUpdate(updatedRecords)
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
      await paymentApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleSubmit = async (values: Partial<Payment>) => {
    try {
      if (editingId) {
        await paymentApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await paymentApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const getPaymentTypeText = (type: number) => {
    const map: Record<number, string> = { 1: '微信支付', 2: '支付宝', 3: '余额支付' }
    return map[type] || '未知'
  }

  const getPaymentTypeTag = (type: number) => {
    const colors: Record<number, string> = { 1: 'green', 2: 'blue', 3: 'orange' }
    return <Tag color={colors[type]}>{getPaymentTypeText(type)}</Tag>
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '待支付', 1: '已支付', 2: '支付失败' }
    return map[status] || '未知'
  }

  const getStatusTag = (status: number) => {
    const colors: Record<number, string> = { 0: 'default', 1: 'success', 2: 'error' }
    return <Tag color={colors[status]}>{getStatusText(status)}</Tag>
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: Payment) => (
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
    { title: '订单ID', dataIndex: 'order_id', key: 'order_id', width: 60 },
    { title: '用户ID', dataIndex: 'user_id', key: 'user_id', width: 60 },
    { title: '金额', dataIndex: 'amount', key: 'amount', render: (a: number) => `¥${a.toFixed(2)}` },
    { title: '支付方式', dataIndex: 'payment_type', key: 'payment_type', render: (t: number) => getPaymentTypeTag(t) },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusTag(s) },
    { title: '交易号', dataIndex: 'transaction_no', key: 'transaction_no', ellipsis: true },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at', render: (t: string) => formatDateTime(t) },
    { title: '操作', key: 'action', render: (_: any, record: Payment) => (
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
              { key: 'payment_type', label: '支付方式', type: 'select', options: [
                { label: '微信支付', value: '1' },
                { label: '支付宝', value: '2' },
                { label: '余额支付', value: '3' }
              ]},
              { key: 'status', label: '状态', type: 'select', options: [
                { label: '待支付', value: '0' },
                { label: '已支付', value: '1' },
                { label: '支付失败', value: '2' }
              ]}
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>支付管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="支付管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchDelete={handleBatchDelete}
            onBatchEdit={handleBatchEdit}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加支付记录</Button>
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
        title={editingId ? '编辑支付记录' : '添加支付记录'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
          <Form.Item name="order_id" label="订单ID" rules={[{ required: true, message: '请输入订单ID' }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="user_id" label="用户ID" rules={[{ required: true, message: '请输入用户ID' }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="amount" label="金额" rules={[{ required: true, message: '请输入金额' }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="payment_type" label="支付方式">
            <Select defaultValue={1}>
              <Select.Option value={1}>微信支付</Select.Option>
              <Select.Option value={2}>支付宝</Select.Option>
              <Select.Option value={3}>余额支付</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="status" label="状态">
            <Select defaultValue={0}>
              <Select.Option value={0}>待支付</Select.Option>
              <Select.Option value={1}>已支付</Select.Option>
              <Select.Option value={2}>支付失败</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="transaction_no" label="交易号">
            <Input />
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
        records={selectedPayments}
        fields={[
          { key: 'payment_type', label: '支付方式', type: 'select', options: [{ label: '微信支付', value: 1 }, { label: '支付宝', value: 2 }, { label: '余额支付', value: 3 }] },
          { key: 'status', label: '状态', type: 'select', options: [{ label: '待支付', value: 0 }, { label: '已支付', value: 1 }, { label: '支付失败', value: 2 }] },
          { key: 'transaction_no', label: '交易号', type: 'input' }
        ]}
        title="批量编辑支付记录"
      />
    </div>
  )
}
