// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add Add article
module.exports.add = (data = {}, opt = {}) => {
  const url = opt.url ||  '/api/kafka/add'
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

// get Get kafka info
module.exports.get = (data = {}, opt = {}) => {
  let url = opt.url || '/api/kafka/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}
