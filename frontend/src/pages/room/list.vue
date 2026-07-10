<template>
  <view class="container">
    <!-- 搜索栏 -->
    <view class="search-bar">
      <view class="search-input-wrap">
        <uv-icon name="search" color="#999" size="28" />
        <input class="search-input" placeholder="搜索包间名称" v-model="searchName" @confirm="handleSearch" />
      </view>
    </view>
    <!-- 包间筛选标签 -->
    <view class="filter-tabs">
      <scroll-view scroll-x class="tabs-scroll">
        <view class="tabs">
          <view 
            class="tab-item" 
            :class="{ active: selectedType === 0 }"
            @click="selectedType = 0; loadRooms()"
          >
            全部
          </view>
          <view 
            class="tab-item" 
            :class="{ active: selectedType === type.id }"
            v-for="type in roomTypes" 
            :key="type.id"
            @click="selectedType = type.id; loadRooms()"
          >
            {{ type.name }}
          </view>
        </view>
      </scroll-view>
    </view>
    <!-- 包间列表卡片 -->
    <view class="room-list">
      <view class="room-card" v-for="room in rooms" :key="room.id" @click="goToDetail(room.id)">
        <image :src="getRoomImage(room.images, room.id)" mode="aspectFill" class="room-img" />
        <view class="room-info">
          <view class="room-header">
            <text class="room-name">{{ room.name }}</text>
            <text class="room-status" :class="getStatusClass(room.status)">{{ getStatusText(room.status) }}</text>
          </view>
          <text class="room-type">{{ room.type.name }} · {{ room.floor }} · {{ room.capacity }}人</text>
          <text class="room-equipment">{{ getEquipmentText(room.equipment) }}</text>
          <view class="room-footer">
            <text class="room-price">¥{{ room.type.base_price }}/小时</text>
            <view class="book-btn" :class="{ disabled: room.status !== 1 }" @click.stop="handleBook(room)">
              {{ room.status === 1 ? '立即预约' : '暂不可用' }}
            </view>
          </view>
        </view>
      </view>
    </view>
    <!-- 包间列表卡片 -->
    <view v-if="rooms.length === 0" class="empty">
      <uv-icon name="home" color="#ccc" size="80" />
      <text class="empty-text">暂无包间</text>
    </view>
    <!-- 底部导航栏 -->
    <uv-tabbar :value="1" @change="handleTabChange">
      <uv-tabbar-item text="首页" icon="home" />
      <uv-tabbar-item text="包间" icon="integral" />
      <uv-tabbar-item text="订单" icon="order" />
      <uv-tabbar-item text="会员" icon="empty-coupon" />
      <uv-tabbar-item text="我的" icon="account" />
    </uv-tabbar>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { roomApi } from '@/api'
import type { Room, RoomType } from '@/types'
import { RoomStatus } from '@/types'
import { getRandomImage } from '@/utils/image'

const rooms = ref<Room[]>([])
const roomTypes = ref<RoomType[]>([])
const selectedType = ref(0)
const searchName = ref('')

onMounted(async () => {
  await loadRoomTypes()
  await loadRooms()
})

const handleTabChange = (index: number) => {
  switch (index) {
    case 0:
      uni.redirectTo({ url: '/pages/index/index' })
      break
    case 1:
      uni.redirectTo({ url: '/pages/room/list' })
      break
    case 2:
      uni.redirectTo({ url: '/pages/order/list' })
      break
    case 3:
      uni.redirectTo({ url: '/pages/member/index' })
      break
    default:
      uni.redirectTo({ url: '/pages/user/index' })
      break
  }
}

const loadRoomTypes = async () => {
  const res = await roomApi.getTypes()
  roomTypes.value = res.data
}

const loadRooms = async () => {
  const params: any = {
    page: 1,
    page_size: 20
  }
  if (selectedType.value > 0) {
    params.type_id = selectedType.value
  }
  if (searchName.value) {
    params.name = searchName.value
  }
  const res = await roomApi.getList(params)
  rooms.value = res.data
}

const handleSearch = () => {
  loadRooms()
}

const getRoomImage = (images: string, id: number) => {
  try {
    const imgList = JSON.parse(images)
    return imgList[0] || getRandomImage(`room${id}`, 280, 280)
  } catch {
    return getRandomImage(`room${id}`, 280, 280)
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

const getEquipmentText = (equipment: string) => {
  try {
    const eqList = JSON.parse(equipment)
    return eqList.slice(0, 3).join('、') + (eqList.length > 3 ? '...' : '')
  } catch {
    return equipment
  }
}

const goToDetail = (id: number) => {
  uni.navigateTo({ url: `/pages/room/detail?id=${id}` })
}

const handleBook = (room: Room) => {
  if (room.status !== 1) {
    uni.showToast({ title: '该包间暂不可预约', icon: 'none' })
    return
  }
  uni.navigateTo({ url: `/pages/booking/index?room_id=${room.id}` })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 120rpx;
}

.search-bar {
  background: #fff;
  padding: 20rpx;
}

.search-input-wrap {
  display: flex;
  align-items: center;
  background: #f5f5f5;
  border-radius: 40rpx;
  padding: 0 24rpx;
}

.search-input {
  flex: 1;
  height: 72rpx;
  font-size: 28rpx;
  padding-left: 16rpx;
}

.filter-tabs {
  background: #fff;
  padding: 20rpx 0;
  border-bottom: 1rpx solid #f0f0f0;
}

.tabs-scroll {
  white-space: nowrap;
}

.tabs {
  display: inline-flex;
  padding: 0 20rpx;
}

.tab-item {
  padding: 12rpx 32rpx;
  font-size: 28rpx;
  color: #666;
  border-radius: 32rpx;
  margin-right: 16rpx;
  background: #f5f5f5;

  &.active {
    background: #4CAF50;
    color: #fff;
  }
}

.room-list {
  padding: 20rpx;
}

.room-card {
  display: flex;
  background: #fff;
  border-radius: 16rpx;
  overflow: hidden;
  margin-bottom: 20rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.06);
}

.room-img {
  width: 280rpx;
  height: 280rpx;
  flex-shrink: 0;
}

.room-info {
  flex: 1;
  padding: 20rpx;
  display: flex;
  flex-direction: column;
}

.room-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12rpx;
}

.room-name {
  font-size: 32rpx;
  font-weight: 600;
  color: #333;
}

.room-status {
  font-size: 22rpx;
  padding: 4rpx 12rpx;
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

.room-type {
  font-size: 24rpx;
  color: #666;
  margin-bottom: 8rpx;
}

.room-equipment {
  font-size: 22rpx;
  color: #999;
  margin-bottom: 16rpx;
  flex: 1;
}

.room-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.room-price {
  font-size: 32rpx;
  font-weight: 600;
  color: #f44336;
}

.book-btn {
  padding: 12rpx 32rpx;
  background: #4CAF50;
  color: #fff;
  font-size: 26rpx;
  border-radius: 8rpx;

  &.disabled {
    background: #ccc;
  }
}

.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 100rpx 0;
}

.empty-text {
  font-size: 28rpx;
  color: #999;
  margin-top: 20rpx;
}
</style>
