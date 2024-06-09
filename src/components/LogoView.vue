<script setup lang="ts">
const props = defineProps<{
  size: 'small' | 'normal'
}>()
</script>

<template>
  <div>
    <n-h2 style="margin: 0" v-if="props.size === 'small'">
      <span style="color: #70c0e8">C</span><span style="color: #e88080">V</span>Kawa
    </n-h2>
    <n-h1 style="margin: 0" v-else>
      <span style="color: #70c0e8">C</span><span style="color: #e88080">V</span>Kawa
    </n-h1>
  </div>
</template>

<style scoped></style>
