import axios from 'axios'

let base = ''

/**
 * 获取用户列表
 */

export const getUserList = data => fetch('/v1/users/list', data)

/**
 * 获取用户数量
 */

export const getUserCount = data => {
  return axios.get(`${base}/user/listpage`, { params: params })
}
