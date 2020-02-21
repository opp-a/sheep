import request from '@/plugins/axios'

export function GetMShops(data) {
  return request({
    url: '/shop/list',
    method: 'get',
    data
  })
}
export function DeleteMShops(data) {
  return request({
    url: '/shop/delete',
    method: 'delete',
    data
  })
}
export function AddUpdateMShops(data) {
  return request({
    url: '/shop/ucreate',
    method: 'post',
    data
  })
}
