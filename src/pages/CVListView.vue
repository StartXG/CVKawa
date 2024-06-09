<script setup lang="ts">
import { useCVStore } from '@/stores/cvStore.ts'
import { timeToStr } from '@/tools/main'
import { NCard } from 'naive-ui'
import { Delete24Regular, Edit24Regular, Print24Regular, Eye24Regular } from '@vicons/fluent'
import { invoke } from '@tauri-apps/api'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores/appStore.ts'

const { t } = useI18n()
const cvStore = useCVStore()
const appStore = useAppStore()
const router = useRouter()

async function PrintCV(version: string) {
  await invoke('cv_print', { version: version, language: appStore.lang })
}

function PreviewCV(version: string) {
  invoke('cv_preview', { version: version })
}

function EditCV(version: string, name: string) {
  router.push({ name: 'CVKawaEdit', query: { mode: 'edit', v: version, name: name } })
}
</script>

<template>
  <div class="full-container flex-full-center" v-if="cvStore.cv_versions.length === 0">
    <n-empty size="large" description="Nothing here. Go Create!!!">
      <template #extra>
        <n-button
          tertiary
          size="small"
          type="info"
          @click="router.push({ name: 'CVKawaEdit', query: { mode: 'new' } })"
        >
          {{ t('listPage.addNew') }}
        </n-button>
      </template>
    </n-empty>
  </div>
  <div v-else style="width: 100%">
    <n-grid :x-gap="8" :y-gap="8" cols="1 400:1 800:2 1200:3 1600:4">
      <n-grid-item v-for="(v, k) in cvStore.cv_versions" :key="k">
        <n-card hoverable :title="v.name">
          <template #default>
            <n-descriptions label-placement="left" :column="1">
              <n-descriptions-item :label="t('listPage.version')">{{
                v.version
              }}</n-descriptions-item>
              <!--              TODO:open in next hotfix version-->
              <!--              <n-descriptions-item label="parent_version">-->
              <!--                {{ v.parent_version }}-->
              <!--              </n-descriptions-item>-->
              <n-descriptions-item :label="t('listPage.updatedAt')">
                {{ timeToStr(v.updated_at) }}
              </n-descriptions-item>
            </n-descriptions>
          </template>
          <template #header-extra>
            <n-space>
              <n-button size="small" tertiary type="primary" @click="EditCV(v.version, v.name)">
                <template #icon>
                  <n-icon>
                    <Edit24Regular />
                  </n-icon>
                </template>
              </n-button>
              <n-button size="small" tertiary type="primary" @click="PreviewCV(v.version)">
                <template #icon>
                  <n-icon>
                    <Eye24Regular />
                  </n-icon>
                </template>
              </n-button>
              <n-button size="small" tertiary type="info" @click="PrintCV(v.version)">
                <template #icon>
                  <n-icon>
                    <Print24Regular />
                  </n-icon>
                </template>
              </n-button>
              <n-popconfirm
                @positive-click="cvStore.delCV(v.version || '')"
                @negative-click="() => {}"
              >
                Confirm to delete?
                <template #trigger>
                  <n-button size="small" tertiary type="error">
                    <template #icon>
                      <n-icon>
                        <Delete24Regular />
                      </n-icon>
                    </template>
                  </n-button>
                </template>
              </n-popconfirm>
            </n-space>
          </template>
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>
