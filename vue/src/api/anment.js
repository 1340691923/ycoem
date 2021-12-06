import request from '@/utils/request'

export function Notice(data) {
  return request({
    url: '/api/NoticeList',
    method: 'post',
    data
  })
}
