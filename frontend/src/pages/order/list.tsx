import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Tabs, Cell, CellGroup, Empty, Badge, Dialog } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { orderApi, Order } from '@/api'
import CustomTabBar from '@/components/CustomTabBar'

export default function OrderList() {
  const navigate = useNavigate()
  const [orders, setOrders] = useState<Order[]>([])
  const [activeTab, setActiveTab] = useState(0)
  const [showCancelDialog, setShowCancelDialog] = useState(false)
  const [cancelOrderId, setCancelOrderId] = useState<number | null>(null)

  useEffect(() => {
    fetchOrders()
  }, [activeTab])

  const fetchOrders = async () => {
    try {
      const status = activeTab === 0 ? 0 : activeTab === 1 ? 1 : activeTab === 2 ? 2 : 3
      const data = await orderApi.getOrderList(activeTab === 0 ? {} : { status })
      setOrders(data)
    } catch (error) {
      console.error('Failed to fetch orders:', error)
    }
  }

  const handleOrderClick = (order: Order) => {
    navigate(`/orders/${order.id}`)
  }

  const handleCancelOrder = async () => {
    if (!cancelOrderId) return
    try {
      await orderApi.cancelOrder(cancelOrderId)
      showToast({ message: '取消成功', type: 'success' })
      setShowCancelDialog(false)
      fetchOrders()
    } catch (error) {
      console.error('Failed to cancel order:', error)
    }
  }

  const confirmCancelOrder = (orderId: number) => {
    setCancelOrderId(orderId)
    setShowCancelDialog(true)
  }

  const getStatusText = (status: number) => {
    const statusMap: Record<number, string> = {
      0: '待支付',
      1: '使用中',
      2: '已完成',
      3: '已取消',
      4: '退款中',
      5: '已退款'
    }
    return statusMap[status] || '未知'
  }

  const getStatusColor = (status: number) => {
    const colorMap: Record<number, string> = {
      0: '#faad14',
      1: '#52c41a',
      2: '#999999',
      3: '#999999',
      4: '#faad14',
      5: '#999999'
    }
    return colorMap[status] || '#999999'
  }

  const tabs = [
    { label: '全部', value: 0 },
    { label: '待支付', value: 1 },
    { label: '使用中', value: 2 },
    { label: '已完成', value: 3 }
  ]

  return (
    <div className="page">
      <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 120px)' }}>
        <Tabs
          value={activeTab}
          onChange={(value) => setActiveTab(Number(value))}
          activeType="card"
          activeColor="#667eea"
        >
          {tabs.map(tab => (
            <Tabs.TabPane key={tab.value} title={tab.label} />
          ))}
        </Tabs>

        <CellGroup>
          {orders.map(order => (
            <Cell
              key={order.id}
              onClick={() => handleOrderClick(order)}
              className="order-cell"
            >
              <div className="order-card">
                <div className="order-header">
                  <span className="order-no">订单号: {order.order_no}</span>
                  <Badge value={getStatusText(order.status)} color={getStatusColor(order.status)} />
                </div>

                <div className="order-room">
                  <span className="room-name">{order.room.name}</span>
                  <span className="room-type">{order.room.type.name}</span>
                </div>

                <div className="order-time">
                  <span className="time-icon">📅</span>
                  <span>{order.start_time} - {order.end_time}</span>
                </div>

                <div className="order-footer">
                  <span className="order-price">¥{order.total_amount}</span>
                  {order.status === 0 && (
                    <div className="order-actions">
                      <Button type="default" size="small" onClick={(e) => { e.stopPropagation(); confirmCancelOrder(order.id) }}>
                        取消
                      </Button>
                      <Button type="primary" size="small">
                        支付
                      </Button>
                    </div>
                  )}
                </div>
              </div>
            </Cell>
          ))}
        </CellGroup>

        {orders.length === 0 && (
          <Empty description="暂无订单" image="empty" />
        )}
      </div>
      <CustomTabBar />

      <Dialog
        visible={showCancelDialog}
        title="取消订单"
        content="确定要取消这个订单吗？"
        onCancel={() => setShowCancelDialog(false)}
        onConfirm={handleCancelOrder}
      />
    </div>
  )
}
