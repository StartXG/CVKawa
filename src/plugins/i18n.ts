import { createI18n } from 'vue-i18n'
import { type App } from 'vue'
import { useAppStore } from '@/stores/appStore'

const appStore = useAppStore()
const i18nFiles: Record<string, any> = import.meta.glob('../i18n/*.ts', { eager: true })
const messages: Record<string, any> = {}
Object.keys(i18nFiles).forEach((key) => {
  const k = key.split('/')[2].split('.')[0].split('-')[0]
  messages[k] = i18nFiles[key].default
})

const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  global: true,
  locale: appStore.lang,
  fallbackLocale: 'en-US',
  messages
})

export default function setupI18n(app: App) {
  app.use(i18n)
}
export { i18n }
