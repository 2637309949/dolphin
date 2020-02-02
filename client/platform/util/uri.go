package util

import (
	"regexp"
	"strings"
	"time"
)

// DbType db type
type DbType string

// URI db uri
type URI struct {
	DbType  DbType
	Proto   string
	Host    string
	Port    string
	DbName  string
	User    string
	Passwd  string
	Charset string
	Laddr   string
	Raddr   string
	Timeout time.Duration
	Schema  string
}

// Parse uri
func Parse(dataSourceName string) (*URI, error) {
	dsnPattern := regexp.MustCompile(
		`^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` + // [user[:password]@]
			`(?:(?P<net>[^\(]*)(?:\((?P<addr>[^\)]*)\))?)?` + // [net[(addr)]]
			`\/(?P<dbname>.*?)` + // /dbname
			`(?:\?(?P<params>[^\?]*))?$`) // [?param1=value1&paramN=valueN]
	matches := dsnPattern.FindStringSubmatch(dataSourceName)
	// tlsConfigRegister := make(map[string]*tls.Config)
	names := dsnPattern.SubexpNames()
	uri := &URI{}
	for i, match := range matches {
		switch names[i] {
		case "user":
			uri.User = match
		case "passwd":
			uri.Passwd = match
		case "net":
			uri.Laddr = match
		case "dbname":
			uri.DbName = match
		case "params":
			if len(match) > 0 {
				kvs := strings.Split(match, "&")
				for _, kv := range kvs {
					splits := strings.Split(kv, "=")
					if len(splits) == 2 {
						switch splits[0] {
						case "charset":
							uri.Charset = splits[1]
						}
					}
				}
			}
		}
	}
	return uri, nil
}
