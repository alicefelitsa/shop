import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/login'),
        meta: {title: '用户登录'}
    },
    {
        path: '/',
        component: () => import('@/layout/index'),
        redirect: '/product',
        children: [
            {
                path: 'message',
                name: 'message',
                component: () => import('@/views/message/index'),
                meta: {title: '留言'}
            },
            {
                path: 'product',
                name: 'product',
                component: () => import('@/views/product/index'),
                meta: {title: '产品列表'}
            },
            {
                path: 'category',
                name: 'category',
                component: () => import('@/views/category/index'),
                meta: {title: '分类'}
            },
            {
                path: 'contact',
                name: 'contact',
                component: () => import('@/views/contact/index'),
                meta: {title: '联系方式'}
            },
            {
                path: 'setting',
                name: 'setting',
                component: () => import('@/views/setting'),
                meta: {title: '设置'}
            },
        ]
    },
    // 通配符路由放在最后
    {
        path: '*',
        name: 'NotFound',
        component: () => import('@/views/NotFound'),
    }
]
// 1. 先创建路由实例
const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

// 2. 全局前置守卫
router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }

    //指定的页面不需要授权验证
    if (to.path === '/login') {
        if (localStorage.getItem("token")) {
            next({path: '/'})
        }
        next()
        return
    }
    if (localStorage.getItem("token")) {
        next()
    } else {
        next({path: '/login'})

    }
})

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}

export default router
