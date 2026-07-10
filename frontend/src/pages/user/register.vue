<template>
  <view class="container">
    <view class="form-card">
      <view class="form-item">
        <uv-icon name="phone" color="#999" size="28" />
        <input class="form-input" placeholder="请输入手机号" v-model="phone" type="number" />
      </view>
      <view class="form-item code-item">
        <view class="code-input-wrap">
          <uv-icon name="email" color="#999" size="28" />
          <input class="form-input" placeholder="请输入验证码" v-model="code" type="number" />
        </view>
        <view class="get-code-btn" :class="{ disabled: countDown > 0 }" @click="getCode">
          {{ countDown > 0 ? `${countDown}s` : '获取验证码' }}
        </view>
      </view>
      <view class="form-item">
        <uv-icon name="lock" color="#999" size="28" />
        <input class="form-input" placeholder="请输入密码" v-model="password" type="password" />
      </view>
      <view class="form-item">
        <uv-icon name="lock" color="#999" size="28" />
        <input class="form-input" placeholder="请再次输入密码" v-model="confirmPassword" type="password" />
      </view>
      <view class="form-item">
        <uv-icon name="account" color="#999" size="28" />
        <input class="form-input" placeholder="请输入昵称" v-model="nickname" />
      </view>
      <view class="agreement">
        <view class="checkbox" :class="{ checked: agreed }" @click="agreed = !agreed">
          <uv-icon name="checkbox-mark" color="#fff" size="24" v-if="agreed" />
        </view>
        <text class="agreement-text">我已阅读并同意</text>
        <text class="agreement-link">《用户服务协议》</text>
        <text class="agreement-text">和</text>
        <text class="agreement-link">《隐私政策》</text>
      </view>
      <view class="register-btn" @click="handleRegister">注册</view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { userApi } from '@/api'

const phone = ref('')
const code = ref('')
const password = ref('')
const confirmPassword = ref('')
const nickname = ref('')
const agreed = ref(false)
const countDown = ref(0)

const handleRegister = async () => {
  if (!phone.value) {
    uni.showToast({ title: '请输入手机号', icon: 'none' })
    return
  }
  if (!code.value) {
    uni.showToast({ title: '请输入验证码', icon: 'none' })
    return
  }
  if (!password.value) {
    uni.showToast({ title: '请输入密码', icon: 'none' })
    return
  }
  if (password.value !== confirmPassword.value) {
    uni.showToast({ title: '两次密码不一致', icon: 'none' })
    return
  }
  if (!nickname.value) {
    uni.showToast({ title: '请输入昵称', icon: 'none' })
    return
  }
  if (!agreed.value) {
    uni.showToast({ title: '请同意用户协议', icon: 'none' })
    return
  }

  try {
    await userApi.register({
      phone: phone.value,
      password: password.value,
      nickname: nickname.value
    })
    uni.showToast({ title: '注册成功', icon: 'success' })
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  } catch (e) {
    console.error('注册失败', e)
  }
}

const getCode = async () => {
  if (countDown.value > 0) return
  if (!phone.value) {
    uni.showToast({ title: '请输入手机号', icon: 'none' })
    return
  }

  try {
    await userApi.sendSms(phone.value)
    countDown.value = 60
    const timer = setInterval(() => {
      countDown.value--
      if (countDown.value <= 0) {
        clearInterval(timer)
      }
    }, 1000)
  } catch (e) {
    console.error('发送验证码失败', e)
  }
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
  margin-top: 100rpx;
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

.form-card {
  background: #fff;
  margin: 20rpx;
  border-radius: 16rpx;
  padding: 30rpx;
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

.agreement {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  margin-bottom: 32rpx;
  padding: 0 8rpx;
}

.checkbox {
  width: 36rpx;
  height: 36rpx;
  border: 2rpx solid #ddd;
  border-radius: 6rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 12rpx;

  &.checked {
    background: #4CAF50;
    border-color: #4CAF50;
  }
}

.agreement-text {
  font-size: 24rpx;
  color: #666;
}

.agreement-link {
  font-size: 24rpx;
  color: #4CAF50;
  margin: 0 4rpx;
}

.register-btn {
  padding: 28rpx;
  text-align: center;
  background: #4CAF50;
  color: #fff;
  font-size: 32rpx;
  border-radius: 8rpx;
}
</style>
