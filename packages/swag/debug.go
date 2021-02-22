package swag

import (
	"github.com/2637309949/dolphin/packages/logrus"
)

const (
	test = iota
	release
)

var swagMode = release

func isRelease() bool {
	return swagMode == release
}

// Println calls Output to print to the standard logger when release mode.
func Println(v ...interface{}) {
	if isRelease() {
		logrus.Info(v...)
	}
}

// Printf calls Output to print to the standard logger when release mode.
func Printf(format string, v ...interface{}) {
	if isRelease() {
		logrus.Infof(format, v...)
	}
}
