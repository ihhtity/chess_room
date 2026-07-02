import { useState, useEffect } from 'react'
import { Cell, CellGroup, Loading, Empty, Badge } from '@nutui/nutui-react'
import { activityApi, Activity } from '@/api'

export default function ActivityPage() {
  const [activities, setActivities] = useState<Activity[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchActivities()
  }, [])

  const fetchActivities = async () => {
    try {
      setLoading(true)
      const data = await activityApi.getActivityList()
      setActivities(data)
    } catch (error) {
      console.error('Failed to fetch activities:', error)
      setActivities([
        { id: 1, name: '新用户专享', description: '首次预订享受8折优惠，限前100名用户', image: '', discount: 0.8, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 },
        { id: 2, name: '工作日特惠', description: '周一至周四全天9折优惠，节假日除外', image: '', discount: 0.9, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 },
        { id: 3, name: '会员日活动', description: '每月8日会员享双倍积分，充值额外赠送10%', image: '', discount: 1, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 },
        { id: 4, name: '周末畅玩', description: '周六日预订满4小时赠送1小时', image: '', discount: 1, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 },
        { id: 5, name: '节日礼包', description: '春节、国庆期间预订享豪华礼包', image: '', discount: 0.85, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 }
      ])
    } finally {
      setLoading(false)
    }
  }

  const getStatusText = (status: number) => {
    return status === 1 ? '进行中' : '已结束'
  }

  const getStatusColor = (status: number) => {
    return status === 1 ? '#52c41a' : '#999999'
  }

  return (
    <div className="page">
      <div className="header">
        <div className="header-content">
          <span className="title">🎁 优惠活动</span>
          <span className="subtitle">超值优惠等你来享</span>
        </div>
      </div>

      <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 180px)' }}>
        {loading ? (
          <Loading type="circular" color="#667eea" style={{ marginTop: '100rpx' }} />
        ) : (
          <CellGroup>
            {activities.map(activity => (
              <Cell key={activity.id} className="activity-cell">
                <div className="activity-card">
                  <div className="activity-left">
                    <div className="activity-icon">🎯</div>
                  </div>
                  <div className="activity-right">
                    <div className="activity-header">
                      <span className="activity-name">{activity.name}</span>
                      <Badge value={getStatusText(activity.status)} color={getStatusColor(activity.status)} />
                    </div>
                    <span className="activity-desc">{activity.description}</span>
                    <div className="activity-footer">
                      {activity.discount < 1 && (
                        <span className="activity-discount">{Math.round((1 - activity.discount) * 10)}折优惠</span>
                      )}
                      <span className="activity-date">有效期: {activity.valid_from} ~ {activity.valid_to}</span>
                    </div>
                  </div>
                </div>
              </Cell>
            ))}
          </CellGroup>
        )}

        {!loading && activities.length === 0 && (
          <Empty description="暂无活动" image="empty" />
        )}
      </div>
    </div>
  )
}
