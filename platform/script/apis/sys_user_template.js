// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加用户模板
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/add' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
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

// batchAdd 添加用户模板
module.exports.batchAdd = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/batch_add' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
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

// del 删除用户模板
module.exports.del = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/del' }, opt).url
  if ((opt.method || 'delete') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
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

// batchDel 删除用户模板
module.exports.batchDel = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/batch_del' }, opt).url
  if ((opt.method || 'delete') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
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

// update 更新用户模板
module.exports.update = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/update' }, opt).url
  if ((opt.method || 'put') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
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

// batchUpdate 更新用户模板
module.exports.batchUpdate = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/batch_update' }, opt).url
  if ((opt.method || 'put') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
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

// page 用户模板分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/page?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取用户模板信息
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/template/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

