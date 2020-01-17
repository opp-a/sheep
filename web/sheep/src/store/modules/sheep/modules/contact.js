// import {Message, MessageBox} from 'element-ui'
// import util from '@/libs/util.js'
// import router from '@/router'
import {NewContact} from '@/api/sys.contact'

export default {
  namespaced: true,
  actions: {
    /**
     * @description 登录
     * @param {Object} context
     * @param {Object} payload username {String} 用户账号
     * @param {Object} payload password {String} 密码
     * @param {Object} payload route {Object} 登录成功后定向的路由对象 任何 vue-router 支持的格式
     */
    newContact({dispatch}, {name = '', email = '', service = '留言', content = ''} = {}) {
      return new Promise((resolve, reject) => {
        // 开始请求登录接口
        NewContact({name, email, service, content})
          .then(async res => {
            this.$notify({
              title: '成功',
              message: '我们已收到您的问题，我们将尽快给予你回馈！！！',
              type: 'success'
            })
            // 结束
            resolve()
          })
          .catch(err => {
            reject(err)
          })
      })
    }
  }
}
