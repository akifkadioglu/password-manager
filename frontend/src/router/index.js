import { createRouter, createWebHashHistory } from 'vue-router'
import LandingPage from '../pages/LandingPage.vue'

const routes = [
  {
    path: '/',
    name: 'Landing',
    component: LandingPage
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
