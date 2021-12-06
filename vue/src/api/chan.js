import request from '@/utils/request'

export function getChanList() {
  return request({
    url: '/api/getChanList',
    method: 'post'
  })
}

export function Channel(data) {
  return request({
    url: '/api/channel',
    method: 'get',
    params: data
  })
}

