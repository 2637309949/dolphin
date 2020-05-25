// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package parser

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// AppParser defined AppParser struct
type AppParser struct {
	*schema.Application
}

// New defined AppParser
func New() *AppParser {
	return &AppParser{}
}

// Parse defined parse xml to Application
func (parser *AppParser) parse(xmlPath string) error {
	reader, err := os.Open(xmlPath)
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(reader)
	var controller *schema.Controller
	var api *schema.API
	var param *schema.Param
	var bean *schema.Bean
	var prop *schema.Prop
	var table *schema.Table
	var column *schema.Column
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			switch true {
			case token.Name.Local == "application":
				parser.parseApplication(xmlPath, token.Attr)
			case token.Name.Local == "controller":
				controller = parser.parseController(xmlPath, token.Attr)
			case token.Name.Local == "api":
				api = parser.parseAPI(xmlPath, token.Attr)
			case token.Name.Local == "success":
				parser.parseSuccess(xmlPath, token.Attr, api)
			case token.Name.Local == "failure":
				parser.parseFailure(xmlPath, token.Attr, api)
			case token.Name.Local == "param":
				param = parser.parseParam(xmlPath, token.Attr)
			case token.Name.Local == "bean":
				bean = parser.parseBean(xmlPath, token.Attr)
			case token.Name.Local == "prop":
				prop = parser.parseProp(xmlPath, token.Attr)
			case token.Name.Local == "table":
				table = parser.parseTable(xmlPath, token.Attr)
			case token.Name.Local == "column":
				column = parser.parseColumn(xmlPath, token.Attr)
			}
		case xml.EndElement:
			switch true {
			case token.Name.Local == "controller":
				parser.Application.Controllers = append(parser.Application.Controllers, controller)
			case token.Name.Local == "api":
				controller.APIS = append(controller.APIS, api)
			case token.Name.Local == "param":
				api.Params = append(api.Params, param)
			case token.Name.Local == "bean":
				parser.Application.Beans = append(parser.Application.Beans, bean)
			case token.Name.Local == "prop":
				bean.Props = append(bean.Props, prop)
			case token.Name.Local == "table":
				parser.Application.Tables = append(parser.Application.Tables, table)
			case token.Name.Local == "column":
				table.Columns = append(table.Columns, column)
			}
			// case xml.CharData:
			// 	content := string([]byte(token))
			// 	if strings.TrimSpace(content) != "" {
			// 		logrus.Warn(fmt.Sprintf("xml.CharData:%v", content))
			// 	}
			// default:
		}
	}
	return nil
}

// Marshal marshal XML to string
func (parser *AppParser) Marshal() []byte {
	data, err := json.Marshal(parser.Application)
	if err != nil {
		panic(err)
	}
	return data
}

// MarshalIndent marshal XML to string
func (parser *AppParser) MarshalIndent(prefix, indent string) []byte {
	data, err := json.MarshalIndent(parser.Application, prefix, indent)
	if err != nil {
		panic(err)
	}
	return data
}

// Walk parse all xml in directory
func (parser *AppParser) Walk(xmlPath string) error {
	var files []string
	if err := filepath.Walk(xmlPath, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".xml") {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return err
	}
	parser.Application = &schema.Application{}
	for _, v := range files {
		if err := parser.parse(v); err != nil {
			return err
		}
	}
	err := validate.Struct(parser.Application)
	if err != nil {
		return err
	}
	return nil
}
