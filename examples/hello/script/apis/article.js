// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add Add article
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/add' }, opt).url
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

// batchAdd Add article
module.exports.batchAdd = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/batch_add' }, opt).url
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

// del Delete article
module.exports.del = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/del' }, opt).url
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

// batchDel Delete article
module.exports.batchDel = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/batch_del' }, opt).url
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

// update Update article
module.exports.update = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/update' }, opt).url
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

// batchUpdate Update article
module.exports.batchUpdate = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/batch_update' }, opt).url
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

// page Article page query
module.exports.page = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/page?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get Get article info
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// todo article todo
module.exports.todo = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/article/todo?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

