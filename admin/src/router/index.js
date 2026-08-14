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
        redirect: '/movie/index',
        children: [
            {
                path: '/movie/index',
                name: 'movie',
                component: () => import('@/views/movie/index'),
                meta: {title: '影片列表'}
            },
            // {
            //     path: 'chat',
            //     name: 'chat',
            //     component: () => import('@/views/chat'),
            //     meta: {title: '消息'}
            // },
            // {
            //     path: 'qr',
            //     name: 'qr',
            //     component: () => import('@/views/qr'),
            //     meta: {title: '二维码'}
            // },
            {
                path: 'user',
                name: 'user',
                component: () => import('@/views/user'),
                meta: {title: '用户'}
            },
            // {
            //     path: 'greet',
            //     name: 'greet',
            //     component: () => import('@/views/greet'),
            //     meta: {title: '打招呼'}
            // },
            // {
            //     path: 'answer',
            //     name: 'answer',
            //     component: () => import('@/views/answer'),
            //     meta: {title: '智能回复'}
            // },
            {
                path: 'statistic',
                name: 'statistic',
                component: () => import('@/views/statistic'),
                meta: {title: '统计'}
            },
            {
                path: 'setting',
                name: 'setting',
                component: () => import('@/views/setting'),
                meta: {title: '设置'}
            },
            {
                path: 'logs',
                name: 'logs',
                component: () => import('@/views/logs'),
                meta: {title: '日志'}
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
