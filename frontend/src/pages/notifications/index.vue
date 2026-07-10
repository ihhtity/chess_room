<template>
  <view class="container">
    <view class="notification-list">
      <!-- 消息操作按钮 -->
      <view class="header">
        <view class="header-btn read-btn" @click="handleMarkAllRead">
          <uv-icon name="checkmark-circle" color="#4CAF50" size="28" />
          <text>全部已读</text>
        </view>
        <view class="header-btn clear-btn" @click="handleClear">
          <uv-icon name="trash" color="#f44336" size="28" />
          <text>清空消息</text>
        </view>
      </view>
      <!-- 消息列表 -->
      <view 
        class="notification-item" 
        :class="{ unread: !notification.read }"
        v-for="notification in notifications" 
        :key="notification.id"
        @click="handleClick(notification)"
      >
        <view class="notification-icon">
          <uv-icon :name="getIcon(notification.type)" :color="getIconColor(notification.type)" size="36" />
        </view>
        <view class="notification-content">
          <text class="notification-title">{{ notification.title }}</text>
          <text class="notification-desc">{{ notification.content }}</text>
          <text class="notification-time">{{ formatTime(notification.created_at) }}</text>
        </view>
        <view class="unread-dot" v-if="!notification.read"></view>
      </view>
    </view>
    <!-- 消息列表为空时的提示 -->
    <view v-if="notifications.length === 0" class="empty">
      <uv-icon name="bell" color="#ccc" size="80" />
      <text class="empty-text">暂无消息</text>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { notificationApi } from '@/api'
import type { Notification } from '@/types'

const notifications = ref<Notification[]>([])

onMounted(async () => {
  await loadNotifications()
})

const loadNotifications = async () => {
  const res = await notificationApi.getList({ page: 1, page_size: 20 })
  notifications.value = res.data
}

const getIcon = (type: number) => {
  switch (type) {
    case 1:
      return 'list'
    case 2:
      return 'gift'
    case 3:
      return 'setting'
    default:
      return 'bell'
  }
}

const getIconColor = (type: number) => {
  switch (type) {
    case 1:
      return '#4CAF50'
    case 2:
      return '#ff9800'
    case 3:
      return '#2196f3'
    default:
      return '#666'
  }
}

const formatTime = (dateStr: string) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const minutes = Math.floor(diff / (1000 * 60))
  const hours = Math.floor(diff / (1000 * 60 * 60))
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))

  if (minutes < 1) return '刚刚'
  if (minutes < 60) return `${minutes}分钟前`
  if (hours < 24) return `${hours}小时前`
  if (days < 7) return `${days}天前`
  return `${date.getMonth() + 1}/${date.getDate()}`
}

const goBack = () => {
  uni.navigateBack()
}

const handleClick = (notification: Notification) => {
  uni.showToast({ title: notification.content, icon: 'none' })
}

const handleMarkAllRead = () => {
  notifications.value.forEach(n => {
    n.read = true
  })
  uni.showToast({ title: '已全部标记为已读', icon: 'success' })
}

const handleClear = () => {
  uni.showModal({
    title: '提示',
    content: '确定要清空所有消息吗？',
    success: (res) => {
      if (res.confirm) {
        notifications.value = []
        uni.showToast({ title: '清空成功', icon: 'success' })
      }
    }
  })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
}

.header {
  display: flex;
  justify-content: space-between;
  padding: 20rpx 0 30rpx;
}

.header-btn {
  width: 40%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16rpx 30rpx;
  border-radius: 30rpx;
  font-size: 26rpx;

  text {
    margin-left: 8rpx;
  }
}

.read-btn {
  background: #f0fff0;
  color: #4CAF50;
}

.clear-btn {
  background: #fff5f5;
  color: #f44336;
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

.clear-btn {
  font-size: 28rpx;
  color: #f44336;
}

.notification-list {
  padding: 20rpx;
}

.notification-item {
  display: flex;
  align-items: center;
  background: #fff;
  padding: 24rpx;
  border-radius: 16rpx;
  margin-bottom: 16rpx;

  &.unread {
    background: #fafafa;
  }
}

.notification-icon {
  width: 64rpx;
  height: 64rpx;
  background: #f5f5f5;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 20rpx;
}

.notification-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.notification-title {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 8rpx;
}

.notification-desc {
  font-size: 26rpx;
  color: #666;
  margin-bottom: 8rpx;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.notification-time {
  font-size: 22rpx;
  color: #999;
}

.unread-dot {
  width: 16rpx;
  height: 16rpx;
  background: #f44336;
  border-radius: 50%;
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
