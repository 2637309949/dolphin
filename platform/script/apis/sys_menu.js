// Code generated by dol build. DO NOT EDIT.
const axios = require('@/utils/request').default

// add 添加菜单
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/menu/add'
  if ((opt.url || 'post') === 'get') {
    for (var key in data) {
      url += key + '=' + encodeURIComponent(data[key]) + '&'
    }
    return axios({
      url: url,
      method: 'get',
      ...opt
    })
  }
  return axios({
    url: url,
    method: 'post',
    data,
    ...opt
  })
}

// del 删除菜单
module.exports.del = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/menu/del'
  if ((opt.url || 'delete') === 'get') {
    for (var key in data) {
      url += key + '=' + encodeURIComponent(data[key]) + '&'
    }
    return axios({
      url: url,
      method: 'get',
      ...opt
    })
  }
  return axios({
    url: url,
    method: 'delete',
    data,
    ...opt
  })
}

// batchDel 删除菜单
module.exports.batchDel = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/menu/batch_del'
  if ((opt.url || 'delete') === 'get') {
    for (var key in data) {
      url += key + '=' + encodeURIComponent(data[key]) + '&'
    }
    return axios({
      url: url,
      method: 'get',
      ...opt
    })
  }
  return axios({
    url: url,
    method: 'delete',
    data,
    ...opt
  })
}

// update 更新菜单
module.exports.update = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/menu/update'
  if ((opt.url || 'put') === 'get') {
    for (var key in data) {
      url += key + '=' + encodeURIComponent(data[key]) + '&'
    }
    return axios({
      url: url,
      method: 'get',
      ...opt
    })
  }
  return axios({
    url: url,
    method: 'put',
    data,
    ...opt
  })
}

// sidebar 系统菜单
module.exports.sidebar = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/menu/sidebar?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// page 菜单分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/menu/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// tree 菜单树形结构
module.exports.tree = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/menu/tree?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取菜单信息
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/menu/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

