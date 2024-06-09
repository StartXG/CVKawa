<script setup lang="ts">
import { useCVStore } from '@/stores/cvStore.ts'
import { open } from '@tauri-apps/api/dialog'
import { convertFileSrc } from '@tauri-apps/api/tauri'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const cvStore = useCVStore()

async function openHandle() {
  let filePath = await open({
    directory: false,
    multiple: false
  })
  if (filePath) {
    cvStore.cv.cv_base_info.photo_path = filePath.toString()
  }
}

function removeFile() {
  cvStore.cv.cv_base_info.photo_path = null
}
</script>

<template>
  <n-form class="full-container">
    <n-grid :x-gap="8" :y-gap="8">
      <n-grid-item :span="12">
        <n-form-item :label="t('editPage.baseInfo.fullName')">
          <n-input v-model:value="cvStore.cv.cv_base_info.name" clearable />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="12">
        <n-form-item :label="t('editPage.baseInfo.mobile')">
          <n-input-number
            v-model:value="cvStore.cv.cv_base_info.mobile"
            style="width: 100%"
            :show-button="false"
            clearable
          />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="12">
        <n-form-item :label="t('editPage.baseInfo.email')">
          <n-input v-model:value="cvStore.cv.cv_base_info.email" clearable />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="12">
        <n-form-item :label="t('editPage.baseInfo.birthday')">
          <n-date-picker
            type="date"
            v-model:value="cvStore.cv.cv_base_info.birthday"
            style="width: 100%"
            clearable
          />
        </n-form-item>
      </n-grid-item>
      <n-grid-item :span="12">
        <n-form-item :label="t('editPage.baseInfo.oneInchPhoto')">
          <n-card
            v-if="cvStore.cv.cv_base_info.photo_path"
            embedded
            size="small"
            hoverable
            :title="t('editPage.baseInfo.photo.preview')"
            closable
            @close="removeFile"
          >
            <n-image
              :src="
                convertFileSrc(
                  cvStore.cv.cv_base_info.photo_path ? cvStore.cv.cv_base_info.photo_path : ''
                )
              "
              width="200"
              preview-disabled
              lazy
            />
          </n-card>
          <n-button v-else @click="openHandle">{{ t('editPage.baseInfo.photo.upload') }}</n-button>
        </n-form-item>
      </n-grid-item>
      <n-form-item-gi :span="12">
        <n-p depth="3" style="margin: 8px 0 0 0">
          {{ t('editPage.baseInfo.photo.tip') }}
        </n-p>
      </n-form-item-gi>
    </n-grid>
  </n-form>
</template>

<style scoped></style>
