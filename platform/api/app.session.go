package api

import (
	"encoding/gob"
	"net/url"

	"github.com/go-session/session"
	"github.com/spf13/viper"
)

// InitSession defined TODO, just for cas server
func InitSession() {
	var hashKey = []byte(viper.GetString("http.hash"))
	session.InitManager(
		session.SetCookieName("session_id"),
		session.SetStore(
			NewCookieStore(
				SetCookieName("store_id"),
				SetHashKey(hashKey),
			),
		),
	)
	// fix: store.Set("ReturnUri", r.Form) can not be parsed.
	gob.Register(url.Values{})
}
