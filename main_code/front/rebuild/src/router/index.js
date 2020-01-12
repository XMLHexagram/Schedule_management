import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Affairs from '../views/Affairs.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/affairs',
    name: 'affairs',
    component: Affairs
  }
]

const router = new VueRouter({
  routes
})

export default router
