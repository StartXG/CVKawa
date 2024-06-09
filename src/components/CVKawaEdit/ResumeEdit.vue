<script setup lang="ts">
import { useCVStore } from '@/stores/cvStore.ts'
import { reactive } from 'vue'
import { CVResumeSchema, CVResumeStepSchema } from '@/common/cvSchemas.ts'
import { NCard } from 'naive-ui'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const cvStore = useCVStore()

const resumeData = reactive<CVResumeSchema>({
  resume_steps: [],
  company: '',
  achievement: ''
})

function addCompany() {
  cvStore.cv.cv_resume.push(JSON.parse(JSON.stringify(resumeData)))
}

function removeCompany(k: number) {
  return () => {
    cvStore.cv.cv_resume.splice(k, 1)
  }
}

function addResumeStep(k: number) {
  let step: CVResumeStepSchema = {
    department: null,
    position: null,
    start_time: null,
    end_time: null,
    responsibilities: null
  }
  cvStore.cv.cv_resume[k].resume_steps.push(step)
}

function removeResumeStep(k: number, subK: number) {
  return () => {
    cvStore.cv.cv_resume[k].resume_steps.splice(subK, 1)
  }
}
</script>

<template>
  <n-empty
    class="full-container flex-full-center"
    size="small"
    :description="t('editPage.resumeInfo.noItem')"
    v-if="cvStore.cv.cv_resume.length === 0"
  >
    <template #extra>
      <n-button tertiary size="small" type="info" @click="addCompany">
        {{ t('editPage.resumeInfo.addResume') }}
      </n-button>
    </template>
  </n-empty>
  <div v-else class="full-container" style="overflow: scroll">
    <n-card
      v-for="(v, k) in cvStore.cv.cv_resume"
      :key="k"
      style="margin-bottom: 12px"
      size="small"
      closable
      :on-close="removeCompany(k)"
    >
      <n-form size="small">
        <n-form-item :label="t('editPage.resumeInfo.company')">
          <n-input v-model:value="cvStore.cv.cv_resume[k].company" />
        </n-form-item>
        <n-button
          v-if="v.resume_steps.length === 0"
          size="small"
          @click="addResumeStep(k)"
          style="margin-bottom: 1rem"
        >
          {{ t('editPage.resumeInfo.addStep') }}
        </n-button>
        <n-card
          embedded
          :bordered="false"
          v-for="(subV, subK) in v.resume_steps"
          :key="subK"
          style="margin-bottom: 1rem"
          size="small"
          closable
          :on-close="removeResumeStep(k, subK)"
        >
          <n-grid :x-gap="4" :y-gap="4">
            <n-form-item-gi :label="t('editPage.resumeInfo.position')" :span="6">
              <n-input style="width: 100%" v-model:value="subV.position" />
            </n-form-item-gi>
            <n-form-item-gi :label="t('editPage.resumeInfo.department')" :span="6">
              <n-input style="width: 100%" v-model:value="subV.department" />
            </n-form-item-gi>
            <n-form-item-gi :label="t('editPage.resumeInfo.startDate')" :span="6">
              <n-date-picker style="width: 100%" type="date" v-model:value="subV.start_time" />
            </n-form-item-gi>
            <n-form-item-gi :label="t('editPage.resumeInfo.endDate')" :span="6">
              <n-date-picker style="width: 100%" type="date" v-model:value="subV.end_time" />
            </n-form-item-gi>
            <n-form-item-gi :label="t('editPage.resumeInfo.responsibilities')" :span="24">
              <n-input type="textarea" style="width: 100%" v-model:value="subV.responsibilities" />
            </n-form-item-gi>
          </n-grid>
        </n-card>
        <n-button
          v-if="v.resume_steps.length > 0"
          size="small"
          @click="addResumeStep(k)"
          style="margin-bottom: 1rem"
        >
          {{ t('editPage.resumeInfo.addStep') }}
        </n-button>
        <n-form-item :label="t('editPage.resumeInfo.achievement')">
          <n-input type="textarea" v-model:value="cvStore.cv.cv_resume[k].achievement" />
        </n-form-item>
      </n-form>
    </n-card>
    <n-divider>
      <n-button ghost round dashed type="success" @click="addCompany">{{
        t('editPage.resumeInfo.addResume')
      }}</n-button>
    </n-divider>
  </div>
</template>

<style scoped></style>
