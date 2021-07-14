// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加客户端
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/add' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// batchAdd 添加客户端
module.exports.batchAdd = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/batch_add' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// del 删除客户端
module.exports.del = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/del' }, opt).url
  if ((opt.method || 'delete') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// batchDel 删除客户端
module.exports.batchDel = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/batch_del' }, opt).url
  if ((opt.method || 'delete') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// update 更新客户端
module.exports.update = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/update' }, opt).url
  if ((opt.method || 'put') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// batchUpdate 更新客户端
module.exports.batchUpdate = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/batch_update' }, opt).url
  if ((opt.method || 'put') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// page 客户端分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/page?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取客户端信息
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/client/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

