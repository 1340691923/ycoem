import request from '@/utils/request'


let api = "/api/service/"

export function list(data) {
  return request({
    url: api+'list',
    method: 'post',
    data
  })
}

export function NameOptions(data) {
  return request({
    url: api+'NameOptionsAction',
    method: 'post',
    data
  })
}

export function UpdateAction(data) {
  return request({
    url: api+'UpdateAction',
    method: 'post',
    headers:{
      'Content-Type': 'multipart/form-data'
    },
    data:data
  })
}

export function Delete(data) {
  return request({
    url: api+'delete',
    method: 'post',
    data
  })
}
