import { invoke } from '@tauri-apps/api/tauri'

class SystemApi {
  static async CheckAndInitDB(): Promise<boolean> {
    return await invoke('init_system')
  }

  static async GetVersionList(): Promise<string[]> {
    return await invoke('version_list')
  }
}

export default SystemApi
