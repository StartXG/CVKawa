<script setup lang="ts">
import { timeToStr } from '@/tools/main.ts'
import { useCVStore } from '@/stores/cvStore.ts'
const cvStore = useCVStore()
</script>

<template>
  <div v-for="(v, k) in cvStore.cvLoadedData?.cv_edu_info" :key="k">
    <div class="school-title">
      {{ v.school }}
    </div>
    <n-flex align="center" justify="space-between">
      <n-space>
        <div>{{ v.major }}</div>
        <div>{{ v.academic_degree }}</div>
      </n-space>
      <div>
        {{ timeToStr(v?.start_time ?? undefined) }}&nbsp;-&nbsp;{{
          timeToStr(v?.end_time ?? undefined)
        }}
      </div>
    </n-flex>
  </div>
</template>

<style scoped>
.school-title {
  font-size: 12pt;
  font-weight: bold;
}
</style>
