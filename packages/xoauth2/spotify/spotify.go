// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package spotify provides constants for using OAuth2 to access Spotify.
package spotify // import "github.com/2637309949/dolphin/packages/xoauth2/spotify"

import (
	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Endpoint is Spotify's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.spotify.com/authorize",
	TokenURL: "https://accounts.spotify.com/api/token",
}
