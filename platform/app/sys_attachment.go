// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_attachment.go

package app

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util/encode"
	"github.com/2637309949/dolphin/platform/util/file"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// SysAttachmentAdd api implementation
// @Summary 添加附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysAttachment false "附件信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/attachment/add [post]
func SysAttachmentAdd(ctx *Context) {
	var payload model.SysAttachment
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.CreateTime = null.TimeFrom(time.Now().Value())
	payload.Creater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.IsDelete = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentBatchAdd api implementation
// @Summary 添加附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_attachment body []model.SysAttachment false "附件信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/attachment/batch_add [post]
func SysAttachmentBatchAdd(ctx *Context) {
	var payload []model.SysAttachment
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].CreateTime = null.TimeFrom(time.Now().Value())
		payload[i].Creater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
		payload[i].IsDelete = null.IntFrom(0)
	}
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/attachment/upload [post]
func SysAttachmentUpload(ctx *Context) {
	var attachments []model.SysAttachment
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
		uuid := uuid.New().String()
		filePath := path.Join(viper.GetString("http.static"), "files", uuid+filepath.Ext(filename))
		if err := ctx.SaveUploadedFile(f, filePath); err != nil {
			ctx.Fail(err)
			return
		}
		urlPath := path.Join("files", uuid+filepath.Ext(filename))
		item := model.SysAttachment{
			Name:       null.StringFrom(filename),
			UUID:       null.StringFrom(uuid),
			Size:       null.IntFrom(f.Size),
			Ext:        null.StringFrom(filepath.Ext(filename)),
			Hash:       null.StringFrom(string(file.MustHash(filePath))),
			URL:        null.StringFrom(urlPath),
			Path:       null.StringFrom(filePath),
			Type:       null.StringFrom(fileType),
			Durable:    null.IntFrom(0),
			CreateTime: null.TimeFrom(time.Now().Value()),
			Creater:    null.IntFromStr(ctx.GetToken().GetUserID()),
			UpdateTime: null.TimeFrom(time.Now().Value()),
			Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
			IsDelete:   null.IntFrom(0),
		}
		attachments = append(attachments, item)
	}
	if _, err = ctx.DB.Insert(attachments); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	attachs := funk.Map(attachments, func(mst model.SysAttachment) model.Attach {
		return model.Attach{
			ID:   mst.ID.Int64,
			Name: mst.Name.String,
			Hash: mst.Hash.String,
			URL:  mst.URL.String,
		}
	})

	ctx.Success(attachs)
}

// SysAttachmentExport api implementation
// @Summary 附件导出
// @Tags 附件
// @Param Authorization header string false "认证令牌"
// @Param file_name  query  string false "附件名称"
// @Param file_id  query  string false "附件ID"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/attachment/export [get]
func SysAttachmentExport(ctx *Context) {
	filePath := path.Join(viper.GetString("http.static"), "files", ctx.QueryString("file_id"))
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+encode.URIComponent(ctx.QueryString("file_name")))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(filePath)
}

// SysAttachmentDel api implementation
// @Summary 删除附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_attachment body model.SysAttachment false "附件"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/attachment/del [delete]
func SysAttachmentDel(ctx *Context) {
	var payload model.SysAttachment
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.Int64).Update(&model.SysAttachment{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentBatchDel api implementation
// @Summary 删除附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_attachment body []model.SysAttachment false "附件"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/attachment/batch_del [post]
func SysAttachmentBatchDel(ctx *Context) {
	var payload []model.SysAttachment
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysAttachment) int64 { return form.ID.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&model.SysAttachment{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		Updater:    null.IntFromStr(ctx.GetToken().GetUserID()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/attachment/update [put]
func SysAttachmentUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentBatchUpdate api implementation
// @Summary 添加附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_attachment body []model.SysAttachment false "附件信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/attachment/batch_update [post]
func SysAttachmentBatchUpdate(ctx *Context) {
	var payload []model.SysAttachment
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.Int64).Update(&payload[i])
		if err != nil {
			s.Rollback()
			logrus.Error(err)
			ctx.Fail(err)
			return
		}
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	err = s.Commit()
	if err != nil {
		logrus.Error(err)
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/attachment/page [get]
func SysAttachmentPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetRule("sys_attachment_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_attachment", "page", "sys_attachment", q.Value())
	if err != nil {
		logrus.Error(err)
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
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/attachment/get [get]
func SysAttachmentGet(ctx *Context) {
	var entity model.SysAttachment
	err := ctx.ShouldBindQuery(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ext, err := ctx.DB.Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(errors.New("not found"))
		return
	}
	ctx.Success(entity)
}
