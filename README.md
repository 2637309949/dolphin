# Dolphin, Go code generate Framework

<img align="right" width="159px" src="./assets/dolphin.jpeg">

Dolphin is a code generate tools and web Framework written in Go (Golang)

## Quick start

1. The first need [Go](https://golang.org/) installed, then you can use the below Go command to install Dolphin.
```sh
$ go get -u github.com/2637309949/dolphin/cmd/dolphin
```

2. Create project dir and run dolphin

```sh
$ mkdir example && cd example && dolphin init && dolphin build && go run main
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

- Log trace record
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
<controller name="activity" desc="微信活动">
</controller>
```

controller
| LabelName   |      LabelMeaning      |
|----------|:-------------:|
| name |  controller name |
| desc |    controller desc  |
| prefix |    controller desc  |


### api
> api, api func in controller. we has some built-in func such as 'add', 'delete', 'update', 'page', 'get', 'tree', or you can refined if you need.

Example: 

```xml
<api name="list" func="page" table="table" desc="desc" method="get" roles="X3ed" cache="60">
    <param name="page" type="int" desc="页码" value="1" />
    <param name="size" type="int" desc="单页数" value="20" />
    <return>
        <success type="$response" />
        <failure type="$response" />
    </return>
</api>
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
> table, as you khnow, you can defined any table structure as you need

Example: 

```xml
<table name="activity" desc="活动" packages="xx">
	<column name="id" desc="主键" type="string" xorm="varchar(36) notnull unique pk" />
	<column name="title" desc="标题" type="xx.String" xorm="varchar(36)" />
</table>
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
