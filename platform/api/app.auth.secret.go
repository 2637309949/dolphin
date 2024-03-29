// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"time"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/platform/util/encrypt"
)

// Secret defines Common request parameter
type Secret struct {
	AppID      string `form:"app_id" json:"app_id" xml:"app_id" binding:"required"`
	Sign       string `form:"sign" json:"sign" xml:"sign" binding:"required"`
	TimeStamp  int64  `form:"timestamp" json:"timestamp" xml:"timestamp" binding:"required"`
	BizContent string `form:"biz_content" json:"biz_content" xml:"biz_content" binding:"required"`
}

// NewSecret defined TODO
func NewSecret(ctx *Context) (*Secret, error) {
	secret := Secret{}
	err := ctx.ShouldBind(&secret)
	if err != nil {
		return nil, err
	}
	return &secret, nil
}

// parseForm defined TODO
func (secret *Secret) parseForm() (string, error) {
	var puJSON map[string]string
	var puKeys = make([]string, 0, len(puJSON))
	puByte, err := json.Marshal(secret)
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

// sign defined TODO
func (secret *Secret) sign(cli oauth2.ClientInfo) ([]byte, error) {
	uri, err := secret.parseForm()
	if err != nil {
		return []byte{}, err
	}
	ecyt, err := encrypt.AesEncrypt([]byte(uri), []byte(cli.GetSecret()))
	if err != nil {
		return []byte{}, err
	}
	return ecyt, nil
}

// Verify defined TODO
func (secret *Secret) Verify(cli oauth2.ClientInfo) (bool, error) {
	nowTs := time.Now().Unix()
	ts := secret.TimeStamp
	if ts > nowTs || nowTs-ts >= 60 {
		return false, errors.New("timestamp error")
	}
	sn, err := secret.sign(cli)
	if err != nil {
		logrus.Error(context.TODO(), err)
		return false, err
	}
	if string(sn) != secret.Sign {
		return false, errors.New("sign error")
	}
	return true, nil
}

var _ Provider = &TokenProvider{}

type RSAProvider struct {
}

// GetName defined TODO
func (p *RSAProvider) GetName() string {
	return "rsa"
}

// Config defined TODO
func (p *RSAProvider) Config(i *Identity) {
}

// Verify defined TODO
func (p *RSAProvider) Verify(ctx *Context) (TokenInfo, bool) {
	secret, err := NewSecret(ctx)
	if err != nil {
		logrus.Error(context.TODO(), err)
		return nil, false
	}

	cli, err := NewClientStore().GetByID(secret.AppID)
	if err != nil {
		logrus.Error(context.TODO(), err)
		return nil, false
	}

	valid, err := secret.Verify(cli)
	if err != nil {
		logrus.Error(context.TODO(), err)
		return nil, false
	}
	return &Token{ClientID: cli.GetID()}, valid
}

// Verify defined TODO
func (p *RSAProvider) Ticket(userId, extra string, ctx *Context) (TokenInfo, error) {
	return nil, nil
}
