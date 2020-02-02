package app

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/2637309949/dolphin/packages/logrotate"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

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
}
