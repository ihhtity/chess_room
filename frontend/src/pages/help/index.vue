<template>
  <view class="container">
    <view class="search-bar">
      <uv-icon name="search" color="#999" size="28" />
      <input class="search-input" placeholder="搜索帮助内容" v-model="searchText" />
    </view>

    <view class="section-title">常见问题</view>

    <view class="faq-list">
      <view 
        class="faq-item" 
        v-for="(item, index) in filteredFaq" 
        :key="index"
        @click="toggleFaq(index)"
      >
        <view class="faq-header">
          <text class="faq-icon">{{ index + 1 }}</text>
          <text class="faq-title">{{ item.title }}</text>
          <uv-icon 
            :name="expandedIndex === index ? 'arrow-up' : 'arrow-down'" 
            color="#999" 
            size="24" 
          />
        </view>
        <view class="faq-content" :class="{ expanded: expandedIndex === index }">
          <text class="content-text">{{ item.content }}</text>
        </view>
      </view>
    </view>

    <view class="contact-card">
      <text class="card-title">联系客服</text>
      <view class="contact-info">
        <view class="contact-item" @click="makePhoneCall">
          <uv-icon name="phone" color="#4CAF50" size="28" />
          <text class="contact-text">客服热线：400-888-8888</text>
          <uv-icon name="arrow-right" color="#ccc" size="24" />
        </view>
        <view class="contact-item">
          <uv-icon name="clock" color="#2196f3" size="28" />
          <text class="contact-text">服务时间：9:00 - 22:00</text>
        </view>
        <view class="contact-item">
          <uv-icon name="email" color="#ff9800" size="28" />
          <text class="contact-text">邮箱：support@chessroom.com</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const searchText = ref('')
const expandedIndex = ref<number | null>(null)

const faqList = [
  {
    title: '如何预约包间？',
    content: '1. 进入首页或包间列表页面\n2. 选择您喜欢的包间\n3. 点击"立即预约"按钮\n4. 选择预约日期和时间段\n5. 确认订单信息并支付\n6. 预约成功后您将收到短信通知'
  },
  {
    title: '支持哪些支付方式？',
    content: '我们支持以下支付方式：\n- 微信支付\n- 支付宝\n- 会员卡余额支付\n- 优惠券抵扣\n支付时请确保您的账户有足够余额。'
  },
  {
    title: '如何取消预约？',
    content: '1. 进入"我的订单"页面\n2. 找到需要取消的订单\n3. 点击"取消订单"按钮\n4. 确认取消操作\n\n温馨提示：提前2小时取消可全额退款，不足2小时取消将扣除30%手续费。'
  },
  {
    title: '会员有哪些权益？',
    content: '会员享受以下专属权益：\n- 所有包间享受8折优惠\n- 生日当月额外赠送2小时免费时长\n- 优先预约热门包间\n- 专属客服通道\n- 积分兑换礼品'
  },
  {
    title: '包间收费标准是怎样的？',
    content: '包间收费标准如下：\n- 小包间：60元/小时\n- 中包间：80元/小时\n- 大包间：120元/小时\n- VIP包间：200元/小时\n\n会员可享受相应折扣优惠。'
  },
  {
    title: '如何修改个人信息？',
    content: '1. 进入"我的"页面\n2. 点击头像进入个人资料\n3. 修改您需要更新的信息\n4. 点击"保存"按钮\n\n注意：手机号暂不支持修改，如需修改请联系客服。'
  },
  {
    title: '忘记密码怎么办？',
    content: '1. 在登录页面点击"忘记密码"\n2. 输入注册手机号\n3. 获取并输入验证码\n4. 设置新密码\n5. 使用新密码登录\n\n如果收不到验证码，请检查短信是否被拦截。'
  },
  {
    title: '如何使用优惠券？',
    content: '1. 在优惠券页面查看可用优惠券\n2. 预约时选择使用优惠券\n3. 系统自动抵扣相应金额\n\n注意：每张优惠券有使用条件，订单金额需满足最低消费要求。'
  }
]

const filteredFaq = computed(() => {
  if (!searchText.value.trim()) {
    return faqList
  }
  return faqList.filter(item => 
    item.title.includes(searchText.value) || 
    item.content.includes(searchText.value)
  )
})

const toggleFaq = (index: number) => {
  expandedIndex.value = expandedIndex.value === index ? null : index
}

const makePhoneCall = () => {
  uni.makePhoneCall({
    phoneNumber: '400-888-8888'
  })
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

.search-bar {
  display: flex;
  align-items: center;
  background: #fff;
  margin: 20rpx;
  border-radius: 40rpx;
  padding: 0 30rpx;
  height: 80rpx;
}

.search-input {
  flex: 1;
  font-size: 28rpx;
  color: #333;
  padding-left: 20rpx;
}

.section-title {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  padding: 20rpx 30rpx;
}

.faq-list {
  background: #fff;
  margin: 0 20rpx;
  border-radius: 16rpx;
  overflow: hidden;
}

.faq-item {
  border-bottom: 1rpx solid #f0f0f0;
  
  &:last-child {
    border-bottom: none;
  }
}

.faq-header {
  display: flex;
  align-items: center;
  padding: 28rpx 30rpx;
}

.faq-icon {
  width: 40rpx;
  height: 40rpx;
  background: #4CAF50;
  color: #fff;
  font-size: 24rpx;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 20rpx;
}

.faq-title {
  flex: 1;
  font-size: 30rpx;
  color: #333;
}

.faq-content {
  max-height: 0;
  overflow: hidden;
  transition: max-height 0.3s ease;
  background: #fafafa;
  
  &.expanded {
    max-height: 500rpx;
    padding: 0 30rpx 28rpx;
  }
}

.content-text {
  font-size: 26rpx;
  color: #666;
  line-height: 1.8;
  white-space: pre-line;
}

.contact-card {
  background: #fff;
  margin: 30rpx 20rpx;
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

.contact-info {
  display: flex;
  flex-direction: column;
}

.contact-item {
  display: flex;
  align-items: center;
  padding: 20rpx 0;
  border-bottom: 1rpx solid #f0f0f0;
  
  &:last-child {
    border-bottom: none;
  }
}

.contact-text {
  flex: 1;
  font-size: 28rpx;
  color: #666;
  margin-left: 16rpx;
}
</style>