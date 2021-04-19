// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add Add ami
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/redis/mq/add'
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

// get Get ami info
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/redis/mq/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}
