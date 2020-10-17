import {NewContact} from '@/api/sys.contact'

export default {
  namespaced: true,
  actions: {
    /**
     * @description 添加留言
     * @param {Object} context
     * @param {Object} payload name {String} 客户名称
     * @param {Object} payload email {String} 客户邮箱
     * @param {Object} payload content {String} 客户问题
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
