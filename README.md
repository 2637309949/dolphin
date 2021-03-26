# Dolphin, Go code generate Framework
[![2637309949](https://circleci.com/gh/2637309949/dolphin.svg?style=shield)](https://circleci.com/gh/2637309949/dolphin)
[![2637309949](https://img.shields.io/github/release/2637309949/dolphin.svg?style=flat-square)](https://github.com/2637309949/dolphin/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/2637309949/dolphin?status.svg)](https://pkg.go.dev/github.com/2637309949/dolphin?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/2637309949/dolphin)](https://goreportcard.com/report/github.com/2637309949/dolphin)

<img align="right" width="159px" src="./assets/dolphin.jpeg">

Dolphin is a code generate tools and web Framework written in Go (Golang), Will reduce the repetitive workload of adding, deleting, revising, and conducting inspections


# Contents
<!-- TOC -->

- [Dolphin, Go code generate Framework](#dolphin-go-code-generate-framework)
- [Contents](#contents)
    - [Quick start](#quick-start)
    - [Directory structure](#directory-structure)
    - [Features](#features)
    - [CMD](#cmd)
        - [build](#build)
        - [clean](#clean)
        - [init](#init)
    - [API Examples](#api-examples)
        - [dolphin-ui](#dolphin-ui)
    - [XML Label](#xml-label)
        - [application](#application)
        - [bean](#bean)
        - [controller](#controller)
        - [api example](#api-example)
            - [add](#add)
                - [one](#one)
                - [batch](#batch)
            - [delete](#delete)
                - [one](#one-1)
                - [batch](#batch-1)
            - [update](#update)
                - [one](#one-2)
                - [batch](#batch-2)
            - [page](#page)
            - [tree](#tree)
            - [one](#one-3)
            - [other](#other)
        - [table](#table)
            - [packages](#packages)
            - [helpers](#helpers)
        - [rpc](#rpc)
            - [server](#server)
            - [client](#client)
    - [Domain](#domain)
        - [app_name](#app_name)
        - [domain](#domain)
    - [OAuth Server](#oauth-server)
        - [redirect sso](#redirect-sso)
        - [sso auth](#sso-auth)
        - [sso affirm](#sso-affirm)
        - [sso token](#sso-token)
        - [sso callback](#sso-callback)
        - [example](#example)
    - [Workload](#workload)
        - [Add Handler](#add-handler)
        - [Add Job](#add-job)
        - [Fetch Job status](#fetch-job-status)
    - [Cron](#cron)
        - [Add Func](#add-func)
        - [Reflesh Func](#reflesh-func)
        - [Del Func](#del-func)
        - [Try Func](#try-func)
    - [Load User Info](#load-user-info)
        - [GetFieldSliceByName](#getfieldslicebyname)
        - [PatchSliceByField](#patchslicebyfield)
    - [Debug pprof](#debug-pprof)
        - [ModHeader](#modheader)
        - [Debug](#debug)
    - [High-level Example](#high-level-example)
        - [File processing](#file-processing)
            - [Upload File](#upload-file)
            - [Persist File or Remove File](#persist-file-or-remove-file)
        - [Template processing](#template-processing)
            - [RenderFile](#renderfile)
            - [RenderHTML](#renderhtml)
        - [Queue processing](#queue-processing)
            - [Base on Redis](#base-on-redis)
            - [Base on Kafka](#base-on-kafka)
        - [Interface authentication](#interface-authentication)
            - [Never authentication](#never-authentication)
            - [Token authentication](#token-authentication)
            - [Encrypt authentication](#encrypt-authentication)
- [MIT License](#mit-license)

<!-- /TOC -->

## Quick start

<img width="580px" src="./assets/dolphin.gif">

1. The first need [Go](https://golang.org/) installed, then you can use the below Go command to install Dolphin.
```sh
$ go get -u github.com/2637309949/dolphin/cmd/dolphin
```

2. Create project dir and run dolphin

```sh
$ mkdir example && cd example && dolphin init && dolphin build && go run main.go
```

Output:
```sh
time="2020/06/13 11:55:58" level=info msg="grpc listen on port:9081"
time="2020/06/13 11:55:58" level=info msg="http listen on port:8082"
```

## Directory structure
> The quasi-directory structure of the project is shown below， The project structure has been simplified as a guideline, such as managing large-scale projects and recommending new sub-projects

```sh
	.
	├── app
	│   ├── app.auto.go
	│   ├── app.go
	│   ├── article.go
	│   └── article.go.new
	├── app.properties
	├── doc
	│   └── swagger.yaml
	├── go.mod
	├── go.sum
	├── log
	│   └── demo.2020071400
	├── main.go
	├── model
	│   ├── article.auto.go
	│   └── article_info.auto.go
	├── rpc
	│   ├── message.cli.go
	│   ├── message.go
	│   ├── message.go.new
	│   └── proto
	│       ├── message.pb.go
	│       ├── message.proto
	│       └── message.proto.new
	├── script
	│   ├── apis
	│   │   ├── article.js
	│   │   └── index.js
	│   └── axios.js
	├── sql
	│   ├── article
	│   │   ├── article_page_count.tpl
	│   │   └── article_page_select.tpl
	│   └── sqlmap
	│       └── article.xml
	├── srv
	│   ├── article.go
	│   └── worker_hello.go
	├── static
	│   ├── files
	│   │   ├── 6b7ead55-f663-4340-a594-d282d5baf753.xlsx
	│   │   └── 6dc88052-54e0-4aa9-a344-fb2b3c30f9b6.xlsx
	│   └── web
	│       ├── affirm.html
	│       └── login.html
	├── util
	│   └── tool.go
	└── xml
		├── application.xml
		├── bean
		│   └── article_info.xml
		├── controller
		│   └── article.xml
		├── rpc
		│   └── message.xml
		└── table
			└── article.xml
```

## Features

```
- Generates the code base on XML configuration

- Generates doc base on XML configuration

- Generates SQL base on XML configuration

- Handles the serialization null problem

- Multi-tenant support

- Login / Logout, or single sign on

- Permission Authentication

- Quick excel reporting or parsing

- Support routing caching

- Data permission control

- Log trace record

- RPC remote service

- The k8S deployment file is generated by default

- Support database reverse XML generation
```

## CMD

Command options

```sh
dolphin, a cli tools for generate golang code

Usage:
  dolphin [command]

Available Commands:
  build       Build project from xml
  clean       Removing intermediate files
  help        Help about any command
  init        Initialize a empty project

Flags:
  -h, --help   help for dolphin

Use "dolphin [command] --help" for more information about a command.
```

### build
The build command generates the preset function by executing the built-in Pipeline function, You can specify that only a pipeline will be executed via the @ symbol

```sh
dolphin build @table xml/test
```

Existing built-in Pipeline function:

| Function   |      Action      |
|----------|:-------------:|
| main | create main file source |
| app | create engine template source |
| ctr | create controller source |
| proto | create proto3 source |
| srv | create server source |
| model | create model source |
| bean | create bean source |
| auto | create register source |
| tool | create tool source |
| sql | create sql source, .sql to .go |
| sqlmap | create table sqlmap |
| oauth | create oauth h5 template |
| script | create js api |
| deploy | create k8s template |
| doc | create swagger api doc |
| table | create table from datasource |


### clean
The clean command clears temporary files

```sh
dolphin clean
```
### init
The init command, as stated, generates a series of initialization files

```sh
mkdir demo && cd demo && dolphin init
```

## API Examples

### dolphin-ui 

> An erp template that uses dolphin background support.

<img align="right" width="200px" src="./assets/dolphin-ui.png">

You can find a number of ready-to-run examples at [dolphin examples repository.](https://github.com/2637309949/dolphin-ui)


## XML Label

### application
> application label contain app infomation, such as name, package

Example: 

```xml
<?xml version="1.0" encoding="utf-8" ?>
<application name="demo" desc="template" packagename="demo"/>
```

application

| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  required, application name |
| desc |    application desc  |
| packagename |    required, application packagename  |


### bean
> bean, you can declare object in bean, just like spring bean. all bean and model will be placed in the model directory, so you needs another name if the conflict

Example: 

```xml
<bean name="activity_info" desc="desc" packages="xxx" extends="$applet_activity">
    <prop name="code" desc="编码" type="xx.String" />
    <prop name="name" desc="名称" type="xx.String" />
</bean>
```

Generate code:

```go
// Code generated by dol build. DO NOT EDIT.

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ArticleInfo defined 文章信息
type ArticleInfo struct {
	*Article
	// 地址
	URL null.String `json:"url" xml:"url"`
}
```

bean
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  bean name |
| desc |    bean desc  |
| packagename |    third party package name，use "," to split |
| extends |    bean extends  |


prop
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  prop name |
| desc |    prop desc  |
| type |    prop type  |

### controller
> controller, a collect api, you can declare api prefix

Example: 

```xml
<controller name="activity" desc="微信活动" />
```

controller
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  controller name |
| desc |    controller desc  |
| prefix |    controller desc  |


### api example
> api, api func in controller. we has some built-in func such as 'add', 'delete', 'update', 'page', 'get', 'tree', or you can refined if you need.

#### add

##### one

```xml
<api name="add" func="add" table="sys_client" desc="添加客户端" method="post">
    <param name="user" type="$sys_client" desc="客户端信息" />
    <return>
        <success type="$success"/>
        <failure type="$fail"/>
    </return>
</api>
```

Generate code:

```go
// SysClientAdd api implementation
// @Summary 添加客户端
// @Tags 客户端
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysClient false "客户端信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/client/add [post]
func SysClientAdd(ctx *Context) {
	var payload model.SysClient
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFrom(time.Now().Value())
	payload.Creater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.Updater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	payload.AppName = null.StringFrom(viper.GetString("app.name"))
	ret, err := ctx.PlatformDB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

##### batch

> Of course, you can also specify array parameters, which will automatically generate batch added templates.

```xml
<api name="batch_add" func="add" table="sys_role_menu" method="post" desc="添加角色菜单">
	<param name="role_menu" type="[]$sys_role_menu" desc="角色菜单信息" />
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

Generate code:  

```go
// SysRoleMenuBatchAdd api implementation
// @Summary 添加角色菜单
// @Tags 角色菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param role_menu body []model.SysRoleMenu false "角色菜单信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/role/menu/batch_add [post]
func SysRoleMenuBatchAdd(ctx *Context) {
	var payload []*model.SysRoleMenu
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form *model.SysRoleMenu) {
		form.ID = null.StringFromUUID()
		form.CreateTime = null.TimeFrom(time.Now().Value())
		form.Creater = null.StringFrom(ctx.GetToken().GetUserID())
		form.UpdateTime = null.TimeFrom(time.Now().Value())
		form.Updater = null.StringFrom(ctx.GetToken().GetUserID())
		form.IsDelete = null.IntFrom(0)
	})
	payload = funk.Filter(payload, func(form *model.SysRoleMenu) bool {
		ext, _ := ctx.DB.Where("role_id=? and menu_id=?", form.RoleId.String, 
			form.MenuId.String).Exist(new(model.SysRoleMenu))
		return !ext
	}).([]*model.SysRoleMenu)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```


#### delete

##### one

> The system default templates are soft delete logic, this is also highly recommended. If you need hard delete, please do it yourself

```xml
<api name="del" func="delete" table="sys_client" desc="删除客户端" method="delete">
    <param name="sys_client" type="$sys_client" desc="客户端" />
    <return>
        <success type="$success"/>
        <failure type="$fail"/>
    </return>
</api>
```

Generate code:

```go
// SysClientDel api implementation
// @Summary 删除客户端
// @Tags 客户端
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_client body model.SysClient false "客户端"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/client/del [delete]
func SysClientDel(ctx *Context) {
	var payload model.SysClient
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.PlatformDB.In("id", payload.ID.String).Update(&model.SysClient{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		Updater:   null.StringFrom(ctx.GetToken().GetUserID()),
		IsDelete:    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

##### batch

> Of course, you can also specify array parameters, which will automatically generate batch deleted templates.


```xml
<api name="batch_del" func="delete" table="sys_optionset" method="delete" desc="删除字典">
	<param name="user" type="[]$sys_optionset" desc="字典" />
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

Generate code:   

```go
// SysOptionsetBatchDel api implementation
// @Summary 删除字典
// @Tags 字典
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body []model.SysOptionset false "字典"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/optionset/batch_del [delete]
func SysOptionsetBatchDel(ctx *Context) {
	var payload []*model.SysOptionset
	var ids []string
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	funk.ForEach(payload, func(form model.SysOptionset) {
		ids = append(ids, form.ID.String)
	})
	ret, err := ctx.DB.In("id", ids).Update(&model.SysOptionset{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		Updater:   null.StringFrom(ctx.GetToken().GetUserID()),
		IsDelete:    null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```


#### update

##### one

> The entire model field definition is null type, so don't worry about the default value types, as long as you reference packages/ xormplus (native xorm has been modified to fit null packages)  


```xml
<api name="update" func="update" table="sys_client" desc="更新客户端" method="put">
    <param name="user" type="$sys_role" desc="客户端信息" />
    <return>
        <success type="$success"/>
        <failure type="$fail"/>
    </return>
</api>
```

Generate code:

```go
// SysClientUpdate api implementation
// @Summary 更新客户端
// @Tags 客户端
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "客户端信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/client/update [put]
func SysClientUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.Updater = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.PlatformDB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

##### batch

> Of course, you can also specify array parameters, which will automatically generate batch updated templates.

```xml
<api name="batch_update" func="update" table="article" desc="更新文章" method="put">
	<param name="article" type="[]$article" desc="文章信息" />
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

Generate code:

```go
// ArticleBatchUpdate api implementation
// @Summary 更新文章 
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body []model.Article false "文章信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/batch_update [put]
func ArticleBatchUpdate(ctx *Context) {
	var payload []*model.Article
	var err error
	var ret []int64
	var r int64
	if err = ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	funk.ForEach(payload, func(form model.Article) {
		form.Updater = null.StringFrom(ctx.GetToken().GetUserID())
		form.UpdateTime = null.TimeFrom(time.Now().Value())
		r, err = s.ID(form.ID.String).Update(&form)
		ret = append(ret, r)
	})
	if err != nil {
		s.Rollback()
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

#### page

> The most basic paging interface template is also supported by default.

```xml
<api name="page" func="page" table="sys_client" desc="客户端分页查询" method="get">
    <param name="page" type="int" value="1" desc="页码"/>
    <param name="size" type="int"  value="10" desc="单页数"/>
    <param name="app_name" type="string" desc="所属应用"/>
    <return>
        <success type="$success"/>
        <failure type="$fail"/>
    </return>
</api>
```

Generate code:

```go
// SysClientPage api implementation
// @Summary 客户端分页查询
// @Tags 客户端
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/client/page [get]
func SysClientPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("app_name", viper.GetString("app.name"))
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.PlatformDB, "sys_client", "page", "sys_client", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

#### tree

> Tree structure is also a common algorithm. The framework integrates a TreeSearch interface for tree retrieval.

```xml
<api name="page" func="page" table="sys_menu" desc="菜单分页查询" method="get">
	<param name="page" type="int" value="1" desc="页码"/>
	<param name="size" type="int"  value="10" desc="单页数"/>
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

Generate code:

```go
// SysMenuTree api implementation
// @Summary 菜单树形结构
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Fail
// @Router /api/sys/menu/tree [get]
func SysMenuTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetRule("sys_menu_tree")
	q.SetTags()
	ret, err := ctx.TreeSearch(ctx.DB, "sys_menu", "tree", "sys_menu", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

> TreeSearch:   
> Detailed instructions can be found in the TreeSearch source code  

```go
// platform/app/app.ctx.go#TreeSearch
func (ctx *Context) TreeSearch(db *xorm.Engine, controller, api, table string, 
		q map[string]interface{}) (interface{}, error)
```

#### one

> Find a single piece of data by unique key.

```xml
<api name="get" func="one" table="sys_client" desc="获取客户端信息" method="get">
    <param name="id" type="string" desc="客户端id" />
    <return>
        <success type="$success"/>
        <failure type="$fail"/>
    </return>
</api>
```

Generate code:

```go
// SysClientGet api implementation
// @Summary 获取客户端信息
// @Tags 客户端
// @Param Authorization header string false "认证令牌"
// @Param id query string false "客户端id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/client/get [get]
func SysClientGet(ctx *Context) {
	var entity model.SysClient
	id := ctx.Query("id")
	_, err := ctx.PlatformDB.ID(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
```

#### other

> In addition to the default interface, custom interfaces are also supported.

```xml
<api name="payment" method="post" desc="文章付费">
	<param name="article" type="$article_info" desc="文章"/>
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

Generate code:

```go
// ArticlePayment api implementation
// @Summary 文章分页查询
// @Tags 文章
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param article body model.ArticleInfo false "文章"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/article/payment [post]
func ArticlePayment(ctx *Context) {
	var payload model.ArticleInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := srv.ArticleAction(payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

api
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  api name |
| desc |    api desc   |
| func | built-in func, 'add', 'delete', 'update', 'page', 'get' |
| table |    table name if you use built-in func   |
| method |    http method   |
| roles |    roles middles   |
| cache |    cache middles   |

param
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  param name |
| desc |    param desc  |
| type | param type |
| value | default value |

return

| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| success |  success tag |
| failure |    success tag  |


### table
> table, as you khnow, you can defined any table structure as you need. and you should use `null` type if you wan to accept form data that avoid null value problems in golang.


table
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  table name |
| desc |    table desc   |
| packages | third party package name，use "," to split |

column
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  column name |
| desc |    column desc  |
| type | column type |
| xorm | xorm tag, please refer to XORM for details |


Example: 

```xml
<table name="article" desc="文章" packages="github.com/2637309949/dolphin/packages/null">
	<column name="id" desc="主键" type="null.String" xorm="varchar(36) notnull unique pk" />
	<column name="type" desc="类别" type="null.String" xorm="varchar(36)" />

	<column name="creater" desc="创建人" type="null.String" xorm="varchar(36)" />
	<column name="create_time" desc="创建时间" type="null.Time" xorm="datetime" />
	<column name="updater" desc="最后更新人" type="null.String" xorm="varchar(36)" />
	<column name="update_time" desc="最后更新时间" type="null.Time" xorm="datetime" />
	<column name="is_delete" desc="删除标记" type="null.Int" xorm="notnull" />
	<column name="remark" desc="备注" type="null.String" xorm="varchar(200)" />
</table>
```

Generate code:

```go
// Code generated by dol build. DO NOT EDIT.

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Article defined 文章
type Article struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk 'id'" json:"id" xml:"id"`
	// 类别
	Type null.String `xorm:"varchar(36) 'type'" json:"type" xml:"type"`
	
	// 创建人
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime 'create_time'" json:"create_time" xml:"create_time"`
	// 最后更新人
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime 'update_time'" json:"update_time" xml:"update_time"`
	// 删除标记
	IsDelete null.Int `xorm:"notnull 'is_delete'" json:"is_delete" xml:"is_delete"`
	// 备注
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" xml:"remark"`
}

// TableName table name of defined Article
func (m *Article) TableName() string {
	return "article"
}
```

####  packages

> The system adds the null and decimal packages by default, null is used to support data type default issues, and decimal is used to support financial data calculations.  

Example: 
```xml
<table 
	name="article" 
	desc="文章" 
	packages="github.com/2637309949/dolphin/packages/null,
	github.com/2637309949/dolphin/packages/decimal">
	<column name="id" desc="主键" type="null.String" xorm="varchar(36) notnull unique pk" />
	<column name="reward" desc="打赏" type="decimal.Decimal" xorm="decimal(6,2)" />

	<column name="creater" desc="创建人" type="null.String" xorm="varchar(36)" />
	<column name="create_time" desc="创建时间" type="null.Time" xorm="datetime" />
	<column name="updater" desc="最后更新人" type="null.String" xorm="varchar(36)" />
	<column name="update_time" desc="最后更新时间" type="null.Time" xorm="datetime" />
	<column name="is_delete" desc="删除标记" type="null.Int" xorm="notnull" />
	<column name="remark" desc="备注" type="null.String" xorm="varchar(200)" />
</table>
```


####  helpers

> It is recommended to write the helper tool class in the corresponding model directory.

Example: 
```go
//platform/model/sys_user.helper.go
// SetPassword Method to set salt and hash the password for a user
func (m *SysUser) SetPassword(password string) {
	b := util.RandString(16, util.RandNumChar)
	m.Salt = null.StringFrom(b)
	dk, err := scrypt.Key([]byte(password), []byte(m.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	m.Password = null.StringFrom(fmt.Sprintf("%x", dk))
}

// ValidPassword Method to check the entered password is correct or not
func (m *SysUser) ValidPassword(password string) bool {
	dk, err := scrypt.Key([]byte(password), []byte(m.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	return m.Password.String == fmt.Sprintf("%x", dk)
}
```


### rpc

> rpc, as a microservice interaction, the basic proto file can be generated here in rpc dir, as well as automatic registration in auto file.

```xml
<service name="message" desc="消息">
    <rpc name="send_mail" desc="发送邮件">
        <request type="$article" desc="文章信息"/>
        <reply type="$success" desc="文章信息"/>
    </rpc>
</service>
```

Generate code:

proto file:
```proto3
// Code generated by dol build. Only Generate by tools if not existed, 
// your can rewrite platform.App default action
// source: MessageSrv.proto

syntax = "proto3";

package proto;

// MessageSrv defined
service MessageSrv {
  rpc SendMail (MessageMail) returns (MessageReply) {}
}


// MessageMail defined
message MessageMail {}

// MessageReply defined
message MessageReply {}
```

#### server

> rpc service is automatically generated

rpc srv impl:
```go
// Code generated by dol build. Only Generate by tools if not existed.
// source: MessageSrv.go

package rpc

import (
	"demo/rpc/proto"

	"golang.org/x/net/context"
)

// MessageSrv defined
type MessageSrv struct{}

// SendMail defined
func (srv *MessageSrv) SendMail(
	ctx context.Context, 
	in *proto.MessageMail) (*proto.MessageReply, error) {
	return &proto.MessageReply{}, nil
}
```

#### client

> rpc client is automatically generated

rpc cli endpoint:
```go
// Code generated by dol build. Only Generate by tools if not existed.
// source: MessageSrv.cli.go

package rpc

import (
	"demo/rpc/proto"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"google.golang.org/grpc"
)

// MessageSrvClient defined
var MessageSrvClient proto.MessageSrvClient

func init() {
	opt := grpc.WithInsecure()
	conn, err := grpc.Dial(viper.GetString("rpc.message_srv"), opt)
	if err != nil {
		logrus.Error("grpc dial failed: %v", err)
	}
	MessageSrvClient = proto.NewMessageSrvClient(conn)
}
```

rpc server will be automatically registered in app/app.auto.go
```go
// MessageSrv defined
func MessageSrvService(engine *Engine) {
	proto.RegisterMessageSrvServer(engine.GRPC, &rpc.MessageSrv{})
}

// SyncService defined
func SyncService() error {
	MessageSrvService(App)
	return nil
}

```

The quasi-directory structure of the rpc is shown below:
```sh
	├── message.cli.go
	├── message.go
	├── message.go.new
	└── proto
		├── message.pb.go
		├── message.proto
		└── message.proto.new
```


## Domain

> Domain, a model of multi-tenant support core. Application splitting is also supported.

```xml
<table name="sys_domain" packages="xx/null" bind="platform">
	<column name="id" type="null.String" xorm="varchar(36) notnull unique pk" />
	<column name="name" type="null.String" xorm="varchar(36) notnull" />
	<column name="app_name" type="null.String" xorm="varchar(36) notnull" />
	<column name="domain" type="null.String" xorm="varchar(36) notnull" />
	<column name="full_name" type="null.String" xorm="varchar(36)" />
	<column name="contact_name" type="null.String" xorm="varchar(36)" />
	<column name="contact_email" type="null.String" xorm="varchar(50) " />
	<column name="contact_mobile" type="null.String" xorm="varchar(50) " />
	<column name="data_source" type="null.String" xorm="varchar(200) notnull" />
	<column name="driver_name" type="null.String" xorm="varchar(50) notnull" />
	<column name="login_url" type="null.String" xorm="varchar(200)" />
	<column name="api_url" type="null.String" xorm="varchar(200)" />
	<column name="static_url" type="null.String" xorm="varchar(200)" />
	<column name="theme" type="null.String" xorm="varchar(50) " />
	<column name="type" type="null.Int" xorm="notnull" />
	<column name="status" type="null.Int" xorm="notnull" />
	<column name="auth_mode" type="null.Int" xorm="notnull" />
	<column name="is_sync" type="null.Int" xorm="notnull" />

	<column name="creater" type="null.String" xorm="varchar(36) notnull" />
	<column name="create_time" type="null.Time" xorm="datetime notnull" />
	<column name="updater" type="null.String" xorm="varchar(36) notnull" />
	<column name="update_time" type="null.Time" xorm="datetime notnull" />
	<column name="is_delete" type="null.Int" xorm="notnull" />
	<column name="remark" type="null.String" xorm="varchar(200)" />
</table>
```

### app_name

> app_name, desynchronize the model as a tag. if you connect same datasource url from localhost, and you would find all tables in `data_source` of same app_name datasource would be created

### domain

> Identify different tenants, the logged in user will use the matching domain to find the DB

As shown in the code below

```go
// Auth middles
func Auth(ctx *Context) {
	if !ctx.Auth(ctx.Request) {
		ctx.Fail(util.ErrInvalidAccessToken, 401)
		ctx.Abort()
		return
	}
	if ctx.DB = App.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
		ctx.Fail(util.ErrInvalidDomain)
		ctx.Abort()
		return
	}
	ctx.Set("DB", ctx.DB)
	ctx.Set("AuthInfo", ctx.AuthInfo)
	ctx.Next()
}
```

if you want to get datasource of `xxx`, you can do the following.

```go
App.Manager.GetBusinessDB("xxx")
```

## OAuth Server

> All projects that inherit the platform support single sign-on by default, you can deploy independently or directly as SSO Server

	Your FrontEnd Project      Your BackEnd Project                     SSO
		||                              ||                              ||
		||                              ||                              ||
		||  1.     fetch api            ||                              ||
		||  -------------------------\  ||                              ||
		||      unauthorized            ||                              ||
		|| /-------------------------   ||                              ||
		||                              ||                              ||
		||  2. fetch oauth url          ||                              ||
		||  ------------------------\   ||                              ||
		|| /-------------------------   ||                              ||
		||                              ||                              ||
		||  3. goto sso oauth           ||     goto sso oauth           ||
		||  -------------------------   ||  -------------------------\  ||
		||                              ||     goto client with code    ||
		||                              || /-------------------------   ||
		||                              ||                              ||
		||                              ||                              ||
		||  4. redirect and set cookie  ||     get token by code        ||
		||                              || -------------------------\   ||
		|| /-------------------------   || /-------------------------   ||
		||                           	||                              ||

### redirect sso

Code segment in platform, You can carry the status if needed.

```go
// SysCasURL api implementation
// @Summary 授权地址
// @Tags 认证中心
// @Param redirect_uri query string false "定向URL"
// @Param state query string false "状态"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/cas/url [get]
func SysCasURL(ctx *Context)
```

### sso auth

Code segment in platform, authentication logic

```go
// SysCasLogin api implementation
// @Summary 用户认证
// @Tags 认证中心
// @Accept multipart/form-data
// @Param username formData string false "用户名称"
// @Param password formData string false "用户密码"
// @Param domain formData string false "用户域"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/cas/login [post]
func SysCasLogin(ctx *Context)
```

### sso affirm

Code segment in platform, you can rewrite this way if you want to skip Affirm

```go
// SysCasAffirm api implementation
// @Summary 用户授权
// @Tags 认证中心
// @Accept application/json
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/cas/affirm [post]
func SysCasAffirm(ctx *Context)
```


### sso token

Code segment in platform, Generate Token by code

```go
// SysCasToken api implementation
// @Summary 获取令牌
// @Tags 认证中心
// @Accept application/json
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/cas/token [post]
func SysCasToken(ctx *Context)
```

### sso callback

Code segment in client, Fetch token from platform and set cookie

```go
// SysCasOauth2 api implementation
// @Summary 授权回调
// @Tags 认证中心
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/cas/oauth2 [get]
func SysCasOauth2(ctx *Context)
```

### example

For more information, please check out dolphin-ui [dolphin-ui](https://github.com/2637309949/dolphin-ui)  


If you want to go sso auth then you can set `auth_mode` to `1`  

```json
"auth_mode": 1
```

public/domain.js  
```js
window.Domain = {
    "name": "localhost",
    "full_name": "localhost",
    "contact_name": null,
    "contact_email": null,
    "contact_mobile": null,
    "login_url": "localhost",
    "api_url": "http://localhost:8082",
    "static_url": null,
    "theme": "default",
    "auth_mode": 1
}
```

## Workload

> High concurrent requests are processed with built-in load interfaces

### Add Handler

```go
// @Summary AddJobHandler
// @Tags worker
func (d *DefaultWorker) AddJobHandler(code string, funk func(model.Worker) (interface{}, error)) {
```

Example:  
demo/srv/article.go  
```go
func init() {
	// add hello topic handler
	pApp.App.Manager
	.Worker()
	.AddJobHandler("hello", func(args pModel.Worker) (interface{}, error) {
		fmt.Printf("topic=%v, payload=%v", "hello", args.Payload)
		return map[string]interface{}{
			"score": 99,
		}, nil
	})
}
```

### Add Job

```go
// SysWorkerAdd api implementation
// @Summary 添加worker
// @Tags worker
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param worker body model.Worker false "worker信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/worker/add [post]
func SysWorkerAdd(ctx *Context) {
```

Example:  

Request:  
```sh
POST /api/sys/worker/add HTTP/1.1
Host: localhost:8082
Content-Type: application/json
token: f4c9f457c82c9ee51e3dc50fea74562b64dd9269
Authorization: Bearer BY3KDUJNMWCN-NQJLKQVAW
Cache-Control: no-cache
Postman-Token: 0a39bb7e-d9bf-607d-120d-ffa59102dab8

{
	"name": "hello",
	"payload": { "user_id": "sdhfusd9f"}
}
```

Reponse:  
```json
{
    "code": 200,
    "data": {
        "code": "fb9b4a91-c918-4d11-a8f7-878e4dd94f70",
        "name": "hello",
        "status": 100
    }
}
```


### Fetch Job status

```go
// SysWorkerGet api implementation
// @Summary 获取worker信息
// @Tags worker
// @Param Authorization header string false "认证令牌"
// @Param code  query  string false "worker code"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/worker/get [get]
func SysWorkerGet(ctx *Context) {
```

Example:  
```sh
GET /api/sys/worker/get?code=a4d13b27-4836-4a1b-b6fe-63473716bc4c HTTP/1.1
Host: localhost:8082
Authorization: Bearer BY3KDUJNMWCN-NQJLKQVAW
Cache-Control: no-cache
Postman-Token: 4f1595de-583b-fac6-06bc-2eafad956d40
```

Response:  
```json
{
    "code": 200,
    "data": {
        "code": "a4d13b27-4836-4a1b-b6fe-63473716bc4c",
        "name": "hello",
        "result": {
            "score": 99
        },
        "status": 103
    }
}
```


## Cron

Cron is not enabled by default, You can do this by modifying the configuration`app.cron = true` if you need. But Cron is stateful, so it must be deployed separately when applying load.

### Add Func

Add a timer task via AddFunc.

```go
AddFunc(string, func()) (int, error)
```

```go
pApp.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})
```

### Reflesh Func

Reflesh a timer task via RefreshFunc.

```go
RefreshFunc(int, string) (int, error)
```

```go
id, _ := pApp.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})
pApp.App.Manager.Cron().RefreshFunc(id, "*/3 * * * * *")
```

### Del Func

Remove a timer task via DelFunc.

```go
DelFunc(int) error
```

```go
id, _ := pApp.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})

pApp.App.Manager.Cron().DelFunc(id)
```

### Try Func

Try a timer task via TryFunc.

```go
TryFunc(int) error
```

```go
id, _ := pApp.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})
pApp.App.Manager.Cron().TryFunc(id)
```

## Load User Info

Because user info db and the primary business db are separated, so user information needs to be loaded separately.  

```go
ctx.PlatformDB 
```
You can load platform info from `PlatformDB`, such as sys_user, sys_client, sys_domain    


```go
ctx.DB 
```
Else, You should load business info from ctx.DB 


### GetFieldSliceByName
Create a new array from inside an array based on the field name.  

```go
uids := slice.GetFieldSliceByName(ret.Data, "id", "'%v'").([]string)

roles, err := srv.SysUserGetUserRolesByUID(ctx.DB, strings.Join(uids, ","))
if err != nil {
	logrus.Error(err)
	ctx.Fail(err)
	return
}
```

### PatchSliceByField
Patch the field value to another array.  

Example:

```go
err = slice.PatchSliceByField(ret.Data, roles, "id", "user_id", "role_name")(&ret.Data)
if err != nil {
	logrus.Error(err)
	ctx.Fail(err)
	return
}
```

## Debug pprof

> The framework itself provides the functionality of pprof, note that only the admin role is visible

### ModHeader

Install ModHeader in chrome, and add token header for any request.

### Debug

Goto http://127.0.0.1:8082/debug/pprof/

```text
/debug/pprof/

Types of profiles available:
Count	Profile
3	allocs
0	block
0	cmdline
48	goroutine
3	heap
0	mutex
0	profile
16	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc 
GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After 
you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the 
seconds GET parameter. After you get the trace file, use the go tool trace command to 
investigate the trace.
```

## High-level Example

### File processing

#### Upload File
```shell
curl -X POST \
  http://127.0.0.1:8081/api/sys/attachment/upload \
  -H 'authorization: Bearer MAGFJEYNMPI9XAIK8GQKLA' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: 3ddea434-b9ee-a3e8-516c-0914facc6517' \
  -F files=@k8s-eschool-1.json
```

#### Persist File or Remove File

```go
// KafkaGet api implementation
// @Summary Get kafka info
// @Tags Kafka controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "kafka id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/kafka/get [get]
func KafkaGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	q.Value()
	ctx.Persist(ctx.DB, "0586e250-5b6c-4a79-9f4a-767a742b7890")
	ctx.Remove(ctx.DB, "6ebce24f-6887-4d6d-a62a-2a706fcf1c3f")
	ret, err := srv.KafkaConsumer(ctx.Raw(), ctx.DB, q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

### Template processing
#### RenderFile

```go
// KafkaGet api implementation
// @Summary Get kafka info
// @Tags Kafka controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "kafka id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/kafka/get [get]
func KafkaGet(ctx *Context) {
	ctx.RenderFile("examples/kafka/static/files/5066.tmpl", "test.txt", 
	map[string]interface{
		"test": 100,
	})
}
```
#### RenderHTML

```go
// KafkaGet api implementation
// @Summary Get kafka info
// @Tags Kafka controller
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "kafka id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/kafka/get [get]
func KafkaGet(ctx *Context) {
	ctx.RenderHTML("examples/kafka/static/files/5066.tmpl",
	map[string]interface{
		"test": 100,
	})
}
```
### Queue processing

#### Base on Redis 

```go
// examples/ami
// AmiAdd api implementation
// @Summary Add ami
// @Tags Ami controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param ami_info body model.AmiInfo false "Ami info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/ami/submit [post]
func AmiAdd(ctx *Context) {
	var payload model.AmiInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := srv.AmiProducer(ctx.Raw(), ctx.DB, payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// examples/ami
// AmiProducer defined srv
func AmiProducer(ctx *gin.Context, 
	db *xorm.Engine, params model.AmiInfo) (interface{}, error) {
	aiStr, err := json.Marshal(params)
	if err != nil {
		logrus.Error("failed to marshal:", err)
		return nil, err
	}
	AmiProducerConn.Send(string(aiStr))
	return nil, nil
}

// examples/ami
// AmiConsumer defined srv
func AmiConsumer(ctx *gin.Context, 
	db *xorm.Engine, params map[string]interface{}) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err.(error), 3)
			fmt.Print(string(goErr.Stack()))
		}
	}()
	var items []model.AmiInfo
	c := AmiConsumerConn.Start()
	cwt, cancel := context.WithCancel(context.TODO())
	go func(cwt context.Context) {
		for {
			select {
			case m, ok := <-c:
				if !ok {
					cancel()
					return
				}
				AmiConsumerConn.Ack(m)
				value := model.AmiInfo{}
				if err := json.Unmarshal([]byte(m.Body), &value); err != nil {
					logrus.Error("failed to unmarshal:", err)
				}
				items = append(items, value)
			case <-time.After(3 * time.Second):
				cancel()
				return
			}
		}
	}(cwt)
	<-cwt.Done()
	// AmiConsumerConn.Stop()
	return items, nil
}
```


#### Base on Kafka 

```go
// examples/kafka
// KafkaAdd api implementation
// @Summary Add article
// @Tags Kafka controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param kafka body model.KafkaInfo false "Kafka info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/kafka/add [post]
func KafkaAdd(ctx *Context) {
	var payload model.KafkaInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := srv.KafkaProducer(ctx.Raw(), ctx.DB, payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// examples/kafka
// KafkaProducer defined srv
func KafkaProducer(ctx *gin.Context, 
	db *xorm.Engine, params model.KafkaInfo) (interface{}, error) {
	kpStr, err := json.Marshal(&params)
	if err != nil {
		logrus.Error("failed to marshal:", err)
		return nil, err
	}
	n, err := kafkaConn.WriteMessages(kafka.Message{Value: kpStr})
	if err != nil {
		logrus.Error("failed to write messages:", err)
		return nil, err
	}
	return n, nil
}

// examples/kafka
func KafkaConsumer(ctx *gin.Context, 
	db *xorm.Engine, params map[string]interface{}) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			goErr := errors.Wrap(err.(error), 3)
			fmt.Print(string(goErr.Stack()))
		}
	}()
	batch := kafkaConn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	defer batch.Close()
	var buffer bytes.Buffer
	var items []model.KafkaInfo
	for {
		b := make([]byte, 10e3) // 10KB max per message
		n, err := batch.Read(b)
		buffer.Write(b[:n])
		if err != nil {
			break
		}
		if bte := buffer.Bytes(); len(bte) > 0 {
			value := model.KafkaInfo{}
			if err := json.Unmarshal(bte, &value); err != nil {
				logrus.Error("failed to unmarshal:", err)
				return nil, err
			}
			items = append(items, value)
		}
	}
	return items, nil
}

```


### Interface authentication

#### Never authentication

```xml
<api name="xml" method="get" desc="XML显示" auth="never">
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

#### Token authentication

The default for all interfaces is token authentication  

```xml
<api name="add" func="add" table="article" method="post" desc="添加文章">
	<param name="article" type="$article" desc="文章信息"/>
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

```go
// ArticleRoutes defined
func ArticleRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/article/add", Auth("token"), ArticleInstance.Add)
}
```

```sh
curl -X POST \
  http://127.0.0.1:8081/api/article/add \
  -H 'authorization: Bearer MAGFJEYNMPI9XAIK8GQKLA' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 458a9045-f7bb-01e1-8369-afb4dd303607' \
  -d '{
	"name": "info",
}'
```

#### Encrypt authentication

In addition to token authentication, a symmetrical encryption is built in

```xml
<api name="add" desc="Add Encrypt" method="post" auth="encrypt">
	<param name="encrypt_info" type="$encrypt_info" desc="Encrypt info" />
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

```go
// EncryptRoutes defined
func EncryptRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/encrypt/add", Auth("encrypt"), EncryptInstance.Add)
}
```

```sh
curl -X POST \
  'http://localhost:8082/api/encrypt/add?age=3' \
  -H 'authorization: Bearer ZH6_VZQBNTUPYSXR88AZXA' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 458a9045-f7bb-01e1-8369-afb4dd303607' \
  -d '{
"app_id": "Y76U9344RABF4",
"timestamp": 1614734783,
"biz_content": "a=1&b=2",
"sign": "b26sMTyLgK5IaPq0LYBYkPPWmz8BCovGiMfCH1sMZPudS9TFILB7O5eT924s60ALPINHRrSoyuo6rOHNXz0LJA=="
}'
```

# MIT License

Copyright (c) 2018-2022 Double

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
