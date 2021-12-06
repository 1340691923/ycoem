import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/api/OrderList',
    method: 'get',
    params: data
  })
}

export function issueOrder(data) {
  return request({
    url: '/api/IssueOrder',
    method: 'get',
    params: data
  })
}
