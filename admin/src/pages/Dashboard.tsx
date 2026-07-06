import { useState, useEffect } from 'react'
import { Card, Row, Col, Statistic, Table, Tag, Progress, Spin } from 'antd'
import {
  HomeOutlined,
  ShoppingCartOutlined,
  UserOutlined,
  DollarOutlined,
  ClockCircleOutlined,
  CheckCircleOutlined
} from '@ant-design/icons'
import { roomApi, orderApi, membershipApi } from '@/api'
import { Room, Order, Membership } from '@/types'
import { formatDateTime } from '@/utils'

export default function Dashboard() {
  const [stats, setStats] = useState({
    totalRooms: 0,
    activeRooms: 0,
    availableRooms: 0,
    todayOrders: 0,
    totalOrders: 0,
    completedOrders: 0,
    totalMembers: 0,
    todayRevenue: 0,
    totalRevenue: 0,
    activeMembers: 0
  })
  const [recentOrders, setRecentOrders] = useState<Order[]>([])
  const [roomStatus, setRoomStatus] = useState<Room[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchStats()
  }, [])

  const isToday = (dateStr: string) => {
    if (!dateStr) return false
    const date = new Date(dateStr)
    const today = new Date()
    return date.toDateString() === today.toDateString()
  }

  const fetchStats = async () => {
    try {
      setLoading(true)
      const [roomsRes, ordersRes, membersRes] = await Promise.all([
        roomApi.getList({ page: '0', page_size: '0' }),
        orderApi.getList({ page: '0', page_size: '0' }),
        membershipApi.getList({ page: '0', page_size: '0' })
      ])

      const rooms = Array.isArray(roomsRes) ? roomsRes : (roomsRes as { data: Room[] }).data
      const orders = Array.isArray(ordersRes) ? ordersRes : (ordersRes as { data: Order[] }).data
      const members = Array.isArray(membersRes) ? membersRes : (membersRes as { data: Membership[] }).data

      const activeRooms = rooms.filter((r: Room) => r.status === 2).length
      const availableRooms = rooms.filter((r: Room) => r.status === 1).length
      const todayOrders = orders.filter((o: Order) => isToday(o.created_at)).length
      const completedOrders = orders.filter((o: Order) => o.status === 2).length
      const todayRevenue = orders.filter((o: Order) => isToday(o.created_at)).reduce((sum: number, o: Order) => sum + o.paid_amount, 0)
      const totalRevenue = orders.reduce((sum: number, o: Order) => sum + o.paid_amount, 0)
      const activeMembers = members.filter((m: Membership) => m.membership_status === 1).length

      setStats({
        totalRooms: rooms.length,
        activeRooms,
        availableRooms,
        todayOrders,
        totalOrders: orders.length,
        completedOrders,
        totalMembers: members.length,
        todayRevenue,
        totalRevenue,
        activeMembers
      })

      setRecentOrders(orders.slice(0, 8))
      setRoomStatus(rooms.slice(0, 10))
    } catch (error) {
      console.error('Failed to fetch stats:', error)
    } finally {
      setLoading(false)
    }
  }

  const getRoomStatusText = (status: number) => {
    const map: Record<number, string> = { 0: '维护中', 1: '空闲', 2: '使用中', 3: '已预约' }
    return map[status] || '未知'
  }

  const getRoomStatusTag = (status: number) => {
    const colors: Record<number, string> = { 0: 'default', 1: 'green', 2: 'blue', 3: 'orange' }
    return <Tag color={colors[status]}>{getRoomStatusText(status)}</Tag>
  }

  const getOrderStatusText = (status: number) => {
    const map: Record<number, string> = {
      0: '待支付', 1: '使用中', 2: '已完成', 3: '已取消', 4: '退款中', 5: '已退款'
    }
    return map[status] || '未知'
  }

  const getOrderStatusTag = (status: number) => {
    const colors: Record<number, string> = {
      0: 'orange', 1: 'blue', 2: 'green', 3: 'default', 4: 'red', 5: 'purple'
    }
    return <Tag color={colors[status]}>{getOrderStatusText(status)}</Tag>
  }

  const statCards = [
    {
      title: '今日订单',
      value: stats.todayOrders,
      icon: <ClockCircleOutlined />,
      color: '#1890ff',
      description: `共 ${stats.totalOrders} 单`
    },
    {
      title: '使用中包间',
      value: stats.activeRooms,
      icon: <HomeOutlined />,
      color: '#52c41a',
      description: `${stats.availableRooms} 间空闲`
    },
    {
      title: '会员总数',
      value: stats.activeMembers,
      icon: <UserOutlined />,
      color: '#faad14',
      description: `共 ${stats.totalMembers} 人`
    },
    {
      title: '今日营收',
      value: `¥${stats.todayRevenue.toFixed(2)}`,
      icon: <DollarOutlined />,
      color: '#ff4d4f',
      description: `累计 ¥${stats.totalRevenue.toFixed(2)}`
    }
  ]

  const orderColumns = [
    { title: '订单号', dataIndex: 'order_no', key: 'order_no', width: 150, ellipsis: true },
    { title: '用户', dataIndex: 'user', key: 'user', render: (u: any) => u?.nickname || '-', width: 100 },
    { title: '包间', dataIndex: 'room', key: 'room', render: (r: { name: string }) => r.name, width: 100 },
    { title: '金额', dataIndex: 'total_amount', key: 'total_amount', render: (v: number) => <span style={{ color: '#ff4d4f', fontWeight: 'bold' }}>¥{v}</span>, width: 100 },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getOrderStatusTag(s), width: 80 },
    { title: '创建时间', dataIndex: 'created_at', key: 'created_at', width: 160, render: (t: string) => formatDateTime(t) }
  ]

  const roomColumns = [
    { title: '包间名称', dataIndex: 'name', key: 'name', width: 120 },
    { title: '类型', dataIndex: 'type', key: 'type', render: (t: any) => t?.name || '-', width: 100 },
    { title: '楼层', dataIndex: 'floor', key: 'floor', width: 80 },
    { title: '状态', dataIndex: 'status', key: 'status', render: (s: number) => getRoomStatusTag(s), width: 80 }
  ]

  return (
    <div>
      <h2 style={{ marginBottom: 24 }}>仪表盘</h2>
      <Spin spinning={loading}>
        <Row gutter={[16, 16]} style={{ marginBottom: 24 }}>
          {statCards.map((card, index) => (
            <Col xs={24} sm={12} lg={6} key={index}>
              <Card hoverable>
                <Statistic
                  title={
                    <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                      <span style={{ color: card.color }}>{card.icon}</span>
                      <span>{card.title}</span>
                    </div>
                  }
                  value={card.value}
                  valueStyle={{ color: card.color, fontSize: '28px', fontWeight: 'bold' }}
                  suffix={<span style={{ fontSize: '12px', color: '#999' }}>{card.description}</span>}
                />
              </Card>
            </Col>
          ))}
        </Row>

        <Row gutter={[16, 16]}>
          <Col xs={24} lg={14}>
            <Card title={
              <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                <ShoppingCartOutlined style={{ color: '#1890ff' }} />
                <span>最近订单</span>
              </div>
            }>
              <Table
                dataSource={recentOrders}
                columns={orderColumns}
                rowKey="id"
                pagination={false}
                size="small"
              />
            </Card>
          </Col>

          <Col xs={24} lg={10}>
            <Card title={
              <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                <HomeOutlined style={{ color: '#52c41a' }} />
                <span>包间状态</span>
              </div>
            }>
              <Table
                dataSource={roomStatus}
                columns={roomColumns}
                rowKey="id"
                pagination={false}
                size="small"
              />
            </Card>

            <Card title={
              <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                <CheckCircleOutlined style={{ color: '#faad14' }} />
                <span>运营数据概览</span>
              </div>
            } style={{ marginTop: 16 }}>
              <div style={{ marginBottom: 16 }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 8 }}>
                  <span style={{ color: '#666' }}>包间使用率</span>
                  <span style={{ fontWeight: 'bold', color: '#1890ff' }}>
                    {stats.totalRooms > 0 ? Math.round((stats.activeRooms / stats.totalRooms) * 100) : 0}%
                  </span>
                </div>
                <Progress
                  percent={stats.totalRooms > 0 ? Math.round((stats.activeRooms / stats.totalRooms) * 100) : 0}
                  strokeColor="#1890ff"
                  size="default"
                />
              </div>
              <div style={{ marginBottom: 16 }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 8 }}>
                  <span style={{ color: '#666' }}>订单完成率</span>
                  <span style={{ fontWeight: 'bold', color: '#52c41a' }}>
                    {stats.totalOrders > 0 ? Math.round((stats.completedOrders / stats.totalOrders) * 100) : 0}%
                  </span>
                </div>
                <Progress
                  percent={stats.totalOrders > 0 ? Math.round((stats.completedOrders / stats.totalOrders) * 100) : 0}
                  strokeColor="#52c41a"
                  size="default"
                />
              </div>
              <div>
                <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 8 }}>
                  <span style={{ color: '#666' }}>会员活跃度</span>
                  <span style={{ fontWeight: 'bold', color: '#faad14' }}>
                    {stats.totalMembers > 0 ? Math.round((stats.activeMembers / stats.totalMembers) * 100) : 0}%
                  </span>
                </div>
                <Progress
                  percent={stats.totalMembers > 0 ? Math.round((stats.activeMembers / stats.totalMembers) * 100) : 0}
                  strokeColor="#faad14"
                  size="default"
                />
              </div>
            </Card>
          </Col>
        </Row>
      </Spin>
    </div>
  )
}