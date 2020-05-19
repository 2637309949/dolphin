package app

import (
	"fmt"
	"path"

	"github.com/2637309949/dolphin/packages/excelize"
	"github.com/2637309949/dolphin/packages/uuid"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
)

// ExcelConfig defined
type ExcelConfig struct {
	FileName  string
	SheetName string
	Rows      []map[string]interface{}
	Header    []map[string]interface{}
	TmplPath  string
	Format    func(interface{}) func(interface{}) interface{}
	Handler   func(map[string]interface{}) map[string]interface{}
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
		optionsets, err := (&Context{}).GetOptions(db, fmt.Sprintf("%v", format))
		return func(source interface{}) interface{} {
			if err == nil {
				target := optionsets[fmt.Sprintf("%v", format)]
				if target != nil {
					for k, v := range target {
						if fmt.Sprintf("%v", source) == fmt.Sprintf("%v", v) {
							return k
						}
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

// NewExcelConfig defined
func NewExcelConfig(rows []map[string]interface{}, header ...[]map[string]interface{}) ExcelConfig {
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

// BuildExcel defined
func BuildExcel(cfg ExcelConfig) (model.ExportInfo, error) {
	f := excelize.NewFile()
	uuid := uuid.MustString()
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
						value := v1["prop"]
						if v1["formatter"] != nil {
							if formats[v1["formatter"]] == nil {
								formats[v1["formatter"]] = cfg.Format(v1["formatter"])
							}
							value = formats[v1["formatter"]](datav[fmt.Sprintf("%v", value)])
						}
						cells = append(cells, value)
					}
				}
			}
			f.SetSheetRow(cfg.SheetName, fmt.Sprintf("A%v", i+2), &cells)
			cells = []interface{}{}
		}
	}
	err := f.SaveAs(filePath)
	return model.ExportInfo{FileId: uuid, FilePath: filePath}, err
}
