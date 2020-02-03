import request from '@/plugins/axios'
// import axios from 'axios'

export function GetShops(data) {
  return new Promise((resolve, reject) => {
    const pageNumber = data.pageNumber
    const shopsArray = []
    // console.log(data.page)
    if (data.page > 3) {
      resolve([])
    }
    for (let i = 0; i < pageNumber; i++) {
      shopsArray.push({
        shopName: 'shop' + data.page + '-' + i,
        image: 'http://',
        price: 999,
        num: 10
      })
    }
    resolve(shopsArray)
  })

  // return request({
  //   url: '/shops',
  //   method: 'get',
  //   data
  // })
}

export function GetCar(data) {
  return request({
    url: '/car',
    method: 'get',
    data
  })
}

export function DoOrder(data) {
  return request({
    url: '/doorder',
    method: 'post',
    data
  })
}

export function SaveOrder(data) {
  return request({
    url: '/saveorder',
    method: 'post',
    data
  })
}
