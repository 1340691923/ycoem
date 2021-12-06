import request from '@/utils/request'

export function searchShareRule(data) {
  return request({
    url: '/api/searchShareRule',
    method: 'get',
    params: data
  })
}

export function updateShareRule(data) {
  return request({
    url: '/api/updateShareRule',
    method: 'get',
    params: data
  })
}

export function delShareRule(data) {
  return request({
    url: '/api/delShareRule',
    method: 'get',
    params: data
  })
}
