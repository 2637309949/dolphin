package parser

import (
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func (parser *AppParser) parseApplication(xmlPath string, attr []xml.Attr) {
	if parser.Application == nil {
		parser.Application = &Application{}
	}
	parser.Application.Path = xmlPath
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			parser.Application.Name = attrValue
		case attrName == "desc":
			parser.Application.Desc = attrValue
		case attrName == "packagename":
			parser.Application.PackageName = attrValue
		}
	}
}

func (parser *AppParser) parseController(xmlPath string, attr []xml.Attr) *Controller {
	controller := &Controller{}
	controller.Path = xmlPath
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			controller.Name = attrValue
		case attrName == "desc":
			controller.Desc = attrValue
		case attrName == "prefix":
			controller.Prefix = attrValue
		}
	}
	return controller
}

func (parser *AppParser) parseAPI(xmlPath string, attr []xml.Attr) *API {
	api := &API{Auth: []string{viper.GetString("app.auth")}, Return: &Return{Success: &Success{}, Failure: &Failure{}}}
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			api.Name = attrValue
		case attrName == "desc":
			api.Desc = attrValue
		case attrName == "method":
			api.Method = attrValue
		case attrName == "func":
			api.Func = attrValue
		case attrName == "table":
			api.Table = attrValue
		case attrName == "version":
			api.Version = attrValue
		case attrName == "path":
			api.Path = attrValue
		case attrName == "roles":
			api.Roles = strings.Split(attrValue, ",")
		case attrName == "cache":
			reg := regexp.MustCompile("^([0-9]*)([smh])?$")
			base := reg.FindAllStringSubmatch(attrValue, -1)
			sed, _ := strconv.ParseUint(base[0][1], 10, 64)
			switch base[0][2] {
			case "m":
				sed = sed * 60
			case "h":
				sed = sed * 60 * 60
			}
			api.Cache = sed
		case attrName == "auth":
			api.Auth = strings.Split(attrValue, ",")
			keys := map[string]interface{}{}
			for i := range api.Auth {
				keys[api.Auth[i]] = ""
			}
			api.Auth = []string{}
			for k := range keys {
				if k != "" && k != "never" {
					api.Auth = append(api.Auth, k)
				}
			}
		}
	}
	return api
}

func (parser *AppParser) parseSuccess(xmlPath string, attr []xml.Attr, api *API) {
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "type":
			api.Return.Success.Type = attrValue
		}
	}
}

func (parser *AppParser) parseFailure(xmlPath string, attr []xml.Attr, api *API) {
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "type":
			api.Return.Failure.Type = attrValue
		}
	}
}

func (parser *AppParser) parseParam(xmlPath string, attr []xml.Attr) *Param {
	param := &Param{}
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			param.Name = attrValue
		case attrName == "desc":
			param.Desc = attrValue
		case attrName == "type":
			param.Type = attrValue
		case attrName == "value":
			param.Value = attrValue
		}
	}
	return param
}

func (parser *AppParser) parseRPC(xmlPath string, attr []xml.Attr) *RPC {
	rpc := &RPC{Request: &Request{}, Reply: &Reply{}}
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			rpc.Name = attrValue
		case attrName == "desc":
			rpc.Desc = attrValue
		}
	}
	return rpc
}

func (parser *AppParser) parseRequest(xmlPath string, attr []xml.Attr, rpc *RPC) {
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "desc":
			rpc.Request.Desc = attrValue
		case attrName == "type":
			rpc.Request.Type = attrValue
		}
	}
}

func (parser *AppParser) parseReply(xmlPath string, attr []xml.Attr, rpc *RPC) {
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "desc":
			rpc.Reply.Desc = attrValue
		case attrName == "type":
			rpc.Reply.Type = attrValue
		}
	}
}

func (parser *AppParser) parseBean(xmlPath string, attr []xml.Attr) *Bean {
	bean := &Bean{}
	bean.Path = xmlPath
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			bean.Name = attrValue
		case attrName == "desc":
			bean.Desc = attrValue
		case attrName == "packages":
			bean.Packages = attrValue
		case attrName == "extends":
			bean.Extends = attrValue
		}
	}
	return bean
}

func (parser *AppParser) parseProp(xmlPath string, attr []xml.Attr) *Prop {
	prop := &Prop{}
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			prop.Name = attrValue
		case attrName == "desc":
			prop.Desc = attrValue
		case attrName == "type":
			prop.Type = attrValue
		case attrName == "json":
			prop.JSON = attrValue
		case attrName == "form":
			prop.Form = attrValue
		case attrName == "example":
			prop.Example = attrValue
		}
	}
	return prop
}

func (parser *AppParser) parseTable(xmlPath string, attr []xml.Attr) *Table {
	table := &Table{}
	table.Path = xmlPath
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			table.Name = attrValue
		case attrName == "bind":
			table.Bind = attrValue
		case attrName == "desc":
			table.Desc = attrValue
		case attrName == "packages":
			table.Packages = attrValue
		case attrName == "extends":
			table.Extends = attrValue
		}
	}
	return table
}

func (parser *AppParser) parseColumn(xmlPath string, attr []xml.Attr) *Column {
	column := &Column{}
	for _, attr := range attr {
		attrName := attr.Name.Local
		attrValue := attr.Value
		if strings.TrimSpace(attrValue) == "" {
			continue
		}
		switch true {
		case attrName == "name":
			column.Name = attrValue
		case attrName == "desc":
			column.Desc = attrValue
		case attrName == "xorm":
			column.Xorm = attrValue
		case attrName == "type":
			column.Type = attrValue
		case attrName == "json":
			column.JSON = attrValue
		case attrName == "example":
			column.Example = attrValue
		}
	}
	return column
}
