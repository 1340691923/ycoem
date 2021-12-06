/* Layout */
import Layout from '@/layout'

// 动态路由列表
export const asyncRoutes = [
  {
    path: '/permission',
    component: 'layout',
    redirect: '/permission/index',
    alwaysShow: true,
    meta: {
      title: '权限',
      icon: 'el-icon-user-solid',

    },
    children: [
      {
        path: 'role',
        component: 'views/permission/role',
        name: 'RolePermission',
        meta: {
          title: '角色管理',

          icon: 'el-icon-s-check'
        }
      },
      {
        path: 'user',
        component: 'views/permission/user',
        name: 'user',
        meta: {
          title: '用户管理',

          icon: 'el-icon-user'
        }
      }
    ]
  },
  {
    path: '/cases',
    component: 'layout',
    alwaysShow: true,
    redirect: '/cases/index',
    meta: {
      title: '案例管理',
      icon: 'el-icon-trophy',
    },
    children: [
      {
        path: 'index',
        component: 'views/cases/index',
        name: 'index',
        meta: { title: '案例列表', icon: 'el-icon-medal-1'}
      }
    ]
  },
  {
    path: '/contact',
    component: 'layout',
    alwaysShow: true,
    redirect: '/contact/index',
    meta: {
      title: '联系我们',
      icon: 'el-icon-phone-outline\n',
    },
    children: [
      {
        path: 'index',
        component: 'views/contact/index',
        name: 'index',
        meta: { title: '事业部信息管理', icon: 'el-icon-map-location'}
      },
      {
        path: 'logistice',
        component: 'views/contact/logistice',
        name: 'logistice',
        meta: { title: '物理信息管理', icon: 'el-icon-discover'}
      },
      {
        path: 'office',
        component: 'views/contact/office',
        name: 'office',
        meta: { title: '分公司信息管理', icon: 'el-icon-location-outline'}
      },
      {
        path: 'info',
        component: 'views/seo/info',
        name: 'info',
        meta: { title: '联系方式', icon: 'el-icon-mobile-phone'}
      },
      {
        path: 'show',
        component: 'views/contact/show',
        name: 'show',
        meta: { title: '首页展示信息', icon: 'el-icon-price-tag'}
      }
    ]
  },
  {
    path: '/firm',
    component: 'layout',
    alwaysShow: true,
    redirect: '/firm/index',
    meta: {
      title: '关于公司',
      icon: 'el-icon-house',
    },
    children: [
      {
        path: 'index',
        component: 'views/firm/index',
        name: 'index',
        meta: { title: '关于公司', icon: 'el-icon-house'}
      },
      {
        path: 'conpany',
        component: 'views/firm/conpany',
        name: 'conpany',
        meta: { title: '公司介绍', icon: 'el-icon-chat-line-round'}
      }
    ]
  },
  {
    path: '/product',
    component: 'layout',
    alwaysShow: true,
    redirect: '/product/index',
    meta: {
      title: '产品管理',
      icon: 'el-icon-goods',
    },
    children: [
      {
        path: 'index',
        component: 'views/product/index',
        name: 'index',
        meta: { title: '产品管理', icon: 'el-icon-goods'}
      }
    ]
  },
  {
    path: '/service',
    component: 'layout',
    alwaysShow: true,
    redirect: '/service/index',
    meta: {
      title: '服务支持管理',
      icon: 'el-icon-service',
    },
    children: [
      {
        path: 'index',
        component: 'views/service/index',
        name: 'index',
        meta: { title: '服务支持管理', icon: 'el-icon-service'}
      },
      {
        path: 'download',
        component: 'views/service/download',
        name: 'download',
        meta: { title: '下载管理', icon: 'el-icon-download'}
      },
      {
        path: 'answer',
        component: 'views/service/answer',
        name: 'answer',
        meta: { title: '解答管理', icon: 'el-icon-microphone'}
      }
    ]
  },
  {
    path: '/solut',
    component: 'layout',
    alwaysShow: true,
    redirect: '/solut/index',
    meta: {
      title: '解决方案管理',
      icon: 'el-icon-magic-stick',
    },
    children: [
      {
        path: 'index',
        component: 'views/solut/index',
        name: 'index',
        meta: { title: '解决方案管理', icon: 'el-icon-magic-stick'}
      }
    ]
  },
  {
    path: '/seo',
    component: 'layout',
    alwaysShow: true,
    redirect: '/seo/index',
    meta: {
      title: 'Seo管理',
      icon: 'el-icon-search',
    },
    children: [
      {
        path: 'index',
        component: 'views/seo/index',
        name: 'index',
        meta: { title: 'SEO', icon: 'el-icon-search'}
      },
      {
        path: 'webinfo',
        component: 'views/seo/webinfo',
        name: 'webinfo',
        meta: { title: '网站备案信息', icon: 'el-icon-collection'}
      },
    ]
  },
  {
    path: '/lunbo',
    component: 'layout',
    redirect: '/lunbo/index',
    meta: {
      title: '模块展示图管理',
      icon: 'el-icon-picture',
    },
    children: [
      {
        path: 'index',
        component: 'views/lunbo/index',
        name: 'index',
        meta: { title: '模块展示图管理', icon: 'el-icon-picture'}
      },
    ]
  }
]

// 路由组件 映射 map
export const RoutesComponentmaps = {
  'layout': Layout,
  'views/dashboard/index': () => import('@/views/dashboard/index'),                     //主页
  'views/permission/role': () => import('@/views/permission/role'),                     //角色管理
  'views/permission/user': () => import('@/views/permission/user'),                     //用户管理
  'views/cases/index':() => import('@/views/cases/index'),
  'views/contact/index':() => import('@/views/contact/index'),
  'views/contact/logistice':() => import('@/views/contact/logistice'),
  'views/contact/office':() => import('@/views/contact/office'),
  'views/contact/show':() => import('@/views/contact/show'),
  'views/firm/index':() => import('@/views/firm/index'),
  'views/product/index':() => import('@/views/product/index'),
  'views/service/index':() => import('@/views/service/index'),
  'views/solut/index':() => import('@/views/solut/index'),
  'views/seo/index':() => import('@/views/seo/index'),
  'views/seo/info':() => import('@/views/seo/info'),
  'views/seo/webinfo':() => import('@/views/seo/webinfo'),
  'views/lunbo/index':() => import('@/views/lunbo/index'),
  'views/service/download':() => import('@/views/service/download'),
  'views/service/answer':() => import('@/views/service/answer'),
  'views/firm/conpany':() => import('@/views/firm/conpany'),
}



