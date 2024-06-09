<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
import BaseInfoEdit from '@/components/CVKawaEdit/BaseInfoEdit.vue'
import { onMounted, ref } from 'vue'
import EduInfoEdit from '@/components/CVKawaEdit/EduInfoEdit.vue'
import ResumeEdit from '@/components/CVKawaEdit/ResumeEdit.vue'
import SkillsEdit from '@/components/CVKawaEdit/SkillsEdit.vue'
import { useCVStore } from '@/stores/cvStore.ts'
import { useNotification } from 'naive-ui'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const cvStore = useCVStore()
const cvName = ref<String>('')
const notification = useNotification()
const router = useRouter()

async function saveCV() {
  let result = await cvStore.saveCV(cvName)
  if (result) {
    notification.success({
      content: `${t('execResult.editState.save')}${t('execResult.success')}`,
      duration: 2500,
      keepAliveOnHover: true
    })
    await router.push({ name: 'CVList' })
  } else {
    notification.error({
      content: `${t('execResult.editState.save')}${t('execResult.error')}`,
      duration: 2500,
      keepAliveOnHover: true
    })
  }
}

onMounted(() => {
  if (route.query.mode === 'new') {
    cvStore.newCV()
  } else {
    let ver = route.query.v
    if (!ver) {
      notification.error({
        content: `${t('execResult.editState.variable')}${t('execResult.error')}`,
        duration: 2500,
        keepAliveOnHover: true
      })
      router.push({ name: 'CVList' })
    } else {
      cvStore.loadCVEdit(ver)
    }
  }
  if (route.query.name) {
    cvName.value = route.query.name as string
  }
})
</script>

<template>
  <div class="full-container">
    <n-form inline size="small" label-placement="left" style="height: 48px">
      <n-form-item :label="t('editPage.cvName')">
        <n-input v-model:value="cvName" clearable />
      </n-form-item>
      <n-form-item>
        <n-button type="primary" @click="saveCV()">{{ t('editPage.cvSave') }}</n-button>
        <!--        <n-button v-if="route.query.name" type="primary" @click="saveCV()">As New</n-button>-->
      </n-form-item>
    </n-form>
    <n-tabs
      default-value="Base Info"
      type="segment"
      animated
      placement="left"
      class="full-container"
      pane-wrapper-class="full-container"
      pane-class="full-container"
      style="height: calc(100% - 48px)"
    >
      <n-tab-pane :tab="t('editPage.baseInfo.title')" name="Base Info">
        <BaseInfoEdit />
      </n-tab-pane>
      <n-tab-pane name="Education" :tab="t('editPage.eduInfo.title')">
        <EduInfoEdit />
      </n-tab-pane>
      <n-tab-pane name="Resume" :tab="t('editPage.resumeInfo.title')">
        <ResumeEdit />
      </n-tab-pane>
      <n-tab-pane name="Skills" :tab="t('editPage.skills.title')">
        <SkillsEdit />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<style scoped></style>
