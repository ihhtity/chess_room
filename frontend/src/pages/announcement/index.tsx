import { useState, useEffect } from 'react'
import { Cell, CellGroup, Loading, Empty } from '@nutui/nutui-react'
import { announcementApi, Announcement } from '@/api'

export default function AnnouncementPage() {
  const [announcements, setAnnouncements] = useState<Announcement[]>([])
  const [loading, setLoading] = useState(true)
  const [selectedId, setSelectedId] = useState<number | null>(null)

  useEffect(() => {
    fetchAnnouncements()
  }, [])

  const fetchAnnouncements = async () => {
    try {
      setLoading(true)
      const data = await announcementApi.getAnnouncementList()
      setAnnouncements(data)
    } catch (error) {
      console.error('Failed to fetch announcements:', error)
      setAnnouncements([
        { id: 1, title: '春节期间营业时间调整通知', content: '尊敬的顾客：\n\n春节期间（2026年1月28日-2月4日），本店营业时间调整为10:00-22:00，敬请谅解。\n\n祝大家春节快乐！', type: 1, created_at: '2026-01-20' },
        { id: 2, title: '会员系统升级公告', content: '尊敬的会员：\n\n为了给您提供更好的服务体验，我们将于2026年2月1日对会员系统进行升级维护，届时会员功能将暂时不可用。\n\n感谢您的理解与支持！', type: 2, created_at: '2026-01-18' },
        { id: 3, title: '新增棋牌包间开业', content: '好消息！本店新增3间豪华棋牌包间，配备全新麻将机和舒适座椅，欢迎各位新老顾客前来体验。', type: 1, created_at: '2026-01-15' },
        { id: 4, title: '安全须知', content: '为了确保您的安全，请遵守以下规定：\n1. 禁止在包间内吸烟\n2. 禁止携带易燃易爆物品\n3. 请妥善保管贵重物品\n4. 如遇紧急情况请联系工作人员', type: 3, created_at: '2026-01-10' }
      ])
    } finally {
      setLoading(false)
    }
  }

  const getTypeText = (type: number) => {
    const typeMap: Record<number, string> = {
      1: '通知',
      2: '公告',
      3: '须知'
    }
    return typeMap[type] || '其他'
  }

  const getTypeColor = (type: number) => {
    const colorMap: Record<number, string> = {
      1: '#667eea',
      2: '#faad14',
      3: '#52c41a'
    }
    return colorMap[type] || '#999999'
  }

  return (
    <div className="page">
      <div className="header">
        <div className="header-content">
          <span className="title">📢 公告通知</span>
          <span className="subtitle">及时了解最新动态</span>
        </div>
      </div>

      <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 180px)' }}>
        {loading ? (
          <Loading type="circular" color="#667eea" style={{ marginTop: '100rpx' }} />
        ) : (
          <CellGroup>
            {announcements.map(announcement => (
              <div key={announcement.id}>
                <Cell
                  onClick={() => setSelectedId(selectedId === announcement.id ? null : announcement.id)}
                  className="announcement-cell"
                >
                  <div className="announcement-card">
                    <div className="announcement-header">
                      <span className="announcement-type" style={{ backgroundColor: getTypeColor(announcement.type) }}>
                        {getTypeText(announcement.type)}
                      </span>
                      <span className="announcement-title">{announcement.title}</span>
                      <span className="announcement-arrow">{selectedId === announcement.id ? '▼' : '▶'}</span>
                    </div>
                    <span className="announcement-date">{announcement.created_at}</span>
                  </div>
                </Cell>
                {selectedId === announcement.id && (
                  <div className="announcement-content">
                    <Cell>
                      <div className="content-text">{announcement.content}</div>
                    </Cell>
                  </div>
                )}
              </div>
            ))}
          </CellGroup>
        )}

        {!loading && announcements.length === 0 && (
          <Empty description="暂无公告" image="empty" />
        )}
      </div>
    </div>
  )
}
