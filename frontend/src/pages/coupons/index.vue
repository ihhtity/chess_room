<template>
  <view class="container">
    <view class="filter-tabs">
      <view 
        class="tab-item" 
        :class="{ active: selectedTab === 0 }"
        @click="selectedTab = 0"
      >
        未使用
      </view>
      <view 
        class="tab-item" 
        :class="{ active: selectedTab === 1 }"
        @click="selectedTab = 1"
      >
        已使用
      </view>
      <view 
        class="tab-item" 
        :class="{ active: selectedTab === 2 }"
        @click="selectedTab = 2"
      >
        已过期
      </view>
    </view>

    <view class="coupon-list">
      <view 
        class="coupon-card" 
        :class="{ used: selectedTab === 1, expired: selectedTab === 2 }"
        v-for="coupon in coupons" 
        :key="coupon.id"
        @click="handleUse(coupon)"
      >
        <view class="coupon-left">
          <text class="coupon-amount">¥{{ coupon.amount }}</text>
          <text class="coupon-condition">满{{ coupon.min_amount }}可用</text>
        </view>
        <view class="coupon-right">
          <text class="coupon-name">{{ coupon.name }}</text>
          <text class="coupon-valid">有效期至 {{ formatDate(coupon.valid_to) }}</text>
          <view class="coupon-btn" v-if="selectedTab === 0">立即使用</view>
        </view>
      </view>
    </view>

    <view v-if="coupons.length === 0" class="empty">
      <uv-icon name="coupon" color="#ccc" size="80" />
      <text class="empty-text">暂无优惠券</text>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import type { Coupon } from '@/types'
import { couponApi } from '@/api'

const selectedTab = ref(0)
const allCoupons = ref<Coupon[]>([])

onMounted(async () => {
  await loadCoupons()
})

const loadCoupons = async () => {
  const res = await couponApi.getList()
  allCoupons.value = res.data
}

const coupons = computed(() => {
  if (selectedTab.value === 0) {
    return allCoupons.value.filter(c => c.status === 0)
  } else if (selectedTab.value === 1) {
    return allCoupons.value.filter(c => c.status === 1)
  } else {
    return allCoupons.value.filter(c => c.status === 2)
  }
})

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}月${date.getDate()}日`
}

const handleUse = (coupon: Coupon) => {
  if (coupon.status === 0) {
    uni.navigateTo({ url: '/pages/booking/index' })
  }
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

.filter-tabs {
  display: flex;
  background: #fff;
  padding: 20rpx 0;
}

.tab-item {
  flex: 1;
  text-align: center;
  font-size: 28rpx;
  color: #666;
  padding: 16rpx 0;

  &.active {
    color: #f44336;
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
      background: #f44336;
      border-radius: 2rpx;
    }
  }
}

.coupon-list {
  padding: 20rpx;
}

.coupon-card {
  display: flex;
  background: linear-gradient(135deg, #fff3e0 0%, #ffe0b2 100%);
  border-radius: 16rpx;
  margin-bottom: 20rpx;
  overflow: hidden;
  border: 2rpx solid #ff9800;

  &.used, &.expired {
    background: #f5f5f5;
    border-color: #ddd;
    opacity: 0.6;
  }
}

.coupon-left {
  width: 200rpx;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 30rpx 0;
  border-right: 2rpx dashed rgba(255, 255, 255, 0.5);

  .used &, .expired & {
    border-right-color: #ddd;
  }
}

.coupon-amount {
  font-size: 56rpx;
  font-weight: 600;
  color: #f44336;
}

.coupon-condition {
  font-size: 22rpx;
  color: #f44336;
  margin-top: 8rpx;
}

.coupon-right {
  flex: 1;
  padding: 24rpx;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.coupon-name {
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 8rpx;
}

.coupon-valid {
  font-size: 22rpx;
  color: #999;
  margin-bottom: 16rpx;
}

.coupon-btn {
  align-self: flex-end;
  padding: 12rpx 32rpx;
  background: #f44336;
  color: #fff;
  font-size: 24rpx;
  border-radius: 8rpx;

  .used &, .expired & {
    display: none;
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
