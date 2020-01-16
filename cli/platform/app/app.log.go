package app

import (
	"github.com/2637309949/dolphin/cli/packages/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

func init() {
	if !terminal.IsTerminal(unix.Stdout) {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006/01/02 15:04:05",
		})
	}
}
