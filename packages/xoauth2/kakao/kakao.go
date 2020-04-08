// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package kakao provides constants for using OAuth2 to access Kakao.
package kakao // import "github.com/2637309949/dolphin/packages/xoauth2/kakao"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Kakao's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://kauth.kakao.com/oauth/authorize",
	TokenURL: "https://kauth.kakao.com/oauth/token",
}
