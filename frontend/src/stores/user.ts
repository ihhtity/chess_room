import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { User, Membership } from '@/types'
import { userApi, membershipApi } from '@/api'

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const token = ref<string>('')
  const membership = ref<Membership | null>(null)

  // 检查用户是否已登录
  const isLoggedIn = () => {
    return !!token.value || !!uni.getStorageSync('user_token')
  }
  // 登录
  const login = async (data: { phone?: string; password?: string; open_id?: string; nickname?: string; avatar?: string; gender?: number }) => {
    const res = await userApi.login(data)
    user.value = res.user
    token.value = res.token
    uni.setStorageSync('user', res.user)
    uni.setStorageSync('user_token', res.token)
    await loadMembership()
  }
  // 加载用户信息
  const loadUserInfo = async () => {
    try {
      const res = await userApi.getInfo()
      user.value = res
    } catch (e) {
      console.error('加载用户信息失败', e)
    }
  }
  // 加载会员信息
  const loadMembership = async () => {
    try {
      const res = await membershipApi.getInfo()
      membership.value = res
    } catch (e) {
      console.error('加载会员信息失败', e)
    }
  }
  // 退出登录
  const logout = () => {
    user.value = null
    token.value = ''
    membership.value = null
    uni.removeStorageSync('user_token')
    uni.navigateTo({ url: '/pages/user/login' })
  }
  // 更新用户信息
  const updateUser = async (data: { nickname?: string; avatar?: string; gender?: number; realname?: string }) => {
    const res = await userApi.updateInfo(data)
    user.value = res
  }
  // 修改密码
  const changePassword = async (data: { old_password: string; new_password: string }) => {
    await userApi.changePassword(data)
  }
  // 初始化用户状态
  const init = async () => {
    const savedToken = uni.getStorageSync('user_token')
    if (savedToken) {
      token.value = savedToken
      await Promise.all([loadUserInfo(), loadMembership()])
    }
  }

  return {
    user,
    token,
    membership,
    isLoggedIn,
    login,
    loadUserInfo,
    loadMembership,
    logout,
    updateUser,
    changePassword,
    init
  }
})
