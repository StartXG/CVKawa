import { type Router } from 'vue-router'
import { useCVStore } from '@/stores/cvStore.ts'

const cvStore = useCVStore()
export default function beforeRouter(router: Router) {
  router.beforeEach((_to, _from, next) => {
    cvStore.loadAllVersion().then(() => {
      console.log('load all version success')
    })
    next()
  })
}
