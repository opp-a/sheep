import request from '@/plugins/axios'

export function GetMyOrders(data) {
  return request({
    url: '/myorders',
    method: 'get',
    data
  })
}
