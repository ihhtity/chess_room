<template>
  <view class="container">
    <view v-if="loading" class="loading-container">
      <uv-loading-icon mode="flower" size="48" color="#4CAF50" />
      <text class="loading-text">加载中...</text>
    </view>

    <view v-else class="content">
      <view class="card">
        <text class="card-title">选择日期</text>
        <view class="date-picker">
          <picker mode="date" :value="selectedDate" @change="handleDateChange">
            <view class="picker-value">
              <text>{{ selectedDate }}</text>
              <view class="date-tag" :class="dateTypeClass">{{ dateTypeText }}</view>
              <uv-icon name="arrow-right" color="#999" size="24" />
            </view>
          </picker>
        </view>
      </view>

      <view class="card">
        <text class="card-title">选择时段</text>
        <view class="time-slots">
          <view 
            class="time-slot-item" 
            :class="{ active: selectedTimeSlot === slot.id, unavailable: !isSlotAvailable(slot) }"
            v-for="slot in timeSlots" 
            :key="slot.id"
            @click="selectTimeSlot(slot)"
          >
            <text class="slot-name">{{ slot.name }}</text>
            <text class="slot-time">{{ slot.start_time }} - {{ slot.end_time }}</text>
            <view class="slot-price-row">
              <text class="slot-price">¥{{ getSlotPrice(slot) }}</text>
              <text v-if="getSlotPrice(slot) !== slot.weekday_price" class="price-tag">{{ dateTypeText }}</text>
            </view>
          </view>
        </view>
      </view>

      <view class="card">
        <text class="card-title">选择包间</text>
        <view class="room-list">
          <view 
            class="room-item" 
            :class="{ active: selectedRoom === room.id, unavailable: room.status !== 1 }"
            v-for="room in availableRooms" 
            :key="room.id"
            @click="selectRoom(room)"
          >
            <view class="room-info">
              <text class="room-name">{{ room.name }}</text>
              <text class="room-type">{{ room.type.name }} · {{ room.capacity }}人</text>
            </view>
            <view class="room-status-badge" :class="getStatusClass(room.status)">
              {{ getStatusText(room.status) }}
            </view>
            <view class="check-icon" v-if="selectedRoom === room.id">
              <uv-icon name="check-circle" color="#4CAF50" size="32" />
            </view>
          </view>
          <view v-if="availableRooms.length === 0" class="empty-tip">
            <text>暂无可用包间</text>
          </view>
        </view>
      </view>

      <view class="card">
        <text class="card-title">订单备注</text>
        <textarea class="remark-input" v-model="remark" placeholder="请输入备注信息" maxlength="200" />
        <text class="remark-count">{{ remark.length }}/200</text>
      </view>

      <view class="summary-card">
        <view class="summary-item">
          <text class="summary-label">预约日期</text>
          <text class="summary-value">{{ selectedDate }} <text class="date-type">{{ dateTypeText }}</text></text>
        </view>
        <view class="summary-item">
          <text class="summary-label">预约时段</text>
          <text class="summary-value">{{ selectedTimeSlotName }}</text>
        </view>
        <view class="summary-item">
          <text class="summary-label">选择包间</text>
          <text class="summary-value">{{ selectedRoomName }}</text>
        </view>
        <view class="summary-item">
          <text class="summary-label">时段时长</text>
          <text class="summary-value">{{ selectedDuration }}小时</text>
        </view>
        <view class="summary-item total">
          <text class="summary-label">总金额</text>
          <text class="summary-value">¥{{ totalAmount }}</text>
        </view>
      </view>

      <view class="bottom-bar">
        <view class="price-display">
          <text class="price-label">合计</text>
          <text class="price-amount">¥{{ totalAmount }}</text>
        </view>
        <view class="submit-btn" :class="{ disabled: !canSubmit }" @click="handleSubmit">
          确认预约
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { roomApi, timeSlotApi, orderApi } from '@/api'
import type { Room, TimeSlot } from '@/types'
import { RoomStatus } from '@/types'

const selectedDate = ref('')
const selectedTimeSlot = ref(0)
const selectedRoom = ref(0)
const remark = ref('')
const timeSlots = ref<TimeSlot[]>([])
const availableRooms = ref<Room[]>([])
const rooms = ref<Room[]>([])
const loading = ref(true)
const prespecifiedRoomId = ref(0)

const today = new Date()
const todayStr = `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}-${String(today.getDate()).padStart(2, '0')}`
selectedDate.value = todayStr

onMounted(async () => {
  await loadData()
})

const loadData = async () => {
  loading.value = true
  
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1] as any
  const options = currentPage?.options || {}
  if (options.room_id) {
    prespecifiedRoomId.value = parseInt(options.room_id as string) || 0
  }

  try {
    const [timeSlotRes, roomRes] = await Promise.all([
      timeSlotApi.getList(),
      roomApi.getList({ status: 1, page: 1, page_size: 20 })
    ])
    timeSlots.value = timeSlotRes.data
    rooms.value = roomRes.data
    availableRooms.value = roomRes.data

    if (prespecifiedRoomId.value > 0) {
      selectedRoom.value = prespecifiedRoomId.value
    }
  } catch (e) {
    console.error('加载数据失败', e)
    uni.showToast({ title: '加载数据失败', icon: 'none' })
  } finally {
    loading.value = false
  }
}

const handleDateChange = (e: any) => {
  selectedDate.value = e.detail.value
  selectedTimeSlot.value = 0
}

const selectTimeSlot = (slot: TimeSlot) => {
  if (!isSlotAvailable(slot)) {
    uni.showToast({ title: '该时段暂不可用', icon: 'none' })
    return
  }
  selectedTimeSlot.value = slot.id
}

const selectRoom = (room: Room) => {
  if (room.status !== 1) {
    uni.showToast({ title: '该包间暂不可用', icon: 'none' })
    return
  }
  selectedRoom.value = room.id
}

const dateType = computed(() => {
  if (!selectedDate.value) return 'weekday'
  
  const date = new Date(selectedDate.value)
  const dayOfWeek = date.getDay()
  
  if (dayOfWeek === 0 || dayOfWeek === 6) {
    return 'weekend'
  }
  
  return 'weekday'
})

const dateTypeText = computed(() => {
  switch (dateType.value) {
    case 'weekend':
      return '周末'
    case 'holiday':
      return '节假日'
    default:
      return '工作日'
  }
})

const dateTypeClass = computed(() => dateType.value)

const selectedTimeSlotName = computed(() => {
  const slot = timeSlots.value.find(s => s.id === selectedTimeSlot.value)
  return slot ? `${slot.name} (${slot.start_time}-${slot.end_time})` : ''
})

const selectedRoomName = computed(() => {
  const room = availableRooms.value.find(r => r.id === selectedRoom.value)
  return room ? room.name : ''
})

const selectedDuration = computed(() => {
  const slot = timeSlots.value.find(s => s.id === selectedTimeSlot.value)
  if (!slot) return 0
  
  const [startHour] = slot.start_time.split(':').map(Number)
  const [endHour] = slot.end_time.split(':').map(Number)
  let duration = endHour - startHour
  if (duration < 0) duration += 24
  return duration
})

const getSlotPrice = (slot: TimeSlot) => {
  switch (dateType.value) {
    case 'weekend':
      return slot.weekend_price || slot.price || slot.weekday_price || 0
    case 'holiday':
      return slot.holiday_price || slot.price || slot.weekday_price || 0
    default:
      return slot.weekday_price || slot.price || 0
  }
}

const isSlotAvailable = (slot: TimeSlot) => {
  return slot.status === 1
}

const totalAmount = computed(() => {
  const slot = timeSlots.value.find(s => s.id === selectedTimeSlot.value)
  if (!slot) return 0
  
  const price = getSlotPrice(slot)
  return price * selectedDuration.value
})

const canSubmit = computed(() => {
  return selectedDate.value && selectedTimeSlot.value > 0 && selectedRoom.value > 0
})

const getStatusText = (status: number) => {
  switch (status) {
    case RoomStatus.Available:
      return '可预约'
    case RoomStatus.InUse:
      return '使用中'
    case RoomStatus.Reserved:
      return '已预约'
    case RoomStatus.Maintenance:
      return '维护中'
    default:
      return '未知'
  }
}

const getStatusClass = (status: number) => {
  switch (status) {
    case RoomStatus.Available:
      return 'available'
    case RoomStatus.InUse:
      return 'in-use'
    case RoomStatus.Reserved:
      return 'reserved'
    case RoomStatus.Maintenance:
      return 'maintenance'
    default:
      return ''
  }
}

const handleSubmit = async () => {
  if (!canSubmit.value) {
    uni.showToast({ title: '请完整填写预约信息', icon: 'none' })
    return
  }

  const slot = timeSlots.value.find(s => s.id === selectedTimeSlot.value)
  if (!slot) return

  const startTime = `${selectedDate.value} ${slot.start_time}:00`
  const endTime = `${selectedDate.value} ${slot.end_time}:00`

  uni.showLoading({ title: '提交中...' })

  try {
    await orderApi.create({
      room_id: selectedRoom.value,
      start_time: startTime,
      end_time: endTime,
      remark: remark.value
    })
    uni.hideLoading()
    uni.showToast({ title: '预约成功', icon: 'success' })
    setTimeout(() => {
      uni.switchTab({ url: '/pages/order/list' })
    }, 1500)
  } catch (e) {
    uni.hideLoading()
    console.error('预约失败', e)
  }
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
}

.loading-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
  padding: 40rpx;
}

.loading-text {
  margin-top: 20rpx;
  font-size: 28rpx;
  color: #999;
}

.content {
  padding-bottom: 140rpx;
}

.card {
  background: #fff;
  padding: 24rpx;
  margin-bottom: 20rpx;
}

.card-title {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 20rpx;
  display: block;
}

.date-picker {
  background: #f5f5f5;
  border-radius: 8rpx;
}

.picker-value {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24rpx;
  font-size: 28rpx;
  color: #333;
}

.date-tag {
  font-size: 22rpx;
  padding: 4rpx 12rpx;
  border-radius: 4rpx;
  margin-right: 16rpx;

  &.weekday {
    background: #e8f5e9;
    color: #4CAF50;
  }

  &.weekend {
    background: #fff3e0;
    color: #ff9800;
  }

  &.holiday {
    background: #fce4ec;
    color: #e91e63;
  }
}

.time-slots {
  display: flex;
  flex-wrap: wrap;
}

.time-slot-item {
  width: calc(50% - 10rpx);
  margin-right: 20rpx;
  margin-bottom: 20rpx;
  padding: 20rpx;
  background: #f5f5f5;
  border-radius: 12rpx;
  border: 2rpx solid transparent;
  transition: all 0.3s;

  &:nth-child(2n) {
    margin-right: 0;
  }

  &.active {
    border-color: #4CAF50;
    background: #e8f5e9;
  }

  &.unavailable {
    opacity: 0.5;
    pointer-events: none;
    background: #eee;
  }
}

.slot-name {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 8rpx;
}

.slot-time {
  font-size: 24rpx;
  color: #666;
  display: block;
  margin-bottom: 8rpx;
}

.slot-price-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.slot-price {
  font-size: 26rpx;
  color: #f44336;
  font-weight: 600;
}

.price-tag {
  font-size: 20rpx;
  padding: 2rpx 8rpx;
  background: #fff3e0;
  color: #ff9800;
  border-radius: 4rpx;
}

.room-list {
  display: flex;
  flex-direction: column;
}

.room-item {
  display: flex;
  align-items: center;
  padding: 20rpx;
  background: #f5f5f5;
  border-radius: 8rpx;
  margin-bottom: 16rpx;
  border: 2rpx solid transparent;
  transition: all 0.3s;

  &:last-child {
    margin-bottom: 0;
  }

  &.active {
    border-color: #4CAF50;
    background: #e8f5e9;
  }

  &.unavailable {
    opacity: 0.6;
    pointer-events: none;
  }
}

.room-info {
  flex: 1;
}

.room-name {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 4rpx;
}

.room-type {
  font-size: 24rpx;
  color: #666;
}

.room-status-badge {
  font-size: 22rpx;
  padding: 4rpx 12rpx;
  border-radius: 4rpx;
  margin-right: 16rpx;

  &.available {
    background: #e8f5e9;
    color: #4CAF50;
  }

  &.in-use {
    background: #fff3e0;
    color: #ff9800;
  }

  &.reserved {
    background: #e3f2fd;
    color: #2196f3;
  }

  &.maintenance {
    background: #f5f5f5;
    color: #999;
  }
}

.check-icon {
  margin-left: auto;
}

.empty-tip {
  width: 100%;
  text-align: center;
  padding: 40rpx;
  font-size: 26rpx;
  color: #999;
}

.remark-input {
  width: 100%;
  height: 160rpx;
  background: #f5f5f5;
  border-radius: 8rpx;
  padding: 20rpx;
  font-size: 26rpx;
  color: #333;
  box-sizing: border-box;
}

.remark-count {
  display: block;
  text-align: right;
  font-size: 22rpx;
  color: #999;
  margin-top: 10rpx;
}

.summary-card {
  background: #fff;
  padding: 24rpx;
  margin-bottom: 20rpx;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  padding: 16rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }

  &.total {
    padding-top: 24rpx;
    
    .summary-label {
      font-size: 28rpx;
      font-weight: 600;
    }

    .summary-value {
      font-size: 36rpx;
      font-weight: 600;
      color: #f44336;
    }
  }
}

.summary-label {
  font-size: 26rpx;
  color: #666;
}

.summary-value {
  font-size: 26rpx;
  color: #333;
}

.date-type {
  font-size: 22rpx;
  color: #ff9800;
  margin-left: 8rpx;
}

.bottom-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20rpx 30rpx;
  padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
  background: #fff;
  border-top: 1rpx solid #f0f0f0;
}

.price-display {
  display: flex;
  flex-direction: column;
}

.price-label {
  font-size: 22rpx;
  color: #999;
}

.price-amount {
  font-size: 36rpx;
  color: #f44336;
  font-weight: 600;
}

.submit-btn {
  padding: 24rpx 80rpx;
  background: #4CAF50;
  color: #fff;
  font-size: 30rpx;
  border-radius: 8rpx;
  transition: background 0.3s;

  &.disabled {
    background: #ccc;
  }
}
</style>