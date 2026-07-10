<template>
  <view class="container">
    <view class="header">
      <view class="back-btn" @click="goBack">
        <uv-icon name="arrow-left" color="#333" size="36" />
      </view>
      <text class="header-title">语言设置</text>
      <view class="placeholder"></view>
    </view>

    <view class="card">
      <view
        v-for="lang in languages"
        :key="lang.value"
        class="language-item"
        :class="{ active: setting.language === lang.value }"
        @click="selectLanguage(lang.value)"
      >
        <view class="lang-info">
          <text class="lang-name">{{ lang.name }}</text>
          <text class="lang-native">{{ lang.native }}</text>
        </view>
        <view v-if="setting.language === lang.value" class="check-icon">
          <uv-icon name="check-circle" color="#1890ff" size="32" />
        </view>
        <view v-else class="check-placeholder"></view>
      </view>
    </view>

    <view class="tip">
      <text class="tip-text">设置后需重启应用生效</text>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { settingApi } from '@/api'
import type { UserSetting } from '@/types'

const setting = ref<UserSetting>({
  id: 0,
  user_id: 0,
  notifications: 1,
  sound: 1,
  vibrate: 1,
  language: 'zh-CN',
  theme: 'light',
  allow_push: 1,
  allow_marketing: 1,
  created_at: '',
  updated_at: ''
})

const languages = [
  { value: 'zh-CN', name: '简体中文', native: '简体中文' },
  { value: 'zh-TW', name: '繁体中文', native: '繁體中文' },
  { value: 'en', name: 'English', native: 'English' }
]

onMounted(() => {
  fetchSettings()
})

const fetchSettings = async () => {
  try {
    const res = await settingApi.get()
    setting.value = res
  } catch (error) {
    console.error('获取设置失败:', error)
  }
}

const selectLanguage = async (lang: string) => {
  if (setting.value.language === lang) return
  
  try {
    const res = await settingApi.update({ language: lang })
    setting.value = res
    uni.showToast({ title: '设置成功', icon: 'success' })
  } catch (error) {
    console.error('更新语言失败:', error)
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
  padding: 20rpx 30rpx;
  padding-top: calc(20rpx + env(safe-area-inset-top));
  background: #fff;
}

.back-btn {
  width: 60rpx;
  height: 60rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-title {
  font-size: 34rpx;
  font-weight: 600;
  color: #333;
}

.placeholder {
  width: 60rpx;
}

.card {
  background: #fff;
  margin: 20rpx;
  border-radius: 16rpx;
  overflow: hidden;
}

.language-item {
  display: flex;
  align-items: center;
  padding: 32rpx 30rpx;
  border-bottom: 1rpx solid #f0f0f0;
  transition: background 0.2s;

  &:last-child {
    border-bottom: none;
  }

  &.active {
    background: #f0f7ff;
  }
}

.lang-info {
  flex: 1;
}

.lang-name {
  font-size: 30rpx;
  color: #333;
  display: block;
}

.lang-native {
  font-size: 24rpx;
  color: #999;
  margin-top: 8rpx;
  display: block;
}

.check-icon {
  width: 48rpx;
  height: 48rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.check-placeholder {
  width: 48rpx;
  height: 48rpx;
}

.tip {
  text-align: center;
  padding: 30rpx;
}

.tip-text {
  font-size: 24rpx;
  color: #999;
}
</style>
