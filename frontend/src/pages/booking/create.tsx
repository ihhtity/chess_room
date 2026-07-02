import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Button, Cell, CellGroup, Loading, TextArea, Calendar, Picker } from '@nutui/nutui-react'
import { showToast } from '@/components/Toast'
import { roomApi, orderApi, Room } from '@/api'
import './create.scss'

export default function CreateBooking() {
  const { room_id } = useParams<{ room_id: string }>()
  const navigate = useNavigate()
  const [room, setRoom] = useState<Room | null>(null)
  const [selectedDate, setSelectedDate] = useState('')
  const [startTime, setStartTime] = useState('')
  const [endTime, setEndTime] = useState('')
  const [remark, setRemark] = useState('')
  const [totalAmount, setTotalAmount] = useState(0)
  const [showCalendar, setShowCalendar] = useState(false)
  const [showStartPicker, setShowStartPicker] = useState(false)
  const [showEndPicker, setShowEndPicker] = useState(false)

  useEffect(() => {
    if (room_id) {
      fetchRoomDetail(parseInt(room_id))
    }
    const today = new Date()
    setSelectedDate(today.toISOString().split('T')[0])
  }, [room_id])

  const fetchRoomDetail = async (id: number) => {
    try {
      const data = await roomApi.getRoomDetail(id)
      setRoom(data)
    } catch (error) {
      console.error('Failed to fetch room detail:', error)
    }
  }

  useEffect(() => {
    if (startTime && endTime && room) {
      const start = new Date(`${selectedDate}T${startTime}`)
      const end = new Date(`${selectedDate}T${endTime}`)
      if (end > start) {
        const hours = (end.getTime() - start.getTime()) / (1000 * 60 * 60)
        setTotalAmount(Math.round(hours * room.type.base_price * 100) / 100)
      } else {
        setTotalAmount(0)
      }
    }
  }, [startTime, endTime, room, selectedDate])

  const timeSlots = [
    { text: '08:00', value: '08:00' },
    { text: '09:00', value: '09:00' },
    { text: '10:00', value: '10:00' },
    { text: '11:00', value: '11:00' },
    { text: '12:00', value: '12:00' },
    { text: '13:00', value: '13:00' },
    { text: '14:00', value: '14:00' },
    { text: '15:00', value: '15:00' },
    { text: '16:00', value: '16:00' },
    { text: '17:00', value: '17:00' },
    { text: '18:00', value: '18:00' },
    { text: '19:00', value: '19:00' },
    { text: '20:00', value: '20:00' },
    { text: '21:00', value: '21:00' },
    { text: '22:00', value: '22:00' },
    { text: '23:00', value: '23:00' },
    { text: '00:00', value: '00:00' }
  ]

  const handleSubmit = async () => {
    if (!room || !startTime || !endTime) {
      showToast({ message: '请选择完整的预约信息', type: 'warning' })
      return
    }

    const start = new Date(`${selectedDate}T${startTime}`)
    const end = new Date(`${selectedDate}T${endTime}`)
    if (end <= start) {
      showToast({ message: '结束时间必须晚于开始时间', type: 'warning' })
      return
    }

    try {
      await orderApi.createOrder({
        room_id: room.id,
        start_time: start.toISOString().replace('T', ' ').slice(0, 19),
        end_time: end.toISOString().replace('T', ' ').slice(0, 19),
        remark
      })
      showToast({ message: '预约成功', type: 'success' })
      setTimeout(() => {
        navigate('/orders')
      }, 1500)
    } catch (error) {
      console.error('Failed to create order:', error)
    }
  }

  if (!room) {
    return (
      <div className="loading-wrapper">
        <Loading type="circular" color="#667eea" />
      </div>
    )
  }

  return (
    <div className="container">
      <div className="room-card">
        <div className="room-info-header">
          <span className="room-name">{room.name}</span>
          <span className="room-type">{room.type.name}</span>
        </div>
        <div className="room-price-hint">
          <span className="price-label">单价</span>
          <span className="price-value">¥{room.type.base_price}/小时</span>
        </div>
      </div>

      <CellGroup className="form-group">
        <Cell
          title="选择日期"
          onClick={() => setShowCalendar(true)}
          description={selectedDate || '请选择日期'}
          extra={<span className="cell-arrow">›</span>}
        />

        <Cell
          title="开始时间"
          onClick={() => setShowStartPicker(true)}
          description={startTime || '请选择开始时间'}
          extra={<span className="cell-arrow">›</span>}
        />

        <Cell
          title="结束时间"
          onClick={() => setShowEndPicker(true)}
          description={endTime || '请选择结束时间'}
          extra={<span className="cell-arrow">›</span>}
        />

        <Cell title="备注">
          <TextArea
            placeholder="请输入备注信息（选填）"
            value={remark}
            onChange={(value) => setRemark(value)}
            rows={3}
          />
        </Cell>
      </CellGroup>

      <CellGroup className="summary-group">
        <Cell title="预约信息">
          <div className="booking-summary">
            <div className="summary-item">
              <span className="summary-label">包间</span>
              <span className="summary-value">{room.name} ({room.type.name})</span>
            </div>
            <div className="summary-item">
              <span className="summary-label">日期</span>
              <span className="summary-value">{selectedDate}</span>
            </div>
            <div className="summary-item">
              <span className="summary-label">时间</span>
              <span className="summary-value">{startTime} - {endTime}</span>
            </div>
            <div className="summary-item total">
              <span className="summary-label">预计费用</span>
              <span className="summary-value price">¥{totalAmount}</span>
            </div>
          </div>
        </Cell>
      </CellGroup>

      <div className="bottom-bar">
        <div className="price-info">
          <span className="price-label">预计费用</span>
          <span className="total-price">¥{totalAmount}</span>
        </div>
        <Button type="primary" block onClick={handleSubmit}>
          提交预约
        </Button>
      </div>

      <Calendar
        visible={showCalendar}
        onClose={() => setShowCalendar(false)}
        onConfirm={(date) => {
          setSelectedDate(date)
          setShowCalendar(false)
        }}
      />

      <Picker
        visible={showStartPicker}
        title="选择开始时间"
        options={timeSlots}
        onCancel={() => setShowStartPicker(false)}
        onConfirm={(_, value) => {
          setStartTime(value[0] as string)
          setShowStartPicker(false)
        }}
      />

      <Picker
        visible={showEndPicker}
        title="选择结束时间"
        options={timeSlots}
        onCancel={() => setShowEndPicker(false)}
        onConfirm={(_, value) => {
          setEndTime(value[0] as string)
          setShowEndPicker(false)
        }}
      />
    </div>
  )
}