<script setup lang="ts">
import { onLaunch, onShow, onHide } from '@dcloudio/uni-app'

onLaunch(() => {
  console.log('App Launch')
})

onShow(() => {
  console.log('App Show')
})

onHide(() => {
  console.log('App Hide')
})
</script>

<style lang="scss">
@import '@climblee/uv-ui/index.scss';

page {
  background-color: #f5f5f5;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.container {
  padding: 20rpx;
}

.card {
  background: #fff;
  border-radius: 16rpx;
  padding: 24rpx;
  margin-bottom: 20rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.06);
}

.btn-primary {
  background: #4CAF50;
  color: #fff;
  border: none;
  border-radius: 8rpx;
  padding: 24rpx;
  font-size: 28rpx;
  text-align: center;
}

.btn-primary:active {
  background: #45a049;
}

.text-primary {
  color: #4CAF50;
}

.text-muted {
  color: #999;
}

.text-success {
  color: #4CAF50;
}

.text-warning {
  color: #ff9800;
}

.text-danger {
  color: #f44336;
}

.flex {
  display: flex;
}

.flex-col {
  flex-direction: column;
}

.items-center {
  align-items: center;
}

.justify-between {
  justify-content: space-between;
}

.flex-1 {
  flex: 1;
}

.mt-20 {
  margin-top: 20rpx;
}

.mb-20 {
  margin-bottom: 20rpx;
}

.w-full {
  width: 100%;
}
</style>
