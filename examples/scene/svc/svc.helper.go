package svc

import (
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type SvcHepler struct {
	bucket *oss.Bucket
	svc.Svc
}

func NewSvcHepler(svc svc.Svc) Svc {
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
	return &SvcHepler{Svc: svc, bucket: bk}
}
