import request from '@/utils/request'

export function create(data) {
  return request({
    url: '/api/seo/create',
    method: 'post',
    data
  })
}

export function get(data) {
  return request({
    url: '/api/seo/get',
    method: 'get',
    params: data
  })
}

export function getphoneQQ(data) {
  return request({
    url: '/api/seo/getphoneQQ',
    method: 'get',
    params: data
  })
}

export function setphoneQQ(data) {
  return request({
    url: '/api/seo/setphoneQQ',
    method: 'post',
    data
  })
}

export function GetWebinfoAction(data) {
  return request({
    url: '/api/seo/GetWebinfoAction',
    method: 'get',
    params: data
  })
}

export function SetWebinfoAction(data) {
  return request({
    url: '/api/seo/SetWebinfoAction',
    method: 'post',
    data
  })
}

export function GetMobileLinkAction(data) {
  return request({
    url: '/api/seo/GetMobileLinkAction',
    method: 'get',
    params: data
  })
}

export function SetMobileLinkAction(data) {
  return request({
    url: '/api/seo/SetMobileLinkAction',
    method: 'post',
    data
  })
}
