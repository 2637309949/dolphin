// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add Add ami
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/redis/mq/add' }, opt).url
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

// get Get ami info
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/redis/mq/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${key}=${encodeURIComponent(data[key])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

