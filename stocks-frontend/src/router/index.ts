import { createRouter, createWebHistory } from 'vue-router'
import MainPageView from '@/views/MainPageView.vue'
import StockPageView from '@/views/StockPageView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'mainPage',
      component: MainPageView,
    },
    {
      path: '/stocks/:id',
      name: 'stockDetail',
      component: StockPageView
    }
  ],
})

export default router
