import { createPinia, type Pinia } from 'pinia'
import { type App } from 'vue'

const store: Pinia = createPinia()
export default function setupStore(app: App) {
  app.use(store)
}

export { store }
