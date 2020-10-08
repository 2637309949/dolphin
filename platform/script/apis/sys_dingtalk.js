// Code generated by dol build. DO NOT EDIT.
const axios = require('@/utils/request').default

// oauth2 授权回调
module.exports.oauth2 = (data) => {
  let url = '/api/sys/dingtalk/oauth2?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get'
  })
}

