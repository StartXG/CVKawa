// import { type Component, h } from 'vue'
// import { NIcon } from 'naive-ui'
import { useAppStore } from '@/stores/appStore.ts'

const appStore = useAppStore()
function timeToStr(timestamp: number = 0): string {
  const date = new Date(timestamp)
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  if (appStore.lang === 'en-US') {
    return month + '/' + day + '/' + year
  }
  return year + '年' + month + '月' + day + '日'
}

// function getItem(key: string): string | null {
//   return <string | null>localStorage.getItem(key)
// }

// function setItem(key: string, value: string): boolean {
//   try {
//     localStorage.setItem(key, value)
//     return true
//   } catch (e) {
//     console.log(e)
//     return false
//   }
// }
// function renderIcon(icon: Component) {
//   return () => {
//     return h(NIcon, null, {
//       default: () => h(icon)
//     })
//   }
// }

export {
  timeToStr
  // getItem,
  // setItem,
  // renderIcon
}
