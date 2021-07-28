package svc

import (
	"time"

	"github.com/2637309949/dolphin/platform/svc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type SvcHepler struct {
	dialer *kafka.Dialer
	bucket *oss.Bucket
	svc.Svc
}

func NewSvcHepler(rds redis.Cmdable) Svc {
	bucket := viper.GetString("oss.bucket")
	endpoint := viper.GetString("oss.endpoint")
	accesskeyid := viper.GetString("oss.accesskeyid")
	accesskeysecret := viper.GetString("oss.accesskeysecret")
	client, err := oss.New(endpoint, accesskeyid, accesskeysecret)
	if err != nil {
		logrus.Errorln(err)
	}

	bk, err := client.Bucket(bucket)
	if err != nil {
		logrus.Errorln(err)
	}

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	return &SvcHepler{Svc: svc.NewSvcHepler(rds), bucket: bk, dialer: dialer}
}
