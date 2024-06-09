<script setup lang="ts">
import { ArrowRight32Regular } from '@vicons/fluent'
import LogoView from '@/components/LogoView.vue'
import { useRouter } from 'vue-router'
import SystemApi from '@/api/SystemApi.ts'
import { onMounted } from 'vue'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const { t } = useI18n()
onMounted(() => {
  SystemApi.CheckAndInitDB().then((res: boolean) => {
    if (res) {
      console.log('DB is ready')
    } else {
      console.log('DB is not ready, it will be initialized automatically')
    }
  })
})
</script>

<template>
  <n-flex align="center" justify="center" style="height: 100%; flex-direction: column">
    <n-flex align="center" justify="center">
      <n-h1 style="margin: 0">{{ t('welcomePage.title') }}</n-h1>
      <LogoView size="normal" />
    </n-flex>
    <n-button size="large" circle @click="router.push({ name: 'LayoutView' })">
      <template #icon>
        <n-icon>
          <ArrowRight32Regular />
        </n-icon>
      </template>
    </n-button>
  </n-flex>
</template>

<style scoped></style>
