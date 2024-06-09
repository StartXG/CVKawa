<script setup lang="ts">
import { ref } from 'vue'
import { useCVStore } from '@/stores/cvStore.ts'
import { Edit24Regular, Delete24Regular, Save24Regular } from '@vicons/fluent'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const cvStore = useCVStore()
const skillItem = ref<string | null>(null)
const editIndex = ref<number | null>(null)

function addItem() {
  skillItem.value !== null && cvStore.cv.cv_skills.push(skillItem.value) && (skillItem.value = '')
}
</script>

<template>
  <n-list bordered size="small" hoverable clickable>
    <n-list-item @keyup.enter="addItem">
      <n-input v-model:value="skillItem" :placeholder="t('editPage.skills.inputPlaceholder')" />
    </n-list-item>
    <n-list-item v-for="(v, k) in cvStore.cv.cv_skills" :key="k">
      <template #suffix>
        <n-button-group>
          <n-button size="small" tertiary type="info" @click="editIndex = k" v-if="editIndex !== k">
            {{ t('editPage.skills.action.edit') }}
            <template #icon>
              <n-icon>
                <Edit24Regular />
              </n-icon>
            </template>
          </n-button>
          <n-button size="small" tertiary type="info" @click="editIndex = null" v-else>
            {{ t('editPage.skills.action.save') }}
            <template #icon>
              <n-icon>
                <Save24Regular />
              </n-icon>
            </template>
          </n-button>
          <n-button size="small" tertiary type="error" @click="cvStore.cv.cv_skills.splice(k, 1)">
            {{ t('editPage.skills.action.delete') }}
            <template #icon>
              <n-icon>
                <Delete24Regular />
              </n-icon>
            </template>
          </n-button>
        </n-button-group>
      </template>
      <template #default>
        <n-input
          v-if="editIndex === k"
          size="small"
          clearable
          v-model:value="cvStore.cv.cv_skills[k]"
          @keyup.enter="editIndex = null"
        />
        <div v-else>{{ v }}</div>
      </template>
    </n-list-item>
  </n-list>
</template>

<style scoped></style>
