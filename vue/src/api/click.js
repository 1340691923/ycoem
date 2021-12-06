import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/api/ClickListAction',
    method: 'get',
    params: data
  })
}

export function getChanList() {
  return request({
    url: 'api/GetClickChanListAction',
    method: 'post'
  })
}

export function ClickList2Action(data) {
  return request({
    url: '/api/ClickList2Action',
    method: 'get',
    params: data
  })
}
export function ClickList2ExcelAction(data) {
  return request({
    responseType: 'arraybuffer', // 必填
    url: '/api/ClickList2Action',
    method: 'get',
    params: data
  })
}
export function GetCountList(data) {
  return request({
    url: 'api/GetCountList',
    method: 'get',
    params: data
  })
}

export function GetCountListExcel(data) {
  return request({
    responseType: 'arraybuffer', // 必填
    url: '/api/GetCountList',
    method: 'get',
    params: data
  })
}

