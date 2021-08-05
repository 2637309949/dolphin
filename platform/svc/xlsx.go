package svc

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/file"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Xlsx defined TODO
type Xlsx struct {
	zipIndex   int
	zipPath    string
	zipFile    *zip.Writer
	xlsxFile   *excelize.File
	xlsxSize   int
	SheetIndex interface{}
	FileName   string
	SheetName  string
	Header     []map[string]interface{}
	TmplPath   string
	Format     func(interface{}) func(interface{}) interface{}
	Handler    func(map[string]interface{}) map[string]interface{}
}

// DumpRows defined TODO
func (xlsx *Xlsx) DumpRows(rows ...map[string]interface{}) *Xlsx {
	formats := map[interface{}]func(interface{}) interface{}{}
	titles := []interface{}{}

	for i, datav := range rows {
		datav = xlsx.Handler(datav)
		cells := []interface{}{}
		// key as title
		if len(xlsx.Header) == 0 {
			if len(titles) == 0 {
				for k := range datav {
					titles = append(titles, k)
				}
				// replace title
				xlsx.xlsxFile.SetSheetRow(xlsx.SheetName, fmt.Sprintf("A%v", i+1), &titles)
			}
			for _, k := range titles {
				cells = append(cells, datav[fmt.Sprintf("%v", k)])
			}
			xlsx.xlsxFile.SetSheetRow(xlsx.SheetName, fmt.Sprintf("A%v", i+2), &cells)
			cells = []interface{}{}
		} else {
			// key from header
			if len(titles) == 0 {
				for _, v := range xlsx.Header {
					titles = append(titles, v["label"])
				}
				xlsx.xlsxFile.SetSheetRow(xlsx.SheetName, fmt.Sprintf("A%v", i+1), &titles)
			}
			for _, k := range titles {
				// replace title
				for _, v1 := range xlsx.Header {
					if v1["label"] == k {
						value := datav[fmt.Sprintf("%v", v1["prop"])]
						if v1["formatter"] != nil {
							if formats[v1["formatter"]] == nil {
								formats[v1["formatter"]] = xlsx.Format(v1["formatter"])
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
			xlsx.xlsxFile.SetSheetRow(xlsx.SheetName, fmt.Sprintf("A%v", i+2), &cells)
			cells = []interface{}{}
		}
	}
	return xlsx.zipPacked(false)
}

// zipPacked defined TODO
func (xlsx *Xlsx) zipPacked(last ...bool) *Xlsx {
	isOver, isLast := xlsx.isSheetOverSize(), util.SomeOne(last, false).(bool)
	if isOver || isLast {
		xlsx.resetSheet()
	}
	return xlsx
}

// ExportInfo defined TODO
func (xlsx *Xlsx) ExportInfo() *types.ExportInfo {
	defer xlsx.resetZip()
	xlsx.zipPacked(true)
	item := types.ExportInfo{
		FileId:   file.RemoveExt(path.Base(xlsx.zipPath)),
		FilePath: xlsx.zipPath,
		FileName: xlsx.FileName,
	}
	return &item
}

// isSheetOverSize defined TODO
func (xlsx *Xlsx) isSheetOverSize() bool {
	return len(xlsx.xlsxFile.GetRows(xlsx.SheetName)) > xlsx.xlsxSize
}

// resetSheet defined TODO
func (xlsx *Xlsx) resetSheet() {
	w, _ := xlsx.zipFile.Create(fmt.Sprintf("%v.xlsx", xlsx.zipIndex))
	xlsx.xlsxFile.WriteTo(w)
	xlsx.xlsxFile.DeleteSheet(xlsx.SheetName)
	index := xlsx.xlsxFile.NewSheet(xlsx.SheetName)
	xlsx.xlsxFile.SetActiveSheet(index)
	xlsx.zipIndex += 1
}

// resetZip defined TODO
func (xlsx *Xlsx) resetZip() {
	xlsx.zipIndex = 1
	xlsx.zipFile.Close()
	xlsx.zipPath = ""
}

// ParseExcel defined TODO
func (xlsx *Xlsx) ParseExcel(r io.Reader, header ...[]map[string]interface{}) ([]map[string]string, error) {
	eFile, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	xlsx.Header = util.SomeOne(header, xlsx.Header).([]map[string]interface{})
	sheetName := ""
	switch sn := xlsx.SheetIndex.(type) {
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

	if len(xlsx.Header) > 0 {
		nd := []map[string]string{}
		h := xlsx.Header
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

// NewXlsx defined TODO
func NewXlsx(header ...[]map[string]interface{}) *Xlsx {
	xlsx := Xlsx{
		SheetName: "Sheet1",
		Format:    DefaultFormat,
		Handler:   DefaultHandler,
		TmplPath:  path.Join(viper.GetString("http.static"), "files"),
	}
	xlsx.Header = util.SomeOne(header, []map[string]interface{}{}).([]map[string]interface{})
	xlsx.xlsxFile = excelize.NewFile()
	index := xlsx.xlsxFile.NewSheet(xlsx.SheetName)
	xlsx.xlsxFile.SetActiveSheet(index)
	xlsx.xlsxSize = 10000
	uuid := uuid.New().String()
	filePath := path.Join(xlsx.TmplPath, fmt.Sprintf("%v.zip", uuid))
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		logrus.Errorln(err)
	}
	zipFile, err := os.Create(filePath)
	if err != nil {
		logrus.Errorln(err)
	}
	xlsx.zipFile = zip.NewWriter(zipFile)
	xlsx.zipPath = filePath
	xlsx.zipIndex = 1
	return &xlsx
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
		optionsets, err := new(SvcHepler).GetOptions(db, fmt.Sprintf("%v", format))
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
