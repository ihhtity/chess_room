<template>
  <view class="container">
    <!-- 关于我们卡片 -->
    <view class="about-card">
      <view class="logo-section">
        <view class="logo">
          <uv-icon name="home" color="#4CAF50" size="80" />
        </view>
        <text class="app-name">棋艺室</text>
        <text class="app-version">版本 1.0.0</text>
      </view>
    </view>
    <!-- 公司简介卡片 -->
    <view class="card">
      <text class="card-title">公司简介</text>
      <text class="description">
        棋艺室是一家专注于棋牌娱乐服务的互联网平台，致力于为用户提供高品质的棋牌体验。我们拥有专业的团队和优质的服务，让每一位用户都能享受到舒适、便捷的棋牌时光。
      </text>
    </view>
    <!-- 服务理念卡片 -->
    <view class="card">
      <text class="card-title">服务理念</text>
      <view class="service-list">
        <view class="service-item">
          <uv-icon name="star" color="#ff9800" size="28" />
          <text>优质服务</text>
        </view>
        <view class="service-item">
          <uv-icon name="bag" color="#4CAF50" size="28" />
          <text>安全可靠</text>
        </view>
        <view class="service-item">
          <uv-icon name="clock" color="#2196f3" size="28" />
          <text>便捷高效</text>
        </view>
        <view class="service-item">
          <uv-icon name="heart" color="#f44336" size="28" />
          <text>用心服务</text>
        </view>
      </view>
    </view>
    <!-- 隐私政策卡片 -->
    <view class="card">
      <view class="about-item" @click="goToPrivacy">
        <text class="about-text">隐私政策</text>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>
      <view class="about-item" @click="goToTerms">
        <text class="about-text">用户协议</text>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>
    </view>
    <!-- 版权卡片 -->
    <view class="copyright">
      <text>© 2024 棋艺室 版权所有</text>
    </view>
  </view>
</template>

<script setup lang="ts">
const goToPrivacy = () => {
  uni.showToast({ title: '隐私政策开发中', icon: 'none' })
}

const goToTerms = () => {
  uni.showToast({ title: '用户协议开发中', icon: 'none' })
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

.about-card {
  background: linear-gradient(135deg, #4CAF50, #45a049);
  margin: 20rpx;
  border-radius: 16rpx;
  padding: 60rpx 30rpx;
}

.logo-section {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.logo {
  width: 120rpx;
  height: 120rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20rpx;
}

.app-name {
  font-size: 40rpx;
  font-weight: 600;
  color: #fff;
  margin-bottom: 12rpx;
}

.app-version {
  font-size: 26rpx;
  color: rgba(255, 255, 255, 0.8);
}

.card {
  background: #fff;
  margin: 20rpx;
  border-radius: 16rpx;
  padding: 24rpx;
}

.card-title {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 20rpx;
  display: block;
}

.description {
  font-size: 28rpx;
  color: #666;
  line-height: 1.8;
}

.service-list {
  display: flex;
  flex-wrap: wrap;
}

.service-item {
  width: calc(50% - 10rpx);
  display: flex;
  align-items: center;
  padding: 20rpx 0;

  text {
    font-size: 28rpx;
    color: #333;
    margin-left: 12rpx;
  }
}

.about-item {
  display: flex;
  align-items: center;
  padding: 28rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.about-text {
  flex: 1;
  font-size: 30rpx;
  color: #333;
}

.copyright {
  text-align: center;
  padding: 60rpx 0;

  text {
    font-size: 24rpx;
    color: #999;
  }
}
</style>
