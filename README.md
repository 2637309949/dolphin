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
    - [CMD](#cmd)
        - [build](#build)
        - [clean](#clean)
        - [init](#init)
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
    - [OAuth Server](#oauth-server)
        - [redirect sso](#redirect-sso)
        - [sso auth](#sso-auth)
        - [sso affirm](#sso-affirm)
        - [sso token](#sso-token)
        - [sso callback](#sso-callback)
    - [Workload](#workload)
        - [Add Handler](#add-handler)
        - [Add Job](#add-job)
        - [Fetch Job status](#fetch-job-status)

<!-- /TOC -->

## Quick start

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
dolphin init
```

## API Examples
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
func (srv *MessageSrv) SendMail(ctx context.Context, in *proto.MessageMail) (*proto.MessageReply, error) {
	return &proto.MessageReply{}, nil
}
```

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
	pApp.App.Manager.Worker().AddJobHandler("hello", func(args pModel.Worker) (interface{}, error) {
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
Host: 127.0.0.1:8082
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
Host: 127.0.0.1:8082
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