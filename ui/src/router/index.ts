import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import Welcome from '../views/Welcome.vue'
import ThankYou from '../views/ThankYou.vue'
import Nothing from '../views/Nothing.vue'
import NotFound from '../views/NotFound.vue'
import Graph002 from '../views/Graph002.vue'

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
    path: '/-graph',
    name: 'Graph002',
    component: Graph002
  },
  {
    path: '/nothing',
    name: 'Nothing',
    component: Nothing
  },
  {
    path: '/404',
    name: 'NotFound',
    component: NotFound
  },
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
