// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request').default

// add 添加组织
module.exports.add = (data) => {
  const url = '/api/sys/org/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// del 删除组织
module.exports.del = (data) => {
  const url = '/api/sys/org/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// batchDel 删除组织
module.exports.batchDel = (data) => {
  const url = '/api/sys/org/batch_del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// update 更新组织
module.exports.update = (data) => {
  const url = '/api/sys/org/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// page 组织分页查询
module.exports.page = (data) => {
  let url = '/api/sys/org/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// tree 菜单树形结构
module.exports.tree = (data) => {
  let url = '/api/sys/org/tree?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取组织信息
module.exports.get = (data) => {
  let url = '/api/sys/org/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}
