import request from '@/plugins/axios'

export function GetMyOrders(data) {
  return request({
    url: '/historyorder/list',
    method: 'post',
    data
  })
}
