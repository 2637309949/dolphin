// Code generated by dol build. DO NOT EDIT.
const axios = require('../request').default

// add 添加附件
module.exports.add = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/add' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// batchAdd 添加附件
module.exports.batchAdd = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/batch_add' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// upload 上传附件
module.exports.upload = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/upload' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// export 附件导出
module.exports.export = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/export?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// del 删除附件
module.exports.del = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/del' }, opt).url
  if ((opt.method || 'delete') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
    return axios({
      url: url,
      method: 'get',
      ...opt
    })
  }
  return axios({
    url: url,
    method: 'delete',
    data,
    ...opt
  })
}

// batchDel 删除附件
module.exports.batchDel = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/batch_del' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// update 更新附件
module.exports.update = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/update' }, opt).url
  if ((opt.method || 'put') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
    return axios({
      url: url,
      method: 'get',
      ...opt
    })
  }
  return axios({
    url: url,
    method: 'put',
    data,
    ...opt
  })
}

// batchUpdate 添加附件
module.exports.batchUpdate = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/batch_update' }, opt).url
  if ((opt.method || 'post') === 'get') {
    url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
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

// page 附件分页查询
module.exports.page = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/page?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

// get 获取附件信息
module.exports.get = (data = {}, opt = {}) => {
  let url = Object.assign({ url: '/api/sys/attachment/get?' }, opt).url
  url = Object.keys(data).reduce((acc, curr) => `${acc}${curr}=${encodeURIComponent(data[curr])}&` ,url)
  return axios({
    url: url,
    method: 'get',
    ...opt
  })
}

