import request from '@/plugins/axios'

export function AccountRegister(data) {
  return request({
    url: '/register',
    method: 'post',
    data
  })
}
