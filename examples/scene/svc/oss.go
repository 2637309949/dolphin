package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// oss defined defined
type Oss interface {
	PutObjectFromFile(target string, src string) error
	GetObjectToFile(target string, src string) error
	ListObjects() ([]string, error)
	DeleteObject(string) error
}

type XOss struct {
	bucket *oss.Bucket
}

// PutObjectFromFile defined TODO
func (x *XOss) PutObjectFromFile(target string, src string) error {
	err := x.bucket.PutObjectFromFile(target, src)
	return err
}

// PutObjectFromFile defined TODO
func (x *XOss) GetObjectToFile(target string, src string) error {
	err := x.bucket.GetObjectToFile(target, src)
	return err
}

// PutObjectFromFile defined TODO
func (x *XOss) ListObjects() (keys []string, err error) {
	lsRes, err := x.bucket.ListObjects()
	if err != nil {
		return []string{}, err
	}
	for _, object := range lsRes.Objects {
		keys = append(keys, object.Key)
	}
	return keys, err
}

// PutObjectFromFile defined TODO
func (x *XOss) DeleteObject(target string) error {
	err := x.bucket.DeleteObject(target)
	return err
}

func NewXOss() *XOss {
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
	return &XOss{bucket: bk}
}
