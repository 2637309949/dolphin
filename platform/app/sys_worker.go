// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_worker.go

package app

import (
	"encoding/json"

	"github.com/2637309949/dolphin/platform/model"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Job defined
type Job model.Worker

// Do do something
func (p Job) Do() error {
	var (
		w   model.Worker
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(&p); err != nil {
		logrus.Error(err)
		return err
	}
	if err := json.Unmarshal(bs, &w); err != nil {
		logrus.Error(err)
		return err
	}

	//Do a job sample
	worker := App.Manager.Worker()
	b, err := worker.Find(p.Code)
	if err != nil {
		logrus.Error(err)
		return err
	}
	b.Status = model.WorkerStatusProccessing
	if err := worker.Update(b); err != nil {
		logrus.Error(err)
		return err
	}
	// Run a worker
	ret, err := worker.GetJobHandler(p.Name)(w)
	if err != nil {
		logrus.Error(err)
		b.Status = model.WorkerStatusInterrupt
		b.Error = err.Error()
		if err := worker.Update(b); err != nil {
			logrus.Error(err)
			return err
		}
		return err
	}
	b.Status = model.WorkerStatusFinish
	b.Result = ret
	if err := worker.Update(b); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

// SysWorkerAdd api implementation
// @Summary 添加worker
// @Tags worker
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param worker body model.Worker false "worker信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/worker/add [post]
func SysWorkerAdd(ctx *Context) {
	var (
		payload model.Worker
		w       Job
		bs      []byte
		err     error
	)
	if err = ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Code = uuid.New().String()
	_, err = ctx.LoginInInfo(&payload.User)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.Status = model.WorkerStatusInitial
	if bs, err = json.Marshal(&payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if err := json.Unmarshal(bs, &w); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	worker := App.Manager.Worker()
	worker.AddJob(w)
	worker.Save(payload)
	ctx.Success(&map[string]interface{}{
		"name":   w.Name,
		"code":   w.Code,
		"status": w.Status,
	})
}

// SysWorkerGet api implementation
// @Summary 获取worker信息
// @Tags worker
// @Param Authorization header string false "认证令牌"
// @Param code  query  string false "worker code"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/worker/get [get]
func SysWorkerGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("code")
	worker := App.Manager.Worker()
	w, err := worker.Find(q.GetString("code"))
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret := map[string]interface{}{
		"name":   w.Name,
		"code":   w.Code,
		"status": w.Status,
	}
	if w.Status == model.WorkerStatusFinish {
		ret["result"] = w.Result
	} else if w.Status == model.WorkerStatusInterrupt {
		ret["error"] = w.Error
	}
	ctx.Success(&ret)
}
