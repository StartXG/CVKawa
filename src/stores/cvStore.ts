import { defineStore } from 'pinia'
import { reactive, Ref, ref, UnwrapRef } from 'vue'
import { store } from '@/plugins/store.ts'
import { invoke } from '@tauri-apps/api/tauri'
import { CVDataSchema } from '@/common/cvSchemas.ts'
import { LocationQueryValue } from 'vue-router'

interface versionSchema {
  id?: string
  name: string
  version: string
  parent_version?: string
  created_at?: number
  updated_at?: number
  deleted_at?: number
}

export const CVStore = defineStore('cvStore', () => {
  const cv_versions = ref<versionSchema[]>([])
  const cvLoadedData = ref<CVDataSchema | null>(null)
  const cv = reactive<CVDataSchema>({
    cv_base_info: {
      name: null,
      birthday: null,
      mobile: null,
      email: null,
      photo_path: null
    },
    cv_edu_info: [],
    cv_resume: [],
    cv_skills: []
  })

  async function loadCV(version?: string | null) {
    cvLoadedData.value = await invoke('load_cv', { version: version })
  }

  async function loadAllVersion() {
    cv_versions.value = await invoke('version_list')
  }

  async function saveCV(cv_name: Ref<UnwrapRef<String>>): Promise<boolean> {
    return await invoke('save_cv', { cv: cv, name: cv_name.value, parent_version: '123' })
  }

  async function loadCVEdit(version: string | LocationQueryValue[]) {
    const result: CVDataSchema = await invoke('load_cv', { version: version })
    cv.cv_base_info = result.cv_base_info
    cv.cv_edu_info = result.cv_edu_info
    cv.cv_resume = result.cv_resume
    cv.cv_skills = result.cv_skills
  }

  async function delCV(version: string) {
    await invoke('soft_del', { version: version })
      .then(() => {
        cv_versions.value.forEach((v, i) => {
          if (v.version === version) {
            cv_versions.value.splice(i, 1)
          }
        })
      })
      .catch((e) => {
        console.error(e)
      })
  }

  async function newCV() {
    cv.cv_base_info = {
      name: null,
      birthday: null,
      mobile: null,
      email: null,
      photo_path: null
    }
    cv.cv_edu_info = []
    cv.cv_resume = []
    cv.cv_skills = []
  }

  return {
    loadCV,
    cvLoadedData,
    loadAllVersion,
    cv_versions,
    cv,
    saveCV,
    delCV,
    newCV,
    loadCVEdit
  }
})

export function useCVStore() {
  return CVStore(store)
}
