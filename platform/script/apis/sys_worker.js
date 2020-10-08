// Code generated by dol build. DO NOT EDIT.
const axios = require('@/utils/request').default

// add 添加worker
module.exports.add = (data) => {
  const url = '/api/sys/worker/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// get 获取worker信息
module.exports.get = (data) => {
  let url = '/api/sys/worker/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

