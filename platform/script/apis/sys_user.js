// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加用户
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/add' }, opt).url
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

// batchAdd 添加用户
module.exports.batchAdd = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/batch_add' }, opt).url
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

// del 删除用户
module.exports.del = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/del' }, opt).url
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

// batchDel 删除用户
module.exports.batchDel = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/batch_del' }, opt).url
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

// update 更新用户
module.exports.update = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/update' }, opt).url
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

// batchUpdate 更新用户
module.exports.batchUpdate = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/batch_update' }, opt).url
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

// page 用户分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/page?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取用户信息
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// login 用户认证
module.exports.login = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/login' }, opt).url
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

// logout 用户退出登录
module.exports.logout = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/user/logout?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

