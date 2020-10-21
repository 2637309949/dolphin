// Code generated by dol build. DO NOT EDIT.
const axios = require('@/utils/request').default

// add 添加站内消息
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/notification/add'
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

// del 删除站内消息
module.exports.del = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/notification/del'
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

// update 更新站内消息
module.exports.update = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/sys/notification/update'
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

// page 站内消息分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/notification/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取站内消息信息
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/sys/notification/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

