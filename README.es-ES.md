# Dolphin, una herramienta de generación de código para desarrolladores web
[![2637309949](https://circleci.com/gh/2637309949/dolphin.svg?style=shield)](https://circleci.com/gh/2637309949/dolphin)
[![2637309949](https://img.shields.io/github/release/2637309949/dolphin.svg?style=flat-square)](https://github.com/2637309949/dolphin/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/2637309949/dolphin?status.svg)](https://pkg.go.dev/github.com/2637309949/dolphin?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/2637309949/dolphin)](https://goreportcard.com/report/github.com/2637309949/dolphin)

<img align="right" width="159px" src="https://hub.fastgit.org/2637309949/dolphin/blob/master/assets/dolphin.jpeg">

Dolphin, una herramienta de generación de código en Golang basada en gin y xormplus, puede cohesionar rápidamente marcos de funciones incógnitos y reducir eficazmente la lógica de código repetitiva.
  
<img align="center" width="220px" src="https://hub.fastgit.org/2637309949/dolphin/blob/master/assets/dolphin-ui.jpeg">
<img align="center" width="220px" src="https://hub.fastgit.org/2637309949/dolphin/blob/master/assets/docs.png">

### Inicio rápido

- Instalar dolphin
```sh
$ go get -u github.com/2637309949/dolphin/cmd/dolphin
$ dolphin 
Code generation tool for golang

Usage:
  dolphin [command]

Available Commands:
  assert      Generate Go code that statically implements input filesystem
  build       Build from the configuration file
  clean       Remove temp file, such as *.go.new
  help        Help about any command
  more        Add controller and table
  new         New a empty project
  reverse     Inversion of the data model
  serve       Serve api document

Flags:
  -h, --help   help for dolphin

Use "dolphin [command] --help" for more information about a command.
```

- Nuevo proyecto

```sh
$ dolphin new hello && cd hello && dolphin build && go run main.go
```

## Características

```
- Genera código basado en configuración XML

- Genera documentación basada en configuración XML

- Genera SQL basado en configuración XML

- Maneja el problema de nulidad en la serialización

- Soporte multi-inquilino (multi-tenant)

- Inicio de sesión / Cierre de sesión, o inicio de sesión único (SSO)

- Autenticación de permisos

- Reportes o análisis rápidos de Excel

- Soporte para caché de rutas

- Control de permisos de datos

- Registro de trazabilidad de logs (trace record)

- Servicio remoto RPC

- El archivo de despliegue de k8S se genera por defecto

- Soporte para generación de XML inversa desde la base de datos
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
