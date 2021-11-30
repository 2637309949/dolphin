// Code generated by dol build. Only Generate by tools if not existed.
// source: xlsx.go

package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/util/math"
)

// XlsxImport api implementation
// @Summary import xlsx
// @Tags xlsx controller
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/xlsx/import [post]
func (ctr *Xlsx) XlsxImport(ctx *Context) {
	var payload struct{}
	if err := ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}

	var err error
	file, headers, _ := ctx.Request().FormFile("file")
	defer file.Close()

	isXlsx := strings.Contains(headers.Filename, ".xlsx")
	if !isXlsx {
		ctx.Fail(errors.New("上传格式必须是xlsx格式"))
		return
	}

	sheet, err := ctr.Srv.Report.ParseExcel(file, 1)
	for i := range sheet {
		item := sheet[i]
		for k, v := range item {
			switch k {
			case "已录收单":
				fmt.Println("已录收单:", math.MustInt(v))
			case "无效收单":
				fmt.Println("无效收单:", math.MustInt(v))
			case "有效收单":
				fmt.Println("有效收单:", math.MustInt(v))
			case "日期":
				fmt.Println("日期:", math.MustTime(v, "2006-01-02"))
			}
		}
	}
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ctx.Success(sheet)
}
