package app

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/logrotate"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/util/slice"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

var receiver chan *plugin.LogFormatterParams
var bucket slice.SyncSlice

// Tracker defined tracker recorder
func Tracker(e *Engine) func(ctx *gin.Context, p *plugin.LogFormatterParams) {
	return func(ctx *gin.Context, p *plugin.LogFormatterParams) {
		token, _ := e.OAuth2.BearerAuth(ctx.Request)
		p.Token = token
		tokenInfo, err := e.OAuth2.Manager.LoadAccessToken(token)
		if err == nil {
			p.Domain = tokenInfo.GetDomain()
		}
		receiver <- p
	}
}

// collector collect logs every 5 second
func collector() {
	receiver = make(chan *plugin.LogFormatterParams, 100)
	go func() {
		for {
			select {
			case <-time.Tick(10 * time.Second):
				logs := bucket.Values()
				if len(logs) > 0 {
					// collect to db
				}
				bucket.Clear()
			case info := <-receiver:
				bucket.Append(info)
			}
		}
	}()
}

func init() {
	var writer io.Writer
	if !terminal.IsTerminal(unix.Stdout) {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006/01/02 15:04:05",
		})
		writer = os.Stdout
	}
	dir := path.Join(viper.GetString("dir.log"), viper.GetString("app.name"))
	logf, err := logrotate.New(
		dir+".%Y%m%d%H",
		logrotate.WithLinkName(dir),
		logrotate.WithMaxAge(24*time.Hour),
		logrotate.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Printf("failed to create rotatelogs: %s", err)
		return
	}
	if writer != nil {
		writer = io.MultiWriter(writer, logf)
	} else {
		writer = logf
	}
	logrus.SetOutput(writer)
	collector()
}
