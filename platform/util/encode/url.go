package encode

import (
	"net/url"
	"strings"
)

// URIComponent defined
func URIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}
