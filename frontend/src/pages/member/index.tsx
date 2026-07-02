import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Tabs, Cell, CellGroup, Loading, Empty } from '@nutui/nutui-react'
import { membershipApi, Membership } from '@/api'
import { getUserAvatar } from '@/utils/image'
import './index.scss'

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
      setMembership({
        id: 1, user_id: 1, user: { id: 1, openid: '', phone: '', nickname: '用户A', realname: '张三', avatar: '', gender: 1 }, level: 3, balance: 1500, total_recharged: 5000, discount: 0.85, points: 2500, total_consumed: 3500
      })
    }
  }

  const fetchRechargeRecords = async () => {
    try {
      const data = await membershipApi.getRechargeRecords()
      setRechargeRecords(data)
    } catch (error) {
      console.error('Failed to fetch recharge records:', error)
      setRechargeRecords([
        { id: 1, membership_id: 1, type: 'recharge', amount: 500, balance_before: 1000, balance_after: 1500, created_at: '2026-01-20 14:30:00' },
        { id: 2, membership_id: 1, type: 'consume', amount: 200, balance_before: 1200, balance_after: 1000, created_at: '2026-01-18 20:15:00' },
        { id: 3, membership_id: 1, type: 'recharge', amount: 1000, balance_before: 200, balance_after: 1200, created_at: '2026-01-15 10:00:00' }
      ])
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

  const getCardGradient = (level: number) => {
    const gradients: Record<number, string> = {
      1: 'linear-gradient(135deg, #666666 0%, #999999 100%)',
      2: 'linear-gradient(135deg, #c0c0c0 0%, #e8e8e8 100%)',
      3: 'linear-gradient(135deg, #ffd700 0%, #ffaa00 100%)',
      4: 'linear-gradient(135deg, #b9f2ff 0%, #87ceeb 100%)'
    }
    return gradients[level] || gradients[1]
  }

  if (!membership) {
    return (
      <div className="loading-wrapper">
        <Loading type="circular" color="#667eea" />
      </div>
    )
  }

  return (
    <div className="container">
      <div className="member-card" style={{ background: getCardGradient(membership.level) }}>
        <div className="card-watermark">VIP</div>
        <div className="member-header">
          <div className="member-avatar">
            <img src={getUserAvatar(membership.user?.avatar)} alt="会员头像" className="avatar-img" />
          </div>
          <div className="member-info">
            <span className="member-name">{membership.user?.realname || '会员'}</span>
            <span className="member-level">{getLevelText(membership.level)}</span>
          </div>
        </div>

        <div className="member-balance">
          <span className="balance-label">账户余额</span>
          <span className="balance-value">¥{membership.balance.toFixed(2)}</span>
        </div>

        <div className="member-stats">
          <div className="stat-item">
            <span className="stat-value">¥{membership.total_recharged}</span>
            <span className="stat-label">累计充值</span>
          </div>
          <div className="stat-divider"></div>
          <div className="stat-item">
            <span className="stat-value">{(membership.discount || 1) * 10}折</span>
            <span className="stat-label">会员折扣</span>
          </div>
          <div className="stat-divider"></div>
          <div className="stat-item">
            <span className="stat-value">{membership.points}</span>
            <span className="stat-label">积分</span>
          </div>
        </div>

        <Button type="primary" block className="recharge-btn" onClick={handleRecharge}>
          立即充值
        </Button>
      </div>

      <Tabs
        value={selectedTab}
        onChange={(value) => setSelectedTab(Number(value))}
        activeType="card"
      >
        <Tabs.TabPane key={0} title="充值记录">
          {rechargeRecords.filter(r => r.type === 'recharge').length > 0 ? (
            <CellGroup className="record-group">
              {rechargeRecords.filter(r => r.type === 'recharge').map(record => (
                <Cell key={record.id} className="record-cell">
                  <div className="record-card">
                    <div className="record-header">
                      <div className="record-icon recharge-icon">💵</div>
                      <span className="record-type">充值</span>
                      <span className="record-amount positive">+¥{record.amount}</span>
                    </div>
                    <div className="record-info">
                      <span className="record-time">{record.created_at}</span>
                      <span className="record-balance">余额: ¥{record.balance_after}</span>
                    </div>
                  </div>
                </Cell>
              ))}
            </CellGroup>
          ) : (
            <Empty description="暂无充值记录" image="empty" />
          )}
        </Tabs.TabPane>
        <Tabs.TabPane key={1} title="消费记录">
          {rechargeRecords.filter(r => r.type === 'consume').length > 0 ? (
            <CellGroup className="record-group">
              {rechargeRecords.filter(r => r.type === 'consume').map(record => (
                <Cell key={record.id} className="record-cell">
                  <div className="record-card">
                    <div className="record-header">
                      <div className="record-icon consume-icon">📋</div>
                      <span className="record-type">消费</span>
                      <span className="record-amount negative">-¥{record.amount}</span>
                    </div>
                    <div className="record-info">
                      <span className="record-time">{record.created_at}</span>
                      <span className="record-balance">余额: ¥{record.balance_after}</span>
                    </div>
                  </div>
                </Cell>
              ))}
            </CellGroup>
          ) : (
            <Empty description="暂无消费记录" image="empty" />
          )}
        </Tabs.TabPane>
      </Tabs>
    </div>
  )
}