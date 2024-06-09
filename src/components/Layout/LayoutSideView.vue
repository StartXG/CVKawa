<script setup lang="ts">
import { Add24Regular } from '@vicons/fluent'
import CVCardList from '@/components/CVCardList.vue'
import { useRouter } from 'vue-router'
import { useCVStore } from '@/stores/cvStore.ts'

const router = useRouter()
const cvStore = useCVStore()
</script>

<template>
  <div class="full-container">
    <n-button
      size="large"
      secondary
      block
      type="success"
      @click="router.push({ name: 'CVKawaEdit', query: { mode: 'new', v: null } })"
    >
      Create New
      <template #icon>
        <n-icon>
          <Add24Regular />
        </n-icon>
      </template>
    </n-button>
    <div class="cv-list">
      <n-scrollbar>
        <n-list hoverable clickable>
          <CVCardList
            v-for="(v, k) in cvStore.cv_versions"
            :key="k"
            :id="v.id"
            :name="v.name"
            :updated_at="v.updated_at"
            :version="v.version"
          />
        </n-list>
      </n-scrollbar>
    </div>
  </div>
</template>

<style scoped>
.cv-list {
  height: calc(100% - 40px);
  width: 100%;
  margin-top: 24px;
}
</style>
