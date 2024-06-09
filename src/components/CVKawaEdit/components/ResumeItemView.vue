<script setup lang="ts">
import { CVResumeSchema, CVResumeStepSchema } from '@/common/cvSchemas.ts'
import { computed } from 'vue'
import { timeToStr } from '@/tools/main.ts'

const props = defineProps<{ data: CVResumeSchema; i: number }>()
const emits = defineEmits<{
  (e: 'edit', id: number): void
  (e: 'delete', id: number): void
}>()

const positionList = computed(() => {
  return props.data.resume_steps.map((item: CVResumeStepSchema) => {
    return item.position
  })
})

const timeRange = computed(() => {
  const timeList: Array<number> = []
  props.data.resume_steps.forEach((item: CVResumeStepSchema) => {
    if (item.start_time !== null) {
      timeList.push(item.start_time)
    }
    if (item.end_time !== null) {
      timeList.push(item.end_time)
    }
  })
  return `${timeToStr(Math.min(...timeList))} - ${timeToStr(Math.max(...timeList))}`
})
</script>

<template>
  <n-card :title="props.data.company" size="small" hoverable>
    <template #default>
      <n-text>{{ timeRange }}</n-text>
      <n-space style="margin-top: 1rem">
        <n-tag v-for="(v, k) in positionList" :key="k" size="small" :bordered="false" type="info">{{
          v
        }}</n-tag>
      </n-space>
    </template>
    <template #header-extra>
      <n-space>
        <n-button size="small" tertiary type="info" @click="emits('edit', i)">Edit</n-button>
        <n-button size="small" tertiary type="error" @click="emits('delete', i)">Delete</n-button>
      </n-space>
    </template>
  </n-card>
</template>

<style scoped></style>
