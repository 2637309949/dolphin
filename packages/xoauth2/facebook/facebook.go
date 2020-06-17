// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package facebook provides constants for using OAuth2 to access Facebook.
package facebook // import "github.com/2637309949/dolphin/packages/xoauth2/facebook"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Facebook's OAuth 2.0 endpoint.
var Endpoint = xoauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/v3.2/dialog/oauth",
	TokenURL: "https://graph.facebook.com/v3.2/oauth/access_token",
}
