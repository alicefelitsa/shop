import Vue from 'vue'
import VueRouter from 'vue-router'
import {checkAccess, accessState} from '@/utils/access'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },
  {
    path: '/products',
    name: 'Products',
    component: () => import('../views/Products.vue')
  },
  {
    path: '/products/:id',
    name: 'ProductDetail',
    component: () => import('../views/ProductDetail.vue')
  },
  {
    path: '/contact',
    name: 'Contact',
    component: () => import('../views/Contact.vue')
  },
  {
    path: '/cart',
    name: 'Cart',
    component: () => import('../views/Cart.vue')
  },
  {
    path: '/notice',
    name: 'Notice',
    component: () => import('../views/Notice.vue')
  },
  {
    path: '*',
    name: 'NotFound',
    component: () => import('../views/NotFound.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) return savedPosition
    return { x: 0, y: 0 }
  }
})

// 访问方式拦截：与后台配置的 access_mode 不匹配时按展示类型原地空白或跳转提示页
router.beforeEach(async (to, from, next) => {
  if (to.name === 'Notice') {
    accessState.ready = true
    return next()
  }
  const {allowed, mode, displayType} = await checkAccess()
  if (!allowed) {
    if (displayType === 'blank') {
      // 不改变 URL，仅置空白状态由 App.vue 渲染空页
      accessState.blank = true
      accessState.ready = true
      return next()
    }
    accessState.ready = true
    return next({name: 'Notice', query: {mode}})
  }
  accessState.ready = true
  next()
})

export default router
