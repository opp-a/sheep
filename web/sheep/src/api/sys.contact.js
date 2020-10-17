import request from '@/plugins/axios'

export function NewContact(data) {
  return request({
    url: '/contact',
    method: 'post',
    data
  })
}
