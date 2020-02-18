import request from '@/plugins/axios'

export function AccountLogin(data) {
  // return {name: 'admin', token: 'mytoken'}
  return request({
    url: '/api/login',
    method: 'post',
    data
  })
}
