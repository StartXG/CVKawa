<script setup lang="ts">
import { Delete24Regular, Edit24Regular } from '@vicons/fluent'
import { timeToStr } from '@/tools/main.ts'
import { useCVStore } from '@/stores/cvStore.ts'

const cvStore = useCVStore()
const props = defineProps<{
  name: string
  version: string
  id?: string
  updated_at?: number
}>()
</script>

<template>
  <n-list-item @click="cvStore.loadCV(props.version)">
    <n-thing :title="props.name">
      <template #description>
        <n-tag :bordered="false" type="info" size="small">
          last update: {{ timeToStr(props.updated_at) }}
        </n-tag>
      </template>
    </n-thing>
    <template #suffix>
      <n-button-group>
        <n-button size="small" ghost type="info">
          <template #icon>
            <n-icon>
              <Edit24Regular />
            </n-icon>
          </template>
        </n-button>
        <n-popconfirm @positive-click="cvStore.delCV(props.version)" @negative-click="() => {}">
          Confirm to delete?
          <template #trigger>
            <n-button size="small" ghost type="error">
              <template #icon>
                <n-icon>
                  <Delete24Regular />
                </n-icon>
              </template>
            </n-button>
          </template>
        </n-popconfirm>
      </n-button-group>
    </template>
  </n-list-item>
</template>

<style scoped></style>
