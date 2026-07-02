import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Button, Cell, CellGroup, Loading, Badge, Divider, Dialog } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { orderApi, Order } from '@/api'

export default function OrderDetail() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const [order, setOrder] = useState<Order | null>(null)
  const [showCancelDialog, setShowCancelDialog] = useState(false)

  useEffect(() => {
    if (id) {
      fetchOrderDetail(parseInt(id))
    }
  }, [id])

  const fetchOrderDetail = async (id: number) => {
    try {
      const data = await orderApi.getOrderDetail(id)
      setOrder(data)
    } catch (error) {
      console.error('Failed to fetch order detail:', error)
    }
  }

  const handleCancelOrder = async () => {
    if (!order) return
    try {
      await orderApi.cancelOrder(order.id)
      showToast({ message: '取消成功', type: 'success' })
      setTimeout(() => {
        navigate('/orders')
      }, 1500)
    } catch (error) {
      console.error('Failed to cancel order:', error)
    }
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

  if (!order) {
    return (
      <Loading type="circular" color="#667eea" style={{ marginTop: '100rpx' }} />
    )
  }

  return (
    <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 100px)' }}>
      <div className="status-bar" style={{ background: getStatusColor(order.status) + '20' }}>
        <Badge value={getStatusText(order.status)} color={getStatusColor(order.status)} style={{ fontSize: '20px', padding: '4px 12px' }} />
      </div>

      <CellGroup>
        <Cell title="订单号" description={order.order_no} />
        <Cell title="下单时间" description={order.created_at} />

        <Cell title="包间信息">
          <div className="room-info">
            <div className="room-name">{order.room.name}</div>
            <div className="room-type">{order.room.type.name}</div>
            <div className="room-floor">楼层: {order.room.floor}</div>
          </div>
        </Cell>

        <Cell title="预约时间">
          <div className="time-info">
            <div className="time-row">
              <span className="time-label">开始时间</span>
              <span className="time-value">{order.start_time}</span>
            </div>
            <div className="time-row">
              <span className="time-label">结束时间</span>
              <span className="time-value">{order.end_time}</span>
            </div>
            <div className="time-row">
              <span className="time-label">时长</span>
              <span className="time-value">{order.duration}分钟</span>
            </div>
          </div>
        </Cell>

        <Cell title="费用明细">
          <div className="price-info">
            <div className="price-row">
              <span className="price-label">订单金额</span>
              <span className="price-value">¥{order.total_amount}</span>
            </div>
            <div className="price-row">
              <span className="price-label">已支付</span>
              <span className="price-value">¥{order.paid_amount}</span>
            </div>
            <div className="price-row total">
              <span className="price-label">待支付</span>
              <span className="price-value">¥{(order.total_amount - order.paid_amount).toFixed(2)}</span>
            </div>
          </div>
        </Cell>

        {order.remark && (
          <Cell title="备注" description={order.remark} />
        )}
      </CellGroup>

      <Divider />

      <div className="bottom-bar">
        {order.status === 0 && (
          <>
            <Button type="default" block onClick={() => setShowCancelDialog(true)}>
              取消订单
            </Button>
            <Button type="primary" block>
              立即支付
            </Button>
          </>
        )}
        {order.status === 1 && (
          <Button type="primary" block>
            确认结束
          </Button>
        )}
      </div>

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
