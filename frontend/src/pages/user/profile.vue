<template>
  <view class="container">
    <view class="header">
      <text class="back-btn" @click="goBack">←</text>
      <text class="header-title">编辑资料</text>
      <text class="save-btn" @click="handleSave" :class="{ disabled: loading }">{{ loading ? '保存中...' : '保存' }}</text>
    </view>

    <view class="card">
      <view class="profile-item avatar-item">
        <text class="item-label">头像</text>
        <view class="avatar-wrap" @click="chooseAvatar">
          <image :src="form.avatar || defaultAvatar" mode="aspectFill" class="avatar-img" />
          <view class="avatar-edit">
            <uv-icon name="camera" color="#fff" size="28" />
          </view>
        </view>
      </view>

      <view class="profile-item">
        <text class="item-label">昵称</text>
        <input class="item-input" placeholder="请输入昵称" v-model="form.nickname" />
      </view>

      <view class="profile-item">
        <text class="item-label">手机号</text>
        <text class="item-value">{{ userStore.user?.phone || '' }}</text>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>

      <view class="profile-item">
        <text class="item-label">性别</text>
        <picker mode="selector" :range="genderOptions" :value="genderIndex" @change="handleGenderChange">
          <view class="picker-value">
            <text>{{ genderOptions[genderIndex] }}</text>
            <uv-icon name="arrow-right" color="#ccc" size="24" />
          </view>
        </picker>
      </view>

      <view class="profile-item">
        <text class="item-label">真实姓名</text>
        <input class="item-input" placeholder="请输入真实姓名" v-model="form.realname" />
      </view>
    </view>

    <view class="card">
      <view class="profile-item" @click="goToChangePassword">
        <text class="item-label">修改密码</text>
        <text class="item-value">修改登录密码</text>
        <uv-icon name="arrow-right" color="#ccc" size="24" />
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { getAvatar } from '@/utils/image'

const userStore = useUserStore()
const defaultAvatar = getAvatar('default')
const loading = ref(false)

const form = reactive({
  nickname: '',
  avatar: '',
  realname: ''
})

const genderOptions = ['请选择', '男', '女']
const genderIndex = ref(0)

onMounted(() => {
  const user = userStore.user
  if (user) {
    form.nickname = user.nickname || ''
    form.avatar = user.avatar || ''
    form.realname = user.realname || ''
    genderIndex.value = user.gender || 0
  }
})

const chooseAvatar = () => {
  uni.chooseImage({
    count: 1,
    sizeType: ['compressed'],
    sourceType: ['album', 'camera'],
    success: (res) => {
      form.avatar = res.tempFilePaths[0]
    },
    fail: () => {
      uni.showToast({ title: '选择图片失败', icon: 'none' })
    }
  })
}

const handleGenderChange = (e: any) => {
  genderIndex.value = e.detail.value
}

const handleSave = async () => {
  if (!form.nickname.trim()) {
    uni.showToast({ title: '请输入昵称', icon: 'none' })
    return
  }
  
  loading.value = true
  
  try {
    await userStore.updateUser({
      nickname: form.nickname.trim(),
      avatar: form.avatar,
      gender: genderIndex.value,
      realname: form.realname.trim()
    })
    uni.showToast({ title: '保存成功', icon: 'success' })
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  } catch (e: any) {
    uni.showToast({ title: e?.message || '保存失败', icon: 'none' })
  } finally {
    loading.value = false
  }
}

const goBack = () => {
  uni.navigateBack()
}

const goToChangePassword = () => {
  uni.navigateTo({ url: '/pages/user/change-password' })
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

.save-btn {
  font-size: 30rpx;
  color: #4CAF50;
  
  &.disabled {
    color: #ccc;
  }
}

.card {
  background: #fff;
  margin: 20rpx;
  border-radius: 16rpx;
  overflow: hidden;
}

.profile-item {
  display: flex;
  align-items: center;
  padding: 28rpx 30rpx;
  border-bottom: 1rpx solid #f0f0f0;

  &:last-child {
    border-bottom: none;
  }
}

.item-label {
  font-size: 30rpx;
  color: #333;
  width: 140rpx;
}

.item-input {
  flex: 1;
  font-size: 30rpx;
  color: #333;
  text-align: right;
}

.item-value {
  flex: 1;
  font-size: 30rpx;
  color: #999;
  text-align: right;
}

.avatar-item {
  align-items: flex-start;
}

.avatar-wrap {
  position: relative;
  width: 120rpx;
  height: 120rpx;
  border-radius: 50%;
  overflow: hidden;
  margin-left: auto;
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
  height: 40rpx;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.picker-value {
  flex: 1;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 30rpx;
  color: #333;
}
</style>
