import request from '@/utils/request'

export function list(data) {
  return request({
    url: '/api/SeverList',
    method: 'get',
    params: data
  })
}
