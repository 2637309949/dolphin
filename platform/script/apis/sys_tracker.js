// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// page 日志分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tracker/page?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取日志信息
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/tracker/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

