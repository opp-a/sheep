import request from '@/plugins/axios'
// import axios from 'axios'

export function GetShops(data) {
  return request({
    url: '/shop/list',
    method: 'post',
    data
  })
}

export function GetCar(data) {
  return request({
    url: '/car/get',
    method: 'get',
    data
  })
}

export function DoOrder(data) {
  return request({
    url: '/car/pay',
    method: 'post',
    data
  })
}

export function SaveOrder(data) {
  return request({
    url: '/car/modify',
    method: 'post',
    data
  })
}
