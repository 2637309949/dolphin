// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package twitch provides constants for using OAuth2 to access Twitch.
package twitch // import "github.com/2637309949/dolphin/packages/xoauth2/twitch"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Twitch's OAuth 2.0 endpoint.
//
// For more information see:
// https://dev.twitch.tv/docs/authentication
var Endpoint = xoauth2.Endpoint{
	AuthURL:  "https://id.twitch.tv/oauth2/authorize",
	TokenURL: "https://id.twitch.tv/oauth2/token",
}
