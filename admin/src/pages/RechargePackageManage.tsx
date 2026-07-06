import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, Select, message, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { rechargePackageApi } from '@/api'
import { RechargePackage } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'

export default function RechargePackageManage() {
  const [data, setData] = useState<RechargePackage[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedPackages, setSelectedPackages] = useState<RechargePackage[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await rechargePackageApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch recharge packages:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: RechargePackage) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await rechargePackageApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchDelete = async (ids: string[]) => {
    try {
      await rechargePackageApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(p => ids.includes(String(p.id)))
    setSelectedPackages(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await rechargePackageApi.batchUpdate(updatedRecords)
      message.success('批量编辑成功')
      fetchData()
      setSelectedRowKeys([])
      setBatchEditVisible(false)
    } catch (error) {
      message.error('批量编辑失败')
    }
  }

  const handleSubmit = async (values: Partial<RechargePackage>) => {
    try {
      if (editingId) {
        await rechargePackageApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await rechargePackageApi.create(values)
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

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: RechargePackage) => (
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
    { title: '套餐名称', dataIndex: 'name', key: 'name' },
    { title: '充值金额', dataIndex: 'amount', key: 'amount', render: (v: number) => `¥${v}` },
    { title: '赠送金额', dataIndex: 'gift_amount', key: 'gift_amount', render: (v: number) => `¥${v}` },
    { title: '赠送积分', dataIndex: 'gift_points', key: 'gift_points' },
    { title: '描述', dataIndex: 'description', key: 'description' },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusText(s) },
    { title: '操作', key: 'action', render: (_: any, record: RechargePackage) => (
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
              { key: 'name', label: '套餐名称', type: 'input', placeholder: '请输入套餐名称' },
              { key: 'package_status', label: '状态', type: 'select', options: [
                { label: '启用', value: '1' },
                { label: '禁用', value: '0' }
              ]}
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>充值套餐管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="充值套餐管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchDelete={handleBatchDelete}
            onBatchEdit={handleBatchEdit}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加套餐</Button>
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
        title={editingId ? '编辑充值套餐' : '添加充值套餐'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="name" rules={[{ required: true, message: '请输入套餐名称' }]}>
            <Input placeholder="套餐名称" />
          </Form.Item>
          <Form.Item name="amount" rules={[{ required: true, message: '请输入充值金额' }]}>
            <InputNumber placeholder="充值金额" prefix="¥" min={0} />
          </Form.Item>
          <Form.Item name="gift_amount">
            <InputNumber placeholder="赠送金额" prefix="¥" min={0} />
          </Form.Item>
          <Form.Item name="gift_points">
            <InputNumber placeholder="赠送积分" min={0} />
          </Form.Item>
          <Form.Item name="description">
            <Input.TextArea placeholder="套餐描述" />
          </Form.Item>
          <Form.Item name="status">
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
        records={selectedPackages}
        fields={[
          { key: 'name', label: '套餐名称', type: 'input' },
          { key: 'amount', label: '充值金额', type: 'number' },
          { key: 'bonus', label: '赠送金额', type: 'number' },
          { key: 'description', label: '套餐描述', type: 'textarea' },
          { key: 'status', label: '状态', type: 'select', options: [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }] }
        ]}
        title="批量编辑充值套餐"
      />
    </div>
  )
}