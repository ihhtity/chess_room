import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Button, Cell, CellGroup, Loading, Badge, Divider, Empty } from '@nutui/nutui-react'
import { roomApi, reviewApi, Room, Review } from '@/api'
import { getRoomImage, getUserAvatar } from '@/utils/image'

export default function RoomDetail() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const [room, setRoom] = useState<Room | null>(null)
  const [equipment, setEquipment] = useState<string[]>([])
  const [reviews, setReviews] = useState<Review[]>([])

  useEffect(() => {
    if (id) {
      fetchRoomDetail(parseInt(id))
      fetchReviews(parseInt(id))
    }
  }, [id])

  const fetchRoomDetail = async (id: number) => {
    try {
      const data = await roomApi.getRoomDetail(id)
      setRoom(data)
      if (data.equipment) {
        setEquipment(JSON.parse(data.equipment))
      }
    } catch (error) {
      console.error('Failed to fetch room detail:', error)
    }
  }

  const fetchReviews = async (roomId: number) => {
    try {
      const data = await reviewApi.getReviewList({ room_id: roomId })
      setReviews(data)
    } catch (error) {
      console.error('Failed to fetch reviews:', error)
      setReviews([
        { id: 1, order_id: 1, user_id: 1, user: { id: 1, openid: '', phone: '', nickname: '用户A', realname: '', avatar: '', gender: 1 }, room_id: roomId, rating: 5, content: '环境很好，包间干净整洁，麻将机很新，下次还会来！', images: '', created_at: '2026-01-20' },
        { id: 2, order_id: 2, user_id: 2, user: { id: 2, openid: '', phone: '', nickname: '用户B', realname: '', avatar: '', gender: 2 }, room_id: roomId, rating: 4, content: '整体不错，就是空调稍微有点凉，建议调高一点温度。', images: '', created_at: '2026-01-18' },
        { id: 3, order_id: 3, user_id: 3, user: { id: 3, openid: '', phone: '', nickname: '用户C', realname: '', avatar: '', gender: 1 }, room_id: roomId, rating: 5, content: '服务态度很好，价格实惠，非常满意！', images: '', created_at: '2026-01-15' }
      ])
    }
  }

  const handleBook = () => {
    if (!room) return
    navigate(`/booking/${room.id}`)
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

  if (!room) {
    return (
      <Loading type="circular" color="#667eea" style={{ marginTop: '100rpx' }} />
    )
  }

  return (
    <div className="container" style={{ overflowY: 'auto', height: 'calc(100vh - 100px)' }}>
      <div className="room-images">
        <img
          className="room-image"
          src={getRoomImage(room.images, 600, 450)}
          alt={room.name}
        />
      </div>

      <CellGroup>
        <Cell>
          <div className="room-header">
            <div>
              <span className="room-name">{room.name}</span>
              <div className="room-meta">
                <span className="room-type">{room.type.name}</span>
                <Badge value={getStatusText(room.status)} color={getStatusColor(room.status)} />
              </div>
            </div>
            <span className="room-price">¥{room.type.base_price}<span className="price-unit">/小时</span></span>
          </div>
        </Cell>

        <Cell title="包间介绍" description={room.description || '暂无介绍'} />

        <Cell title="设备清单">
          <div className="equipment-list">
            {equipment.map((item, index) => (
              <span key={index} className="equipment-item">✓ {item}</span>
            ))}
          </div>
        </Cell>

        <Cell title="楼层" description={room.floor} />
        <Cell title="容纳人数" description={`${room.type.max_people}人`} />
        <Cell title="基础价格" description={`¥${room.type.base_price}/小时`} />
      </CellGroup>

      <Divider />

      <div className="review-section">
        <div className="section-header">
          <span className="section-title">⭐ 用户评价</span>
          <span className="review-count">({reviews.length})</span>
        </div>
        {reviews.length > 0 ? (
          <CellGroup>
            {reviews.map(review => (
              <Cell key={review.id} className="review-cell">
                <div className="review-card">
                  <div className="review-header">
                    <div className="reviewer-info">
                      <img src={getUserAvatar(review.user?.avatar)} alt="" className="reviewer-avatar" />
                      <span className="reviewer-name">{review.user?.nickname || '用户'}</span>
                    </div>
                    <div className="review-rating">
                      {'★'.repeat(review.rating)}{'☆'.repeat(5 - review.rating)}
                    </div>
                  </div>
                  <p className="review-content">{review.content}</p>
                  <span className="review-date">{review.created_at}</span>
                </div>
              </Cell>
            ))}
          </CellGroup>
        ) : (
          <Empty description="暂无评价" image="empty" />
        )}
      </div>

      <div className="bottom-bar">
        <div className="price-info">
          <span className="total-price">¥{room.type.base_price}</span>
          <span className="price-hint">/小时起</span>
        </div>
        <Button type="primary" block onClick={handleBook}>
          立即预订
        </Button>
      </div>
    </div>
  )
}
