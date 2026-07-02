import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Tabs, Cell, CellGroup, Loading, Empty } from '@nutui/nutui-react'
import { membershipApi, Membership } from '@/api'
import CustomTabBar from '@/components/CustomTabBar'
import { getUserAvatar } from '@/utils/image'

export default function MemberPage() {
  const navigate = useNavigate()
  const [membership, setMembership] = useState<Membership | null>(null)
  const [rechargeRecords, setRechargeRecords] = useState<any[]>([])
  const [selectedTab, setSelectedTab] = useState(0)

  useEffect(() => {
    fetchMembership()
    fetchRechargeRecords()
  }, [])

  const fetchMembership = async () => {
    try {
      const data = await membershipApi.getMembership()
      setMembership(data)
    } catch (error) {
      console.error('Failed to fetch membership:', error)
    }
  }

  const fetchRechargeRecords = async () => {
    try {
      const data = await membershipApi.getRechargeRecords()
      setRechargeRecords(data)
    } catch (error) {
      console.error('Failed to fetch recharge records:', error)
    }
  }

  const handleRecharge = () => {
    navigate('/recharge')
  }

  const getLevelText = (level: number) => {
    const levelMap: Record<number, string> = {
      1: '普通会员',
      2: '银卡会员',
      3: '金卡会员',
      4: '钻石会员'
    }
    return levelMap[level] || '普通会员'
  }

  const getLevelColor = (level: number) => {
    const colorMap: Record<number, string> = {
      1: '#999999',
      2: '#c0c0c0',
      3: '#ffd700',
      4: '#b9f2ff'
    }
    return colorMap[level] || '#999999'
  }

  if (!membership) {
    return (
      <Loading type="circular" color="#667eea" style={{ marginTop: '100rpx' }} />
    )
  }

  return (
    <div className="page">
      <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 120px)' }}>
        <div className="member-card" style={{ borderColor: getLevelColor(membership.level) }}>
          <div className="member-header">
            <div className="member-avatar">
              <img src={getUserAvatar(membership.user?.avatar)} alt="会员头像" className="avatar-img" />
            </div>
            <div className="member-info">
              <span className="member-name">{membership.user?.realname || '会员'}</span>
              <span className="member-level" style={{ color: getLevelColor(membership.level) }}>
                {getLevelText(membership.level)}
              </span>
            </div>
          </div>

          <div className="member-balance">
            <span className="balance-label">账户余额</span>
            <span className="balance-value">¥{membership.balance}</span>
          </div>

          <div className="member-stats">
            <div className="stat-item">
              <span className="stat-value">{membership.total_recharged}</span>
              <span className="stat-label">累计充值</span>
            </div>
            <div className="stat-item">
              <span className="stat-value">{membership.discount}</span>
              <span className="stat-label">会员折扣</span>
            </div>
            <div className="stat-item">
              <span className="stat-value">{membership.points}</span>
              <span className="stat-label">积分</span>
            </div>
          </div>

          <Button type="primary" block onClick={handleRecharge}>
            立即充值
          </Button>
        </div>

        <Tabs
          value={selectedTab}
          onChange={(value) => setSelectedTab(Number(value))}
          activeType="card"
          activeColor="#667eea"
        >
          <Tabs.TabPane key={0} title="充值记录" />
          <Tabs.TabPane key={1} title="使用记录" />
        </Tabs>

        <CellGroup>
          {rechargeRecords.map(record => (
            <Cell key={record.id}>
              <div className="record-card">
                <div className="record-header">
                  <span className="record-type">充值</span>
                  <span className="record-amount">+¥{record.amount}</span>
                </div>
                <div className="record-info">
                  <span className="record-time">{record.created_at}</span>
                </div>
              </div>
            </Cell>
          ))}
        </CellGroup>

        {rechargeRecords.length === 0 && (
          <Empty description="暂无记录" image="empty" />
        )}
      </div>
      <CustomTabBar />
    </div>
  )
}
