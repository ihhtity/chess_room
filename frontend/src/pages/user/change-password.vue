<template>
  <view class="container">
    <view class="header">
      <text class="back-btn" @click="goBack">←</text>
      <text class="header-title">修改密码</text>
      <view class="placeholder"></view>
    </view>

    <view class="form-card">
      <view class="form-item">
        <uv-icon name="lock" color="#999" size="28" />
        <input class="form-input" placeholder="请输入旧密码" v-model="oldPassword" type="password" />
      </view>
      <view class="form-item">
        <uv-icon name="lock" color="#999" size="28" />
        <input class="form-input" placeholder="请输入新密码（至少6位）" v-model="newPassword" type="password" />
      </view>
      <view class="form-item">
        <uv-icon name="lock" color="#999" size="28" />
        <input class="form-input" placeholder="请再次输入新密码" v-model="confirmPassword" type="password" />
      </view>
      <view class="submit-btn" @click="handleSubmit" :class="{ disabled: loading }">{{ loading ? '修改中...' : '确认修改' }}</view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

const oldPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const loading = ref(false)

const handleSubmit = async () => {
  if (!oldPassword.value.trim()) {
    uni.showToast({ title: '请输入旧密码', icon: 'none' })
    return
  }
  if (!newPassword.value.trim()) {
    uni.showToast({ title: '请输入新密码', icon: 'none' })
    return
  }
  if (newPassword.value.length < 6) {
    uni.showToast({ title: '密码长度不少于6位', icon: 'none' })
    return
  }
  if (!/^[a-zA-Z0-9]{6,20}$/.test(newPassword.value)) {
    uni.showToast({ title: '密码只能包含字母和数字', icon: 'none' })
    return
  }
  if (newPassword.value !== confirmPassword.value) {
    uni.showToast({ title: '两次密码不一致', icon: 'none' })
    return
  }
  if (oldPassword.value === newPassword.value) {
    uni.showToast({ title: '新密码不能与旧密码相同', icon: 'none' })
    return
  }

  loading.value = true

  try {
    await userStore.changePassword({
      old_password: oldPassword.value,
      new_password: newPassword.value
    })
    uni.showToast({ title: '修改成功', icon: 'success' })
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  } catch (e: any) {
    uni.showToast({ title: e?.message || '修改失败', icon: 'none' })
  } finally {
    loading.value = false
  }
}

const goBack = () => {
  uni.navigateBack()
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

.submit-btn {
  padding: 28rpx;
  text-align: center;
  background: #4CAF50;
  color: #fff;
  font-size: 32rpx;
  border-radius: 8rpx;
  margin-top: 20rpx;
  
  &.disabled {
    background: #ccc;
  }
}
</style>
