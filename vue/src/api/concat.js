import request from '@/utils/request'

let api = "/api/concat/"

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
    data
  })
}

export function Delete(data) {
  return request({
    url: api+'delete',
    method: 'post',
    data
  })
}

export function SetLogisticeAction(data) {
  return request({
    url: api+'SetLogisticeAction',
    method: 'post',
    data
  })
}

export function GetLogisticeAction(data) {
  return request({
    url: api+'GetLogisticeAction',
    method: 'post',
    data
  })
}

export function SetOfficeAction(data) {
  return request({
    url: api+'SetOfficeAction',
    method: 'post',
    data
  })
}
export function GetOfficeAction(data) {
  return request({
    url: api+'GetOfficeAction',
    method: 'post',
    data
  })
}

export function SetContactAction(data) {
  return request({
    url: api+'SetContactAction',
    method: 'post',
    data
  })
}

export function GetContactAction(data) {
  return request({
    url: api+'GetContactAction',
    method: 'post',
    data
  })
}

