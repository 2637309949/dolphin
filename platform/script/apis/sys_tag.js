// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加标签
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/add' }, opt).url
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

// batchAdd 添加标签
module.exports.batchAdd = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/batch_add' }, opt).url
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

// del 删除标签
module.exports.del = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/del' }, opt).url
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

// batchDel 删除标签
module.exports.batchDel = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/batch_del' }, opt).url
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

// update 更新标签
module.exports.update = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/update' }, opt).url
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

// batchUpdate 更新标签
module.exports.batchUpdate = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/batch_update' }, opt).url
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

// page 标签分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/page?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取标签信息
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tag/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

