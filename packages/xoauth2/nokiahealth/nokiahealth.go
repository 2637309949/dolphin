// Copyright 2018 The oauth2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package nokiahealth provides constants for using OAuth2 to access the Nokia Health Mate API.
package nokiahealth

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Nokia Health Mate's OAuth 2.0 endpoint.
var Endpoint = xoauth2.Endpoint{
	AuthURL:  "https://account.health.nokia.com/oauth2_user/authorize2",
	TokenURL: "https://account.health.nokia.com/oauth2/token",
}
