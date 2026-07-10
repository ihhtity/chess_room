<template>
  <view class="container">
    <view class="header">
      <view class="back-btn" @click="goBack">
        <uv-icon name="arrow-left" color="#333" size="36" />
      </view>
      <text class="header-title">意见反馈</text>
      <view class="placeholder"></view>
    </view>

    <view class="card">
      <view class="form-item">
        <text class="label">反馈类型</text>
        <view class="type-options">
          <view
            v-for="type in feedbackTypes"
            :key="type.value"
            class="type-option"
            :class="{ active: selectedType === type.value }"
            @click="selectedType = type.value"
          >
            <uv-icon :name="type.icon" :color="selectedType === type.value ? '#fff' : '#999'" size="24" />
            <text class="type-text">{{ type.label }}</text>
          </view>
        </view>
      </view>

      <view class="form-item">
        <text class="label">反馈内容</text>
        <textarea
          v-model="content"
          class="textarea"
          placeholder="请详细描述您的问题或建议..."
          :maxlength="500"
        />
        <text class="word-count">{{ content.length }}/500</text>
      </view>

      <view class="form-item">
        <text class="label">联系方式</text>
        <input
          v-model="contact"
          class="input"
          placeholder="请输入手机号或邮箱（选填）"
        />
      </view>
    </view>

    <view class="submit-btn" :class="{ disabled: !content.trim() }" @click="submitFeedback">
      <text class="btn-text">提交反馈</text>
    </view>

    <view v-if="feedbackList.length > 0" class="card">
      <view class="card-title">我的反馈</view>
      <view v-for="item in feedbackList" :key="item.id" class="feedback-item">
        <view class="feedback-header">
          <view class="type-badge" :class="getTypeClass(item.type)">
            {{ getTypeText(item.type) }}
          </view>
          <view class="status-badge" :class="item.status === 1 ? 'resolved' : 'pending'">
            {{ item.status === 1 ? '已回复' : '待处理' }}
          </view>
        </view>
        <text class="feedback-content">{{ item.content }}</text>
        <text class="feedback-time">{{ formatTime(item.created_at) }}</text>
        <view v-if="item.reply" class="reply-section">
          <text class="reply-label">回复：</text>
          <text class="reply-content">{{ item.reply }}</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { feedbackApi } from '@/api'
import type { Feedback } from '@/types'

const selectedType = ref(0)
const content = ref('')
const contact = ref('')
const feedbackList = ref<Feedback[]>([])
const submitting = ref(false)

const feedbackTypes = [
  { value: 0, label: '功能建议', icon: 'lightbulb' },
  { value: 1, label: 'Bug反馈', icon: 'bug' },
  { value: 2, label: '意见投诉', icon: 'alert-circle' },
  { value: 3, label: '其他', icon: 'more' }
]

onMounted(() => {
  fetchFeedbackList()
})

const fetchFeedbackList = async () => {
  try {
    const res = await feedbackApi.getList({ page: 1, page_size: 10 })
    feedbackList.value = res.data
  } catch (error) {
    console.error('获取反馈列表失败:', error)
  }
}

const submitFeedback = async () => {
  if (!content.value.trim() || submitting.value) return
  
  submitting.value = true
  try {
    await feedbackApi.create({
      content: content.value.trim(),
      contact: contact.value.trim(),
      type: selectedType.value
    })
    
    uni.showToast({ title: '提交成功', icon: 'success' })
    content.value = ''
    contact.value = ''
    selectedType.value = 0
    fetchFeedbackList()
  } catch (error) {
    console.error('提交反馈失败:', error)
  } finally {
    submitting.value = false
  }
}

const getTypeText = (type: number) => {
  const map = feedbackTypes.find(t => t.value === type)
  return map ? map.label : '其他'
}

const getTypeClass = (type: number) => {
  const classes = ['type-suggest', 'type-bug', 'type-complaint', 'type-other']
  return classes[type] || classes[3]
}

const formatTime = (timeStr: string) => {
  if (!timeStr) return ''
  const date = new Date(timeStr)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hour = String(date.getHours()).padStart(2, '0')
  const minute = String(date.getMinutes()).padStart(2, '0')
  return `${year}-${month}-${day} ${hour}:${minute}`
}

const goBack = () => {
  uni.navigateBack()
}
</script>

<style lang="scss" scoped>
.container {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 120rpx;
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
  padding: 30rpx;
}

.card-title {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 24rpx;
  padding-bottom: 16rpx;
  border-bottom: 1rpx solid #f0f0f0;
}

.form-item {
  margin-bottom: 30rpx;

  &:last-child {
    margin-bottom: 0;
  }
}

.label {
  font-size: 28rpx;
  color: #333;
  margin-bottom: 16rpx;
  display: block;
}

.type-options {
  display: flex;
  flex-wrap: wrap;
  gap: 16rpx;
}

.type-option {
  display: flex;
  align-items: center;
  padding: 16rpx 24rpx;
  border-radius: 30rpx;
  background: #f5f5f5;
  border: 2rpx solid transparent;
  transition: all 0.3s;

  &.active {
    background: #1890ff;
    border-color: #1890ff;
  }
}

.type-text {
  font-size: 26rpx;
  color: #666;
  margin-left: 8rpx;

  .active & {
    color: #fff;
  }
}

.textarea {
  width: 100%;
  height: 200rpx;
  padding: 20rpx;
  background: #f5f5f5;
  border-radius: 12rpx;
  font-size: 28rpx;
  color: #333;
  box-sizing: border-box;
}

.word-count {
  font-size: 24rpx;
  color: #999;
  text-align: right;
  margin-top: 8rpx;
  display: block;
}

.input {
  width: 100%;
  height: 88rpx;
  padding: 0 20rpx;
  background: #f5f5f5;
  border-radius: 12rpx;
  font-size: 28rpx;
  color: #333;
  box-sizing: border-box;
}

.submit-btn {
  position: fixed;
  bottom: 0;
  left: 20rpx;
  right: 20rpx;
  height: 88rpx;
  background: linear-gradient(135deg, #1890ff, #096dd9);
  border-radius: 44rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: calc(20rpx + env(safe-area-inset-bottom));
  box-shadow: 0 4rpx 16rpx rgba(24, 144, 255, 0.3);

  &.disabled {
    background: #ccc;
    box-shadow: none;
  }
}

.btn-text {
  font-size: 32rpx;
  font-weight: 600;
  color: #fff;
}

.feedback-item {
  padding: 24rpx 0;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.feedback-header {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 12rpx;
}

.type-badge {
  font-size: 22rpx;
  padding: 6rpx 16rpx;
  border-radius: 20rpx;

  &.type-suggest { background: #e6f7ff; color: #1890ff; }
  &.type-bug { background: #fff2f0; color: #f5222d; }
  &.type-complaint { background: #fff7e6; color: #fa8c16; }
  &.type-other { background: #f9f0ff; color: #722ed1; }
}

.status-badge {
  font-size: 22rpx;
  padding: 6rpx 16rpx;
  border-radius: 20rpx;
  background: #fff7e6;
  color: #fa8c16;

  &.resolved {
    background: #f6ffed;
    color: #52c41a;
  }
}

.feedback-content {
  font-size: 28rpx;
  color: #333;
  line-height: 1.6;
  display: block;
}

.feedback-time {
  font-size: 24rpx;
  color: #999;
  margin-top: 12rpx;
  display: block;
}

.reply-section {
  background: #f6ffed;
  padding: 16rpx;
  border-radius: 8rpx;
  margin-top: 16rpx;
}

.reply-label {
  font-size: 24rpx;
  color: #52c41a;
}

.reply-content {
  font-size: 26rpx;
  color: #333;
}
</style>
