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
      },
      {
        path: '/equipment/table',
        name: 'table',
        meta: { title: '信息展示' },
        component: () => import('../views/table.vue')
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
  },
  {
    path: '/signup',
    name: 'Signup',
    component: () => import('../views/signup/index.vue')
  }
]

const router = new VueRouter({
  routes
})
router.beforeEach((to, from, next) => {
  let noauth = to.path == '/login' || to.path == '/signup'
  let hastoken = localStorage.getItem('Authorization') != null

  if (noauth && hastoken) { //已登录，访问需要权限的页面
    next('/home')
  } else if (!noauth && !hastoken) { // 未登录，访问无需auth
    next('/login')
  } 
  next()
})
export default router
