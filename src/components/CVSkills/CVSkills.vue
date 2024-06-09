<script setup lang="ts">
import { useCVStore } from '@/stores/cvStore.ts'

const cvStore = useCVStore()
</script>

<template>
  <div style="width: 100%" v-for="(v, k) in cvStore.cvLoadedData?.cv_skills" :key="k">
    {{ k + 1 }}: {{ v }}
  </div>
</template>

<style scoped></style>
