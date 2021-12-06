import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/api/ClickListAction',
    method: 'get',
    params: data
  })
}

export function tempList(data) {
  return request({
    url: '/api/tempList',
    method: 'get',
    params: data
  })
}
export function SaveTempMsgAction(data) {
  return request({
    url: '/api/SaveTempMsgAction',
    method: 'get',
    params: data
  })
}
export function SubscribeMsgSendAction(data) {
  return request({
    url: '/api/SubscribeMsgSendAction',
    method: 'get',
    params: data
  })
}
export function TempMsgDelAction(data) {
  return request({
    url: '/api/TempMsgDelAction',
    method: 'get',
    params: data
  })
}
