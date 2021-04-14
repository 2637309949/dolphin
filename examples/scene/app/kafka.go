// Code generated by dol build. Only Generate by tools if not existed.
// source: kafka.go

package app

import (
	"context"
	"scene/model"
	"scene/srv"

	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

// KafkaAdd api implementation
// @Summary Add article
// @Tags Kafka controller
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param kafka body model.KafkaInfo false "Kafka info"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/kafka/add [post]
func KafkaAdd(ctx *Context) {
	var payload model.KafkaInfo
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if err := srv.KafkaProducer(context.Background(), payload); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success("ok")
}

// KafkaGet api implementation
// @Summary Get kafka info
// @Tags Kafka controller
// @Param id  query  string false "kafka id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/kafka/get [get]
func KafkaGet(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("id")
	q.Value()
	// ctx.Persist(ctx.DB, "0586e250-5b6c-4a79-9f4a-767a742b7890")
	// ctx.Remove(ctx.DB, "6ebce24f-6887-4d6d-a62a-2a706fcf1c3f")
	// ret, err := srv.KafkaConsumer(context.Background(), q.Value())
	// if err != nil {
	// 	logrus.Error(err)
	// 	ctx.Fail(err)
	// 	return
	// }
	ctx.Success("ok")
}
