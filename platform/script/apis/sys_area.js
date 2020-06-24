// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request').default

// add 添加区域
module.exports.add = (data) => {
  const url = '/api/sys/area/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// del 删除区域
module.exports.del = (data) => {
  const url = '/api/sys/area/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// update 更新区域
module.exports.update = (data) => {
  const url = '/api/sys/area/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// page 区域分页查询
module.exports.page = (data) => {
  let url = '/api/sys/area/page?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// get 获取区域信息
module.exports.get = (data) => {
  let url = '/api/sys/area/get?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}
