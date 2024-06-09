<script setup lang="ts">
import { useAppStore } from './stores/appStore.ts'
import { RouterView } from 'vue-router'
import { onMounted } from 'vue'
import SystemApi from './api/SystemApi.ts'
import { useI18n } from 'vue-i18n'

const appStore = useAppStore()
const i18n = useI18n()

async function setLang() {
  const systemLanguage: string = navigator.language || navigator.language
  let lang = await appStore.setLang(systemLanguage)
  i18n.locale.value = lang
  console.log('Set Lang:', lang)
}

function initDB() {
  SystemApi.CheckAndInitDB()
    .then(() => {
      console.log('DB Init Success')
    })
    .catch((err) => {
      console.log('DB Init Failed', err)
    })
}

onMounted(() => {
  setLang()
  initDB()
})
</script>

<template>
  <n-config-provider
    :theme="appStore.theme"
    style="width: 100%; height: 100%"
    :locale="appStore.UILangSys.l"
    :date-locale="appStore.UILangSys.dl"
  >
    <n-dialog-provider>
      <n-message-provider>
        <n-modal-provider>
          <n-notification-provider>
            <RouterView />
          </n-notification-provider>
        </n-modal-provider>
      </n-message-provider>
      <n-global-style />
    </n-dialog-provider>
  </n-config-provider>
</template>
