import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, Input, InputNumber, message, Checkbox } from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined } from '@ant-design/icons'
import { roomTypeApi } from '@/api'
import { RoomType } from '@/types'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'

export default function RoomTypeManage() {
  const [data, setData] = useState<RoomType[]>([])
  const [open, setOpen] = useState(false)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedRoomTypes, setSelectedRoomTypes] = useState<RoomType[]>([])
  const [form] = Form.useForm()

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const result = await roomTypeApi.getList({ ...params, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch room types:', error)
    }
  }

  const handleAdd = () => {
    setEditingId(null)
    form.resetFields()
    setOpen(true)
  }

  const handleEdit = (record: RoomType) => {
    setEditingId(record.id)
    form.setFieldsValue(record)
    setOpen(true)
  }

  const handleDelete = async (id: number) => {
    try {
      await roomTypeApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(r => ids.includes(String(r.id)))
    setSelectedRoomTypes(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      await roomTypeApi.batchUpdate(updatedRecords)
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
      await roomTypeApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleSubmit = async (values: Partial<RoomType>) => {
    try {
      if (editingId) {
        await roomTypeApi.update(editingId, values)
        message.success('更新成功')
      } else {
        await roomTypeApi.create(values)
        message.success('创建成功')
      }
      setOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to submit:', error)
    }
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: RoomType) => (
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
    { title: '名称', dataIndex: 'name', key: 'name' },
    { title: '描述', dataIndex: 'description', key: 'description' },
    { title: '基础价格', dataIndex: 'base_price', key: 'base_price', render: (v: number) => `¥${v}/小时` },
    { title: '最大人数', dataIndex: 'max_people', key: 'max_people' },
    { title: '操作', key: 'action', render: (_: any, record: RoomType) => (
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
              { key: 'name', label: '名称', type: 'input', placeholder: '请输入包间类型名称' }
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>包间类型管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="包间类型管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchEdit={handleBatchEdit}
            onBatchDelete={handleBatchDelete}
          />
          <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加类型</Button>
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
        title={editingId ? '编辑包间类型' : '添加包间类型'}
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit}>
          <Form.Item name="name" rules={[{ required: true, message: '请输入名称' }]}>
            <Input placeholder="名称" />
          </Form.Item>
          <Form.Item name="description">
            <Input placeholder="描述" />
          </Form.Item>
          <Form.Item name="base_price" rules={[{ required: true, message: '请输入基础价格' }]}>
            <InputNumber placeholder="基础价格" />
          </Form.Item>
          <Form.Item name="max_people" rules={[{ required: true, message: '请输入最大人数' }]}>
            <InputNumber placeholder="最大人数" />
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
        records={selectedRoomTypes}
        fields={[
          { key: 'name', label: '名称', type: 'input' },
          { key: 'description', label: '描述', type: 'input' },
          { key: 'base_price', label: '基础价格', type: 'number' },
          { key: 'max_people', label: '最大人数', type: 'number' }
        ]}
        title="批量编辑包间类型"
      />
    </div>
  )
}