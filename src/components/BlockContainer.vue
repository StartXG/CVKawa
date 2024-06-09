<script setup lang="ts">
const props = defineProps<{
  title?: string
  width?: string
  minHeight?: string
  standBlock?: boolean
}>()
</script>

<template>
  <div class="block-container">
    <n-h6 class="block-title" v-if="props.title">{{ props.title }}</n-h6>
    <n-divider v-if="props.standBlock" style="margin-top: 0" />
    <slot name="default" />
  </div>
</template>

<style scoped>
.block-container {
  margin: 1rem 0 2rem 0;
}

.block-title {
  margin: 0 0 1pt 0;
  padding: 0 0 1pt 0;
}
</style>
