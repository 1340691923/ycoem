import Vue from 'vue'

export default new Vuex.Store({
  state: {
    websock: null,
    eventlist: []
  },
  getters: {
    onEvent(state) {
      return function() {
        console.log('11111')
        if (state.eventlist.length > 0) {
          const result = Object.assign({}, state.eventlist[0])
          state.eventlist.splice(0, 1)
          console.log(result.data)
          return result.data
        }
        return null
      }
    }
  },
  mutations: {
    WEBSOCKET_INIT(state, url) {
      state.websock = new WebSocket(url)
      state.websock.onopen = function() {
        console.log('连接成功！')
      }
      state.websock.onmessage = function(callBack) {
        console.log('ws接收！')
        console.log(callBack.data)

        var datas = JSON.parse(callBack.data)

        if (datas.msg == 'pong') {
          console.log('收到服务器心跳包！')
          return
        }
        console.log('datas', datas)
        state.eventlist.push({
          data: datas
        })
      }
      state.websock.οnerrοr = function(e) { // 错误
        console.log('ws错误!')
        console.log(e)
      }
      state.websock.onclose = function(e) { // 关闭
        console.log('ws关闭！')
        console.log(e)
      }
      // 发送心跳包
      setInterval(function() {
        console.log('ws发送心跳！')
        var heart = {
          type: 5,
          method: '',
          jsonData: {}
        }
        state.websock.send('ping')
      }, 30000)
    },
    WEBSOCKET_SEND(state, p) {
      p = JSON.stringify(p)
      console.log('ws发送！', p)
      state.websock.send(p)
    },
    WEBSOCKET_CLOSE(state, p) {
      state.websock.close()
    }
  },
  actions: {
    WEBSOCKET_INIT({ commit }, url) {
      commit('WEBSOCKET_INIT', url)
    },
    WEBSOCKET_SEND({ commit }, p) {
      commit('WEBSOCKET_SEND', p)
    },
    WEBSOCKET_CLOSE({ commit }, p) {
      commit('WEBSOCKET_CLOSE', p)
    }

  }
})

function makeByte(p) {
  var bytesMethod = stringToByte(p.method)

  var bytes = stringToByte(JSON.stringify(p.jsonData))

  var headLenth = 6
  var bodyLenth = bytes.length + bytesMethod.length + 2 + 2
  var buffer = new ArrayBuffer(headLenth + bodyLenth)
  var view = new DataView(buffer)

  var count = 0
  // 包头
  view.setInt16(count, p.type)// 消息类型2
  count += 2
  view.setUint32(count, bodyLenth)// 包体大小4
  count += 4
  // 包体

  view.setInt16(count, 1)// 序列化类型2
  count += 2
  view.setInt16(count, bytesMethod.length)// 方法字符大小2
  count += 2
  for (var i = 0; i < bytesMethod.length; i++) {
    view.setUint8(count, bytesMethod[i])
    count++
  }
  for (var i = 0; i < bytes.length; i++) {
    view.setUint8(count, bytes[i])
    count++
  }
  return view
}

function stringToByte(str) {
  var bytes = new Array()
  var len, c
  len = str.length
  for (var i = 0; i < len; i++) {
    c = str.charCodeAt(i)
    if (c >= 0x010000 && c <= 0x10FFFF) {
      bytes.push(((c >> 18) & 0x07) | 0xF0)
      bytes.push(((c >> 12) & 0x3F) | 0x80)
      bytes.push(((c >> 6) & 0x3F) | 0x80)
      bytes.push((c & 0x3F) | 0x80)
    } else if (c >= 0x000800 && c <= 0x00FFFF) {
      bytes.push(((c >> 12) & 0x0F) | 0xE0)
      bytes.push(((c >> 6) & 0x3F) | 0x80)
      bytes.push((c & 0x3F) | 0x80)
    } else if (c >= 0x000080 && c <= 0x0007FF) {
      bytes.push(((c >> 6) & 0x1F) | 0xC0)
      bytes.push((c & 0x3F) | 0x80)
    } else {
      bytes.push(c & 0xFF)
    }
  }
  return bytes
}

function byteToString(arr) {
  if (typeof arr === 'string') {
    return arr
  }
  var str = ''
  var _arr = arr
  for (var i = 0; i < _arr.length; i++) {
    var one = _arr[i].toString(2)
    var v = one.match(/^1+?(?=0)/)
    if (v && one.length == 8) {
      var bytesLength = v[0].length
      var store = _arr[i].toString(2).slice(7 - bytesLength)
      for (var st = 1; st < bytesLength; st++) {
        store += _arr[st + i].toString(2).slice(2)
      }
      str += String.fromCharCode(parseInt(store, 2))
      i += bytesLength - 1
    } else {
      str += String.fromCharCode(_arr[i])
    }
  }
  return str
}
