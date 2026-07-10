import { createSSRApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'
import App from './App.vue'
import uvUI from '@climblee/uv-ui'
import zhCN from './locales/zh-CN'

const i18n = createI18n({
  legacy: false,
  locale: 'zh-CN',
  messages: {
    'zh-CN': zhCN
  }
})

export function createApp() {
  const app = createSSRApp(App)
  const pinia = createPinia()
  app.use(pinia)
  app.use(i18n)
  app.use(uvUI)
  return {
    app,
    pinia,
    i18n
  }
}
