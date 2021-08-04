# Dolphin, a code generation tool for web developer
[![2637309949](https://circleci.com/gh/2637309949/dolphin.svg?style=shield)](https://circleci.com/gh/2637309949/dolphin)
[![2637309949](https://img.shields.io/github/release/2637309949/dolphin.svg?style=flat-square)](https://github.com/2637309949/dolphin/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/2637309949/dolphin?status.svg)](https://pkg.go.dev/github.com/2637309949/dolphin?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/2637309949/dolphin)](https://goreportcard.com/report/github.com/2637309949/dolphin)

<img align="right" width="159px" src="https://hub.fastgit.org/2637309949/dolphin/blob/master/assets/dolphin.jpeg">

Dolphin, a Golang code generation tool, based on the gin and xormplus, can quickly cohesion incognito function framework, and effectively reduce repeat code logic.
  
<img align="center" width="220px" src="https://hub.fastgit.org/2637309949/dolphin/blob/master/assets/dolphin-ui.jpeg">
<img align="center" width="220px" src="https://hub.fastgit.org/2637309949/dolphin/blob/master/assets/docs.png">

### Quick start

- Install dolphin
```sh
$ go get -u github.com/2637309949/dolphin/cmd/dolphin
$ dolphin 
Code generation tool for golang

Usage:
  dolphin [command]

Available Commands:
  build       Build from the configuration file
  clean       Remove temp file, such as *.go.new
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  more        Add controller and table
  new         New a empty project
  reverse     Inversion of the data model
  serve       Serve api document

Flags:
  -h, --help   help for dolphin

Use "dolphin [command] --help" for more information about a command.
```

- New project

```sh
$ dolphin new hello && cd hello && dolphin build && go run main.go
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

## Tutorial

- [快速入门](./docs/快速入门.md)
- [配置XML](./docs/配置XML.md)
- [数据权限](./docs/数据权限.md)
- [全局配置](./docs/全局配置.md)
- [内置指令](./docs/内置指令.md)
- [远程调用](./docs/远程调用.md)
- [集成平台](./docs/集成平台.md)
- [用户认证](./docs/用户认证.md)
- [单点认证](./docs/单点认证.md)
- [容器平台](./docs/容器平台.md)
- [微服架构](./docs/微服架构.md)
- [服务网格](./docs/服务网格.md)
- [解放服务](./docs/解放服务.md)
- [链路跟踪](./docs/链路跟踪.md)

# MIT License

Copyright (c) 2018-2022 Double

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
