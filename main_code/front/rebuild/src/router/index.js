import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
// import Affairs from '../views/Affairs.vue'
import Login from '../views/Login.vue'
import Regist from '../views/Regist.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'login',
    component: Login
  },
  {
    path: '/regist',
    name: 'regist',
    component: Regist,
  },
  {
    path: '/home',
    name: 'home',
    component: Home,
  }
]

const router = new VueRouter({
  routes
})

export default router
