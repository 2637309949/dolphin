// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stackoverflow provides constants for using OAuth2 to access Stack Overflow.
package stackoverflow // import "github.com/2637309949/dolphin/packages/xoauth2/stackoverflow"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Stack Overflow's OAuth 2.0 endpoint.
var Endpoint = xoauth2.Endpoint{
	AuthURL:  "https://stackoverflow.com/oauth",
	TokenURL: "https://stackoverflow.com/oauth/access_token",
}
