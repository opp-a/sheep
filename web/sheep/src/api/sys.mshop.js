import request from '@/plugins/axios'

export function GetMShops(data) {
  return request({
    url: '/mshops',
    method: 'get',
    data
  })
}
export function DeleteMShops(data) {
  return request({
    url: '/mshops',
    method: 'delete',
    data
  })
}
export function AddUpdateMShops(data) {
  return request({
    url: '/mshops',
    method: 'post',
    data
  })
}
