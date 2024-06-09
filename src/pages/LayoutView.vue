<script setup lang="ts">
import LayoutHeaderView from '@/components/Layout/LayoutHeaderView.vue'
</script>

<template>
  <n-layout class="base" position="absolute">
    <n-layout-header style="padding: 0 24px; height: 48px" bordered>
      <LayoutHeaderView />
    </n-layout-header>
    <n-layout has-sider position="absolute" style="top: 48px" content-style="padding: 24px;">
      <RouterView />
    </n-layout>
  </n-layout>
</template>

<style scoped>
.base {
  height: 100vh;
  transition: 0.3s;
}
</style>
