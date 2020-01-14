// Vue
import Vue from 'vue'
import App from './App'

// 菜单和路由设置
import router from './router'
// store
import store from '@/store/index'
// 核心插件
import sheep from '@/plugins/sheep'
import '@/plugins/js/element.js'

// 核心插件
Vue.use(sheep)

Vue.config.productionTip = false
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
