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

### <application/>

```xml
<?xml version="1.0" encoding="utf-8" ?>
<application name="demo" desc="template" packagename="demo"/>
```

### <bean/>

```xml
<bean name="activity_info" desc="desc" packages="xxx" extends="$applet_activity">
    <prop name="code" desc="编码" type="xx.String" />
    <prop name="name" desc="名称" type="xx.String" />
</bean>
```

### <controller/>

```xml
<controller name="activity" desc="微信活动">
</controller>
```

### <api/>

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

### <table/>

```xml
<table name="activity" desc="活动" packages="xx">
	<column name="id" desc="主键" type="string" xorm="varchar(36) notnull unique pk" />
	<column name="title" desc="标题" type="xx.String" xorm="varchar(36)" />
</table>
```