// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"encoding/json"
	"errors"
	"sort"
	"time"

	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/platform/util/encrypt"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

// EncryptForm defines Common request parameter
type EncryptForm struct {
	AppID      string `form:"app_id" json:"app_id" xml:"app_id" binding:"required"`
	Sign       string `form:"sign" json:"sign" xml:"sign" binding:"required"`
	TimeStamp  int64  `form:"timestamp" json:"timestamp" xml:"timestamp" binding:"required"`
	BizContent string `form:"biz_content" json:"biz_content" xml:"biz_content" binding:"required"`
}

// NewEncryptForm defined TODO
func NewEncryptForm() *EncryptForm {
	return &EncryptForm{}
}

// ParseForm defined
func (ec *EncryptForm) ParseForm(ctx *Context) (*EncryptForm, error) {
	puData := EncryptForm{}
	if ctx.Request.Method != "POST" {
		if err := ctx.BindQuery(&puData); err != nil {
			return nil, err
		}
	} else {
		if err := ctx.ShouldBindBodyWith(&puData, binding.JSON); err != nil {
			return nil, err
		}
	}
	return &puData, nil
}

// form2Uri defined
func (ec *EncryptForm) parseForm() (string, error) {
	var puJSON map[string]string
	var puKeys = make([]string, 0, len(puJSON))
	puByte, err := json.Marshal(ec)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(puByte, &puJSON)
	if err != nil {
		return "", err
	}
	for k := range puJSON {
		if k != "sign" {
			puKeys = append(puKeys, k)
		}
	}
	sort.Strings(puKeys)
	var signString = ""
	for i := range puKeys {
		if signString != "" {
			signString = signString + "&" + puKeys[i] + "=" + puJSON[puKeys[i]]
		} else {
			signString = signString + puKeys[i] + "=" + puJSON[puKeys[i]]
		}
	}
	return signString, nil
}

func (ec *EncryptForm) sign(cli oauth2.ClientInfo) ([]byte, error) {
	uri, err := ec.parseForm()
	if err != nil {
		return []byte{}, err
	}
	ecyt, err := encrypt.AesEncrypt([]byte(uri), []byte(cli.GetSecret()))
	if err != nil {
		return []byte{}, err
	}
	return ecyt, nil
}

// Verify defined
func (ec *EncryptForm) Verify(cli oauth2.ClientInfo) (bool, error) {
	nowTs := time.Now().Unix()
	ts := ec.TimeStamp
	if ts > nowTs || nowTs-ts >= 60 {
		return false, errors.New("timestamp error")
	}
	sn, err := ec.sign(cli)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	if string(sn) != ec.Sign {
		return false, errors.New("sign error")
	}
	return true, nil
}
