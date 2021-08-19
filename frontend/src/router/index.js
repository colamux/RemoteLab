import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/home/index.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    children: [
      {
        path: '/equipment',
        name: 'equipment',
        meta: { title: '设备' },
        component: () => import('../views/home/equipment/index.vue')
      },
      {
        path: '/equipment/list',
        name: 'list',
        meta: { title: '列表展示' },
        component: () => import('../views/list.vue')
      }
    ]
    // children: [
    //   {
    //     path: '/index',
    //     // component: Home
    //   },
    //   
    //   },
    //   {
    //     path: '/wms',
    //     component: () => import('../views/wms.vue')
    //   }
    // ]
  },
  {
    path: '/login',
    name: 'Login',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/login/index.vue')
  }
]

const router = new VueRouter({
  routes
})
router.beforeEach((to, from, next) => {
  if (!sessionStorage.getItem('username')) {
    if (to.path != '/login') {
      next('/login')
    }
  }
  next()
})
export default router
