// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_worker.go

package api

import (
	"context"
	"encoding/json"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/google/uuid"
)

// Job defined
type Job types.Worker

// Do do something
func (p Job) Do() error {
	var (
		w   types.Worker
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(&p); err != nil {
		logrus.Error(context.TODO(), err)
		return err
	}
	if err := json.Unmarshal(bs, &w); err != nil {
		logrus.Error(context.TODO(), err)
		return err
	}

	//Do a job sample
	worker := App.Manager.Worker()
	b, err := worker.Find(p.Code)
	if err != nil {
		logrus.Error(context.TODO(), err)
		return err
	}
	b.Status = types.WorkerStatusProccessing
	if err := worker.Update(b); err != nil {
		logrus.Error(context.TODO(), err)
		return err
	}
	// Run a worker
	ret, err := worker.GetJobHandler(p.Name)(w)
	if err != nil {
		logrus.Error(context.TODO(), err)
		b.Status = types.WorkerStatusInterrupt
		b.Error = err.Error()
		if err := worker.Update(b); err != nil {
			logrus.Error(context.TODO(), err)
			return err
		}
		return err
	}
	b.Status = types.WorkerStatusFinish
	b.Result = ret
	if err := worker.Update(b); err != nil {
		logrus.Error(context.TODO(), err)
		return err
	}
	return nil
}

// SysWorkerAdd api implementation
// @Summary 添加worker
// @Tags worker
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param worker body types.Worker false "worker信息"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/worker/add [post]
func (ctr *SysWorker) SysWorkerAdd(ctx *Context) {
	var (
		payload types.Worker
		w       Job
		bs      []byte
		err     error
	)
	if err = ctx.ShouldBindWith(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	payload.Code = uuid.New().String()
	_, err = ctx.LoginInInfo(&payload.User)
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	payload.Status = types.WorkerStatusInitial
	if bs, err = json.Marshal(&payload); err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	if err := json.Unmarshal(bs, &w); err != nil {
		logrus.Error(ctx, err)
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
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/sys/worker/get [get]
func (ctr *SysWorker) SysWorkerGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("code")
	worker := App.Manager.Worker()
	w, err := worker.Find(q.GetString("code"))
	if err != nil {
		logrus.Error(ctx, err)
		ctx.Fail(err)
		return
	}
	ret := map[string]interface{}{
		"name":   w.Name,
		"code":   w.Code,
		"status": w.Status,
	}
	if w.Status == types.WorkerStatusFinish {
		ret["result"] = w.Result
	} else if w.Status == types.WorkerStatusInterrupt {
		ret["error"] = w.Error
	}
	ctx.Success(&ret)
}
