package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-errors/errors"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

// SearchFileInDirWithSuffix defined
func SearchFileInDirWithSuffix(wd string, suffix string, cb func(string) bool) string {
	var file string
	filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, suffix) {
			ext := cb(path)
			if ext {
				file = path
				return errors.New("has found")
			}
		}
		return nil
	})
	return file
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

// NetWorkStatus defined
func NetWorkStatus() bool {
	if err := exec.Command("ping", "baidu.com", "-c", "1", "-W", "5").Run(); err != nil {
		return false
	}
	return true
}

// InstallPackages defined
func InstallPackages(pkgs ...string) error {
	for i := range pkgs {
		if err := exec.Command("go", "get", pkgs[i]).Run(); err != nil && err != exec.ErrNotFound {
			return err
		}
	}
	return nil
}

func HasBin(bins ...string) bool {
	for i := range bins {
		if err := exec.Command(bins[i]).Run(); err == exec.ErrNotFound {
			return false
		}
	}
	return true
}

// EnsureLeft defined return left
func EnsureLeft(left interface{}, err error) interface{} {
	if err != nil {
		panic(fmt.Errorf("%v", string(errors.Wrap(err, 3).Stack())))
	}
	return left
}

// EnsureRight defined return right
func EnsureRight(err error, right interface{}) interface{} {
	if err != nil {
		panic(fmt.Errorf("%v", string(errors.Wrap(err, 3).Stack())))
	}
	return right
}

// Ensure defined
func Ensure(err error) {
	if err != nil {
		panic(fmt.Errorf("%v", string(errors.Wrap(err, 3).Stack())))
	}
}
