import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { darkTheme, dateEnUS, dateZhCN, enUS, useOsTheme, zhCN, jaJP, dateJaJP } from 'naive-ui'
import { store } from '@/plugins/store.ts'
import { invoke } from '@tauri-apps/api'

export const appStore = defineStore('appStore', () => {
  const osThemeRef = useOsTheme()
  const lang = ref<string>('')
  const langOptions = [
    { label: '中文', value: 'zh-CN' },
    { label: 'English', value: 'en-US' },
    { label: '日本語', value: 'ja-JP' }
  ]
  const theme = computed(() => {
    return osThemeRef.value === 'dark' ? darkTheme : osThemeRef
  })

  async function setLang(value: string): Promise<string> {
    const result: string = await invoke('get_config', { key: 'lang' })
    lang.value = result || value
    return lang.value
  }

  async function changeLang() {
    await invoke('set_config', { key: 'lang', value: lang.value })
  }

  const UILangSys = computed(() => {
    switch (lang.value) {
      case 'zh-CN':
        return { l: zhCN, dl: dateZhCN }
      case 'en-US':
        return { l: enUS, dl: dateEnUS }
      case 'ja-JP':
        return { l: jaJP, dl: dateJaJP }
      default:
        return { l: enUS, dl: dateEnUS }
    }
  })

  return {
    theme,
    lang,
    setLang,
    UILangSys,
    langOptions,
    changeLang
  }
})

export function useAppStore() {
  return appStore(store)
}
