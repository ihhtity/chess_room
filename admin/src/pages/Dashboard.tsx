import { useState, useEffect } from 'react'
import { Card, Row, Col, Statistic } from 'antd'
import {
  HomeOutlined,
  ShoppingCartOutlined,
  UserOutlined,
  DollarOutlined
} from '@ant-design/icons'
import { roomApi, orderApi, membershipApi } from '@/api'

export default function Dashboard() {
  const [stats, setStats] = useState({
    roomCount: 0,
    orderCount: 0,
    memberCount: 0,
    totalRevenue: 0
  })

  useEffect(() => {
    fetchStats()
  }, [])

  const fetchStats = async () => {
    try {
      const [rooms, orders, members] = await Promise.all([
        roomApi.getList(),
        orderApi.getList(),
        membershipApi.getList()
      ])
      const totalRevenue = orders.reduce((sum, o) => sum + o.paid_amount, 0)
      setStats({
        roomCount: rooms.length,
        orderCount: orders.length,
        memberCount: members.length,
        totalRevenue
      })
    } catch (error) {
      console.error('Failed to fetch stats:', error)
    }
  }

  const statCards = [
    {
      title: '包间总数',
      value: stats.roomCount,
      icon: <HomeOutlined />,
      color: '#1890ff'
    },
    {
      title: '订单总数',
      value: stats.orderCount,
      icon: <ShoppingCartOutlined />,
      color: '#52c41a'
    },
    {
      title: '会员总数',
      value: stats.memberCount,
      icon: <UserOutlined />,
      color: '#faad14'
    },
    {
      title: '总收入',
      value: `¥${stats.totalRevenue.toFixed(2)}`,
      icon: <DollarOutlined />,
      color: '#ff4d4f'
    }
  ]

  return (
    <div>
      <h2>仪表盘</h2>
      <Row gutter={16} style={{ marginTop: 24 }}>
        {statCards.map((card, index) => (
          <Col span={6} key={index}>
            <Card>
              <Statistic
                title={card.title}
                value={card.value}
                prefix={card.icon}
                valueStyle={{ color: card.color }}
              />
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  )
}