import { createRouter, createWebHashHistory, type Router } from 'vue-router'
import { type App } from 'vue'
import CVKawaRoutes from '@/routes/routes.ts'
import beforeRouter from '@/routes/guard.ts'

const router: Router = createRouter({
  history: createWebHashHistory(),
  routes: [...CVKawaRoutes]
})

export default function setupRouter(app: App) {
  beforeRouter(router)
  app.use(router)
}
