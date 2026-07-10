<template>
  <view class="container">
    <view class="status-card" :class="getStatusClass(order?.status || 0)">
      <text class="status-icon">{{ getStatusIcon(order?.status || 0) }}</text>
      <text class="status-text">{{ getStatusText(order?.status || 0) }}</text>
    </view>

    <view class="card">
      <text class="card-title">订单信息</text>
      <view class="info-list">
        <view class="info-item">
          <text class="info-label">订单号</text>
          <text class="info-value">{{ order?.order_no }}</text>
        </view>
        <view class="info-item">
          <text class="info-label">下单时间</text>
          <text class="info-value">{{ formatDateTime(order?.created_at) }}</text>
        </view>
      </view>
    </view>

    <view class="card">
      <text class="card-title">包间信息</text>
      <view class="room-detail" @click="goToRoomDetail(order?.room.id || 0)">
        <image :src="getRoomImage(order?.room.images || '', order?.room.id || 0)" mode="aspectFill" class="room-img" />
        <view class="room-info">
          <text class="room-name">{{ order?.room.name }}</text>
          <text class="room-type">{{ order?.room.type.name }} · {{ order?.room.floor }}</text>
          <text class="room-capacity">{{ order?.room.capacity }}人</text>
        </view>
      </view>
    </view>

    <view class="card">
      <text class="card-title">预约时间</text>
      <view class="time-info">
        <view class="time-item">
          <text class="time-label">开始时间</text>
          <text class="time-value">{{ formatDateTime(order?.start_time || '') }}</text>
        </view>
        <view class="time-divider"></view>
        <view class="time-item">
          <text class="time-label">结束时间</text>
          <text class="time-value">{{ formatDateTime(order?.end_time || '') }}</text>
        </view>
      </view>
      <text class="duration-text">时长：{{ getDuration(order?.start_time, order?.end_time) }}</text>
    </view>

    <view class="card">
      <text class="card-title">费用信息</text>
      <view class="price-list">
        <view class="price-item">
          <text class="price-label">订单金额</text>
          <text class="price-value">¥{{ order?.total_amount }}</text>
        </view>
        <view class="price-item">
          <text class="price-label">已支付</text>
          <text class="price-value">¥{{ order?.paid_amount }}</text>
        </view>
        <view class="price-item total">
          <text class="price-label">待支付</text>
          <text class="price-value">¥{{ (order?.total_amount || 0) - (order?.paid_amount || 0) }}</text>
        </view>
      </view>
    </view>

    <view class="card" v-if="order?.remark">
      <text class="card-title">订单备注</text>
      <text class="remark-text">{{ order?.remark }}</text>
    </view>

    <view class="bottom-bar" v-if="order?.status === 0">
      <view class="cancel-btn" @click="handleCancel">取消订单</view>
      <view class="pay-btn" @click="handlePay">立即支付</view>
    </view>

    <view class="bottom-bar" v-if="order?.status === 1">
      <view class="confirm-btn" @click="handleConfirm">确认使用</view>
      <view class="complete-btn" @click="handleComplete">完成订单</view>
    </view>

    <view class="bottom-bar" v-if="order?.status === 2">
      <view class="evaluate-btn" @click="handleEvaluate">去评价</view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { orderApi } from '@/api'
import type { Order } from '@/types'
import { OrderStatus } from '@/types'
import { getRandomImage } from '@/utils/image'

const order = ref<Order | null>(null)

onMounted(async () => {
  const options = uni.getLaunchOptionsSync() || { query: {} }
  const id = parseInt(options.query.id as string) || parseInt((uni as any).$page?.options?.id as string) || 0
  if (id) {
    await loadOrder(id)
  }
})

const loadOrder = async (id: number) => {
  const res = await orderApi.getDetail(id)
  order.value = res
}

const getStatusText = (status: number) => {
  switch (status) {
    case OrderStatus.Pending:
      return '待支付'
    case OrderStatus.Active:
      return '使用中'
    case OrderStatus.Completed:
      return '已完成'
    case OrderStatus.Cancelled:
      return '已取消'
    case OrderStatus.Refunding:
      return '退款中'
    case OrderStatus.Refunded:
      return '已退款'
    default:
      return '未知'
  }
}

const getStatusClass = (status: number) => {
  switch (status) {
    case OrderStatus.Pending:
      return 'pending'
    case OrderStatus.Active:
      return 'active'
    case OrderStatus.Completed:
      return 'completed'
    case OrderStatus.Cancelled:
      return 'cancelled'
    default:
      return ''
  }
}

const getStatusIcon = (status: number) => {
  switch (status) {
    case OrderStatus.Pending:
      return '⏳'
    case OrderStatus.Active:
      return '✅'
    case OrderStatus.Completed:
      return '🎯'
    case OrderStatus.Cancelled:
      return '❌'
    default:
      return '❓'
  }
}

const getRoomImage = (images: string, id: number) => {
  try {
    const imgList = JSON.parse(images)
    return imgList[0] || getRandomImage(`room${id}`, 160, 160)
  } catch {
    return getRandomImage(`room${id}`, 160, 160)
  }
}

const formatDateTime = (dateStr?: string) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  return `${month}月${day}日 ${hours}:${minutes}`
}

const getDuration = (startTime?: string, endTime?: string) => {
  if (!startTime || !endTime) return ''
  const start = new Date(startTime)
  const end = new Date(endTime)
  const diff = Math.abs(end.getTime() - start.getTime())
  const hours = Math.floor(diff / (1000 * 60 * 60))
  const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60))
  return `${hours}小时${minutes}分钟`
}

const goToRoomDetail = (id: number) => {
  if (id > 0) {
    uni.navigateTo({ url: `/pages/room/detail?id=${id}` })
  }
}

const handleCancel = async () => {
  uni.showModal({
    title: '提示',
    content: '确定要取消订单吗？',
    success: async (res) => {
      if (res.confirm && order.value) {
        await orderApi.cancel(order.value.id)
        uni.showToast({ title: '取消成功', icon: 'success' })
        setTimeout(() => {
          uni.switchTab({ url: '/pages/order/list' })
        }, 1500)
      }
    }
  })
}

const handlePay = () => {
  uni.showToast({ title: '支付功能开发中', icon: 'none' })
}

const handleConfirm = async () => {
  uni.showModal({
    title: '提示',
    content: '确定要确认使用吗？',
    success: async (res) => {
      if (res.confirm && order.value) {
        await orderApi.confirm(order.value.id)
        uni.showToast({ title: '确认成功', icon: 'success' })
        setTimeout(() => {
          uni.switchTab({ url: '/pages/order/list' })
        }, 1500)
      }
    }
  })
}

const handleComplete = async () => {
  uni.showModal({
    title: '提示',
    content: '确定要完成订单吗？',
    success: async (res) => {
      if (res.confirm && order.value) {
        await orderApi.complete(order.value.id)
        uni.showToast({ title: '完成成功', icon: 'success' })
        setTimeout(() => {
          uni.switchTab({ url: '/pages/order/list' })
        }, 1500)
      }
    }
  })
}

const handleEvaluate = () => {
  uni.showToast({ title: '评价功能开发中', icon: 'none' })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 140rpx;
}

.status-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40rpx;
  margin-bottom: 20rpx;

  &.pending {
    background: linear-gradient(135deg, #fff3e0, #ffe0b2);
  }

  &.active {
    background: linear-gradient(135deg, #e8f5e9, #c8e6c9);
  }

  &.completed {
    background: linear-gradient(135deg, #e0e0e0, #bdbdbd);
  }

  &.cancelled {
    background: linear-gradient(135deg, #f5f5f5, #e0e0e0);
  }
}

.status-icon {
  font-size: 80rpx;
  margin-bottom: 16rpx;
}

.status-text {
  font-size: 36rpx;
  font-weight: 600;
  color: #333;
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

.info-list {
  display: flex;
  flex-direction: column;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 16rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.info-label {
  font-size: 26rpx;
  color: #666;
}

.info-value {
  font-size: 26rpx;
  color: #333;
}

.room-detail {
  display: flex;
}

.room-img {
  width: 160rpx;
  height: 160rpx;
  border-radius: 8rpx;
  margin-right: 20rpx;
}

.room-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.room-name {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 8rpx;
}

.room-type {
  font-size: 24rpx;
  color: #666;
  margin-bottom: 8rpx;
}

.room-capacity {
  font-size: 24rpx;
  color: #999;
}

.time-info {
  display: flex;
  justify-content: space-between;
}

.time-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.time-label {
  font-size: 24rpx;
  color: #999;
  margin-bottom: 8rpx;
}

.time-value {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
}

.time-divider {
  width: 1rpx;
  background: #e0e0e0;
  margin: 0 20rpx;
}

.duration-text {
  text-align: center;
  font-size: 24rpx;
  color: #666;
  margin-top: 20rpx;
  display: block;
}

.price-list {
  display: flex;
  flex-direction: column;
}

.price-item {
  display: flex;
  justify-content: space-between;
  padding: 16rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }

  &.total {
    padding-top: 24rpx;
    
    .price-label {
      font-size: 28rpx;
      font-weight: 600;
    }

    .price-value {
      font-size: 36rpx;
      font-weight: 600;
      color: #f44336;
    }
  }
}

.price-label {
  font-size: 26rpx;
  color: #666;
}

.price-value {
  font-size: 26rpx;
  color: #333;
}

.remark-text {
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
  padding: 20rpx 30rpx;
  padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
  background: #fff;
  border-top: 1rpx solid #f0f0f0;
}

.cancel-btn {
  flex: 1;
  padding: 24rpx;
  text-align: center;
  font-size: 30rpx;
  color: #999;
  border: 1rpx solid #eee;
  border-radius: 8rpx;
  margin-right: 20rpx;
}

.pay-btn {
  flex: 1;
  padding: 24rpx;
  text-align: center;
  font-size: 30rpx;
  color: #fff;
  background: #4CAF50;
  border-radius: 8rpx;
}

.confirm-btn {
  flex: 1;
  padding: 24rpx;
  text-align: center;
  font-size: 30rpx;
  color: #4CAF50;
  border: 1rpx solid #4CAF50;
  border-radius: 8rpx;
  margin-right: 20rpx;
}

.complete-btn {
  flex: 1;
  padding: 24rpx;
  text-align: center;
  font-size: 30rpx;
  color: #fff;
  background: #4CAF50;
  border-radius: 8rpx;
}

.evaluate-btn {
  flex: 1;
  padding: 24rpx;
  text-align: center;
  font-size: 30rpx;
  color: #fff;
  background: #ff9800;
  border-radius: 8rpx;
}
</style>
