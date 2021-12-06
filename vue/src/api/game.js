import request from '@/utils/request'

export function getGameList(data) {
  return request({
    url: '/api/game',
    method: 'get',
    params: data
  })
}

export function getGameLoad(data) {
  console.log('data', data)
  return request({
    url: '/api/GameSaveGet',
    method: 'get',
    params: data
  })
}

export function getGameSave(data) {
  return request({
    url: '/api/GameSave',
    method: 'post',
    data
  })
}
