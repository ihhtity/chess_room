export default {
  pages: [
    'pages/index/index',
    'pages/room/detail',
    'pages/booking/create',
    'pages/order/list',
    'pages/order/detail',
    'pages/member/index',
    'pages/user/profile',
    'pages/user/login'
  ],
  globalStyle: {
    navigationBarTextStyle: 'black',
    navigationBarTitleText: '棋牌室',
    navigationBarBackgroundColor: '#ffffff',
    backgroundColor: '#f5f5f5',
    backgroundTextStyle: 'dark'
  },
  tabBar: {
    custom: true,
    color: '#999999',
    selectedColor: '#667eea',
    borderStyle: 'black',
    backgroundColor: '#ffffff',
    list: [
      { pagePath: 'pages/index/index', text: '首页' },
      { pagePath: 'pages/order/list', text: '订单' },
      { pagePath: 'pages/member/index', text: '会员' },
      { pagePath: 'pages/user/profile', text: '我的' }
    ]
  }
}