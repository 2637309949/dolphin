# Dolphin, Go code generate Framework

<img align="right" width="159px" src="./assets/dolphin.jpeg">

Dolphin is a code generate tools and web Framework written in Go (Golang), Will reduce the repetitive workload of adding, deleting, revising, and conducting inspections


# Contents
<!-- TOC -->

- [Dolphin, Go code generate Framework](#dolphin-go-code-generate-framework)
- [Contents](#contents)
    - [Quick start](#quick-start)
    - [Directory structure](#directory-structure)
    - [Features](#features)
    - [API Examples](#api-examples)
    - [XML Label](#xml-label)
        - [application](#application)
        - [bean](#bean)
        - [controller](#controller)
        - [api example](#api-example)
            - [add](#add)
            - [delete](#delete)
            - [update](#update)
            - [page](#page)
            - [tree](#tree)
            - [one](#one)
            - [other](#other)
        - [table](#table)
        - [rpc](#rpc)
    - [Domain](#domain)
        - [app_name](#app_name)
        - [domain](#domain)

<!-- /TOC -->

## Quick start

1. The first need [Go](https://golang.org/) installed, then you can use the below Go command to install Dolphin.
```sh
$ go get -u github.com/2637309949/dolphin/cmd/dolphin
```

2. Create project dir and run dolphin

```sh
$ mkdir example && cd example && dolphin init && dolphin build && go run main
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
```

## API Examples
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
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.DelFlag = null.IntFrom(0)
	payload.AppName = null.StringFrom(viper.GetString("app.name"))
	ret, err := ctx.PlatformDB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

#### delete
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
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

#### update
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
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.PlatformDB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
```

#### page
```xml
<api name="page" func="page" table="sys_client" desc="客户端分页查询" method="get">
    <param name="page" type="int" value="1" desc="页码"/>
    <param name="size" type="int"  value="15" desc="单页数"/>
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
	q.SetInt("size", 15)
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
```xml
<api name="page" func="page" table="sys_menu" desc="菜单分页查询" method="get">
	<param name="page" type="int" value="1" desc="页码"/>
	<param name="size" type="int"  value="15" desc="单页数"/>
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

#### one
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

Example: 

```xml
<table name="article" desc="文章" packages="github.com/2637309949/dolphin/packages/null">
	<column name="id" desc="主键" type="null.String" xorm="varchar(36) notnull unique pk" />
	<column name="type" desc="类别" type="null.String" xorm="varchar(36)" />

	<column name="create_by" desc="创建人" type="null.String" xorm="varchar(36)" />
	<column name="create_time" desc="创建时间" type="null.Time" xorm="datetime" />
	<column name="update_by" desc="最后更新人" type="null.String" xorm="varchar(36)" />
	<column name="update_time" desc="最后更新时间" type="null.Time" xorm="datetime" />
	<column name="del_flag" desc="删除标记" type="null.Int" xorm="notnull" />
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
	CreateBy null.String `xorm:"varchar(36) 'create_by'" json:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime 'create_time'" json:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) 'update_by'" json:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime 'update_time'" json:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull 'del_flag'" json:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" xml:"remark"`
}

// TableName table name of defined Article
func (m *Article) TableName() string {
	return "article"
}
```

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

```sh
rpc/proto/message.proto
rpc/proto/message.pb.go
rpc/message.go
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
	<column name="sync_flag" type="null.Int" xorm="notnull" />

	<column name="create_by" type="null.String" xorm="varchar(36) notnull" />
	<column name="create_time" type="null.Time" xorm="datetime notnull" />
	<column name="update_by" type="null.String" xorm="varchar(36) notnull" />
	<column name="update_time" type="null.Time" xorm="datetime notnull" />
	<column name="del_flag" type="null.Int" xorm="notnull" />
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
	if ctx.DB = ctx.engine.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
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