import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import Welcome from '../views/Welcome.vue'
import ThankYou from '../views/ThankYou.vue'
import NotFound from '../views/NotFound.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/welcome',
    name: 'Welcome',
    component: Welcome
  },
  {
    path: '/thankyou',
    name: 'ThankYou',
    component: ThankYou
  },
  {
    path: '/404',
    name: 'NotFound',
    component: NotFound
  },
  { path: "/:pathMatch(.*)*", component: NotFound },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
