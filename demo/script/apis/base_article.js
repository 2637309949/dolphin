// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request').default

// add 添加文章
module.exports.add = (data) => {
  const url = '/api/base/article/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// del 删除文章
module.exports.del = (data) => {
  const url = '/api/base/article/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// update 更新文章
module.exports.update = (data) => {
  const url = '/api/base/article/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// page 文章分页查询
module.exports.page = (data) => {
  let url = '/api/base/article/page?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取文章信息
module.exports.get = (data) => {
  let url = '/api/base/article/get?'
  for (var key in data) {
    url += key + '=' + data[key] + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

