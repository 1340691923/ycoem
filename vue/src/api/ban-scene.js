import request from '@/utils/request'

export function CleanUser(data) {
  return request({
    url: '/api/CleanUser',
    method: 'get',
    params: data
  })
}

export function GetSceneFlagAction(data) {
  return request({
    url: '/api/GetSceneFlagAction',
    method: 'get',
    params: data
  })
}

export function SetSceneFlagAction(data) {
  return request({
    url: '/api/SetSceneFlagAction',
    method: 'get',
    params: data
  })
}

export function searchBanScene(data) {
  return request({
    url: '/api/searchBanScene',
    method: 'get',
    params: data
  })
}
export function setBanScene(data) {
  return request({
    url: '/api/setBanScene',
    method: 'get',
    params: data
  })
}

export function DelBanScene(data) {
  return request({
    url: '/api/DelBanScene',
    method: 'get',
    params: data
  })
}

export function ProvinceAll(data) {
  return request({
    url: '/api/ProvinceAll',
    method: 'get',
    params: data
  })
}

export function AddProvince(data) {
  return request({
    url: '/api/AddProvince',
    method: 'get',
    params: data
  })
}

export function RemoveProvince(data) {
  return request({
    url: '/api/RemoveProvince',
    method: 'get',
    params: data
  })
}
