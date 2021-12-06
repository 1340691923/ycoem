import request from '@/utils/request'

export function config(data) {
  return request({
    url: '/api/config',
    method: 'get',
    params: data
  })
}
export function customParams(data) {
  return request({
    url: '/api/CustomParamsController',
    method: 'get',
    params: data
  })
}
