// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// login 用户认证
module.exports.login = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/cas/login'
  if ((opt.method || 'post') === 'get') {
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

// logout 注销信息
module.exports.logout = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/logout?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// affirm 用户授权
module.exports.affirm = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/cas/affirm'
  if ((opt.method || 'post') === 'get') {
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

// authorize 用户授权
module.exports.authorize = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/authorize?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// token 获取令牌
module.exports.token = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/cas/token'
  if ((opt.method || 'post') === 'get') {
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

// uRL 授权地址
module.exports.uRL = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/url?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// oauth2 授权回调
module.exports.oauth2 = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/oauth2?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// refresh 刷新令牌
module.exports.refresh = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/refresh?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// check 检验令牌
module.exports.check = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/check?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// profile 用户信息
module.exports.profile = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/profile?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// qrcode 扫码登录(绑定第三方)
module.exports.qrcode = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/cas/qrcode?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

