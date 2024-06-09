<script setup lang="ts">
import { useCVStore } from '@/stores/cvStore.ts'
import { timeToStr } from '@/tools/main.ts'
import { convertFileSrc } from '@tauri-apps/api/tauri'

const cvStore = useCVStore()
</script>

<template>
  <n-flex justify="space-between" align="center" :warp="false">
    <div>
      <div style="font-size: 14pt; font-weight: 800">
        {{ cvStore.cvLoadedData?.cv_base_info.name }}
      </div>
      <n-space>
        <div>
          {{
            timeToStr(
              cvStore.cvLoadedData?.cv_base_info.birthday
                ? cvStore.cvLoadedData?.cv_base_info.birthday
                : 0
            )
          }}
        </div>
        <div>{{ cvStore.cvLoadedData?.cv_base_info.mobile }}</div>
        <div>{{ cvStore.cvLoadedData?.cv_base_info.email }}</div>
      </n-space>
    </div>
    <n-image
      :src="
        convertFileSrc(
          cvStore.cvLoadedData?.cv_base_info.photo_path
            ? cvStore.cvLoadedData?.cv_base_info.photo_path
            : ''
        )
      "
      width="88"
      height="117"
      preview-disabled
      lazy
    />
  </n-flex>
</template>

<style scoped></style>
