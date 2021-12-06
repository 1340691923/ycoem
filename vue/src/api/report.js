import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/api/EventListAction',
    method: 'post',
    data
  })
}
export function getListExcel(data) {
  return request({
    responseType: 'arraybuffer', // 必填
    url: '/api/EventListAction',
    method: 'post',
    data
  })
}

export function EventManagerAction(data) {
  return request({
    url: '/api/EventManagerAction',
    method: 'post',
    data
  })
}

export function EventManagerDelAction(data) {
  return request({
    url: '/api/EventManagerDelAction',
    method: 'post',
    data
  })
}

export function EventParmasListAction(data) {
  return request({
    url: '/api/EventParmasListAction',
    method: 'get',
    params: data
  })
}


export function EventParmasAnalyseAction(data) {
  return request({
    url: '/api/EventParmasAnalyseAction',
    method: 'post',
    data
  })
}


export function EventManagerByTimeAction(data) {
  return request({
    url: '/api/EventManagerByTimeAction',
    method: 'post',
    data
  })
}
