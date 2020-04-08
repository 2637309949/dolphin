// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package yahoo provides constants for using OAuth2 to access Yahoo.
package yahoo // import "github.com/2637309949/dolphin/packages/xoauth2/yahoo"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Yahoo's OAuth 2.0 endpoint.
// See https://developer.yahoo.com/oauth2/guide/
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.login.yahoo.com/oauth2/request_auth",
	TokenURL: "https://api.login.yahoo.com/oauth2/get_token",
}
