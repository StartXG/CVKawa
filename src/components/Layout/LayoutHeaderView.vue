<script setup lang="ts">
import LogoView from '@/components/LogoView.vue'
import {
  // Settings24Regular,
  ArrowLeft24Regular,
  ArrowRight24Regular,
  Add24Regular
} from '@vicons/fluent'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores/appStore.ts'
import { watch } from 'vue'

const appStore = useAppStore()
const i18n = useI18n()
const router = useRouter()
const route = useRoute()

watch(
  () => appStore.lang,
  () => {
    i18n.locale.value = appStore.lang
    appStore.changeLang()
  },
  { immediate: true }
)
</script>

<template>
  <n-flex align="center" justify="space-between" style="height: 100%">
    <LogoView size="small" @click="router.push({ name: 'CVList' })" class="logo" />
    <n-space>
      <n-button @click="router.back()" size="small" circle type="success">
        <template #icon>
          <n-icon>
            <ArrowLeft24Regular />
          </n-icon>
        </template>
      </n-button>
      <n-button @click="router.forward()" size="small" circle type="success">
        <template #icon>
          <n-icon>
            <ArrowRight24Regular />
          </n-icon>
        </template>
      </n-button>
      <n-button
        v-if="route.name !== 'CVKawaEdit'"
        size="small"
        round
        type="success"
        @click="router.push({ name: 'CVKawaEdit', query: { mode: 'new' } })"
      >
        {{ i18n.t('layoutPage.create') }}
        <template #icon>
          <n-icon>
            <Add24Regular />
          </n-icon>
        </template>
      </n-button>
      <n-select
        style="width: 100px"
        v-model:value="appStore.lang"
        size="small"
        :options="appStore.langOptions"
      />
    </n-space>
  </n-flex>
</template>

<style scoped>
.logo:hover {
  cursor: pointer;
}
</style>
