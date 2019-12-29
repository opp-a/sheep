import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'index',
        component: () => import(/* webpackChunkName: "about" */ '../views/Home.vue')
    },
    {
        path: '/home',
        name: 'home',
        redirect: '/',
        component: () => import(/* webpackChunkName: "about" */ '../views/Home.vue')
    },
    {
        path: '/login',
        name: 'login',
        component: () => import(/* webpackChunkName: "about" */ '../views/login/Login.vue')
    },
    {
        path: '/warehouse',
        name: 'warehouse',
        component: () => import(/* webpackChunkName: "about" */ '../views/warehouse/wareHouse.vue')
        // children: [
        //     {
        //         path: '1',
        //         component: () => import(/* webpackChunkName: "about" */ '../views/login/Login.vue')
        //     },
        //     {
        //         path: '2',
        //         component: () => import(/* webpackChunkName: "about" */ '../views/login/Login.vue')
        //     }
        // ]
    },
    {
        path: '/culture',
        name: 'culture',
        redirect: '/login',
        component: () => import(/* webpackChunkName: "about" */ '../views/login/Login.vue')
    },
    {
        path: '/self',
        name: 'self',
        redirect: '/login',
        component: () => import(/* webpackChunkName: "about" */ '../views/login/Login.vue')
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    scrollBehavior (to, from, savedPosition) {
        if (to.hash) {
            return { selector: to.hash }
        } else if (savedPosition) {
            return savedPosition
        } else {
            return { x: 0, y: 0 }
        }
    },
    routes
})

export default router
