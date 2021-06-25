// Code generated by dol build. Only Generate by tools if not existed.
// source: like.go

package srv

import (
	"context"
	"errors"
	"fmt"
	"scene/model"
	"strconv"

	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/app"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

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
				userId, err := app.RedisClient.SPop(context.Background(), fmt.Sprintf("post_user_like_set_%v", postId)).Result()
				if err != nil {
					logrus.Error(err)
					break
				}
				field, err := app.RedisClient.HMGet(context.Background(), fmt.Sprintf("post_user_like_%v%v", postId, userId), "create_time", "type").Result()
				if err != nil {
					logrus.Error(err)
					break
				}
				te, _ := time.Parse("2006-01-02 15:04", fmt.Sprintf("%v", field[0]))
				ulp.UserId = null.IntFromStr(userId)
				ulp.PostId = null.IntFromStr(postId)
				ulp.CreateTime = null.TimeFrom(te)
				ulp.Creater = null.IntFromStr(userId)
				ulp.UpdateTime = null.TimeFrom(te)
				ulp.Updater = null.IntFromStr(userId)
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
		userDomains := map[int64]string{}
		if uids, ok := slice.GetFieldSliceByName(ulpsIncr, "user_id", "%v").([]string); ok {
			users, _ := app.App.PlatformDB.Table("sys_user").In("id", uids).Where("is_delete != ?", 1).Cols("id", "domain").QueryString()
			for i := range users {
				userDomains[null.IntFromStr(users[i]["id"]).Int64] = users[i]["domain"]
			}
		}
		if uids, ok := slice.GetFieldSliceByName(ulpsDecr, "user_id", "%v").([]string); ok {
			users, _ := app.App.PlatformDB.Table("sys_user").In("id", uids).Where("is_delete != ?", 1).Cols("id", "domain").QueryString()
			for i := range users {
				userDomains[null.IntFromStr(users[i]["id"]).Int64] = users[i]["domain"]
			}
		}
		domainIncr := map[string][]model.UserLikePost{}
		domainDecr := map[string][]model.UserLikePost{}
		for i := range ulpsIncr {
			ret := domainIncr[userDomains[ulpsIncr[i].UserId.Int64]]
			domainIncr[userDomains[ulpsIncr[i].UserId.Int64]] = append(ret, ulpsIncr[i])
		}
		for i := range ulpsDecr {
			ret := domainDecr[userDomains[ulpsDecr[i].UserId.Int64]]
			domainDecr[userDomains[ulpsDecr[i].UserId.Int64]] = append(ret, ulpsDecr[i])
		}
		for k, v := range domainDecr {
			db := app.App.Manager.GetBusinessDB(k)
			if db != nil {
				logrus.Println("domainDecr:", len(v))
				sql := "false"
				for i2 := range v {
					sql += " or (user_id=" + fmt.Sprintf("%v", v[i2].UserId.Int64) + " and " + "post_id=" + fmt.Sprintf("%v", v[i2].PostId.Int64) + ")"
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
					sql += " or (user_id=" + fmt.Sprintf("%v", v[i2].UserId.Int64) + " and " + "post_id=" + fmt.Sprintf("%v", v[i2].PostId.Int64) + ")"
				}
				if err := db.Table(new(model.UserLikePost)).Where(sql).Find(&extList); err != nil {
					logrus.Error(err)
					break
				}
				items := funk.Filter(v, func(v1 model.UserLikePost) bool {
					for i3 := range extList {
						if v1.UserId.Int64 == extList[i3].UserId.Int64 &&
							v1.PostId.Int64 == extList[i3].PostId.Int64 {
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

type Vote struct {
}

func NewVote() *Vote {
	return &Vote{}
}

// Like defined srv
func (srv *Vote) Like(ctx context.Context, db *xorm.Engine, params model.VoteInfo) (interface{}, error) {
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
		"create_time", time.Now().Format("2006-01-02 15:04"), "type", params.Type.Int64).Result(); err != nil {
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
		fmt.Sprintf("post%v_counter", params.PostId.String), num, 0).Result(); err != nil {
		return nil, err
	}
	return "ok", nil
}

// TODO defined srv
func (srv *Vote) TODO(ctx context.Context, db *xorm.Engine, params struct{}) (interface{}, error) {
	cwt, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	chi := func(cwt context.Context) chan interface{} {
		chi := make(chan interface{}, 1)
		go func() {
			time.Sleep(1 * time.Second)
			chi <- 100
			cancel()
			close(chi)
		}()
		return chi
	}(cwt)
	for range ticker.C {
		select {
		case <-cwt.Done():
			logrus.Infoln("child process interrupt...")
			return <-chi, cwt.Err()
		default:
			logrus.Infoln("awaiting job...")
		}
	}
	return nil, errors.New("no implementation found")
}
