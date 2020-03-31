// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_attachment.go

package app

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/2637309949/dolphin/platform/util/file"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/packages/uuid"
	"github.com/2637309949/dolphin/packages/viper"
)

// SysAttachmentAdd api implementation
// @Summary 添加附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "附件信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/add [post]
func SysAttachmentAdd(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFromPtr(time.Now().Value())
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentUpload api implementation
// @Summary 上传附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/upload [post]
func SysAttachmentUpload(ctx *Context) {
	var attachments []model.SysAttachment
	var attachs []model.Attach
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println("err", err)
		ctx.Fail(err)
		return
	}
	fileType := ""
	files := form.File["files"]
	if params := form.Value["type"]; len(params) > 0 {
		fileType = params[0]
	}
	file.EnsureDir(path.Join(viper.GetString("http.static"), "files"))
	for _, f := range files {
		filename := filepath.Base(f.Filename)
		uuid := uuid.MustString()
		filePath := path.Join(viper.GetString("http.static"), "files", uuid+filepath.Ext(filename))
		if err := ctx.SaveUploadedFile(f, filePath); err != nil {
			ctx.Fail(err)
			return
		}
		urlPath := path.Join("files", uuid+filepath.Ext(filename))
		item := model.SysAttachment{
			ID:         null.StringFromUUID(),
			Name:       null.StringFrom(filename),
			UUID:       null.StringFrom(uuid),
			Size:       null.IntFrom(f.Size),
			Ext:        null.StringFrom(filepath.Ext(filename)),
			Hash:       null.StringFrom(string(file.MustHash(filePath))),
			URL:        null.StringFrom(urlPath),
			Path:       null.StringFrom(filePath),
			Type:       null.StringFrom(fileType),
			CreateTime: null.TimeFromPtr(time.Now().Value()),
			CreateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
			UpdateTime: null.TimeFromPtr(time.Now().Value()),
			UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		}
		attachments = append(attachments, item)
		attachs = append(attachs, model.Attach{
			Name: item.Name.String,
			Hash: item.Hash.String,
			URL:  item.URL.String,
		})
	}
	if _, err = ctx.DB.Insert(attachments); err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(attachs)
}

// SysAttachmentDel api implementation
// @Summary 删除附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_attachment body model.SysAttachment false "附件"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/del [delete]
func SysAttachmentDel(ctx *Context) {
	var payload model.SysAttachment
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Table(new(model.SysAttachment)).In("id", payload.ID.String).Update(map[string]interface{}{
		"delete_time": null.TimeFromPtr(time.Now().Value()),
		"delete_by":   null.StringFrom(ctx.GetToken().GetUserID()),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentUpdate api implementation
// @Summary 更新附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "附件信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/update [put]
func SysAttachmentUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentPage api implementation
// @Summary 附件分页查询
// @Tags 附件
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/page [get]
func SysAttachmentPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page")
	q.SetInt("size")
	ret, err := ctx.PageSearch(ctx.DB, "sys_attachment", "page", "sys_attachment", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentGet api implementation
// @Summary 获取附件信息
// @Tags 附件
// @Param Authorization header string false "认证令牌"
// @Param id query string false "附件id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/get [get]
func SysAttachmentGet(ctx *Context) {
	var entity model.SysAttachment
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
