// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

//  JWT defined TODO
type JWT struct {
	Secret string
	Expire int64
}

// NewJWT defined TODO
func NewJWT(secret string, expire int64) *JWT {
	return &JWT{
		Secret: secret,
		Expire: expire,
	}
}

// BearerAuth parse bearer token
func (j *JWT) BearerAuth(r *http.Request) (string, bool) {
	prefix, auth := "Bearer ", r.Header.Get("Authorization")
	if strings.HasPrefix(auth, prefix) {
		return auth[len(prefix):], true
	}
	return "", false
}

func (j *JWT) LoadAccessToken(tk string) (*jwt.MapClaims, error) {
	jc := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tk, jc, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return &jc, nil
}

func (j *JWT) GenerateAccessToken(userId, domain string) (string, error) {
	payloads, now, claims := map[string]interface{}{
		"userId": userId,
		"domain": domain,
	}, time.Now().Unix(), jwt.MapClaims{}
	claims["exp"] = now + j.Expire
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(j.Secret))
}
