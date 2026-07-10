<template>
  <view class="container">
    <!-- banner 轮播图 -->
    <swiper class="banner" indicator-dots autoplay circular>
      <swiper-item v-for="item in banners" :key="item.id">
        <image :src="item.image" mode="aspectFill" class="banner-img" />
      </swiper-item>
    </swiper>
    <!-- 快捷入口 -->
    <view class="quick-entry">
      <view class="entry-item" v-for="item in quickEntry" :key="item.key" @click="handleEntry(item.path)">
        <view class="entry-icon">
          <uv-icon :name="item.icon" color="#4CAF50" size="48" />
        </view>
        <text class="entry-text">{{ item.text }}</text>
      </view>
    </view>
    <!-- 推荐包间 -->
    <view class="section">
      <view class="section-header">
        <text class="section-title">推荐包间</text>
        <text class="section-more" @click="goToRoomList">更多</text>
      </view>
      <scroll-view scroll-x class="room-scroll">
        <view class="room-card" v-for="room in rooms" :key="room.id" @click="goToRoomDetail(room.id)">
          <image :src="getRoomImage(room.images, room.id)" mode="aspectFill" class="room-img" />
          <view class="room-info">
            <text class="room-name">{{ room.name }}</text>
            <text class="room-type">{{ room.type.name }}</text>
            <text class="room-price">¥{{ room.type.base_price }}/小时</text>
          </view>
        </view>
      </scroll-view>
    </view>
    <!-- 最新活动 -->
    <view class="section">
      <view class="section-header">
        <text class="section-title">最新活动</text>
        <text class="section-more" @click="goToActivities">更多</text>
      </view>
      <view class="activity-list">
        <view class="activity-item" v-for="activity in activities" :key="activity.id">
          <image :src="activity.image || getActivityImage(activity.id)" mode="aspectFill" class="activity-img" />
          <view class="activity-info">
            <text class="activity-name">{{ activity.name }}</text>
            <text class="activity-discount">折扣 {{ (activity.discount * 10).toFixed(0) }}折</text>
            <text class="activity-time">{{ formatDate(activity.valid_from) }} - {{ formatDate(activity.valid_to) }}</text>
          </view>
        </view>
      </view>
    </view>
    <!-- 公告 -->
    <view class="section">
      <view class="section-header">
        <text class="section-title">公告</text>
        <text class="section-more" @click="goToAnnouncements">更多</text>
      </view>
      <view class="announcement-list">
        <view class="announcement-item" v-for="item in announcements" :key="item.id">
          <text class="announcement-title">{{ item.title }}</text>
          <text class="announcement-time">{{ formatDate(item.created_at) }}</text>
        </view>
      </view>
    </view>
    <!-- 底部导航栏 -->
    <uv-tabbar :value="0" @change="handleTabChange">
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
import { roomApi, activityApi, announcementApi } from '@/api'
import type { Room, Activity, Announcement } from '@/types'
import { getRandomImage, getActivityImage } from '@/utils/image'

const banners = ref([
  { id: 1, image: getRandomImage('banner1', 800, 450) },
  { id: 2, image: getRandomImage('banner2', 800, 450) },
  { id: 3, image: getRandomImage('banner3', 800, 450) }
])

const quickEntry = ref([
  { key: 'room', text: '包间', icon: 'home', path: '/pages/room/list' },
  { key: 'booking', text: '预约', icon: 'calendar', path: '/pages/booking/index' },
  { key: 'order', text: '订单', icon: 'order', path: '/pages/order/list' },
  { key: 'member', text: '会员', icon: 'empty-coupon', path: '/pages/member/index' },
  { key: 'notification', text: '消息', icon: 'bell', path: '/pages/notifications/index' }
])

const rooms = ref<Room[]>([])
const activities = ref<Activity[]>([])
const announcements = ref<Announcement[]>([])

onMounted(async () => {
  await loadData()
})

const loadData = async () => {
  try {
    const [roomRes, activityRes, announcementRes] = await Promise.all([
      roomApi.getList({ page: 1, page_size: 5 }),
      activityApi.getList(),
      announcementApi.getList()
    ])
    rooms.value = roomRes?.data || []
    activities.value = (activityRes?.data || []).slice(0, 3)
    announcements.value = (announcementRes?.data || []).slice(0, 3)
  } catch (e) {
    console.error('加载首页数据失败', e)
    rooms.value = []
    activities.value = []
    announcements.value = []
  }
}

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

const getRoomImage = (images: string, id: number) => {
  try {
    const imgList = JSON.parse(images)
    return imgList[0] || getRandomImage(`room${id}`, 280, 200)
  } catch {
    return getRandomImage(`room${id}`, 280, 200)
  }
}

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}/${date.getDate()}`
}

const handleEntry = (path: string) => {
  uni.switchTab({ url: path }).catch(() => {
    uni.navigateTo({ url: path })
  })
}

const goToRoomList = () => {
  uni.switchTab({ url: '/pages/room/list' })
}

const goToRoomDetail = (id: number) => {
  uni.navigateTo({ url: `/pages/room/detail?id=${id}` })
}

const goToActivities = () => {
  uni.navigateTo({ url: '/pages/activities/index' })
}

const goToAnnouncements = () => {
  uni.navigateTo({ url: '/pages/announcements/index' })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
}

.banner {
  width: 100%;
  height: 360rpx;
  background: #fff;
}

.banner-img {
  width: 100%;
  height: 100%;
}

.quick-entry {
  display: flex;
  justify-content: space-around;
  background: #fff;
  padding: 30rpx 0;
  margin-bottom: 20rpx;
}

.entry-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.entry-icon {
  width: 100rpx;
  height: 100rpx;
  background: #f0f9f0;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 16rpx;
}

.entry-text {
  font-size: 24rpx;
  color: #333;
}

.section {
  background: #fff;
  margin-bottom: 20rpx;
  padding: 24rpx;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20rpx;
}

.section-title {
  font-size: 32rpx;
  font-weight: 600;
  color: #333;
}

.section-more {
  font-size: 26rpx;
  color: #999;
}

.room-scroll {
  white-space: nowrap;
}

.room-card {
  display: inline-block;
  width: 280rpx;
  margin-right: 20rpx;
  background: #fafafa;
  border-radius: 12rpx;
  overflow: hidden;
}

.room-img {
  width: 280rpx;
  height: 200rpx;
}

.room-info {
  padding: 16rpx;
}

.room-name {
  font-size: 26rpx;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 8rpx;
}

.room-type {
  font-size: 22rpx;
  color: #999;
  display: block;
  margin-bottom: 8rpx;
}

.room-price {
  font-size: 24rpx;
  color: #f44336;
  font-weight: 600;
}

.activity-list {
  display: flex;
  flex-direction: column;
}

.activity-item {
  display: flex;
  padding: 16rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.activity-img {
  width: 160rpx;
  height: 120rpx;
  border-radius: 8rpx;
  margin-right: 20rpx;
}

.activity-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.activity-name {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 8rpx;
}

.activity-discount {
  font-size: 24rpx;
  color: #f44336;
  margin-bottom: 8rpx;
}

.activity-time {
  font-size: 22rpx;
  color: #999;
}

.announcement-list {
  display: flex;
  flex-direction: column;
}

.announcement-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.announcement-title {
  font-size: 26rpx;
  color: #333;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.announcement-time {
  font-size: 22rpx;
  color: #999;
  margin-left: 20rpx;
}
</style>
