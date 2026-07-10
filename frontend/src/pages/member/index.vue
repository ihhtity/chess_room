<template>
  <view class="container">
    <!-- 会员信息卡片 -->
    <view class="member-card" v-if="membership">
      <view class="card-bg"></view>
      <view class="card-content">
        <view class="member-header">
          <view class="member-icon">
            <text>{{ membership.level_name.slice(0, 1) }}</text>
          </view>
          <view class="member-info">
            <text class="member-name">{{ membership.level_name }}</text>
            <text class="member-id">会员号：{{ membership.member_no }}</text>
          </view>
          <text class="member-status" :class="{ expired: isExpired }">
            {{ isExpired ? '已过期' : '有效期至 ' + formatDate(membership.expired_at || '') }}
          </text>
        </view>
        <view class="member-stats">
          <view class="stat-item">
            <text class="stat-value">{{ membership.points }}</text>
            <text class="stat-label">积分</text>
          </view>
          <view class="stat-divider"></view>
          <view class="stat-item">
            <text class="stat-value">{{ membership.discount }}</text>
            <text class="stat-label">折扣</text>
          </view>
          <view class="stat-divider"></view>
          <view class="stat-item">
            <text class="stat-value">{{ membership.remaining_hours }}</text>
            <text class="stat-label">剩余时长(小时)</text>
          </view>
        </view>
      </view>
    </view>
    <!-- 会员特权卡片 -->
    <view class="card" v-else>
      <view class="not-member">
        <uv-icon name="empty-coupon" color="#ccc" size="80" />
        <text class="not-member-text">您还不是会员</text>
        <view class="become-member-btn" @click="goToBuy">立即开通</view>
      </view>
    </view>
    <!-- 会员特权卡片 -->
    <view class="card">
      <text class="card-title">会员特权</text>
      <view class="privilege-list">
        <view class="privilege-item" v-for="item in privileges" :key="item.key">
          <view class="privilege-icon">{{ item.icon }}</view>
          <view class="privilege-info">
            <text class="privilege-name">{{ item.name }}</text>
            <text class="privilege-desc">{{ item.desc }}</text>
          </view>
        </view>
      </view>
    </view>
    <!-- 会员套餐卡片 -->
    <view class="card">
      <text class="card-title">会员套餐</text>
      <view class="plan-list">
        <view 
          class="plan-item" 
          :class="{ active: selectedPlan === plan.id }"
          v-for="plan in plans" 
          :key="plan.id"
          @click="selectedPlan = plan.id"
        >
          <view class="plan-header">
            <text class="plan-name">{{ plan.name }}</text>
            <text class="plan-duration">{{ plan.duration }}天</text>
          </view>
          <view class="plan-price">
            <text class="price-symbol">¥</text>
            <text class="price-amount">{{ plan.price }}</text>
          </view>
          <view class="plan-benefits">
            <text class="benefit-item" v-for="benefit in plan.benefits" :key="benefit">
              • {{ benefit }}
            </text>
          </view>
          <view class="select-icon" v-if="selectedPlan === plan.id">
            <uv-icon name="checkmark-circle" color="#4CAF50" size="32" />
          </view>
        </view>
      </view>
      <view class="buy-btn" @click="handleBuy">立即购买</view>
    </view>
    <!-- 积分记录卡片 -->
    <view class="card">
      <text class="card-title">积分记录</text>
      <view class="points-list">
        <view class="points-item" v-for="record in pointsRecords" :key="record.id">
          <view class="points-info">
            <text class="points-type">{{ record.type === 'earn' ? '获得积分' : '消耗积分' }}</text>
            <text class="points-time">{{ formatDate(record.created_at) }}</text>
          </view>
          <text class="points-value" :class="record.type === 'earn' ? 'earn' : 'spend'">
            {{ record.type === 'earn' ? '+' : '-' }}{{ record.points }}
          </text>
        </view>
      </view>
    </view>
    <!-- 底部导航栏 -->
    <uv-tabbar :value="3" @change="handleTabChange">
      <uv-tabbar-item text="首页" icon="home" />
      <uv-tabbar-item text="包间" icon="integral" />
      <uv-tabbar-item text="订单" icon="order" />
      <uv-tabbar-item text="会员" icon="empty-coupon" />
      <uv-tabbar-item text="我的" icon="account" />
    </uv-tabbar>
  </view>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { membershipApi } from '@/api'
import type { Membership } from '@/types'

const membership = ref<Membership | null>(null)
const selectedPlan = ref(0)

const privileges = ref([
  { key: 'discount', icon: '🎫', name: '专属折扣', desc: '享受会员专属折扣优惠' },
  { key: 'points', icon: '⭐', name: '积分累计', desc: '消费可累计积分兑换礼品' },
  { key: 'priority', icon: '🚀', name: '优先预约', desc: '优先预约热门包间' },
  { key: 'gift', icon: '🎁', name: '生日礼品', desc: '生日当月赠送精美礼品' }
])

const plans = ref([
  { id: 1, name: '月度会员', duration: 30, price: 99, benefits: ['9折优惠', '赠送100积分', '赠送2小时时长'] },
  { id: 2, name: '季度会员', duration: 90, price: 268, benefits: ['8.5折优惠', '赠送300积分', '赠送8小时时长'] },
  { id: 3, name: '年度会员', duration: 365, price: 888, benefits: ['8折优惠', '赠送1000积分', '赠送30小时时长'] }
])

const pointsRecords = ref([
  { id: 1, type: 'earn', points: 100, created_at: '2024-01-15' },
  { id: 2, type: 'spend', points: 50, created_at: '2024-01-10' },
  { id: 3, type: 'earn', points: 200, created_at: '2024-01-05' }
])

onMounted(async () => {
  await loadMembership()
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

const loadMembership = async () => {
  try {
    const res = await membershipApi.getInfo()
    membership.value = res
  } catch (e) {
    console.log('用户未开通会员')
  }
}

const isExpired = computed(() => {
  if (!membership.value) return false
  return membership.value.expired_at && new Date(membership.value.expired_at) < new Date()
})

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}

const goToBuy = () => {
  selectedPlan.value = plans.value[0].id
}

const handleBuy = () => {
  if (selectedPlan.value === 0) {
    uni.showToast({ title: '请选择会员套餐', icon: 'none' })
    return
  }
  uni.showToast({ title: '购买功能开发中', icon: 'none' })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
}

.member-card {
  position: relative;
  margin: 20rpx;
  border-radius: 20rpx;
  overflow: hidden;
  padding: 30rpx;
}

.card-bg {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, #4CAF50, #45a049);
}

.card-content {
  position: relative;
  z-index: 1;
}

.member-header {
  display: flex;
  align-items: center;
  margin-bottom: 30rpx;
}

.member-icon {
  width: 100rpx;
  height: 100rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 20rpx;
  font-size: 48rpx;
  color: #fff;
  font-weight: 600;
}

.member-info {
  flex: 1;
}

.member-name {
  font-size: 36rpx;
  font-weight: 600;
  color: #fff;
  display: block;
  margin-bottom: 8rpx;
}

.member-id {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
}

.member-status {
  font-size: 22rpx;
  color: rgba(255, 255, 255, 0.9);
  padding: 6rpx 16rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 20rpx;

  &.expired {
    background: rgba(255, 0, 0, 0.3);
    color: #fff;
  }
}

.member-stats {
  display: flex;
  justify-content: space-around;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 12rpx;
  padding: 24rpx 0;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-value {
  font-size: 40rpx;
  font-weight: 600;
  color: #fff;
  margin-bottom: 8rpx;
}

.stat-label {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
}

.stat-divider {
  width: 1rpx;
  background: rgba(255, 255, 255, 0.2);
}

.card {
  background: #fff;
  padding: 24rpx;
  margin: 20rpx;
  border-radius: 16rpx;
}

.card-title {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 20rpx;
  display: block;
}

.not-member {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 60rpx 0;
}

.not-member-text {
  font-size: 28rpx;
  color: #999;
  margin-top: 20rpx;
  margin-bottom: 30rpx;
}

.become-member-btn {
  padding: 20rpx 60rpx;
  background: #4CAF50;
  color: #fff;
  font-size: 28rpx;
  border-radius: 8rpx;
}

.privilege-list {
  display: flex;
  flex-wrap: wrap;
}

.privilege-item {
  width: calc(50% - 10rpx);
  display: flex;
  align-items: center;
  margin-right: 20rpx;
  margin-bottom: 20rpx;

  &:nth-child(2n) {
    margin-right: 0;
  }
}

.privilege-icon {
  font-size: 48rpx;
  margin-right: 16rpx;
}

.privilege-info {
  flex: 1;
}

.privilege-name {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 4rpx;
}

.privilege-desc {
  font-size: 22rpx;
  color: #999;
}

.plan-list {
  display: flex;
  flex-direction: column;
}

.plan-item {
  display: flex;
  align-items: center;
  padding: 20rpx;
  background: #f5f5f5;
  border-radius: 12rpx;
  margin-bottom: 16rpx;
  border: 2rpx solid transparent;
  position: relative;

  &:last-child {
    margin-bottom: 20rpx;
  }

  &.active {
    border-color: #4CAF50;
    background: #e8f5e9;
  }
}

.plan-header {
  display: flex;
  align-items: center;
  margin-bottom: 8rpx;
}

.plan-name {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  margin-right: 12rpx;
}

.plan-duration {
  font-size: 22rpx;
  color: #999;
  padding: 4rpx 12rpx;
  background: rgba(0, 0, 0, 0.05);
  border-radius: 8rpx;
}

.plan-price {
  display: flex;
  align-items: baseline;
  margin-bottom: 8rpx;
}

.price-symbol {
  font-size: 24rpx;
  color: #f44336;
}

.price-amount {
  font-size: 36rpx;
  font-weight: 600;
  color: #f44336;
}

.plan-benefits {
  display: flex;
  flex-wrap: wrap;
}

.benefit-item {
  font-size: 22rpx;
  color: #666;
  margin-right: 16rpx;
  margin-bottom: 4rpx;
}

.select-icon {
  position: absolute;
  right: 20rpx;
}

.buy-btn {
  padding: 24rpx;
  text-align: center;
  background: #4CAF50;
  color: #fff;
  font-size: 30rpx;
  border-radius: 8rpx;
}

.points-list {
  display: flex;
  flex-direction: column;
}

.points-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.points-info {
  display: flex;
  flex-direction: column;
}

.points-type {
  font-size: 28rpx;
  color: #333;
  margin-bottom: 4rpx;
}

.points-time {
  font-size: 22rpx;
  color: #999;
}

.points-value {
  font-size: 28rpx;
  font-weight: 600;

  &.earn {
    color: #4CAF50;
  }

  &.spend {
    color: #f44336;
  }
}
</style>
