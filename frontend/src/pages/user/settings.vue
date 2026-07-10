<template>
  <view class="container">
    <!-- 通知设置 -->
    <view class="card">
      <view class="card-title">通知设置</view>
      <view class="setting-item">
        <view class="item-info">
          <view class="icon-wrapper bg-green">
            <uv-icon name="bell" color="#fff" size="24" />
          </view>
          <text class="item-text">消息通知</text>
        </view>
        <uv-switch :value="!!setting.notifications" @change="handleToggle('notifications')" />
      </view>
      <view class="setting-item">
        <view class="item-info">
          <view class="icon-wrapper bg-blue">
            <uv-icon name="volume-fill" color="#fff" size="24" />
          </view>
          <text class="item-text">声音提醒</text>
        </view>
        <uv-switch :value="!!setting.sound" @change="handleToggle('sound')" />
      </view>
      <view class="setting-item">
        <view class="item-info">
          <view class="icon-wrapper bg-orange">
            <uv-icon name="error-circle" color="#fff" size="24" />
          </view>
          <text class="item-text">震动提醒</text>
        </view>
        <uv-switch :value="!!setting.vibrate" @change="handleToggle('vibrate')" />
      </view>
    </view>

    <view class="card">
      <view class="setting-item" @click="goToLanguage">
        <view class="item-info">
          <view class="icon-wrapper bg-cyan">
            <uv-icon name="attach" color="#fff" size="24" />
          </view>
          <text class="item-text">语言设置</text>
        </view>
        <text class="item-value">{{ getLanguageText(setting.language) }}</text>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>
      <view class="setting-item" @click="goToClearCache">
        <view class="item-info">
          <view class="icon-wrapper bg-red-light">
            <uv-icon name="trash" color="#fff" size="24" />
          </view>
          <text class="item-text">清除缓存</text>
        </view>
        <text class="cache-size">{{ cacheSize }}</text>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>
      <view class="setting-item" @click="goToFeedback">
        <view class="item-info">
          <view class="icon-wrapper bg-orange-light">
            <uv-icon name="chat" color="#fff" size="24" />
          </view>
          <text class="item-text">意见反馈</text>
        </view>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>
    </view>

    <view class="version-info">
      <text class="version-text">版本号 {{ version }}</text>
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

const cacheSize = ref('5.2MB')
const version = ref('1.0.0')
const loading = ref(false)

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

const handleToggle = async (key: string) => {
  if (loading.value) return
  loading.value = true
  try {
    const res = await settingApi.toggle(key)
    setting.value = res
    uni.showToast({
      title: key === 'notifications' ? (res.notifications ? '已开启消息通知' : '已关闭消息通知') :
        key === 'sound' ? (res.sound ? '已开启声音提醒' : '已关闭声音提醒') :
          '已开启震动提醒',
      icon: 'none'
    })
  } catch (error) {
    uni.showToast({
      title: '操作失败，请重试',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

const getLanguageText = (lang: string) => {
  const map: Record<string, string> = {
    'zh-CN': '简体中文',
    'zh-TW': '繁体中文',
    'en': 'English'
  }
  return map[lang] || '简体中文'
}

const goToLanguage = () => {
  uni.navigateTo({ url: '/pages/user/language' })
}

const goToClearCache = () => {
  uni.showModal({
    title: '提示',
    content: '确定要清除缓存吗？',
    success: (res) => {
      if (res.confirm) {
        cacheSize.value = '0MB'
        uni.showToast({ title: '清除成功', icon: 'success' })
      }
    }
  })
}

const goToFeedback = () => {
  uni.navigateTo({ url: '/pages/user/feedback' })
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

.card-title {
  font-size: 28rpx;
  color: #999;
  padding: 24rpx 30rpx 12rpx;
}

.setting-item {
  display: flex;
  align-items: center;
  padding: 28rpx 30rpx;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.item-info {
  display: flex;
  align-items: center;
  flex: 1;
}

.icon-wrapper {
  width: 56rpx;
  height: 56rpx;
  border-radius: 14rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.bg-green {
  background: linear-gradient(135deg, #4CAF50, #45a049);
}

.bg-blue {
  background: linear-gradient(135deg, #2196f3, #1976d2);
}

.bg-orange {
  background: linear-gradient(135deg, #ff9800, #f57c00);
}

.bg-purple {
  background: linear-gradient(135deg, #9c27b0, #7b1fa2);
}

.bg-red {
  background: linear-gradient(135deg, #f44336, #d32f2f);
}

.bg-cyan {
  background: linear-gradient(135deg, #00bcd4, #0097a7);
}

.bg-red-light {
  background: linear-gradient(135deg, #ef5350, #e53935);
}

.bg-orange-light {
  background: linear-gradient(135deg, #ff7043, #f4511e);
}

.item-text {
  font-size: 30rpx;
  color: #333;
  margin-left: 24rpx;
}

.cache-size {
  font-size: 26rpx;
  color: #999;
  margin-right: 16rpx;
}

.item-value {
  font-size: 26rpx;
  color: #999;
  margin-right: 16rpx;
}

.version-info {
  text-align: center;
  padding: 40rpx 0;
}

.version-text {
  font-size: 24rpx;
  color: #bbb;
}
</style>
