// Code generated by dol build. DO NOT EDIT.
// source: auto.go

const axios = require('../axios')


// Login 登录信息
module.exports.Login = (data) => {
	const url = '/api/sys/oauth2/login'
	return axios({
		url: url,
		method: 'post',
		data
    })
}

// Affirm 用户授权
module.exports.Affirm = (data) => {
	const url = '/api/sys/oauth2/affirm'
	return axios({
		url: url,
		method: 'post',
		data
    })
}

// Authorize 用户授权
module.exports.Authorize = (data) => {
	const url = '/api/sys/oauth2/authorize'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

// Token 获取令牌
module.exports.Token = (data) => {
	const url = '/api/sys/oauth2/token'
	return axios({
		url: url,
		method: 'post',
		data
    })
}

// URL 授权地址
module.exports.URL = (data) => {
	const url = '/api/sys/oauth2/url'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

// Oauth2 授权回调
module.exports.Oauth2 = (data) => {
	const url = '/api/sys/oauth2/oauth2'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

// Refresh 刷新令牌
module.exports.Refresh = (data) => {
	const url = '/api/sys/oauth2/refresh'
	return axios({
		url: url,
		method: 'get',
		data
    })
}

