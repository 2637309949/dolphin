// Code generated by dol build. Only Generate by tools if not existed.
// source: like.go

package srv

import (
	"context"
	"fmt"
	"scene/model"
	"strconv"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/app"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// VoteLike defined srv
func VoteLike(ctx *gin.Context, db *xorm.Engine, params model.VoteInfo) (interface{}, error) {
	// 1.写入post_set
	if _, err := app.RedisClient.SAdd(context.Background(), "post_set", params.PostId.String).Result(); err != nil {
		return nil, err
	}
	// 2.添加点赞,仅仅判断缓存级别的
	// post用户集合 点赞+取消 为了查找key
	if _, err := app.RedisClient.SAdd(context.Background(),
		fmt.Sprintf("post_user_like_set_%v", params.PostId.String), params.UserId.String).Result(); err != nil {
		return nil, err
	}
	// post用户集合 点赞
	if params.Type.Int64 == 101 {
		// 统计文章点赞数目
		if _, err := app.RedisClient.SAdd(context.Background(),
			fmt.Sprintf("post_user_like_set_incr_%v", params.PostId.String), params.UserId.String).Result(); err != nil {
			return nil, err
		}
		// 用户查找列表时, 需要判断该文章是否已经点赞
		if _, err := app.RedisClient.SAdd(context.Background(),
			fmt.Sprintf("user_post_like_set_%v", params.UserId.String), params.PostId.String).Result(); err != nil {
			return nil, err
		}
	} else if params.Type.Int64 == 102 {
		// 统计文章点赞数目
		if _, err := app.RedisClient.SRem(context.Background(),
			fmt.Sprintf("post_user_like_set_incr_%v", params.PostId.String), params.UserId.String).Result(); err != nil {
			return nil, err
		}
		// 用户查找列表时, 需要判断该文章是否已经点赞
		if _, err := app.RedisClient.SRem(context.Background(),
			fmt.Sprintf("user_post_like_set_%v", params.UserId.String), params.PostId.String).Result(); err != nil {
			return nil, err
		}
	}
	// 3.写入post_user_like_{$post_id}{$user_id}
	if _, err := app.RedisClient.HMSet(context.Background(),
		fmt.Sprintf("post_user_like_%v%v", params.PostId.String, params.UserId.String),
		"create_time", time.Now().Value().Format("2006-01-02 15:04"), "type", params.Type.Int64).Result(); err != nil {
		return nil, err
	}
	// 4.新增计数
	num, err := app.RedisClient.SCard(context.Background(),
		fmt.Sprintf("post_user_like_set_incr_%v", params.PostId.String)).Result()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if _, err := app.RedisClient.Set(context.Background(),
		fmt.Sprintf("post%v_counter", num), params.PostId.String, 0).Result(); err != nil {
		return nil, err
	}
	return "ok", nil
}

func init() {
	app.App.Manager.Cron().AddFunc("*/10 * * * * *", func() {
		logrus.Infoln("scan vote")
		num1, err := app.RedisClient.SCard(context.Background(), "post_set").Result()
		if err != nil {
			logrus.Error(err)
			return
		}
		ulpsIncr, ulpsDecr := []model.UserLikePost{}, []model.UserLikePost{}
		for i := int64(0); i < num1; i++ {
			postId, err := app.RedisClient.SPop(context.Background(), "post_set").Result()
			if err != nil {
				logrus.Error(err)
				break
			}
			num2, err := app.RedisClient.SCard(context.Background(),
				fmt.Sprintf("post_user_like_set_%v", postId)).Result()
			if err != nil {
				logrus.Error(err)
				break
			}
			for i := int64(0); i < num2; i++ {
				ulp := model.UserLikePost{}
				userId, err := app.RedisClient.SPop(context.Background(),
					fmt.Sprintf("post_user_like_set_%v", postId)).Result()
				if err != nil {
					logrus.Error(err)
					break
				}
				field, err := app.RedisClient.HMGet(context.Background(),
					fmt.Sprintf("post_user_like_%v%v", postId, userId), "create_time", "type").Result()
				if err != nil {
					logrus.Error(err)
					break
				}
				te, _ := time.Parse("2006-01-02 15:04", fmt.Sprintf("%v", field[0]))
				ulp.ID = null.StringFromUUID()
				ulp.UserId = null.StringFrom(userId)
				ulp.PostId = null.StringFrom(postId)
				ulp.CreateTime = null.TimeFrom(te.Value())
				ulp.Creater = null.StringFrom(userId)
				ulp.UpdateTime = null.TimeFrom(te.Value())
				ulp.Updater = null.StringFrom(userId)
				ulp.IsDelete = null.IntFrom(0)
				tpe, _ := strconv.Atoi(fmt.Sprintf("%v", field[1]))
				if tpe == 101 {
					ulpsIncr = append(ulpsIncr, ulp)
				} else if tpe == 102 {
					ulpsDecr = append(ulpsDecr, ulp)
				}
				if _, err := app.RedisClient.Del(context.Background(),
					fmt.Sprintf("post_user_like_%v%v", postId, userId)).Result(); err != nil {
					logrus.Error(err)
					break
				}
				if _, err := app.RedisClient.Decr(context.Background(),
					fmt.Sprintf("post%v_counter", postId)).Result(); err != nil {
					logrus.Error(err)
					break
				}
			}
		}
		userDomains := map[string]string{}
		if uids, ok := slice.GetFieldSliceByName(ulpsIncr, "user_id", "%v").([]string); ok {
			users, _ := app.App.PlatformDB.Table("sys_user").In("id", uids).Where("is_delete != ?", 1).Cols("id", "domain").QueryString()
			for i := range users {
				userDomains[users[i]["id"]] = users[i]["domain"]
			}
		}
		if uids, ok := slice.GetFieldSliceByName(ulpsDecr, "user_id", "%v").([]string); ok {
			users, _ := app.App.PlatformDB.Table("sys_user").In("id", uids).Where("is_delete != ?", 1).Cols("id", "domain").QueryString()
			for i := range users {
				userDomains[users[i]["id"]] = users[i]["domain"]
			}
		}
		domainIncr := map[string][]model.UserLikePost{}
		domainDecr := map[string][]model.UserLikePost{}
		for i := range ulpsIncr {
			ret := domainIncr[userDomains[ulpsIncr[i].UserId.String]]
			domainIncr[userDomains[ulpsIncr[i].UserId.String]] = append(ret, ulpsIncr[i])
		}
		for i := range ulpsDecr {
			ret := domainDecr[userDomains[ulpsDecr[i].UserId.String]]
			domainDecr[userDomains[ulpsDecr[i].UserId.String]] = append(ret, ulpsDecr[i])
		}
		for k, v := range domainDecr {
			db := app.App.Manager.GetBusinessDB(k)
			if db != nil {
				logrus.Println("domainDecr:", len(v))
				sql := "false"
				for i2 := range v {
					sql += " or (user_id='" + v[i2].UserId.String + "' and " + "post_id='" + v[i2].PostId.String + "')"
				}
				if _, err := db.Table(new(model.UserLikePost)).Where(sql).Update(&map[string]interface{}{
					"is_delete": 1,
				}); err != nil {
					logrus.Error(err)
					break
				}
			}
		}
		for k, v := range domainIncr {
			db := app.App.Manager.GetBusinessDB(k)
			if db != nil {
				logrus.Println("domainIncr:", len(v))
				extList := []model.UserLikePost{}
				sql := "false"
				for i2 := range v {
					sql += " or (user_id='" + v[i2].UserId.String + "' and " + "post_id='" + v[i2].PostId.String + "')"
				}
				if err := db.Table(new(model.UserLikePost)).Where(sql).Find(&extList); err != nil {
					logrus.Error(err)
					break
				}
				items := funk.Filter(v, func(v1 model.UserLikePost) bool {
					for i3 := range extList {
						if v1.UserId.String == extList[i3].UserId.String &&
							v1.PostId.String == extList[i3].PostId.String {
							return false
						}
					}
					return true
				}).([]model.UserLikePost)
				if _, err := db.Table(new(model.UserLikePost)).Insert(&items); err != nil {
					logrus.Error(err)
					break
				}
			}
		}
	})
}