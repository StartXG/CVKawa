<script setup lang="ts">
import { timeToStr } from '@/tools/main.ts'
import { useCVStore } from '@/stores/cvStore.ts'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const cvStore = useCVStore()
</script>

<template>
  <div v-if="cvStore.cvLoadedData && cvStore.cvLoadedData?.cv_resume.length === 0" />
  <div v-else v-for="(v, k) in cvStore.cvLoadedData?.cv_resume" :key="k">
    <div class="res-title">{{ v.company }}</div>
    <div v-for="(subV, subK) in v.resume_steps" :key="subK">
      <n-flex align="center" justify="space-between" style="margin-bottom: 0.5rem">
        <n-space>
          <div>{{ subV.position }}</div>
          <div>{{ subV.department }}</div>
        </n-space>
        <div>
          {{ timeToStr(subV.start_time ?? undefined) }}&nbsp;-&nbsp;{{
            timeToStr(subV.end_time ?? undefined)
          }}
        </div>
      </n-flex>
      <div style="margin-bottom: 1rem">
        <div class="res-sub-title">{{ t('preview.responsibilities') }}</div>
        <div v-html="subV.responsibilities?.replace(/\n/g, '<br />')" />
      </div>
    </div>
    <div>
      <div class="res-sub-title">{{ t('preview.achievement') }}</div>
      <div v-html="v.achievement?.replace(/\n/g, '<br />')" />
    </div>
  </div>
</template>

<style scoped>
.res-title {
  font-size: 12pt;
  font-weight: bold;
}

.res-sub-title {
  font-size: 10pt;
  font-weight: bold;
}
</style>
