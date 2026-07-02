import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, Loading, Empty, Badge, Swiper, SwiperItem, Cell, CellGroup, Grid, GridItem } from '@nutui/nutui-react'
import { roomApi, activityApi, announcementApi, Room, RoomType, Activity } from '@/api'
import CustomTabBar from '@/components/CustomTabBar'
import { getRoomImage } from '@/utils/image'
import './index.scss'

export default function Index() {
  const navigate = useNavigate()
  const [rooms, setRooms] = useState<Room[]>([])
  const [types, setTypes] = useState<RoomType[]>([])
  const [selectedType, setSelectedType] = useState(0)
  const [loading, setLoading] = useState(true)
  const [activities, setActivities] = useState<Activity[]>([])
  const [announcement, setAnnouncement] = useState('')

  useEffect(() => {
    fetchRooms()
    fetchTypes()
    fetchActivities()
    fetchAnnouncement()
  }, [selectedType])

  const fetchRooms = async () => {
    try {
      setLoading(true)
      const data = await roomApi.getRoomList({ type_id: selectedType })
      setRooms(data)
    } catch (error) {
      console.error('Failed to fetch rooms:', error)
    } finally {
      setLoading(false)
    }
  }

  const fetchTypes = async () => {
    try {
      const data = await roomApi.getRoomTypeList()
      setTypes([{ id: 0, name: '全部', description: '', base_price: 0, max_people: 0 }, ...data])
    } catch (error) {
      console.error('Failed to fetch types:', error)
    }
  }

  const fetchActivities = async () => {
    try {
      const data = await activityApi.getActivityList()
      setActivities(data)
    } catch (error) {
      console.error('Failed to fetch activities:', error)
      setActivities([
        { id: 1, name: '新用户专享', description: '首次预订享受8折优惠', image: '', discount: 0.8, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 },
        { id: 2, name: '工作日特惠', description: '周一至周四全天9折', image: '', discount: 0.9, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 },
        { id: 3, name: '会员日活动', description: '每月8日会员享双倍积分', image: '', discount: 1, valid_from: '2026-01-01', valid_to: '2026-12-31', status: 1 }
      ])
    }
  }

  const fetchAnnouncement = async () => {
    try {
      const data = await announcementApi.getAnnouncementList()
      if (data.length > 0) {
        setAnnouncement(data[0].title)
      }
    } catch (error) {
      console.error('Failed to fetch announcements:', error)
      setAnnouncement('春节期间营业时间调整通知')
    }
  }

  const handleRoomClick = (room: Room) => {
    navigate(`/rooms/${room.id}`)
  }

  const handleActivityClick = () => {
    navigate('/activities')
  }

  const handleAnnouncementClick = () => {
    navigate('/announcements')
  }

  const getStatusText = (status: number) => {
    const statusMap: Record<number, string> = {
      0: '维护中',
      1: '空闲',
      2: '使用中',
      3: '已预约'
    }
    return statusMap[status] || '未知'
  }

  const getStatusColor = (status: number) => {
    const colorMap: Record<number, string> = {
      0: '#999999',
      1: '#52c41a',
      2: '#ff4d4f',
      3: '#faad14'
    }
    return colorMap[status] || '#999999'
  }

  const menuItems = [
    { icon: '🏠', label: '包间预订', path: '/' },
    { icon: '🎯', label: '活动中心', path: '/activities' },
    { icon: '📢', label: '公告通知', path: '/announcements' },
    { icon: '👤', label: '会员中心', path: '/member' }
  ]

  return (
    <div className="page">
      <div className="container">
        <div className="header">
          <div className="header-top">
            <span className="logo">♟️</span>
            <div className="header-text">
              <span className="title">棋牌室</span>
              <span className="subtitle">享受休闲时光</span>
            </div>
          </div>
          <div className="search-bar" onClick={() => {}}>
            <span className="search-icon">🔍</span>
            <span className="search-placeholder">搜索包间、活动...</span>
          </div>
        </div>

        <div className="swiper-container">
          <Swiper
            autoPlay={3000}
            indicator={true}
            loop={true}
            duration={500}
          >
            <SwiperItem>
              <div className="swiper-item" style={{ background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)' }}>
                <div className="swiper-decoration swiper-decoration-1"></div>
                <div className="swiper-decoration swiper-decoration-2"></div>
                <div className="swiper-content">
                  <Badge value="限时特惠" className="swiper-badge" />
                  <span className="swiper-title">新用户专享</span>
                  <span className="swiper-desc">首次预订享受8折优惠</span>
                </div>
              </div>
            </SwiperItem>
            <SwiperItem>
              <div className="swiper-item" style={{ background: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)' }}>
                <div className="swiper-decoration swiper-decoration-1"></div>
                <div className="swiper-decoration swiper-decoration-2"></div>
                <div className="swiper-content">
                  <Badge value="周末活动" className="swiper-badge" />
                  <span className="swiper-title">周末特惠</span>
                  <span className="swiper-desc">周六日全天9折优惠</span>
                </div>
              </div>
            </SwiperItem>
            <SwiperItem>
              <div className="swiper-item" style={{ background: 'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)' }}>
                <div className="swiper-decoration swiper-decoration-1"></div>
                <div className="swiper-decoration swiper-decoration-2"></div>
                <div className="swiper-content">
                  <Badge value="会员福利" className="swiper-badge" />
                  <span className="swiper-title">会员日活动</span>
                  <span className="swiper-desc">每月8日双倍积分</span>
                </div>
              </div>
            </SwiperItem>
          </Swiper>
        </div>

        <CellGroup className="notice-group">
          <Cell onClick={handleAnnouncementClick}>
            <div className="notice-bar">
              <span className="notice-icon">📢</span>
              <span className="notice-text">{announcement}</span>
              <span className="notice-arrow">→</span>
            </div>
          </Cell>
        </CellGroup>

        <Grid columns={4} className="menu-grid">
          {menuItems.map((item, index) => (
            <GridItem key={index} onClick={() => navigate(item.path)}>
              <div className="menu-icon-wrapper">
                <span className="menu-icon">{item.icon}</span>
              </div>
              <span className="menu-label">{item.label}</span>
            </GridItem>
          ))}
        </Grid>

        {activities.length > 0 && (
          <div className="activity-section">
            <div className="section-header">
              <span className="section-title">🎁 优惠活动</span>
              <span className="section-more" onClick={handleActivityClick}>查看更多 →</span>
            </div>
            <CellGroup>
              {activities.slice(0, 2).map(activity => (
                <Cell key={activity.id} onClick={handleActivityClick} className="activity-cell">
                  <div className="activity-card">
                    <div className="activity-left">
                      <div className="activity-icon">🎯</div>
                    </div>
                    <div className="activity-right">
                      <div className="activity-header">
                        <span className="activity-name">{activity.name}</span>
                        {activity.discount < 1 && (
                          <Badge value={`${Math.round((1 - activity.discount) * 10)}折`} color="#ff4d4f" />
                        )}
                      </div>
                      <span className="activity-desc">{activity.description}</span>
                    </div>
                  </div>
                </Cell>
              ))}
            </CellGroup>
          </div>
        )}

        <div className="section-header">
          <span className="section-title">🏠 包间列表</span>
        </div>

        <div className="filter-row">
          {types.map(type => (
            <Button
              key={type.id}
              type={selectedType === type.id ? 'primary' : 'default'}
              shape="round"
              size="small"
              className={`filter-btn ${selectedType === type.id ? 'active' : ''}`}
              onClick={() => setSelectedType(type.id)}
            >
              {type.name}
            </Button>
          ))}
        </div>

        {loading ? (
          <div className="loading-wrapper">
            <Loading type="circular" color="#667eea" />
          </div>
        ) : rooms.length > 0 ? (
          <CellGroup>
            {rooms.map(room => (
              <Cell key={room.id} onClick={() => handleRoomClick(room)} className="room-cell">
                <div className="room-card">
                  <div className="room-image-wrapper">
                    <img
                      className="room-image"
                      src={getRoomImage(room.images, 400, 300)}
                      alt={room.name}
                    />
                    <Badge value={getStatusText(room.status)} color={getStatusColor(room.status)} className="room-status-badge" />
                  </div>
                  <div className="room-info">
                    <span className="room-name">{room.name}</span>
                    <span className="room-type">{room.type.name}</span>
                    <div className="room-details">
                      <span className="room-floor">📍 {room.floor}</span>
                      <span className="room-people">👥 {room.type.max_people}人</span>
                    </div>
                    <div className="room-price-row">
                      <span className="room-price">¥{room.type.base_price}<span className="price-unit">/小时</span></span>
                      <Button type="primary" size="small" onClick={(e) => { e.stopPropagation(); handleRoomClick(room) }}>
                        预订
                      </Button>
                    </div>
                  </div>
                </div>
              </Cell>
            ))}
          </CellGroup>
        ) : (
          <Empty description="暂无可用包间" image="empty" />
        )}

        <div className="bottom-spacer"></div>
      </div>
      <CustomTabBar />
    </div>
  )
}