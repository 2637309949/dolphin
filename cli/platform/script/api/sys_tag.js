// Code generated by dol build. DO NOT EDIT.
// source: auto.go

const axios = require('../axios')


// Add 添加标签
module.exports.Add = (data) => {
	const url = '/api/sys/tag/add'
	return axios({
		url: url,
		method: 'post',
		data
    })
}

// Del 删除标签
module.exports.Del = (data) => {
	const url = '/api/sys/tag/del'
	return axios({
		url: url,
		method: 'delete',
		data
    })
}

// Update 更新标签
module.exports.Update = (data) => {
	const url = '/api/sys/tag/update'
	return axios({
		url: url,
		method: 'put',
		data
    })
}

// Page 标签分页查询
module.exports.Page = (data) => {
	const url = '/api/sys/tag/page'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

// Get 获取标签信息
module.exports.Get = (data) => {
	const url = '/api/sys/tag/get'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

