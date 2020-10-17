import request from '@/plugins/axios'

export function GetCulturePictures(data) {
  return request({
    url: '/culturepictures',
    method: 'get',
    data
  })
}
