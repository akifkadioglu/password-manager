import { createRouter, createWebHashHistory } from 'vue-router'
import LandingPage from '../pages/LandingPage.vue'
import PasswordGenerator from '../pages/PasswordGenerator.vue'
import PasswordsList from '../pages/PasswordsList.vue'

const routes = [
  {
    path: '/',
    name: 'Landing',
    component: LandingPage
  },
  {
    path: '/password-generator',
    name: 'PasswordGenerator',
    component: PasswordGenerator
  },
  {
    path: '/passwords',
    name: 'PasswordsList',
    component: PasswordsList
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
