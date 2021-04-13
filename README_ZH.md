# Dolphin 一个Go代码生成框架
[![2637309949](https://circleci.com/gh/2637309949/dolphin.svg?style=shield)](https://circleci.com/gh/2637309949/dolphin)
[![2637309949](https://img.shields.io/github/release/2637309949/dolphin.svg?style=flat-square)](https://github.com/2637309949/dolphin/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/2637309949/dolphin?status.svg)](https://pkg.go.dev/github.com/2637309949/dolphin?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/2637309949/dolphin)](https://goreportcard.com/report/github.com/2637309949/dolphin)

<img align="right" width="159px" src="./assets/dolphin.jpeg">

Dolphin是一个使用Golang代码生成工具, 用于避免日常代码增删改查等重复型的轮子从而减少工具量, 把时间专注在业务层的设计.

# 内容
<!-- TOC -->

- [Dolphin 一个Go代码生成框架](#dolphin-一个go代码生成框架)
- [内容](#内容)
    - [快速入手](#快速入手)
    - [目录结构](#目录结构)
    - [特性](#特性)
    - [命令](#命令)
        - [build](#build)
        - [clean](#clean)
        - [init](#init)
    - [API Examples](#api-examples)
        - [dolphin-ui](#dolphin-ui)
    - [XML 标签](#xml-标签)
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
- [MIT License](#mit-license)

<!-- /TOC -->
## 快速入手

<img width="580px" src="./assets/dolphin.gif">      


1. 首先安装Golang环境 [Go](https://golang.org/), 接着使用下面的命令安装Dolphin
```sh
$ go get -u github.com/2637309949/dolphin/cmd/dolphin
```

2. 使用dolphin命令创建项目

```sh
$ mkdir example && cd example && dolphin init && dolphin build && go run main.go
```

输出:
```sh
time="2020/06/13 11:55:58" level=info msg="grpc listen on port:9081"
time="2020/06/13 11:55:58" level=info msg="http listen on port:8082"
```

## 目录结构

> 准目录结构如下面所示, 这是简化且推荐的目录结构, 当需要使用dolphin管理大型项目时可以复用这个结构去创建更多的子项目.


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

## 特性


```
- 基于XML配置生成代码

- 基于XML配置生成DOC文档

- 基于XML配置生成SQL

- 默认处理NULL空值问题

- 支持多租户分库

- 登录/退出, 或者单点登录

- 支持路由权限认证

- 快速EXCEL报表和EXCEL解析

- 支持路由缓存

- 支持数据权限控制

- 支持日子记录

- 支持RPC远程服务

- 默认生成K8s部署配置

- 支持基于数据库生成XML配置
```

## 命令


### build

build是执行一切内置的步骤命令, 如果需要只执行某个步骤, 例如生成DOC, 你可以使用@并带上内置函数

例如生成XML
```sh
dolphin build @table xml/test
```

已存在内置的函数列表:
| Function   |      Action      |
|----------|:-------------:|
| main | 创建main函数 |
| app | 创建引擎app.go源码 |
| ctr | 创建控制器 |
| proto | 创建PROTO |
| srv | 生成SRV层代码 |
| model | 创建模型源码 |
| bean | 创建BEAN函数 |
| auto | 自动生成注册函数 |
| tool | 创建TOOL工具 |
| sql | 转换SQL成GO变量 |
| sqlmap | 创建SQL文件 |
| oauth | 生成H5认证模板 |
| script | 创建前端JSAPI代码 |
| deploy | 创建k8s部署模板 |
| doc | 生成swagger文档 |
| table | 从数据库链接生成table配置 |

### clean

清理所有临时文件

```sh
dolphin clean
```


### init

初始化命令, 创建一系列初始化文件


```sh
mkdir demo && cd demo && dolphin init
```

## API Examples

### dolphin-ui 

> 一个与dolphin对接的ERP模板

<img align="right" width="200px" src="./assets/dolphin-ui.png">

你可以找到一系列已经准备好的demo [dolphin examples repository.](https://github.com/2637309949/dolphin-ui)


## XML 标签

### application
> application 标签包含app的信息, 例如名字,包名

示列: 

```xml
<?xml version="1.0" encoding="utf-8" ?>
<application name="demo" desc="template" packagename="demo"/>
```

application

| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  必填, 应用名称 |
| desc |   app的描述  |
| packagename |    必填, 应用包名 |

### bean

> 通过bean标签你可以声明对象, 就像spring的bean一样, 所有的bean和模型都会放在model目录下, 所以你需要避免因名字冲突的情况(通常按模块创建不同的子项目bean的数量会控制在合理范围内)

示例: 

```xml
<bean name="activity_info" desc="desc" packages="xxx" extends="$applet_activity">
    <prop name="code" desc="编码" type="xx.String" />
    <prop name="name" desc="名称" type="xx.String" />
</bean>
```

生成代码:

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
| name |  名称 |
| desc |   描述  |
| packagename | 第三方包名, 使用,分割 |
| extends |   bean继承  |


prop
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name | 属性名称 |
| desc | 属性描述 |
| type | 属性类型 |

### controller

> 控制器, 一个API集合, 你可以声明api前缀等

示范: 

```xml
<controller name="activity" desc="微信活动" />
```

controller
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  控制器名称 |
| desc |    描述  |
| prefix |    控制器前缀  |

### api example

> 在api标签中, 你可以定义函数, 而且dolphin包含了一些内置的函数, 例如: 'add', 'delete', 'update', 'page', 'get', 'tree' 或者你可以重新定义

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

生成的代码:

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

> 当然, 你可以指定数组参数形式以生成批量新增的模板

```xml
<api name="batch_add" func="add" table="sys_role_menu" method="post" desc="添加角色菜单">
	<param name="role_menu" type="[]$sys_role_menu" desc="角色菜单信息" />
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```


生成的代码:  

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


> 系统默认的删除是软删除, 这也是极力推荐的. 你可以通过修改模板实现硬删除的功能

```xml
<api name="del" func="delete" table="sys_client" desc="删除客户端" method="delete">
    <param name="sys_client" type="$sys_client" desc="客户端" />
    <return>
        <success type="$success"/>
        <failure type="$fail"/>
    </return>
</api>
```

生成代码:

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

> 当然, 你可以指定数组参数形式以生成批量删除的模板

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

> 整个模型的设计是默认支持NULL类型的, 所以不需要担心序列化控制更新问题


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

> 当然, 你可以指定数组参数形式以生成批量更新的模板

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
    s.Begin()
    defer s.Close()
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

> 分页是最常见的接口, dolphin系统默认是支持的


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

生成代码:

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

> 树形搜索是常见的算法, dolphin也提供了一种树形的接口模板

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

生成代码:

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
> 详细介绍可以在TreeSearch的代码中看到

```go
// platform/app/app.ctx.go#TreeSearch
func (ctx *Context) TreeSearch(db *xorm.Engine, controller, api, table string, 
		q map[string]interface{}) (interface{}, error)
```
#### one

> 通过条件查找一条数据

```xml
<api name="get" func="one" table="sys_client" desc="获取客户端信息" method="get">
    <param name="id" type="string" desc="客户端id" />
    <return>
        <success type="$success"/>
        <failure type="$fail"/>
    </return>
</api>
```

生成代码:

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

> 另外, 我们可以自定自已的接口

```xml
<api name="payment" method="post" desc="文章付费">
	<param name="article" type="$article_info" desc="文章"/>
	<return>
		<success type="$success"/>
		<failure type="$fail"/>
	</return>
</api>
```

生成代码:

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
| name |  API名称 |
| desc |    描述   |
| func | built-in func, 'add', 'delete', 'update', 'page', 'get' |
| table |    如果你使用内置函数则需要指定table   |
| method |    http 方法   |
| roles |    roles 插件   |
| cache |    cache 中间件   |

param
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  参数名称 |
| desc |    参数描述  |
| type | 参数类型 |
| value | 默认值 |

return

| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| success |  success tag |
| failure |    success tag  |

### table
> 通过table标签你可以定义任何表结构, 同时如果想要接受允许空值类型的数据,你应该使用null类型从而避免golang的空值更新问题.


table
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  table 名称 |
| desc |    table 描述   |
| packages | 第三方包名，可以使用 "," 分割 |

column
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  字段名称 |
| desc |    字段描述  |
| type | 字段类型 |
| xorm | xorm标记, 可以参考xormplus的文档 |


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

生成代码:

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

> 系统默认添加了null和decimal包, null是用来支持默认值更新问题, 而decimal是解决金融数据精度问题.  

示例: 
```xml
<table 
	name="article" 
	desc="文章" 
	packages="github.com/2637309949/dolphin/packages/null,
	github.com/shopspring/decimal">
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

> 在model下面写相应的helper函数是很推荐的.

示例: 
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

> rpc, 作为微服定义, 基本的proto文件会在rpc目录下生成, 同时会被auto文件中自动注册.


```xml
<service name="message" desc="消息">
    <rpc name="send_mail" desc="发送邮件">
        <request type="$article" desc="文章信息"/>
        <reply type="$success" desc="文章信息"/>
    </rpc>
</service>
```

生成代码:

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

> 自动生成rpc服务

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

> rpc客户端也是自动生成的

rpc cli endpoint:
```go
// Code generated by dol build. Only Generate by tools if not existed.
// source: MessageSrv.cli.go

package rpc

import (
	"demo/rpc/proto"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

> Domain, 一个支持多组户的模型, 应用分割也是支持.

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

> app_name，将模型取消同步为标记。如果从本地主机连接相同的数据源 URL，并且您会发现将创建`data_source`app_name的所有表

### domain

> 识别不同的租户，登录的用户将使用匹配的域来查找数据库

如下代码所示

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

> 默认情况下，继承平台的所有项目都支持单点登录，您可以独立部署或直接部署为 SSO 服务器

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

平台中的代码段，从cookie获取状态.

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

平台中的代码段，身份验证逻辑:

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

在平台的代码段，您可以重写这种方式，如果你想跳过确认

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

平台中的代码段，按代码生成令牌

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

客户端中的代码段，从平台获取令牌并设置 Cookie

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

有关详细信息，请查看 dolphin-ui [dolphin-ui](https://github.com/2637309949/dolphin-ui)  

如果你想去 sso auth， 那么你可以设置 `auth_mode` 为 `1`  

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

> 使用内置负载接口处理高并发请求

### Add Handler

```go
// @Summary AddJobHandler
// @Tags worker
func (d *DefaultWorker) AddJobHandler(code string, funk func(model.Worker) (interface{}, error)) {
```

例子:  
demo/srv/article.go  
```go
func init() {
	// add hello topic handler
	app.App.Manager
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

例子:  

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

例子:  
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

默认情况下未启用 Cron，如果需要，可以通过修改配置`app.cron = true`来做到这一点。但是 Cron 是有状态的，因此在应用负载时必须单独部署它。

### Add Func

通过 AddFunc 添加计时器任务.

```go
AddFunc(string, func()) (int, error)
```

```go
app.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})
```

### Reflesh Func

通过 RefreshFunc 重新刷新计时器任务.

```go
RefreshFunc(int, string) (int, error)
```

```go
id, _ := app.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})
app.App.Manager.Cron().RefreshFunc(id, "*/3 * * * * *")
```

### Del Func

通过 DelFunc 删除计时器任务.

```go
DelFunc(int) error
```

```go
id, _ := app.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})

app.App.Manager.Cron().DelFunc(id)
```

### Try Func

尝试通过 TryFunc 的计时器任务.

```go
TryFunc(int) error
```

```go
id, _ := app.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
	fmt.Println("hello")
})
app.App.Manager.Cron().TryFunc(id)
```

## Load User Info

由于用户信息 db 和主业务 db 是分开的，因此需要单独加载用户信息

```go
ctx.PlatformDB 
```

可以从"平台数据库"加载平台信息，例如sys_user、sys_client sys_domain

```go
ctx.DB 
```
否则，您应该从 ctx 加载业务信息ctx.DB 

### GetFieldSliceByName

基于字段名称从数组内部创建新数组。

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

将字段值修补到另一个数组.

示范:

```go
err = slice.PatchSliceByField(ret.Data, roles, "id", "user_id", "role_name")(&ret.Data)
if err != nil {
	logrus.Error(err)
	ctx.Fail(err)
	return
}
```
## Debug pprof

> 框架本身提供了 pprof 的功能，请注意，只有管理员角色可见

### ModHeader

在镶边中安装 ModHeader，并添加任何请求的令牌头.

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

# MIT License

Copyright (c) 2018-2022 Double

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
