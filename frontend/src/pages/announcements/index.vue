<template>
  <view class="container">
    <view class="header">
      <text class="back-btn" @click="goBack">←</text>
      <text class="header-title">公告通知</text>
      <view class="placeholder"></view>
    </view>

    <view class="announcement-list">
      <view class="announcement-item" v-for="item in announcements" :key="item.id" @click="goToDetail(item.id)">
        <view class="announcement-icon">
          <uv-icon name="info-circle" color="#2196f3" size="32" />
        </view>
        <view class="announcement-content">
          <text class="announcement-title">{{ item.title }}</text>
          <text class="announcement-time">{{ formatDate(item.created_at) }}</text>
        </view>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>
    </view>

    <view v-if="announcements.length === 0" class="empty">
      <uv-icon name="inbox" color="#ccc" size="80" />
      <text class="empty-text">暂无公告</text>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { announcementApi } from '@/api'
import type { Announcement } from '@/types'

const announcements = ref<Announcement[]>([])

onMounted(async () => {
  await loadAnnouncements()
})

const loadAnnouncements = async () => {
  const res = await announcementApi.getList()
  announcements.value = res.data
}

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}月${date.getDate()}日`
}

const goBack = () => {
  uni.navigateBack()
}

const goToDetail = (id: number) => {
  uni.navigateTo({ url: `/pages/announcements/detail?id=${id}` })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 40rpx 30rpx;
  background: #fff;
}

.back-btn {
  font-size: 40rpx;
  color: #333;
}

.header-title {
  font-size: 34rpx;
  font-weight: 600;
  color: #333;
}

.placeholder {
  width: 40rpx;
}

.announcement-list {
  padding: 20rpx;
}

.announcement-item {
  display: flex;
  align-items: center;
  background: #fff;
  padding: 24rpx;
  border-radius: 16rpx;
  margin-bottom: 16rpx;
}

.announcement-icon {
  width: 56rpx;
  height: 56rpx;
  background: #e3f2fd;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 20rpx;
}

.announcement-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.announcement-title {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 8rpx;
}

.announcement-time {
  font-size: 24rpx;
  color: #999;
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
