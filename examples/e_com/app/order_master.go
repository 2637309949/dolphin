// Code generated by dol build. Only Generate by tools if not existed.
// source: order_master.go

package app

import (
	"e_com/model"
	"errors"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// OrderMasterAdd api implementation
// @Summary 添加order_master
// @Tags order_master接口
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param order_master body model.OrderMaster false "order_master信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/add [post]
func OrderMasterAdd(ctx *Context) {
	var payload model.OrderMaster
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

// OrderMasterBatchAdd api implementation
// @Summary 添加order_master
// @Tags order_master接口
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param order_master body []model.OrderMaster false "order_master信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/batch_add [post]
func OrderMasterBatchAdd(ctx *Context) {
	var payload []model.OrderMaster
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

// OrderMasterDel api implementation
// @Summary 删除order_master
// @Tags order_master接口
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param order_master body model.OrderMaster false "order_master"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/del [delete]
func OrderMasterDel(ctx *Context) {
	var payload model.OrderMaster
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.OrderId.Int64).Update(&model.OrderMaster{
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

// OrderMasterBatchDel api implementation
// @Summary 删除order_master
// @Tags order_master接口
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param order_master body []model.OrderMaster false "order_master"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/batch_del [delete]
func OrderMasterBatchDel(ctx *Context) {
	var payload []model.OrderMaster
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.OrderMaster) int64 { return form.OrderId.Int64 }).([]int64)
	ret, err := ctx.DB.In("id", ids).Update(&model.OrderMaster{
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

// OrderMasterUpdate api implementation
// @Summary 更新order_master
// @Tags order_master接口
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param order_master body model.OrderMaster false "order_master信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/update [put]
func OrderMasterUpdate(ctx *Context) {
	var payload model.OrderMaster
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Updater = null.IntFromStr(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.OrderId.Int64).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// OrderMasterBatchUpdate api implementation
// @Summary 更新order_master
// @Tags order_master接口
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param order_master body []model.OrderMaster false "order_master信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/batch_update [put]
func OrderMasterBatchUpdate(ctx *Context) {
	var payload []model.OrderMaster
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
	s.Begin()
	defer s.Close()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].Updater = null.IntFromStr(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].OrderId.Int64).Update(&payload[i])
		if err != nil {
			logrus.Error(err)
			s.Rollback()
			ctx.Fail(err)
			return
		}
		ret = append(ret, r)
	}
	if err != nil {
		logrus.Error(err)
		s.Rollback()
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

// OrderMasterPage api implementation
// @Summary order_master分页查询
// @Tags order_master接口
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/page [get]
func OrderMasterPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 10)
	q.SetRule("order_master_page")
	q.SetString("creater")
	q.SetString("updater")
	q.SetRange("create_time")
	q.SetRange("update_time")
	q.SetInt("is_delete", 0)()
	q.SetString("sort", "update_time desc")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "order_master", "page", "order_master", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// OrderMasterGet api implementation
// @Summary 获取order_master信息
// @Tags order_master接口
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "order_masterid"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/order/master/get [get]
func OrderMasterGet(ctx *Context) {
	var entity model.OrderMaster
	err := ctx.ShouldBindQuery(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if ext, err := ctx.DB.Get(&entity); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	} else if !ext {
		ctx.Fail(errors.New("not found"))
		return
	}
	ctx.Success(entity)
}
