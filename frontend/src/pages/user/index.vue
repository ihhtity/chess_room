<template>
  <view class="container">
    <view class="user-header">
      <view class="header-bg"></view>
      <view class="user-info">
        <view class="avatar" v-if="user.avatarURL" @click="goToEditProfile">
          <image :src="user.avatarURL || defaultAvatar" mode="aspectFill" class="avatar-img" />
          <view class="avatar-edit">
            <uv-icon name="camera" color="#fff" size="24" />
          </view>
        </view>
        <view class="avatar" v-else @click="goToEditProfile">
          <uv-icon name="account-fill" color="#fff" size="74" />
        </view>
        <view class="user-detail" v-if="user.id" @click="goToEditProfile">
          <text class="user-name">{{ user?.name || '-' }}</text>
          <text class="user-phone">{{ user?.phone || '-' }}</text>
        </view>
        <view class="user-detail" v-else @click="handleLogout">
          <text class="user-name">点击登录</text>
          <text class="user-phone">登录享更多特权</text>
        </view>
        <view class="edit-btn" @click="goToEditProfile">
          <uv-icon name="edit-pen" color="#fff" size="24" />
        </view>
      </view>
      <view class="user-stats" v-if="user.id">
        <view class="stat-item" @click="goToOrderByStatus(0)">
          <text class="stat-value">{{ pendingCount }}</text>
          <text class="stat-label">待确认</text>
        </view>
        <view class="stat-divider"></view>
        <view class="stat-item" @click="goToOrderByStatus(1)">
          <text class="stat-value">{{ activeCount }}</text>
          <text class="stat-label">使用中</text>
        </view>
        <view class="stat-divider"></view>
        <view class="stat-item" @click="goToOrderByStatus(2)">
          <text class="stat-value">{{ completedCount }}</text>
          <text class="stat-label">已完成</text>
        </view>
        <view class="stat-divider"></view>
        <view class="stat-item" @click="goToOrderByStatus(3)">
          <text class="stat-value">{{ cancelledCount }}</text>
          <text class="stat-label">已取消</text>
        </view>
      </view>
    </view>

    <view class="menu-grid">
      <view class="menu-item" @click="goToCoupons">
        <view class="menu-icon-wrap">
          <uv-icon name="coupon" color="#ff9800" size="36" />
        </view>
        <text class="menu-text">优惠券</text>
      </view>
      <view class="menu-item" @click="goToNotifications">
        <view class="menu-icon-wrap">
          <uv-icon name="bell" color="#2196f3" size="36" />
          <view class="badge" v-if="notificationCount > 0">{{ notificationCount }}</view>
        </view>
        <text class="menu-text">消息</text>
      </view>
      <view class="menu-item" @click="goToMember">
        <view class="menu-icon-wrap">
          <uv-icon name="empty-coupon" color="#9c27b0" size="36" />
        </view>
        <text class="menu-text">会员中心</text>
      </view>
      <view class="menu-item" @click="goToSettings">
        <view class="menu-icon-wrap">
          <uv-icon name="setting" color="#666" size="36" />
        </view>
        <text class="menu-text">设置</text>
      </view>
      <view class="menu-item" @click="goToHelp">
        <view class="menu-icon-wrap">
          <uv-icon name="question-circle" color="#4CAF50" size="36" />
        </view>
        <text class="menu-text">帮助中心</text>
      </view>
      <view class="menu-item" @click="goToAbout">
        <view class="menu-icon-wrap">
          <uv-icon name="info-circle" color="#2196f3" size="36" />
        </view>
        <text class="menu-text">关于我们</text>
      </view>
      <view class="menu-item logout-menu" @click="handleLogout" v-if="user.id">
        <view class="menu-icon-wrap">
          <uv-icon name="share-square" color="#f44336" size="36" />
        </view>
        <text class="menu-text">退出登录</text>
      </view>
    </view>

    <uv-tabbar :value="4" @change="handleTabChange">
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
import { orderApi, notificationApi } from '@/api'
import { getAvatar } from '@/utils/image'

const defaultAvatar = getAvatar('default')

const user = uni.getStorageSync('user') || null
const pendingCount = ref(0)
const activeCount = ref(0)
const completedCount = ref(0)
const cancelledCount = ref(0)
const notificationCount = ref(0)

onMounted(async () => {
  await loadData()
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

const loadData = async () => {
  console.log(user)
  if (!user) return
  
  try {
    const [orderRes, notificationRes] = await Promise.all([
      orderApi.getList({ page: 1, page_size: 100 }),
      notificationApi.getList({ page: 1, page_size: 10 })
    ])
    
    const orders = orderRes?.data || []
    pendingCount.value = orders.filter((o: any) => o.status === 0).length
    activeCount.value = orders.filter((o: any) => o.status === 1).length
    completedCount.value = orders.filter((o: any) => o.status === 2).length
    cancelledCount.value = orders.filter((o: any) => o.status === 3).length
    
    const notifications = notificationRes?.data || []
    notificationCount.value = notifications.filter((n: any) => !n.read).length
  } catch (e) {
    console.error('加载数据失败', e)
  }
}

const goToEditProfile = () => {
  if (!user.id) {
    uni.navigateTo({ url: '/pages/user/login' })
    return
  }
  uni.navigateTo({ url: '/pages/user/profile' })
}

const goToCoupons = () => {
  uni.navigateTo({ url: '/pages/coupons/index' })
}

const goToNotifications = () => {
  uni.navigateTo({ url: '/pages/notifications/index' })
}

const goToMember = () => {
  uni.redirectTo({ url: '/pages/member/index' })
}

const goToHelp = () => {
  uni.navigateTo({ url: '/pages/help/index' })
}

const goToSettings = () => {
  uni.navigateTo({ url: '/pages/user/settings' })
}

const goToAbout = () => {
  uni.navigateTo({ url: '/pages/about/index' })
}

const goToOrderByStatus = (status: number) => {
  uni.redirectTo({ url: `/pages/order/list?status=${status}` })
}

const handleLogout = () => {
  uni.showModal({
    title: '提示',
    content: '确定要退出登录吗？',
    success: async (res) => {
      if (res.confirm) {
        uni.removeStorageSync('user')
        uni.showToast({ title: '退出成功', icon: 'success' })
        setTimeout(() => {
          uni.reLaunch({ url: '/pages/user/login' })
        }, 1500)
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

.user-header {
  position: relative;
  padding: 80rpx 30rpx 40rpx;
  overflow: hidden;
}

.header-bg {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 400rpx;
  background: linear-gradient(135deg, #4CAF50, #45a049);
  border-radius: 0 0 50rpx 50rpx;
}

.user-info {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  margin-bottom: 30rpx;
}

.avatar {
  position: relative;
  width: 140rpx;
  height: 140rpx;
  border-radius: 50%;
  overflow: hidden;
  border: 4rpx solid rgba(255, 255, 255, 0.5);
  margin-right: 24rpx;
}

.avatar-img {
  width: 100%;
  height: 100%;
}

.avatar-edit {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 44rpx;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.user-detail {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.user-name {
  font-size: 36rpx;
  font-weight: 600;
  color: #fff;
  margin-bottom: 8rpx;
}

.user-phone {
  font-size: 26rpx;
  color: rgba(255, 255, 255, 0.8);
}

.edit-btn {
  width: 60rpx;
  height: 60rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.user-stats {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: space-around;
  align-items: center;
  background: #fff;
  border-radius: 16rpx;
  padding: 30rpx 0;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.1);
}

.stat-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-value {
  font-size: 40rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 8rpx;
}

.stat-label {
  font-size: 24rpx;
  color: #999;
}

.stat-divider {
  width: 1rpx;
  height: 60rpx;
  background: #f0f0f0;
}

.menu-grid {
  display: flex;
  flex-wrap: wrap;
  background: #fff;
  margin: 20rpx;
  border-radius: 16rpx;
  padding: 20rpx 0;
}

.menu-item {
  width: 25%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 30rpx 0;
}

.menu-icon-wrap {
  position: relative;
  width: 88rpx;
  height: 88rpx;
  background: #f5f5f5;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 16rpx;
}

.menu-text {
  font-size: 24rpx;
  color: #333;
}

.badge {
  position: absolute;
  top: -4rpx;
  right: -4rpx;
  min-width: 32rpx;
  height: 32rpx;
  padding: 0 6rpx;
  background: #f44336;
  color: #fff;
  font-size: 18rpx;
  border-radius: 16rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logout-menu .menu-icon-wrap {
  background: #fff0f0;
}

.logout-menu .menu-text {
  color: #f44336;
}
</style>