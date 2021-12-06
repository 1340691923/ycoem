import request from '@/utils/request'

export function list(data) {
  return request({
    url: '/api/Cdk',
    method: 'get',
    params: data
  })
}

export function download(data) {
  return request({
    url: '/api/Cdk/download',
    method: 'get',
    params: data,
    responseType: 'blob',
  })
}
