// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// message Get message
module.exports.message = (data = {}, opt = {}) => {
  let url = opt.url || '/api/rpc/message?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}
