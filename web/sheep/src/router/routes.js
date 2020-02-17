import layoutHeader from '@/layout'
// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.production')

/**
 * 在主框架内显示
 */
const frameIn = [
  {
    path: '/home',
    name: 'layout',
    redirect: {name: 'index'},
    component: layoutHeader,
    children: [
      {
        path: 'index',
        name: 'index',
        meta: {
          title: '首页',
          auth: false
        },
        component: _import('home')
      },
      {
        path: 'warehouse',
        name: 'warehouse',
        meta: {
          title: '货仓',
          auth: false
        },
        component: _import('warehouse')
      },
      {
        path: 'culture',
        name: 'culture',
        meta: {
          title: '文化墙',
          auth: false
        },
        component: _import('culture')
      },
      {
        path: 'myorder',
        name: 'myorder',
        meta: {
          title: '我的订单',
          auth: false
        },
        component: _import('myorder')
      },
      {
        path: 'manageShop',
        name: 'manage-shop',
        meta: {
          title: '商品管理',
          auth: false
        },
        component: _import('mshop')
      },
      {
        path: 'manageCulture',
        name: 'manage-culture',
        meta: {
          title: '文化管理',
          auth: false
        },
        component: _import('mculture')
      }
    ]
  }
]

/**
 * 在主框架之外显示
 */
const frameIndex = [
  // 登录
  {
    path: '/',
    name: 'home',
    redirect: '/home/index'
  }
]

/**
 * 在主框架之外显示
 */
const frameOut = [
  // 登录
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '登录',
      auth: false
    },
    component: _import('login')
  }
]

/**
 * 错误页面
 */
const errorPage = [
  {
    path: '*',
    name: '404',
    meta: {
      title: '404',
      auth: false
    },
    component: _import('error')
  }
]

// 导出需要显示菜单的
export const frameInRoutes = frameIn

// 重新组织后导出
export default [...frameIn, ...frameIndex, ...frameOut, ...errorPage]
