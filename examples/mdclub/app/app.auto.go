// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package app

import (
	"mdclub/model"

	"github.com/spf13/viper"
)

// Name project
var Name = "mdclub"

// Article defined
type Article struct {
	Name string
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get Route
}

// NewArticle defined
func NewArticle() *Article {
	ctr := &Article{Name: "article"}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/article/add"
	ctr.Add.Handler = ArticleAdd
	ctr.BatchAdd.Method = "POST"
	ctr.BatchAdd.RelativePath = "/article/batch_add"
	ctr.BatchAdd.Handler = ArticleBatchAdd
	ctr.Del.Method = "DELETE"
	ctr.Del.RelativePath = "/article/del"
	ctr.Del.Handler = ArticleDel
	ctr.BatchDel.Method = "DELETE"
	ctr.BatchDel.RelativePath = "/article/batch_del"
	ctr.BatchDel.Handler = ArticleBatchDel
	ctr.Update.Method = "PUT"
	ctr.Update.RelativePath = "/article/update"
	ctr.Update.Handler = ArticleUpdate
	ctr.BatchUpdate.Method = "PUT"
	ctr.BatchUpdate.RelativePath = "/article/batch_update"
	ctr.BatchUpdate.Handler = ArticleBatchUpdate
	ctr.Page.Method = "GET"
	ctr.Page.RelativePath = "/article/page"
	ctr.Page.Handler = ArticlePage
	ctr.Get.Method = "GET"
	ctr.Get.RelativePath = "/article/get"
	ctr.Get.Handler = ArticleGet
	return ctr
}

// ArticleRoutes defined
func ArticleRoutes(engine *Engine) {
	group, instance := engine.Group(viper.GetString("http.prefix")), ArticleInstance
	group.Handle(instance.Add.Method, instance.Add.RelativePath, Auth("token"), instance.Add)
	group.Handle(instance.BatchAdd.Method, instance.BatchAdd.RelativePath, Auth("token"), instance.BatchAdd)
	group.Handle(instance.Del.Method, instance.Del.RelativePath, Auth("token"), instance.Del)
	group.Handle(instance.BatchDel.Method, instance.BatchDel.RelativePath, Auth("token"), instance.BatchDel)
	group.Handle(instance.Update.Method, instance.Update.RelativePath, Auth("token"), instance.Update)
	group.Handle(instance.BatchUpdate.Method, instance.BatchUpdate.RelativePath, Auth("token"), instance.BatchUpdate)
	group.Handle(instance.Page.Method, instance.Page.RelativePath, Auth("token"), instance.Page)
	group.Handle(instance.Get.Method, instance.Get.RelativePath, Auth("token"), instance.Get)
}

// ArticleInstance defined
var ArticleInstance = NewArticle()

// SyncModel defined
func SyncModel(engine *Engine) error {
	mseti := engine.Manager.MSet()
	mseti.Add(new(model.McAnswer))
	mseti.Add(new(model.McArticle))
	mseti.Add(new(model.McCache))
	mseti.Add(new(model.McComment))
	mseti.Add(new(model.McFollow))
	mseti.Add(new(model.McImage))
	mseti.Add(new(model.McInbox))
	mseti.Add(new(model.McNotification))
	mseti.Add(new(model.McOption))
	mseti.Add(new(model.McQuestion))
	mseti.Add(new(model.McReport))
	mseti.Add(new(model.McToken))
	mseti.Add(new(model.McTopic))
	mseti.Add(new(model.McTopicable))
	mseti.Add(new(model.McUser))
	mseti.Add(new(model.McVote))
	return nil
}

// SyncController defined
func SyncController(engine *Engine) error {
	ArticleRoutes(engine)
	return nil
}

// SyncService defined
func SyncService(engine *Engine) error {
	return nil
}
