<template>
  <view class="container">
    <view class="header">
      <text class="back-btn" @click="goBack">←</text>
      <text class="header-title">活动详情</text>
      <view class="placeholder"></view>
    </view>

    <view class="activity-detail" v-if="activity">
      <image :src="activity.image || getActivityImage(activity.id)" mode="aspectFill" class="activity-img" />

      <view class="detail-card">
        <text class="activity-name">{{ activity.name }}</text>
        <view class="activity-meta">
          <text class="discount-tag">折扣 {{ (activity.discount * 10).toFixed(0) }}折</text>
          <text class="valid-time">有效期：{{ formatDate(activity.valid_from) }} - {{ formatDate(activity.valid_to) }}</text>
        </view>
      </view>

      <view class="detail-card">
        <text class="card-title">活动规则</text>
        <text class="rule-text">{{ activity.description }}</text>
      </view>

      <view class="detail-card">
        <text class="card-title">适用范围</text>
        <view class="applicable-list">
          <view class="applicable-item" v-for="(item, index) in applicableItems" :key="index">
            <uv-icon name="check-circle" color="#4CAF50" size="24" />
            <text>{{ item }}</text>
          </view>
        </view>
      </view>

      <view class="detail-card">
        <text class="card-title">注意事项</text>
        <view class="notice-list">
          <text class="notice-item" v-for="(item, index) in noticeItems" :key="index">{{ item }}</text>
        </view>
      </view>

      <view class="bottom-bar">
        <view class="book-btn" @click="goToBooking">立即预约</view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { activityApi } from '@/api'
import type { Activity } from '@/types'
import { getActivityImage } from '@/utils/image'

const activity = ref<Activity | null>(null)

const applicableItems = ref([
  '所有普通包间',
  '所有豪华包间',
  '所有VIP包间'
])

const noticeItems = ref([
  '活动期间每人每天限使用一次',
  '不可与其他优惠叠加使用',
  '最终解释权归商家所有'
])

onMounted(async () => {
  const options = uni.getLaunchOptionsSync() || { query: {} }
  const id = parseInt(options.query.id as string) || parseInt((uni as any).$page?.options?.id as string) || 0
  if (id) {
    await loadActivity(id)
  }
})

const loadActivity = async (id: number) => {
  const res = await activityApi.getDetail(id)
  activity.value = res
}

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`
}

const goBack = () => {
  uni.navigateBack()
}

const goToBooking = () => {
  uni.navigateTo({ url: '/pages/booking/index' })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 140rpx;
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

.activity-img {
  width: 100%;
  height: 400rpx;
}

.detail-card {
  background: #fff;
  padding: 24rpx;
  margin: 20rpx;
  border-radius: 16rpx;
}

.activity-name {
  font-size: 36rpx;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 20rpx;
}

.activity-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 16rpx;
}

.discount-tag {
  padding: 8rpx 20rpx;
  background: rgba(244, 67, 54, 0.1);
  color: #f44336;
  font-size: 24rpx;
  border-radius: 8rpx;
  font-weight: 600;
}

.valid-time {
  font-size: 26rpx;
  color: #666;
}

.card-title {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 20rpx;
  display: block;
}

.rule-text {
  font-size: 28rpx;
  color: #666;
  line-height: 1.8;
}

.applicable-list {
  display: flex;
  flex-wrap: wrap;
}

.applicable-item {
  display: flex;
  align-items: center;
  padding: 12rpx 20rpx;
  background: #f5f5f5;
  border-radius: 8rpx;
  margin: 0 16rpx 16rpx 0;
  font-size: 26rpx;
  color: #666;
}

.notice-list {
  display: flex;
  flex-direction: column;
}

.notice-item {
  font-size: 26rpx;
  color: #999;
  line-height: 1.8;
  padding-left: 24rpx;
  position: relative;

  &::before {
    content: '•';
    position: absolute;
    left: 0;
    color: #ccc;
  }
}

.bottom-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 20rpx 30rpx;
  padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
  background: #fff;
  border-top: 1rpx solid #f0f0f0;
}

.book-btn {
  padding: 28rpx;
  text-align: center;
  background: #4CAF50;
  color: #fff;
  font-size: 32rpx;
  border-radius: 8rpx;
}
</style>
