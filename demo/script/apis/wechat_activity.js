// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request').default

// batchAdd 添加活动
module.exports.batchAdd = (data) => {
  const url = '/api/wechat/activity/batch_add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// add 添加活动
module.exports.add = (data) => {
  const url = '/api/wechat/activity/add'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// batchDel 删除活动
module.exports.batchDel = (data) => {
  const url = '/api/wechat/activity/batch_del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// del 删除活动
module.exports.del = (data) => {
  const url = '/api/wechat/activity/del'
  return axios({
    url: url,
    method: 'delete',
    data
  })
}

// batchUpdate 更新活动
module.exports.batchUpdate = (data) => {
  const url = '/api/wechat/activity/batch_update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// update 更新活动
module.exports.update = (data) => {
  const url = '/api/wechat/activity/update'
  return axios({
    url: url,
    method: 'put',
    data
  })
}

// list 活动分页查询
module.exports.list = (data) => {
  let url = '/api/wechat/activity/list?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// one 获取活动
module.exports.one = (data) => {
  let url = '/api/wechat/activity/one?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

// increase 增加次数
module.exports.increase = (data) => {
  const url = '/api/v1/wechat/activity/increase'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

// increaseV2 增加次数
module.exports.increaseV2 = (data) => {
  const url = '/api/v2/wechat/activity/increase_v2'
  return axios({
    url: url,
    method: 'post',
    data
  })
}

