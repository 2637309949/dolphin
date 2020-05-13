// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

// TmplAxios defined template
var TmplAxios = `// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('axios')

const request = axios.create({
  baseURL: '/',
  timeout: 6000,
  headers: { 'X-Custom-Header': 'xxx' }
});

module.exports = request
`

// TmplAPIS defined template
var TmplAPIS = `// Code generated by dol build. DO NOT EDIT.
// source: auto.go
const axios = require('@/utils/request').default
{{range .Controller.APIS}}
{{- $tn := .ToUpperCase .Name}}
// {{.LcFirst $tn}} {{.Desc}}
module.exports.{{.LcFirst $tn}} = (data) => {
  {{- if eq .Method "get"}}
  let url = '/api{{.APIPrefix .Version}}/{{.APIPath $.Controller.Name .Path}}/{{.Name}}?'
  for (var key in data) {
    url += key + '=' + encodeURIComponent(data[key]) + '&'
  }
  return axios({
    url: url,
    method: '{{.Method}}'
  })
  {{- else}}
  const url = '/api{{.APIPrefix .Version}}/{{.APIPath $.Controller.Name .Path}}/{{.Name}}'
  return axios({
    url: url,
    method: '{{.Method}}',
    data
  })
  {{- end}}
}
{{end}}
`

// TmplAPI defined template
var TmplAPI = `// Code generated by dol build. DO NOT EDIT.
// source: auto.go
{{- range .Controllers}}
{{- $tn := .ToUpperCase .Name}}
// {{.LcFirst $tn}} {{.Desc}}
module.exports.{{.LcFirst $tn}} = require('./{{.Name}}')
{{- end}}
`
