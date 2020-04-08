// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gitlab provides constants for using OAuth2 to access GitLab.
package gitlab // import "github.com/2637309949/dolphin/packages/xoauth2/gitlab"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is GitLab's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://gitlab.com/oauth/authorize",
	TokenURL: "https://gitlab.com/oauth/token",
}
