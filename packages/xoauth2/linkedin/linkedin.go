// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedin provides constants for using OAuth2 to access LinkedIn.
package linkedin // import "github.com/2637309949/dolphin/packages/xoauth2/linkedin"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is LinkedIn's OAuth 2.0 endpoint.
var Endpoint = xoauth2.Endpoint{
	AuthURL:   "https://www.linkedin.com/oauth/v2/authorization",
	TokenURL:  "https://www.linkedin.com/oauth/v2/accessToken",
	AuthStyle: xoauth2.AuthStyleInParams,
}
