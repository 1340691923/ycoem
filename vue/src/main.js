import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets



import Element from 'element-ui'
import './styles/element-variables.scss'
import VDistpicker from 'v-distpicker'

Vue.component('v-distpicker', VDistpicker)
import '@/styles/index.scss' // global css

import VueParticles from 'vue-particles'
Vue.use(VueParticles)

import App from './App'
import store from './store'
import router from './router'

import './icons' // icon
import './permission' // permission control
import './utils/error-log' // error log

import * as filters from './filters' // global filters
import clipboard from 'clipboard' // global filters

Vue.prototype.clipboard = clipboard


Vue.use(Element, {
  size: Cookies.get('size') || 'medium' // set element-ui default size
})

// 注册全局过滤器
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
