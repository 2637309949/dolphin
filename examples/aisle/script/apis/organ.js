// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add Add organ
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/organ/add'
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

// del Delete organ
module.exports.del = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/organ/del'
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

// update Update organ
module.exports.update = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/organ/update'
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

// page Article page query
module.exports.page = (data = {}, opt = {}) => {
  let url = opt.url || '/api/organ/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get Get organ info
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/organ/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

