// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heroku provides constants for using OAuth2 to access Heroku.
package heroku // import "github.com/2637309949/dolphin/packages/xoauth2/heroku"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Heroku's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://id.heroku.com/oauth/authorize",
	TokenURL: "https://id.heroku.com/oauth/token",
}
