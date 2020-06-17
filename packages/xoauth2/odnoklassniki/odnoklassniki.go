// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package odnoklassniki provides constants for using OAuth2 to access Odnoklassniki.
package odnoklassniki // import "github.com/2637309949/dolphin/packages/xoauth2/odnoklassniki"
import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Odnoklassniki's OAuth 2.0 endpoint.
var Endpoint = xoauth2.Endpoint{
	AuthURL:  "https://www.odnoklassniki.ru/oauth/authorize",
	TokenURL: "https://api.odnoklassniki.ru/oauth/token.do",
}
