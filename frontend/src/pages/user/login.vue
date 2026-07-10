<template>
  <view class="container">
    <view class="logo-section">
      <view class="logo">
        <uv-icon name="home" color="#4CAF50" size="80" />
      </view>
      <text class="title">棋艺室</text>
      <text class="subtitle">享受品质棋牌时光</text>
    </view>

    <view class="form-card">
      <view class="form-item">
        <uv-icon name="phone" color="#999" size="28" />
        <input class="form-input" placeholder="请输入手机号" v-model="phone" type="number" />
      </view>
      <view class="form-item">
        <uv-icon name="lock" color="#999" size="28" />
        <input class="form-input" placeholder="请输入密码" v-model="password" type="password" />
      </view>
      <view class="remember-me">
        <view class="checkbox" :class="{ checked: remember }" @click="remember = !remember">
          <uv-icon name="checkbox-mark" color="#fff" size="24" v-if="remember" />
        </view>
        <text class="remember-text">记住我</text>
      </view>
      <view class="login-btn" @click="handleLogin">登录</view>
    </view>

    <view class="other-options">
      <text class="option-text" @click="goToRegister">注册新账号</text>
      <text class="option-divider">|</text>
      <text class="option-text" @click="goToForgotPassword">忘记密码</text>
    </view>

    <view class="social-login">
      <text class="social-text">其他登录方式</text>
      <view class="social-icons">
        <view class="social-icon" @click="handleWechatLogin">
          <uv-icon name="weixin-circle-fill" color="#07C160" size="48" />
        </view>
        <view class="social-icon">
          <uv-icon name="qq-circle-fill" color="#12B7F5" size="48" />
        </view>
        <view class="social-icon">
          <uv-icon name="twitter-circle-fill" color="#E6162D" size="48" />
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const phone = ref('13800138001')
const password = ref('123456')
const remember = ref(false)

const userStore = useUserStore()

const handleLogin = async () => {
  if (!phone.value) {
    uni.showToast({ title: '请输入手机号', icon: 'none' })
    return
  }
  if (!password.value) {
    uni.showToast({ title: '请输入密码', icon: 'none' })
    return
  }

  try {
    await userStore.login({ phone: phone.value, password: password.value })
    uni.showToast({ title: '登录成功', icon: 'success' })
    setTimeout(() => {
      uni.redirectTo({ url: '/pages/index/index' })
    }, 1500)
  } catch (e) {
    console.error('登录失败', e)
  }
}

const goToRegister = () => {
  uni.navigateTo({ url: '/pages/user/register' })
}

const goToForgotPassword = () => {
  uni.navigateTo({ url: '/pages/user/forgot' })
}

const handleWechatLogin = () => {
  uni.showToast({ title: '微信登录开发中', icon: 'none' })
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
  display: flex;
  flex-direction: column;
  padding: 60rpx 40rpx;
}

.logo-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 100rpx;
  margin-bottom: 60rpx;
}

.logo {
  width: 160rpx;
  height: 160rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 30rpx;
}

.title {
  font-size: 48rpx;
  font-weight: 600;
  color: #000;
  margin-bottom: 16rpx;
}

.subtitle {
  font-size: 28rpx;
  color: #000;
}

.form-card {
  background: #fff;
  border-radius: 20rpx;
  padding: 40rpx;
}

.form-item {
  display: flex;
  align-items: center;
  background: #f5f5f5;
  border-radius: 8rpx;
  padding: 0 24rpx;
  margin-bottom: 24rpx;
}

.form-input {
  flex: 1;
  height: 88rpx;
  font-size: 30rpx;
  color: #333;
  padding-left: 20rpx;
}

.code-item {
  display: flex;
  background: transparent;
  padding: 0;

  .code-input-wrap {
    flex: 1;
    display: flex;
    align-items: center;
    background: #f5f5f5;
    border-radius: 8rpx;
    padding: 0 24rpx;
    margin-right: 16rpx;
  }
}

.get-code-btn {
  padding: 0 32rpx;
  height: 88rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #4CAF50;
  color: #fff;
  font-size: 26rpx;
  border-radius: 8rpx;

  &.disabled {
    background: #ccc;
  }
}

.remember-me {
  display: flex;
  align-items: center;
  margin-bottom: 32rpx;
}

.checkbox {
  width: 40rpx;
  height: 40rpx;
  border: 2rpx solid #ddd;
  border-radius: 8rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 12rpx;

  &.checked {
    background: #4CAF50;
    border-color: #4CAF50;
  }
}

.remember-text {
  font-size: 26rpx;
  color: #000;
}

.login-btn {
  padding: 28rpx;
  text-align: center;
  background: #4CAF50;
  color: #fff;
  font-size: 32rpx;
  border-radius: 8rpx;
}

.other-options {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 32rpx;
}

.option-text {
  font-size: 26rpx;
  color: #000;
}

.option-divider {
  margin: 0 24rpx;
  color: #000;
}

.social-login {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 60rpx;
}

.social-text {
  font-size: 26rpx;
  color: #000;
  margin-bottom: 24rpx;
}

.social-icons {
  display: flex;
}

.social-icon {
  width: 88rpx;
  height: 88rpx;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 24rpx;
}
</style>
