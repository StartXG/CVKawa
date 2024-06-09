<script setup lang="ts">
import { useCVStore } from '@/stores/cvStore.ts'
import { CVEduInfoSchema } from '@/common/cvSchemas.ts'
import { type DataTableColumns, NButton, NInput, NDatePicker } from 'naive-ui'

import { h } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const cvStore = useCVStore()
const columns: DataTableColumns<CVEduInfoSchema> = [
  {
    title: t('editPage.eduInfo.school'),
    key: 'school',
    align: 'center',
    render(row, index) {
      return h(NInput, {
        value: row.school,
        size: 'small',
        clearable: true,
        onUpdateValue(v) {
          cvStore.cv.cv_edu_info[index].school = v
        }
      })
    }
  },
  {
    title: t('editPage.eduInfo.major'),
    key: 'major',
    align: 'center',
    render(row, index) {
      return h(NInput, {
        value: row.major,
        size: 'small',
        clearable: true,
        onUpdateValue(v) {
          cvStore.cv.cv_edu_info[index].major = v
        }
      })
    }
  },
  {
    title: t('editPage.eduInfo.degree'),
    key: 'academic_degree',
    align: 'center',
    render(row, index) {
      return h(NInput, {
        value: row.academic_degree,
        size: 'small',
        clearable: true,
        onUpdateValue(v) {
          cvStore.cv.cv_edu_info[index].academic_degree = v
        }
      })
    }
  },
  {
    title: t('editPage.eduInfo.startDate'),
    key: 'start_time',
    align: 'center',
    render(row, index) {
      return h(NDatePicker, {
        value: row.start_time,
        type: 'date',
        size: 'small',
        clearable: true,
        onUpdateValue(v) {
          cvStore.cv.cv_edu_info[index].start_time = v
        }
      })
    }
  },
  {
    title: t('editPage.eduInfo.endDate'),
    key: 'end_time',
    align: 'center',
    sorter: (row1: any = 0, row2: any = 0) => row1.end_time - row2.end_time,
    render(row, index) {
      return h(NDatePicker, {
        value: row.end_time,
        type: 'date',
        size: 'small',
        clearable: true,
        onUpdateValue(v) {
          cvStore.cv.cv_edu_info[index].end_time = v
        }
      })
    }
  },
  {
    title: t('editPage.eduInfo.action.action'),
    key: 'actions',
    align: 'center',
    render: (row: any) => {
      return h(
        NButton,
        {
          round: true,
          tertiary: true,
          size: 'small',
          type: 'error',
          onClick: () => {
            removeItem(row)
          }
        },
        { default: () => t('editPage.eduInfo.action.delete') }
      )
    }
  }
]
function removeItem(row: CVEduInfoSchema) {
  const index = cvStore.cv.cv_edu_info.indexOf(row)
  cvStore.cv.cv_edu_info.splice(index, 1)
}
function addEdu() {
  cvStore.cv.cv_edu_info.push({
    school: '',
    major: '',
    academic_degree: '',
    start_time: null,
    end_time: null
  })
}
</script>

<template>
  <div class="full-container">
    <n-button
      type="success"
      tertiary
      size="small"
      style="margin-bottom: 1rem"
      @click="addEdu"
      block
    >
      {{ t('editPage.eduInfo.addEdu') }}
    </n-button>
    <n-data-table striped :columns="columns" :data="cvStore.cv.cv_edu_info" size="small" />
  </div>
</template>

<style scoped></style>
