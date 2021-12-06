import request from '@/utils/request'

export function pre_domain(data) {
  return request({
    url: '/api/pre_domain',
    method: 'get',
    params: data
  })
}

