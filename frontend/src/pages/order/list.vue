<template>
  <view class="container">
    <!-- 订单筛选标签 -->
    <view class="filter-tabs">
      <view 
        class="tab-item" 
        :class="{ active: selectedStatus === -1 }"
        @click="selectedStatus = -1; loadOrders()"
      >
        全部
      </view>
      <view 
        class="tab-item" 
        :class="{ active: selectedStatus === 0 }"
        @click="selectedStatus = 0; loadOrders()"
      >
        待支付
      </view>
      <view 
        class="tab-item" 
        :class="{ active: selectedStatus === 1 }"
        @click="selectedStatus = 1; loadOrders()"
      >
        使用中
      </view>
      <view 
        class="tab-item" 
        :class="{ active: selectedStatus === 2 }"
        @click="selectedStatus = 2; loadOrders()"
      >
        已完成
      </view>
      <view 
        class="tab-item" 
        :class="{ active: selectedStatus === 3 }"
        @click="selectedStatus = 3; loadOrders()"
      >
        已取消
      </view>
    </view>
    <!-- 订单列表卡片 -->
    <view class="order-list">
      <view class="order-card" v-for="order in orders" :key="order.id" @click="goToDetail(order.id)">
        <view class="order-header">
          <text class="order-no">订单号：{{ order.order_no }}</text>
          <text class="order-status" :class="getStatusClass(order.status)">{{ getStatusText(order.status) }}</text>
        </view>
        <view class="order-room">
          <image :src="getRoomImage(order.room.images, order.room.id)" mode="aspectFill" class="room-img" />
          <view class="room-info">
            <text class="room-name">{{ order.room.name }}</text>
            <text class="room-type">{{ order.room.type.name }} · {{ order.room.floor }}</text>
            <text class="order-time">{{ formatDate(order.start_time) }} {{ order.start_time.split(' ')[1] }} - {{ order.end_time.split(' ')[1] }}</text>
          </view>
        </view>
        <view class="order-footer">
          <text class="order-amount">¥{{ order.total_amount }}</text>
          <view class="order-actions">
            <view 
              class="action-btn cancel" 
              v-if="order.status === 0" 
              @click.stop="handleCancel(order.id)"
            >
              取消订单
            </view>
            <view 
              class="action-btn pay" 
              v-if="order.status === 0" 
              @click.stop="handlePay(order)"
            >
              立即支付
            </view>
            <view 
              class="action-btn confirm" 
              v-if="order.status === 1" 
              @click.stop="handleConfirm(order.id)"
            >
              确认使用
            </view>
            <view 
              class="action-btn complete" 
              v-if="order.status === 1" 
              @click.stop="handleComplete(order.id)"
            >
              完成订单
            </view>
            <view 
              class="action-btn evaluate" 
              v-if="order.status === 2" 
              @click.stop="handleEvaluate(order.id)"
            >
              去评价
            </view>
          </view>
        </view>
      </view>
    </view>
    <!-- 订单列表卡片 -->
    <view v-if="orders.length === 0" class="empty">
      <uv-icon name="order" color="#ccc" size="80" />
      <text class="empty-text">暂无订单</text>
    </view>
    <!-- 底部导航栏 -->
    <uv-tabbar :value="2" @change="handleTabChange">
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
import { orderApi } from '@/api'
import type { Order } from '@/types'
import { OrderStatus } from '@/types'
import { getRandomImage } from '@/utils/image'

const orders = ref<Order[]>([])
const selectedStatus = ref(-1)

onMounted(async () => {
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  const options = (currentPage as any)?.options || {}
  if (options.status !== undefined) {
    selectedStatus.value = parseInt(options.status)
  }
  await loadOrders()
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

const loadOrders = async () => {
  const params: any = {
    page: 1,
    page_size: 20
  }
  if (selectedStatus.value >= 0) {
    params.status = selectedStatus.value
  }
  const res = await orderApi.getList(params)
  orders.value = res.data
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
    case OrderStatus.Refunding:
      return 'refunding'
    case OrderStatus.Refunded:
      return 'refunded'
    default:
      return ''
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

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}月${date.getDate()}日`
}

const goToDetail = (id: number) => {
  uni.navigateTo({ url: `/pages/order/detail?id=${id}` })
}

const handleCancel = async (id: number) => {
  uni.showModal({
    title: '提示',
    content: '确定要取消订单吗？',
    success: async (res) => {
      if (res.confirm) {
        await orderApi.cancel(id)
        uni.showToast({ title: '取消成功', icon: 'success' })
        await loadOrders()
      }
    }
  })
}

const handlePay = (order: Order) => {
  uni.showToast({ title: '支付功能开发中', icon: 'none' })
}

const handleConfirm = async (id: number) => {
  uni.showModal({
    title: '提示',
    content: '确定要确认使用吗？',
    success: async (res) => {
      if (res.confirm) {
        await orderApi.confirm(id)
        uni.showToast({ title: '确认成功', icon: 'success' })
        await loadOrders()
      }
    }
  })
}

const handleComplete = async (id: number) => {
  uni.showModal({
    title: '提示',
    content: '确定要完成订单吗？',
    success: async (res) => {
      if (res.confirm) {
        await orderApi.complete(id)
        uni.showToast({ title: '完成成功', icon: 'success' })
        await loadOrders()
      }
    }
  })
}

const handleEvaluate = (id: number) => {
  uni.showToast({ title: '评价功能开发中', icon: 'none' })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 120rpx;
}

.filter-tabs {
  display: flex;
  background: #fff;
  padding: 20rpx 0;
}

.tab-item {
  flex: 1;
  text-align: center;
  font-size: 26rpx;
  color: #666;
  padding: 16rpx 0;

  &.active {
    color: #4CAF50;
    font-weight: 600;
    position: relative;

    &::after {
      content: '';
      position: absolute;
      bottom: 0;
      left: 50%;
      transform: translateX(-50%);
      width: 40rpx;
      height: 4rpx;
      background: #4CAF50;
      border-radius: 2rpx;
    }
  }
}

.order-list {
  padding: 20rpx;
}

.order-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 24rpx;
  margin-bottom: 20rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.06);
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20rpx;
  padding-bottom: 16rpx;
  border-bottom: 1rpx solid #f0f0f0;
}

.order-no {
  font-size: 24rpx;
  color: #999;
}

.order-status {
  font-size: 26rpx;
  font-weight: 600;

  &.pending {
    color: #ff9800;
  }

  &.active {
    color: #4CAF50;
  }

  &.completed {
    color: #999;
  }

  &.cancelled {
    color: #999;
  }

  &.refunding {
    color: #ff9800;
  }

  &.refunded {
    color: #999;
  }
}

.order-room {
  display: flex;
  margin-bottom: 20rpx;
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

.order-time {
  font-size: 24rpx;
  color: #999;
}

.order-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.order-amount {
  font-size: 32rpx;
  font-weight: 600;
  color: #f44336;
}

.order-actions {
  display: flex;
}

.action-btn {
  padding: 12rpx 24rpx;
  font-size: 24rpx;
  border-radius: 8rpx;
  margin-left: 16rpx;
  border: 1rpx solid #eee;

  &.pay {
    background: #4CAF50;
    color: #fff;
    border-color: #4CAF50;
  }

  &.cancel {
    color: #999;
  }

  &.confirm {
    border-color: #4CAF50;
    color: #4CAF50;
  }

  &.complete {
    background: #4CAF50;
    color: #fff;
    border-color: #4CAF50;
  }

  &.evaluate {
    border-color: #ff9800;
    color: #ff9800;
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
