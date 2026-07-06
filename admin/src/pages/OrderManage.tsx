import { useState, useEffect } from 'react'
import { Table, Button, Modal, Form, DatePicker, InputNumber, Input, message, Card, Tag, Select, Checkbox } from 'antd'
import { EyeOutlined, EditOutlined, CheckOutlined, StopOutlined, DeleteOutlined, PlusOutlined, SearchOutlined } from '@ant-design/icons'
import dayjs from 'dayjs'
import { orderApi, roomApi, userApi } from '@/api'
import { Order, Room, User } from '@/types'
import { usePermission } from '@/context/PermissionContext'
import SearchBar from '@/components/SearchBar'
import BatchActions from '@/components/BatchActions'
import CustomPagination from '@/components/CustomPagination'
import ExportDropdown from '@/components/ExportDropdown'
import BatchEditModal from '@/components/BatchEditModal'
import { formatDateTime } from '@/utils'

export default function OrderManage() {
  const [data, setData] = useState<Order[]>([])
  const [rooms, setRooms] = useState<Room[]>([])
  const [users, setUsers] = useState<User[]>([])
  const [open, setOpen] = useState(false)
  const [editOpen, setEditOpen] = useState(false)
  const [addOpen, setAddOpen] = useState(false)
  const [detail, setDetail] = useState<Order | null>(null)
  const [editingId, setEditingId] = useState<number | null>(null)
  const [searchVisible, setSearchVisible] = useState(false)
  const [searchParams, setSearchParams] = useState<Record<string, string>>({})
  const [selectedRowKeys, setSelectedRowKeys] = useState<string[]>([])
  const [currentPage, setCurrentPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [total, setTotal] = useState(0)
  const [batchEditVisible, setBatchEditVisible] = useState(false)
  const [selectedOrders, setSelectedOrders] = useState<Order[]>([])
  const [form] = Form.useForm()
  const [addForm] = Form.useForm()
  const { hasPermission } = usePermission()

  useEffect(() => {
    fetchData()
    fetchRooms()
    fetchUsers()
  }, [])

  const fetchData = async (params?: Record<string, string>, page: number = currentPage, size: number = pageSize) => {
    try {
      const mergedParams = { ...searchParams, ...params }
      const result = await orderApi.getList({ ...mergedParams, page: String(page), page_size: String(size) })
      if (Array.isArray(result)) {
        setData(result)
        setTotal(result.length)
      } else if (result.data) {
        setData(result.data)
        setTotal(result.total || 0)
      }
    } catch (error) {
      console.error('Failed to fetch orders:', error)
    }
  }

  const fetchRooms = async () => {
    try {
      const result = await roomApi.getList()
      setRooms(Array.isArray(result) ? result : (result as { data: Room[] }).data)
    } catch (error) {
      console.error('Failed to fetch rooms:', error)
    }
  }

  const fetchUsers = async () => {
    try {
      const result = await userApi.getList()
      setUsers(Array.isArray(result) ? result : (result as { data: User[] }).data)
    } catch (error) {
      console.error('Failed to fetch users:', error)
    }
  }

  const handleView = async (id: number) => {
    try {
      const result = await orderApi.getDetail(id)
      setDetail(result)
      setOpen(true)
    } catch (error) {
      console.error('Failed to fetch order:', error)
    }
  }

  const handleEdit = (record: Order) => {
    setEditingId(record.id)
    form.setFieldsValue({
      start_time: record.start_time ? dayjs(record.start_time) : null,
      end_time: record.end_time ? dayjs(record.end_time) : null,
      duration: record.duration,
      remark: record.remark,
      status: record.status
    })
    setEditOpen(true)
  }

  const handleEditSubmit = async (values: any) => {
    try {
      if (!editingId) return
      const data = {
        ...values,
        start_time: values.start_time ? values.start_time.format('YYYY-MM-DD HH:mm:ss') : null,
        end_time: values.end_time ? values.end_time.format('YYYY-MM-DD HH:mm:ss') : null
      }
      await orderApi.update(editingId, data)
      message.success('更新成功')
      setEditOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to update order:', error)
    }
  }

  const getStatusText = (status: number) => {
    const map: Record<number, string> = {
      0: '待支付', 1: '使用中', 2: '已完成', 3: '已取消', 4: '退款中', 5: '已退款'
    }
    return map[status] || '未知'
  }

  const getStatusTag = (status: number) => {
    const colors: Record<number, string> = {
      0: 'orange',
      1: 'blue',
      2: 'green',
      3: 'default',
      4: 'red',
      5: 'purple'
    }
    return <Tag color={colors[status]}>{getStatusText(status)}</Tag>
  }

  const handleConfirm = async (id: number) => {
    try {
      await orderApi.confirm(id)
      message.success('确认成功')
      fetchData()
    } catch (error) {
      console.error('Failed to confirm order:', error)
    }
  }

  const handleComplete = async (id: number) => {
    try {
      await orderApi.complete(id)
      message.success('完成成功')
      fetchData()
    } catch (error) {
      console.error('Failed to complete order:', error)
    }
  }

  const handleDelete = async (id: number) => {
    try {
      await orderApi.delete(id)
      message.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete order:', error)
    }
  }

  const handleBatchEdit = (ids: string[]) => {
    const selected = data.filter(o => ids.includes(String(o.id)))
    setSelectedOrders(selected)
    setBatchEditVisible(true)
  }

  const handleBatchEditSubmit = async (updatedRecords: Record<string, any>[]) => {
    try {
      const items = updatedRecords.map(record => ({
        ...record,
        start_time: record.start_time ? record.start_time.format('YYYY-MM-DD HH:mm:ss') : undefined,
        end_time: record.end_time ? record.end_time.format('YYYY-MM-DD HH:mm:ss') : undefined
      }))
      await orderApi.batchUpdate(items)
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
      await orderApi.batchDelete(ids.map(id => parseInt(id)))
      message.success('批量删除成功')
      fetchData()
      setSelectedRowKeys([])
    } catch (error) {
      message.error('批量删除失败')
    }
  }

  const handleAdd = () => {
    addForm.resetFields()
    setAddOpen(true)
  }

  const handleAddSubmit = async (values: any) => {
    try {
      const data = {
        ...values,
        start_time: values.start_time ? values.start_time.format('YYYY-MM-DD HH:mm:ss') : null,
        end_time: values.end_time ? values.end_time.format('YYYY-MM-DD HH:mm:ss') : null
      }
      await orderApi.create(data)
      message.success('添加成功')
      setAddOpen(false)
      fetchData()
    } catch (error) {
      console.error('Failed to create order:', error)
    }
  }

  const columns = [
    {
      title: '选择',
      key: 'selection',
      width: 60,
      render: (_: any, record: Order) => (
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
    { title: '订单号', dataIndex: 'order_no', key: 'order_no' },
    { title: '用户', dataIndex: 'user', key: 'user', width: 60, render: (u: any) => u?.nickname || '-' },
    { title: '包间', dataIndex: 'room', key: 'room', width: 60, render: (r: { name: string }) => r.name },
    { title: '开始时间', dataIndex: 'start_time', key: 'start_time', render: (t: string) => formatDateTime(t) },
    { title: '结束时间', dataIndex: 'end_time', key: 'end_time', render: (t: string) => formatDateTime(t) },
    { title: '时长', dataIndex: 'duration', key: 'duration', render: (d: number) => `${d}分钟` },
    { title: '金额', dataIndex: 'total_amount', key: 'total_amount', render: (v: number) => <span style={{ color: '#ff4d4f', fontWeight: 'bold' }}>¥{v}</span> },
    { title: '已支付', dataIndex: 'paid_amount', key: 'paid_amount', render: (v: number) => `¥${v}` },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusTag(s) },
    { title: '操作', key: 'action', render: (_: any, record: Order) => (
      <div style={{ display: 'flex', gap: '8px', flexWrap: 'wrap' }}>
        {hasPermission('order_view') && (
          <Button size="small" icon={<EyeOutlined />} onClick={() => handleView(record.id)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>查看</Button>
        )}
        {hasPermission('order_edit') && (
          <Button size="small" icon={<EditOutlined />} onClick={() => handleEdit(record)} style={{ backgroundColor: '#52c41a', color: '#fff', borderColor: '#52c41a' }}>编辑</Button>
        )}
        {hasPermission('order_confirm') && record.status === 0 && (
          <Button size="small" icon={<CheckOutlined />} onClick={() => handleConfirm(record.id)} style={{ backgroundColor: '#722ed1', color: '#fff', borderColor: '#722ed1' }}>确认</Button>
        )}
        {hasPermission('order_confirm') && record.status === 1 && (
          <Button size="small" icon={<StopOutlined />} onClick={() => handleComplete(record.id)} style={{ backgroundColor: '#13c2c2', color: '#fff', borderColor: '#13c2c2' }}>完成</Button>
        )}
        {hasPermission('order_delete') && (
          <Button size="small" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
        )}
      </div>
    )}
  ]

  const handleSearch = (values: Record<string, string>) => {
    const params: Record<string, string> = {}
    Object.keys(values).forEach(key => {
      if (values[key]) {
        params[key] = values[key]
      }
    })
    setSearchParams(params)
    setCurrentPage(1)
    fetchData(params, 1)
  }

  return (
    <div>
      {searchVisible && (
        <div style={{ marginBottom: 16, width: '100%' }}>
          <SearchBar
            fields={[
              { key: 'order_no', label: '订单号', type: 'input', placeholder: '请输入订单号' },
              { key: 'user_id', label: '用户', type: 'select', options: users.map(u => ({ label: u.nickname || u.realname, value: String(u.id) })) },
              { key: 'room_id', label: '包间', type: 'select', options: rooms.map(r => ({ label: r.name, value: String(r.id) })) },
              { key: 'status', label: '状态', type: 'select', options: [
                { label: '待支付', value: '0' },
                { label: '使用中', value: '1' },
                { label: '已完成', value: '2' },
                { label: '已取消', value: '3' },
                { label: '退款中', value: '4' },
                { label: '已退款', value: '5' }
              ]}
            ]}
            onSearch={handleSearch}
          />
        </div>
      )}
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>订单管理</h2>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button type="link" icon={<SearchOutlined />} onClick={() => setSearchVisible(!searchVisible)} style={{ backgroundColor: '#1890ff', color: '#fff', borderColor: '#1890ff' }}>
            {searchVisible ? '收起搜索' : '搜索'}
          </Button>
          <ExportDropdown data={data} filename="订单管理数据" />
          <BatchActions
            selectedRowKeys={selectedRowKeys}
            onBatchEdit={handleBatchEdit}
            onBatchDelete={handleBatchDelete}
          />
          {hasPermission('order_create') && (
            <Button type="primary" icon={<PlusOutlined />} onClick={handleAdd}>添加订单</Button>
          )}
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
        title="订单详情"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        width={500}
      >
        {detail && (
          <div>
            <Card style={{ marginBottom: 16 }}>
              <div style={{ fontSize: '20px', fontWeight: 'bold', color: '#1890ff', marginBottom: 16 }}>
                {detail.order_no}
              </div>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <div>
                  <div style={{ color: '#666', fontSize: '14px' }}>用户</div>
                  <div style={{ fontWeight: 'bold', fontSize: '16px' }}>{detail.user?.nickname || '-'}</div>
                </div>
                <div style={{ textAlign: 'right' }}>
                  <div style={{ color: '#666', fontSize: '14px' }}>状态</div>
                  {getStatusTag(detail.status)}
                </div>
              </div>
            </Card>
            <Card title="预约信息" style={{ marginBottom: 16 }}>
              <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '16px' }}>
                <div>
                  <div style={{ color: '#666', fontSize: '12px' }}>包间名称</div>
                  <div style={{ fontWeight: '500' }}>{detail.room.name}</div>
                </div>
                <div>
                  <div style={{ color: '#666', fontSize: '12px' }}>包间类型</div>
                  <div style={{ fontWeight: '500' }}>{detail.room.type?.name || '-'}</div>
                </div>
                <div>
                  <div style={{ color: '#666', fontSize: '12px' }}>开始时间</div>
                  <div style={{ fontWeight: '500' }}>{detail.start_time}</div>
                </div>
                <div>
                  <div style={{ color: '#666', fontSize: '12px' }}>结束时间</div>
                  <div style={{ fontWeight: '500' }}>{detail.end_time}</div>
                </div>
                <div>
                  <div style={{ color: '#666', fontSize: '12px' }}>时长</div>
                  <div style={{ fontWeight: '500' }}>{detail.duration}分钟</div>
                </div>
              </div>
            </Card>
            <Card title="支付信息" style={{ marginBottom: 16 }}>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <div>
                  <div style={{ color: '#666', fontSize: '12px' }}>订单金额</div>
                  <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#ff4d4f' }}>¥{detail.total_amount}</div>
                </div>
                <div style={{ textAlign: 'right' }}>
                  <div style={{ color: '#666', fontSize: '12px' }}>已支付</div>
                  <div style={{ fontSize: '18px', fontWeight: 'bold', color: '#52c41a' }}>¥{detail.paid_amount}</div>
                </div>
              </div>
            </Card>
            {detail.remark && (
              <Card title="备注信息">
                <div style={{ color: '#333' }}>{detail.remark}</div>
              </Card>
            )}
          </div>
        )}
      </Modal>

      <Modal
        title="编辑订单"
        open={editOpen}
        onCancel={() => setEditOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={form} onFinish={handleEditSubmit} layout="vertical">
          <Form.Item name="start_time" label="开始时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="end_time" label="结束时间">
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="duration" label="时长(分钟)">
            <InputNumber min={0} style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="status" label="状态">
            <select style={{ width: '100%', height: 32, padding: '0 8px', borderRadius: 4, border: '1px solid #d9d9d9' }}>
              <option value={0}>待支付</option>
              <option value={1}>使用中</option>
              <option value={2}>已完成</option>
              <option value={3}>已取消</option>
              <option value={4}>退款中</option>
              <option value={5}>已退款</option>
            </select>
          </Form.Item>
          <Form.Item name="remark" label="备注">
            <Input.TextArea rows={3} />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">保存</Button>
            <Button onClick={() => setEditOpen(false)} style={{ marginLeft: 8 }}>取消</Button>
          </Form.Item>
        </Form>
      </Modal>

      <Modal
        title="添加订单"
        open={addOpen}
        onCancel={() => setAddOpen(false)}
        footer={null}
        width={500}
      >
        <Form form={addForm} onFinish={handleAddSubmit} layout="vertical">
          <Form.Item name="user_id" label="用户" rules={[{ required: true, message: '请选择用户' }]}>
            <Select style={{ width: '100%' }}>
              {users.map(user => (
                <Select.Option key={user.id} value={user.id}>{user.nickname || user.realname}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="room_id" label="包间" rules={[{ required: true, message: '请选择包间' }]}>
            <Select style={{ width: '100%' }}>
              {rooms.map(room => (
                <Select.Option key={room.id} value={room.id}>{room.name}</Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="start_time" label="开始时间" rules={[{ required: true, message: '请选择开始时间' }]}>
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="end_time" label="结束时间" rules={[{ required: true, message: '请选择结束时间' }]}>
            <DatePicker showTime format="YYYY-MM-DD HH:mm:ss" style={{ width: '100%' }} />
          </Form.Item>
          <Form.Item name="remark" label="备注">
            <Input.TextArea rows={3} />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">保存</Button>
            <Button onClick={() => setAddOpen(false)} style={{ marginLeft: 8 }}>取消</Button>
          </Form.Item>
        </Form>
      </Modal>
      <BatchEditModal
        visible={batchEditVisible}
        onCancel={() => setBatchEditVisible(false)}
        onOk={handleBatchEditSubmit}
        records={selectedOrders}
        fields={[
          { key: 'start_time', label: '开始时间', type: 'date' },
          { key: 'end_time', label: '结束时间', type: 'date' },
          { key: 'duration', label: '时长(分钟)', type: 'number' },
          { key: 'remark', label: '备注', type: 'textarea' },
          { key: 'status', label: '状态', type: 'select', options: [{ label: '待支付', value: 0 }, { label: '使用中', value: 1 }, { label: '已完成', value: 2 }, { label: '已取消', value: 3 }, { label: '退款中', value: 4 }, { label: '已退款', value: 5 }] }
        ]}
        title="批量编辑订单"
      />
    </div>
  )
}