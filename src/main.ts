import AppView from './App.vue'
import 'vfonts/Lato.css'
import '@/assets/base.css'
import { type App, createApp } from 'vue'
import setupNaive from '@/plugins/naive.ts'
import setupStore from '@/plugins/store.ts'
import setupRouter from '@/plugins/router.ts'
import setupI18n from '@/plugins/i18n.ts'

function bootstrap(app: App): void {
  setupNaive(app)
  setupStore(app)
  setupRouter(app)
  setupI18n(app)
  app.mount('#app')
}

bootstrap(createApp(AppView))
