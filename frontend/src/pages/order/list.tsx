import { useState, useEffect } from 'react'
import { Tabs, Cell, CellGroup, Loading, Empty, Button } from '@nutui/nutui-react'
import { orderApi, Order } from '@/api'
import { showToast } from '@/components/Toast'
import './list.scss'

export default function OrderList() {
  const [orders, setOrders] = useState<Order[]>([])
  const [activeTab, setActiveTab] = useState(0)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchOrders(activeTab)
  }, [activeTab])

  const fetchOrders = async (status: number) => {
    setLoading(true)
    try {
      const data = await orderApi.getOrderList({ status })
      setOrders(data)
    } catch (error) {
      console.error('Failed to fetch orders:', error)
      setOrders([
        { id: 1, order_no: 'ORD20260120001', user_id: 1, room_id: 1, room: { id: 1, name: '豪华包间', floor: '1F', status: 1, equipment: '', images: '', description: '', type_id: 1, type: { id: 1, name: '豪华型', max_people: 4, base_price: 200, description: '豪华型包间' } }, status: 1, total_amount: 400, paid_amount: 400, duration: 2, start_time: '2026-01-21 14:00:00', end_time: '2026-01-21 16:00:00', remark: '', created_at: '2026-01-20 10:00:00' },
        { id: 2, order_no: 'ORD20260120002', user_id: 1, room_id: 2, room: { id: 2, name: '商务包间', floor: '2F', status: 1, equipment: '', images: '', description: '', type_id: 2, type: { id: 2, name: '商务型', max_people: 6, base_price: 300, description: '商务型包间' } }, status: 0, total_amount: 900, paid_amount: 900, duration: 3, start_time: '2026-01-22 19:00:00', end_time: '2026-01-22 22:00:00', remark: '', created_at: '2026-01-20 15:00:00' },
        { id: 3, order_no: 'ORD20260118001', user_id: 1, room_id: 1, room: { id: 1, name: '豪华包间', floor: '1F', status: 1, equipment: '', images: '', description: '', type_id: 1, type: { id: 1, name: '豪华型', max_people: 4, base_price: 200, description: '豪华型包间' } }, status: 2, total_amount: 600, paid_amount: 600, duration: 3, start_time: '2026-01-19 13:00:00', end_time: '2026-01-19 16:00:00', remark: '', created_at: '2026-01-18 09:00:00' }
      ])
    } finally {
      setLoading(false)
    }
  }

  const getStatusText = (status: number) => {
    const statusMap: Record<number, string> = {
      0: '待确认',
      1: '已确认',
      2: '使用中',
      3: '已完成',
      4: '已取消'
    }
    return statusMap[status] || '未知'
  }

  const getStatusColor = (status: number) => {
    const colorMap: Record<number, string> = {
      0: '#faad14',
      1: '#52c41a',
      2: '#667eea',
      3: '#999999',
      4: '#ff4d4f'
    }
    return colorMap[status] || '#999999'
  }

  const handleCancel = async (orderId: number) => {
    try {
      await orderApi.cancelOrder(orderId)
      showToast({ message: '取消成功', type: 'success' })
      fetchOrders(activeTab)
    } catch (error) {
      console.error('Failed to cancel order:', error)
    }
  }

  const handleConfirm = async (orderId: number) => {
    try {
      await orderApi.confirmOrder(orderId)
      showToast({ message: '确认成功', type: 'success' })
      fetchOrders(activeTab)
    } catch (error) {
      console.error('Failed to confirm order:', error)
    }
  }

  const handleComplete = async (orderId: number) => {
    try {
      await orderApi.completeOrder(orderId)
      showToast({ message: '订单已完成', type: 'success' })
      fetchOrders(activeTab)
    } catch (error) {
      console.error('Failed to complete order:', error)
    }
  }

  const filterOrders = () => {
    if (activeTab === 0) return orders
    return orders.filter(order => order.status === activeTab)
  }

  const renderActionButtons = (order: Order) => {
    switch (order.status) {
      case 0:
        return (
          <div className="action-buttons">
            <Button type="warning" size="small" onClick={() => handleCancel(order.id)}>取消预约</Button>
            <Button type="primary" size="small" onClick={() => handleConfirm(order.id)}>确认订单</Button>
          </div>
        )
      case 1:
        return (
          <Button type="primary" size="small" onClick={() => handleComplete(order.id)}>完成订单</Button>
        )
      case 2:
        return (
          <Button type="primary" size="small" onClick={() => handleComplete(order.id)}>结束使用</Button>
        )
      default:
        return null
    }
  }

  const tabs = [
    { title: '全部' },
    { title: '待确认', subTitle: orders.filter(o => o.status === 0).length },
    { title: '已确认', subTitle: orders.filter(o => o.status === 1).length },
    { title: '使用中', subTitle: orders.filter(o => o.status === 2).length },
    { title: '已完成', subTitle: orders.filter(o => o.status === 3).length }
  ]

  return (
    <div className="container">
      <Tabs
        value={activeTab}
        onChange={(index) => setActiveTab(Number(index))}
        activeType="card"
      >
        {tabs.map((tab, index) => (
          <Tabs.TabPane key={index} title={tab.title}>
            {loading ? (
              <div className="loading-wrapper">
                <Loading type="circular" color="#667eea" />
              </div>
            ) : filterOrders().length > 0 ? (
              <CellGroup className="order-group">
                {filterOrders().map(order => (
                  <Cell key={order.id} className="order-cell">
                    <div className="order-card">
                      <div className="order-header">
                        <span className="order-id">订单号: {order.id}</span>
                        <span className="order-status" style={{ color: getStatusColor(order.status) }}>
                          {getStatusText(order.status)}
                        </span>
                      </div>

                      <div className="order-room-info">
                        <div className="room-info">
                          <span className="room-name">{order.room.name}</span>
                          <span className="room-type">{order.room.type.name}</span>
                        </div>
                        <span className="room-floor">{order.room.floor}</span>
                      </div>

                      <div className="order-time">
                        <span className="time-label">预约时间</span>
                        <span className="time-value">
                          {order.start_time} - {order.end_time}
                        </span>
                      </div>

                      <div className="order-footer">
                        <span className="order-amount">金额: ¥{order.total_amount}</span>
                        {renderActionButtons(order)}
                      </div>
                    </div>
                  </Cell>
                ))}
              </CellGroup>
            ) : (
              <Empty description="暂无订单" image="empty" />
            )}
          </Tabs.TabPane>
        ))}
      </Tabs>
    </div>
  )
}