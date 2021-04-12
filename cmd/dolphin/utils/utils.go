package utils

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/2637309949/dolphin/packages/cobra"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/thoas/go-funk"
)

var timeFormat = "2006/01/02 15:04:05"

// ISBlank defined
func ISBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// ViperSetDefault defined
func ViperSetDefault(key string, value interface{}, standby ...interface{}) {
	isBlank := ISBlank(reflect.ValueOf(value))
	if isBlank {
		value = standby[0]
	}
	viper.SetDefault(key, value)
}

// WalkFileInDirWithSuffix defined
func WalkFileInDirWithSuffix(wd string, suffix string) ([]string, error) {
	var files []string
	if err := filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, suffix) {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return []string{}, err
	}
	return files, nil
}

// RemoveFileWithSuffix defined
func RemoveFileWithSuffix(wd string, suffix string) {
	files, _ := WalkFileInDirWithSuffix(wd, suffix)
	funk.ForEach(files, func(file string) {
		os.Remove(file)
	})
}

// SetFormatter defined
func SetFormatter(isTerminal bool) {
	if !isTerminal {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeFormat,
		})
	}
}

// SetLevel defined
func SetLevel(cmd *cobra.Command) {
	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// OpenFile defiend
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if err := os.MkdirAll(path.Dir(name), os.ModePerm); err != nil {
		return nil, err
	}
	return os.OpenFile(name, flag, perm)
}

// InstallPackages defined
func InstallPackages(pkgs ...string) error {
	for i := range pkgs {
		if err := exec.Command("go", "get", pkgs[i]).Run(); err != nil && err != exec.ErrNotFound {
			logrus.Error(err)
			return err
		}
	}
	return nil
}
