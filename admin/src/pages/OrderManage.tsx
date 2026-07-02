import { useState, useEffect } from 'react'
import { Table, Button, Modal, message } from 'antd'
import { EyeOutlined, CheckOutlined, StopOutlined, DeleteOutlined } from '@ant-design/icons'
import { orderApi } from '@/api'
import { Order } from '@/types'

export default function OrderManage() {
  const [data, setData] = useState<Order[]>([])
  const [open, setOpen] = useState(false)
  const [detail, setDetail] = useState<Order | null>(null)

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    try {
      const result = await orderApi.getList()
      setData(result)
    } catch (error) {
      console.error('Failed to fetch orders:', error)
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

  const getStatusText = (status: number) => {
    const map: Record<number, string> = {
      0: '待支付', 1: '使用中', 2: '已完成', 3: '已取消', 4: '退款中', 5: '已退款'
    }
    return map[status] || '未知'
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

  const columns = [
    { title: '订单号', dataIndex: 'order_no', key: 'order_no' },
    { title: '用户', dataIndex: 'user', key: 'user', render: (u: any) => u?.nickname || '-' },
    { title: '包间', dataIndex: 'room', key: 'room', render: (r: { name: string }) => r.name },
    { title: '开始时间', dataIndex: 'start_time', key: 'start_time' },
    { title: '结束时间', dataIndex: 'end_time', key: 'end_time' },
    { title: '时长', dataIndex: 'duration', key: 'duration', render: (d: number) => `${d}分钟` },
    { title: '金额', dataIndex: 'total_amount', key: 'total_amount', render: (v: number) => `¥${v}` },
    { title: '已支付', dataIndex: 'paid_amount', key: 'paid_amount', render: (v: number) => `¥${v}` },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getStatusText(s) },
    { title: '操作', key: 'action', render: (_: any, record: Order) => (
      <div>
        <Button type="text" icon={<EyeOutlined />} onClick={() => handleView(record.id)}>查看</Button>
        {record.status === 0 && (
          <Button type="text" icon={<CheckOutlined />} onClick={() => handleConfirm(record.id)}>确认</Button>
        )}
        {record.status === 1 && (
          <Button type="text" icon={<StopOutlined />} onClick={() => handleComplete(record.id)}>完成</Button>
        )}
        <Button type="text" danger icon={<DeleteOutlined />} onClick={() => handleDelete(record.id)}>删除</Button>
      </div>
    )}
  ]

  return (
    <div>
      <h2>订单管理</h2>
      <Table dataSource={data} columns={columns} rowKey="id" />

      <Modal
        title="订单详情"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
      >
        {detail && (
          <div>
            <p><strong>订单号:</strong> {detail.order_no}</p>
            <p><strong>包间:</strong> {detail.room.name}</p>
            <p><strong>类型:</strong> {detail.room.type.name}</p>
            <p><strong>开始时间:</strong> {detail.start_time}</p>
            <p><strong>结束时间:</strong> {detail.end_time}</p>
            <p><strong>时长:</strong> {detail.duration}分钟</p>
            <p><strong>金额:</strong> ¥{detail.total_amount}</p>
            <p><strong>已支付:</strong> ¥{detail.paid_amount}</p>
            <p><strong>状态:</strong> {getStatusText(detail.status)}</p>
            <p><strong>备注:</strong> {detail.remark || '-'}</p>
          </div>
        )}
      </Modal>
    </div>
  )
}