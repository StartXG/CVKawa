import { RouteRecordRaw } from 'vue-router'

const CVKawaRoutes: Array<RouteRecordRaw> = [
  {
    path: '/print_view',
    name: 'printPreview',
    component: () => import('@/pages/CVPrintPreview.vue')
  },
  {
    path: '/',
    name: 'WelcomeView',
    component: () => import('@/pages/WelcomeView.vue')
  },

  {
    path: '/inner',
    name: 'LayoutView',
    component: () => import('@/pages/LayoutView.vue'),
    redirect: { name: 'CVList' },
    children: [
      {
        path: 'edit',
        name: 'CVKawaEdit',
        component: () => import('@/pages/CVKawaEdit.vue')
      },
      {
        path: 'list',
        name: 'CVList',
        component: () => import('@/pages/CVListView.vue')
      }
    ]
  }
]

export default CVKawaRoutes
