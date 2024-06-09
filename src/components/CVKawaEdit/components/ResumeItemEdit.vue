<script setup lang="ts">
import { CVResumeStepSchema } from '@/common/cvSchemas.ts'
import { PropType } from 'vue'

const model = defineModel({
  type: Object as PropType<CVResumeStepSchema>,
  default: {
    department: '',
    position: '',
    startTime: '',
    endTime: '',
    responsibilities: ''
  }
})
</script>

<template>
  <n-card :bordered="false" class="resume-item" title="Step1" size="small" closable hoverable>
    <template></template>
    <n-grid :x-gap="12" :y-gap="8">
      <n-grid-item :span="6">
        <n-form-item label="department">
          <n-input v-model:value="model.department" />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="6">
        <n-form-item label="position">
          <n-input v-model:value="model.position" />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="6">
        <n-form-item label="start time">
          <n-date-picker style="width: 100%" type="date" v-model:value="model.start_time" />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="6">
        <n-form-item label="end time">
          <n-date-picker style="width: 100%" type="date" v-model:value="model.end_time" />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="24">
        <n-form-item label="responsibilities">
          <n-input type="textarea" v-model:value="model.responsibilities" />
        </n-form-item>
      </n-grid-item>
    </n-grid>
  </n-card>
</template>

<style scoped></style>
