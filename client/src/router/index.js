import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Modal from '../views/Modal.vue'
import FeedbackModal from '@/views/FeedbackModal.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/presensi',
      name: 'presensi',
      component: Modal
    },
    {
      path: '/feedback',
      name: 'feedback',
      component: FeedbackModal
    }
  ]
})

export default router
