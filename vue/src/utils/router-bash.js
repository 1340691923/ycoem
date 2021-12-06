//自动生成 路由映射文件 脚本
var asyncRoutes = [
  {
    path: '/permission',
    component: 'layout',
    redirect: '/permission/index',
    alwaysShow: true,
    meta: {
      title: '权限',
      icon: 'el-icon-user-solid',
      roles: ['admin', 'editor']
    },
    children: [
      {
        path: 'role',
        component: 'views/permission/role',
        name: 'RolePermission',
        meta: {
          title: '角色管理',
          roles: ['admin'],
          icon: 'el-icon-s-check'
        }
      },
      {
        path: 'user',
        component: 'views/permission/user',
        name: 'user',
        meta: {
          title: '用户管理',
          roles: ['admin'],
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
      icon: 'el-icon-thumb',
    },
    children: [
      {
        path: 'index',
        component: 'views/cases/index',
        name: 'index',
        meta: { title: '案例列表', icon: 'el-icon-s-order', affix: true }
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
      icon: 'el-icon-thumb',
    },
    children: [
      {
        path: 'index',
        component: 'views/contact/index',
        name: 'index',
        meta: { title: '事业部信息管理', icon: 'el-icon-s-order', affix: true }
      },
      {
        path: 'logistice',
        component: 'views/contact/logistice',
        name: 'logistice',
        meta: { title: '物理信息管理', icon: 'el-icon-s-order', affix: true }
      },
      {
        path: 'office',
        component: 'views/contact/office',
        name: 'office',
        meta: { title: '分公司信息管理', icon: 'el-icon-s-order', affix: true }
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
      icon: 'el-icon-thumb',
    },
    children: [
      {
        path: 'index',
        component: 'views/firm/index',
        name: 'index',
        meta: { title: '关于公司', icon: 'el-icon-s-order', affix: true }
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
      icon: 'el-icon-thumb',
    },
    children: [
      {
        path: 'index',
        component: 'views/product/index',
        name: 'index',
        meta: { title: '产品管理', icon: 'el-icon-s-order', affix: true }
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
      icon: 'el-icon-thumb',
    },
    children: [
      {
        path: 'index',
        component: 'views/service/index',
        name: 'index',
        meta: { title: '服务支持管理', icon: 'el-icon-s-order', affix: true }
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
      icon: 'el-icon-thumb',
    },
    children: [
      {
        path: 'index',
        component: 'views/solut/index',
        name: 'index',
        meta: { title: '解决方案管理', icon: 'el-icon-s-order', affix: true }
      }
    ]
  },
  {
    path: '/seo',
    component: 'layout',
    alwaysShow: true,
    redirect: '/seo/index',
    meta: {
      title: 'SEO管理',
      icon: 'el-icon-thumb',
    },
    children: [
      {
        path: 'index',
        component: 'views/seo/index',
        name: 'index',
        meta: { title: 'SEO管理', icon: 'el-icon-s-order', affix: true }
      }
    ]
  },
]

function filterAsyncRoutes(routes) {
  routes.forEach(route => {
    const tmp = { ...route }
    if (tmp.children) {
      tmp.children = filterAsyncRoutes(tmp.children)
    }
    if (tmp.component != "layout") {
      console.log("'"+tmp.component+"':() => import('@/"+tmp.component+"'),")
    }
  })
}

filterAsyncRoutes(asyncRoutes);

