import request from '@/plugins/axios'

export function AccountLogin(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}
