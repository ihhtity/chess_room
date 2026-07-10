<template>
  <view class="container">
    <view v-if="loading" class="loading-container">
      <uv-loading-icon mode="flower" size="48" color="#4CAF50" />
      <text class="loading-text">加载中...</text>
    </view>

    <view v-else-if="error" class="error-container">
      <uv-icon name="close-circle" color="#f44336" size="64" />
      <text class="error-text">{{ error }}</text>
      <view class="retry-btn" @click="loadRoom(currentRoomId)">重新加载</view>
    </view>

    <view v-else class="content">
      <swiper class="image-swiper" indicator-dots autoplay :circular="roomImages.length > 1">
        <swiper-item v-for="(img, index) in roomImages" :key="index">
          <image :src="img" mode="aspectFill" class="room-img" />
        </swiper-item>
      </swiper>

      <view class="room-info-card">
        <view class="room-title">
          <text class="room-name">{{ room?.name }}</text>
          <text class="room-status" :class="getStatusClass(room?.status || 0)">{{ getStatusText(room?.status || 0) }}</text>
        </view>
        <view class="room-basic">
          <view class="basic-item">
            <uv-icon name="home" color="#999" size="24" />
            <text>{{ room?.type.name }}</text>
          </view>
          <view class="basic-item">
            <uv-icon name="list" color="#999" size="24" />
            <text>{{ room?.floor }}</text>
          </view>
          <view class="basic-item">
            <uv-icon name="account" color="#999" size="24" />
            <text>{{ room?.capacity }}人</text>
          </view>
        </view>
      </view>

      <view class="card">
        <text class="card-title">设备清单</text>
        <view class="equipment-list">
          <view class="equipment-item" v-for="(item, index) in equipmentList" :key="index">
            <uv-icon name="check-circle" color="#4CAF50" size="24" />
            <text>{{ item }}</text>
          </view>
          <view v-if="equipmentList.length === 0" class="empty-tip">
            <text>暂无设备信息</text>
          </view>
        </view>
      </view>

      <view class="card">
        <text class="card-title">价格信息</text>
        <view class="price-info">
          <view class="price-item">
            <text class="price-label">基础价格</text>
            <text class="price-value">¥{{ room?.type.base_price }}/小时</text>
          </view>
          <view class="price-item">
            <text class="price-label">最大人数</text>
            <text class="price-value">{{ room?.type.max_people }}人</text>
          </view>
        </view>
      </view>

      <view class="card">
        <text class="card-title">包间描述</text>
        <text class="description-text">{{ room?.description || '暂无描述' }}</text>
      </view>

      <view class="bottom-bar">
        <view class="price-display">
          <text class="price-label">当前价格</text>
          <text class="price-amount">¥{{ room?.type.base_price }}/小时</text>
        </view>
        <view 
          class="book-btn" 
          :class="{ disabled: room?.status !== 1 }"
          @click="handleBook"
        >
          {{ room?.status === 1 ? '立即预约' : '暂不可用' }}
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { roomApi } from '@/api'
import type { Room } from '@/types'
import { RoomStatus } from '@/types'
import { getRandomImage } from '@/utils/image'

const room = ref<Room | null>(null)
const roomImages = ref<string[]>([])
const equipmentList = ref<string[]>([])
const loading = ref(true)
const error = ref('')
const currentRoomId = ref(0)

onMounted(async () => {
  const id = getRoomId()
  if (id) {
    currentRoomId.value = id
    await loadRoom(id)
  } else {
    loading.value = false
    error.value = '未找到房间ID'
  }
})

const getRoomId = (): number => {
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1] as any
  const options = currentPage?.options || {}
  
  if (options.id) {
    return parseInt(options.id as string) || 0
  }
  
  const launchOptions = uni.getLaunchOptionsSync()
  if (launchOptions.query?.id) {
    return parseInt(launchOptions.query.id as string) || 0
  }
  
  return 0
}

const loadRoom = async (id: number) => {
  loading.value = true
  error.value = ''
  
  try {
    const res = await roomApi.getDetail(id)
    room.value = res
    
    try {
      roomImages.value = JSON.parse(res.images) || []
    } catch {
      roomImages.value = [getRandomImage(`room${id}`, 800, 600)]
    }

    try {
      equipmentList.value = JSON.parse(res.equipment) || []
    } catch {
      equipmentList.value = []
    }
  } catch (e: any) {
    error.value = e?.message || '加载房间信息失败'
    console.error('加载房间失败', e)
  } finally {
    loading.value = false
  }
}

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

const handleBook = () => {
  if (room.value?.status !== 1) {
    uni.showToast({ title: '该包间暂不可预约', icon: 'none' })
    return
  }
  uni.navigateTo({ url: `/pages/booking/index?room_id=${room.value.id}` })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
}

.loading-container,
.error-container {
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

.error-text {
  margin-top: 20rpx;
  font-size: 28rpx;
  color: #f44336;
  text-align: center;
}

.retry-btn {
  margin-top: 30rpx;
  padding: 20rpx 48rpx;
  background: #4CAF50;
  color: #fff;
  font-size: 28rpx;
  border-radius: 8rpx;
}

.content {
  padding-bottom: 140rpx;
}

.image-swiper {
  width: 100%;
  height: 500rpx;
}

.room-img {
  width: 100%;
  height: 100%;
}

.room-info-card {
  background: #fff;
  padding: 30rpx;
  margin-bottom: 20rpx;
}

.room-title {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20rpx;
}

.room-name {
  font-size: 36rpx;
  font-weight: 600;
  color: #333;
}

.room-status {
  font-size: 24rpx;
  padding: 6rpx 16rpx;
  border-radius: 8rpx;

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

.room-basic {
  display: flex;
  justify-content: space-around;
}

.basic-item {
  display: flex;
  align-items: center;
  font-size: 26rpx;
  color: #666;

  text {
    margin-left: 8rpx;
  }
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

.equipment-list {
  display: flex;
  flex-wrap: wrap;
}

.equipment-item {
  display: flex;
  align-items: center;
  padding: 12rpx 20rpx;
  background: #f5f5f5;
  border-radius: 8rpx;
  margin: 0 16rpx 16rpx 0;
  font-size: 24rpx;
  color: #666;
}

.empty-tip {
  width: 100%;
  text-align: center;
  padding: 40rpx;
  font-size: 26rpx;
  color: #999;
}

.price-info {
  display: flex;
  flex-direction: column;
}

.price-item {
  display: flex;
  justify-content: space-between;
  padding: 12rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.price-label {
  font-size: 26rpx;
  color: #666;
}

.price-value {
  font-size: 26rpx;
  color: #f44336;
  font-weight: 600;
}

.description-text {
  font-size: 26rpx;
  color: #666;
  line-height: 1.6;
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

.book-btn {
  padding: 24rpx 64rpx;
  background: #4CAF50;
  color: #fff;
  font-size: 30rpx;
  border-radius: 8rpx;

  &.disabled {
    background: #ccc;
  }
}
</style>