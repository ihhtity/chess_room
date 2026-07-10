<template>
  <view class="container">
    <view class="header">
      <text class="back-btn" @click="goBack">←</text>
      <text class="header-title">公告详情</text>
      <view class="placeholder"></view>
    </view>

    <view class="detail-content" v-if="announcement">
      <text class="announcement-title">{{ announcement.title }}</text>
      <text class="announcement-time">{{ formatDate(announcement.created_at) }}</text>
      <view class="announcement-body">
        <text>{{ announcement.content }}</text>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { announcementApi } from '@/api'
import type { Announcement } from '@/types'

const announcement = ref<Announcement | null>(null)

onMounted(async () => {
  const options = uni.getLaunchOptionsSync() || { query: {} }
  const id = parseInt(options.query.id as string) || parseInt((uni as any).$page?.options?.id as string) || 0
  if (id) {
    await loadAnnouncement(id)
  }
})

const loadAnnouncement = async (id: number) => {
  const res = await announcementApi.getDetail(id)
  announcement.value = res
}

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日 ${date.getHours()}:${String(date.getMinutes()).padStart(2, '0')}`
}

const goBack = () => {
  uni.navigateBack()
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

.detail-content {
  background: #fff;
  margin: 20rpx;
  padding: 30rpx;
  border-radius: 16rpx;
}

.announcement-title {
  font-size: 36rpx;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 20rpx;
}

.announcement-time {
  font-size: 26rpx;
  color: #999;
  display: block;
  margin-bottom: 30rpx;
}

.announcement-body {
  font-size: 28rpx;
  color: #666;
  line-height: 1.8;
}
</style>
