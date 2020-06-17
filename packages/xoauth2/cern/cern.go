// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cern provides constants for using OAuth2 to access CERN services.
package cern // import "github.com/2637309949/dolphin/packages/xoauth2/cern"
import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is CERN's OAuth 2.0 endpoint.
var Endpoint = xoauth2.Endpoint{
	AuthURL:  "https://oauth.web.cern.ch/OAuth/Authorize",
	TokenURL: "https://oauth.web.cern.ch/OAuth/Token",
}
