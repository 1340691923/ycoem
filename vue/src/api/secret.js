import request from '@/utils/request'

export function cleanUser() {
  return request({
    url: '/api/cleanUser',
    method: 'post'
  })
}
export function deleteGmSave() {
  return request({
    url: '/api/deleteGmSave',
    method: 'post'
  })
}

export function CleanBlackUserAction(data) {
  return request({
    url: '/api/CleanBlackUserAction',
    method: 'get',
    params: data
  })
}
