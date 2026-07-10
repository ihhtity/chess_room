<template>
  <view class="container">
    <view class="header">
      <text class="back-btn" @click="goBack">←</text>
      <text class="header-title">活动中心</text>
      <view class="placeholder"></view>
    </view>

    <view class="activity-list">
      <view class="activity-card" v-for="activity in activities" :key="activity.id" @click="goToDetail(activity.id)">
        <image :src="activity.image || getActivityImage(activity.id)" mode="aspectFill" class="activity-img" />
        <view class="activity-overlay">
          <text class="activity-tag">限时活动</text>
        </view>
        <view class="activity-info">
          <text class="activity-name">{{ activity.name }}</text>
          <text class="activity-discount">折扣 {{ (activity.discount * 10).toFixed(0) }}折</text>
          <text class="activity-time">有效期：{{ formatDate(activity.valid_from) }} - {{ formatDate(activity.valid_to) }}</text>
        </view>
      </view>
    </view>

    <view v-if="activities.length === 0" class="empty">
      <uv-icon name="inbox" color="#ccc" size="80" />
      <text class="empty-text">暂无活动</text>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { activityApi } from '@/api'
import type { Activity } from '@/types'
import { getActivityImage } from '@/utils/image'

const activities = ref<Activity[]>([])

onMounted(async () => {
  await loadActivities()
})

const loadActivities = async () => {
  const res = await activityApi.getList()
  activities.value = res.data
}

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}月${date.getDate()}日`
}

const goBack = () => {
  uni.navigateBack()
}

const goToDetail = (id: number) => {
  uni.navigateTo({ url: `/pages/activities/detail?id=${id}` })
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

.activity-list {
  padding: 20rpx;
}

.activity-card {
  background: #fff;
  border-radius: 16rpx;
  overflow: hidden;
  margin-bottom: 20rpx;
  position: relative;
}

.activity-img {
  width: 100%;
  height: 300rpx;
}

.activity-overlay {
  position: absolute;
  top: 20rpx;
  left: 20rpx;
}

.activity-tag {
  padding: 8rpx 20rpx;
  background: rgba(244, 67, 54, 0.9);
  color: #fff;
  font-size: 22rpx;
  border-radius: 8rpx;
}

.activity-info {
  padding: 24rpx;
}

.activity-name {
  font-size: 32rpx;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 12rpx;
}

.activity-discount {
  font-size: 28rpx;
  color: #f44336;
  font-weight: 600;
  display: block;
  margin-bottom: 8rpx;
}

.activity-time {
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
