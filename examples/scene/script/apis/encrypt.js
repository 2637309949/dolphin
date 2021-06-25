// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add Add Encrypt
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/encrypt/add' }, opt).url
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

// info skip auth
module.exports.info = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/encrypt/info?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

