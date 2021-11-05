// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
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

// BearerAuth defined TODO
func (j *JWT) BearerAuth(req *http.Request) (string, bool) {
	prefix, auth := "Bearer ", req.Header.Get("Authorization")
	if strings.HasPrefix(auth, prefix) {
		return auth[len(prefix):], true
	}
	return "", false
}

// LoadAccessToken defined TODO
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

// GenerateAccessToken defined TODO
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

var _ Provider = &TokenProvider{}

type JWTProvider struct {
	jwt *JWT
}

// GetName defined TODO
func (p *JWTProvider) GetName() string {
	return "jwt"
}

// Config defined TODO
func (p *JWTProvider) Config(i *Identity) {
	p.jwt = i.jwt
}

// parseToken defined TODO
func (p *JWTProvider) parseJWTToken(t *jwt.MapClaims) TokenInfo {
	return &Token{
		UserID: (*t)["userId"].(string),
		Domain: (*t)["domain"].(string),
	}
}

// Verify defined TODO
func (p *JWTProvider) Verify(ctx *Context) (TokenInfo, bool) {
	if bearer, ok := p.jwt.BearerAuth(ctx.Request()); ok {
		tk, err := p.jwt.LoadAccessToken(bearer)
		if err != nil {
			logrus.Error(err)
			return nil, false
		}
		return p.parseJWTToken(tk), true
	}
	return nil, false
}

// Verify defined TODO
func (p *JWTProvider) Ticket(userId, extra string, ctx *Context) (TokenInfo, error) {
	tk, err := p.jwt.GenerateAccessToken(userId, extra)
	if err != nil {

		return nil, err
	}
	return &Token{Access: tk}, nil
}
