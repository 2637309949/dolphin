// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// ExcelConfig defined
type ExcelConfig struct {
	FileReader io.Reader
	SheetIndex interface{}
	FileName   string
	SheetName  string
	Rows       []map[string]interface{}
	Header     []map[string]interface{}
	TmplPath   string
	Format     func(interface{}) func(interface{}) interface{}
	Handler    func(map[string]interface{}) map[string]interface{}
}

// AddRows defined
func (e *ExcelConfig) AddRows(rows ...map[string]interface{}) *ExcelConfig {
	e.Rows = append(e.Rows, rows...)
	return e
}

// DefaultFormat defined
func DefaultFormat(interface{}) func(interface{}) interface{} {
	return func(source interface{}) interface{} {
		return source
	}
}

// OptionsetsFormat defiend
func OptionsetsFormat(db *xorm.Engine) func(interface{}) func(interface{}) interface{} {
	return func(format interface{}) func(source interface{}) interface{} {
		optionsets, err := new(Context).GetOptions(db, fmt.Sprintf("%v", format))
		return func(source interface{}) interface{} {
			if err == nil {
				target := optionsets[fmt.Sprintf("%v", format)]
				for k, v := range target {
					if fmt.Sprintf("%v", source) == fmt.Sprintf("%v", v) {
						return k
					}
				}
			}
			return source
		}
	}
}

// DefaultHandler defined
func DefaultHandler(source map[string]interface{}) map[string]interface{} {
	return source
}

// NewBuildExcelConfig defined
func NewBuildExcelConfig(rows []map[string]interface{}, header ...[]map[string]interface{}) ExcelConfig {
	ecg := ExcelConfig{
		SheetName: "Sheet1",
		Rows:      rows,
		Format:    DefaultFormat,
		Handler:   DefaultHandler,
		TmplPath:  path.Join(viper.GetString("http.static"), "files"),
	}
	if len(header) > 0 {
		ecg.Header = header[0]
	}
	return ecg
}

// NewParseExcelConfig defined
func NewParseExcelConfig(file io.Reader, sheet interface{}, header ...[]map[string]interface{}) ExcelConfig {
	ecg := ExcelConfig{
		Format:  DefaultFormat,
		Handler: DefaultHandler,
	}
	ecg.FileReader = file
	ecg.SheetIndex = sheet
	if len(header) > 0 {
		ecg.Header = header[0]
	}
	return ecg
}

// BuildExcel defined
func BuildExcel(cfg ExcelConfig) (model.ExportInfo, error) {
	f := excelize.NewFile()
	uuid := uuid.New().String()
	index := f.NewSheet("Sheet1")
	filePath := path.Join(cfg.TmplPath, fmt.Sprintf("%v.xlsx", uuid))
	f.SetActiveSheet(index)
	formats := map[interface{}]func(interface{}) interface{}{}
	titles := []interface{}{}

	for i, datav := range cfg.Rows {
		datav = cfg.Handler(datav)
		cells := []interface{}{}
		// key as title
		if len(cfg.Header) == 0 {
			if len(titles) == 0 {
				for k := range datav {
					titles = append(titles, k)
				}
				// replace title
				f.SetSheetRow(cfg.SheetName, fmt.Sprintf("A%v", i+1), &titles)
			}
			for _, k := range titles {
				cells = append(cells, datav[fmt.Sprintf("%v", k)])
			}
			f.SetSheetRow(cfg.SheetName, fmt.Sprintf("A%v", i+2), &cells)
			cells = []interface{}{}
		} else {
			// key from header
			if len(titles) == 0 {
				for _, v := range cfg.Header {
					titles = append(titles, v["label"])
				}
				f.SetSheetRow(cfg.SheetName, fmt.Sprintf("A%v", i+1), &titles)
			}
			for _, k := range titles {
				// replace title
				for _, v1 := range cfg.Header {
					if v1["label"] == k {
						value := datav[fmt.Sprintf("%v", v1["prop"])]
						if v1["formatter"] != nil {
							if formats[v1["formatter"]] == nil {
								formats[v1["formatter"]] = cfg.Format(v1["formatter"])
							}
							value = formats[v1["formatter"]](value)
						}
						switch vt := value.(type) {
						case time.Time:
							value = vt.Format("2006-01-02 15:04:05")
						}
						cells = append(cells, value)
					}
				}
			}
			f.SetSheetRow(cfg.SheetName, fmt.Sprintf("A%v", i+2), &cells)
			cells = []interface{}{}
		}
	}
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return model.ExportInfo{}, err
	}
	err := f.SaveAs(filePath)
	return model.ExportInfo{FileId: uuid, FilePath: filePath}, err
}

// ParseExcel defined
func ParseExcel(cfg ExcelConfig) ([]map[string]string, error) {
	eFile, err := excelize.OpenReader(cfg.FileReader)
	if err != nil {
		return nil, err
	}
	sheetName := ""
	switch sn := cfg.SheetIndex.(type) {
	case int:
		sheetName = eFile.GetSheetName(sn)
	case string:
		sheetName = sn
	}
	rows := eFile.GetRows(sheetName)

	data := []map[string]string{}
	iTitle := map[int]string{}
	for ri, row := range rows {
		if ri == 0 {
			for ci, cell := range row {
				iTitle[ci] = cell
			}
			continue
		}
		r := map[string]string{}
		for ic, iv := range row {
			r[iTitle[ic]] = iv
		}
		data = append(data, r)
	}

	if len(cfg.Header) > 0 {
		nd := []map[string]string{}
		h := cfg.Header
		for _, dv := range data {
			ndItem := map[string]string{}
			for dk, dvv := range dv {
				for _, v := range h {
					if v["label"] == dk {
						ndItem[fmt.Sprintf("%v", v["prop"])] = dvv
					}
				}
			}
			nd = append(nd, ndItem)
		}
		return nd, nil
	}
	return data, nil
}
