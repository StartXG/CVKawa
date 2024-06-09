<script setup lang="ts">
import { getLayoutOptions, LayoutOptions, PageOption } from '@/components/DocLayout/LayoutOptions'
import { computed, ComputedRef, onMounted, ref } from 'vue'
import BaseInfo from '@/components/CVBaseInfo/CVBaseInfo.vue'
import BlockContainer from '@/components/BlockContainer.vue'
import CVEduInfo from '@/components/CVEduInfo/CVEduInfo.vue'
import CVResume from '@/components/CVResume/CVResume.vue'
import CVSkills from '@/components/CVSkills/CVSkills.vue'
import { useCVStore } from '@/stores/cvStore'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const props = withDefaults(defineProps<{ mode?: PageOption }>(), { mode: 'A4' })

const pageConfig: ComputedRef<Readonly<LayoutOptions>> = computed(() => {
  return getLayoutOptions(props.mode)
})

const cvStore = useCVStore()
const route = useRoute()

const printRef = ref<HTMLElement | null>(null)

function printDiv() {
  console.log(pageConfig)
}

onMounted(() => {
  let v = route.query.version as string
  cvStore.loadCV(v).then(() => {
    console.log(t('preview.cvLoaded'))
    setTimeout(() => {
      printDiv()
    }, 1000)
  })
})
</script>

<template>
  <n-empty v-if="cvStore.cvLoadedData === null" size="large" :description="t('preview.empty')" />
  <div
    ref="printRef"
    v-else
    :style="{
      width: `${pageConfig.width}px`,
      height: `${pageConfig.height}px`,
      boxSizing: 'border-box',
      padding: `${pageConfig.padding}px`,
      fontSize: `${pageConfig.fontSize}px`,
      marginBottom: `${pageConfig.blockGap}px`
    }"
    class="pc-mb"
  >
    <BlockContainer :stand-block="false">
      <BaseInfo />
    </BlockContainer>
    <BlockContainer :title="t('editPage.eduInfo.title')" :stand-block="true">
      <CVEduInfo />
    </BlockContainer>
    <BlockContainer :title="t('editPage.resumeInfo.title')" :stand-block="true">
      <CVResume />
    </BlockContainer>
    <BlockContainer :title="t('editPage.skills.title')" :stand-block="true">
      <CVSkills />
    </BlockContainer>
  </div>
</template>

<style scoped>
.pc-mb {
  transition: all 0.3s;
  width: 100%;
  height: 100%;
}
</style>
