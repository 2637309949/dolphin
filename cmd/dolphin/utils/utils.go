package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

var timeFormat = "2006/01/02 15:04:05"

// DeepCopyInterface defined TODO
type DeepCopyInterface interface {
	DeepCopy() interface{}
}

// Copy defined TODO
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	original := reflect.ValueOf(src)
	cpy := reflect.New(original.Type()).Elem()
	copyRecursive(original, cpy)

	return cpy.Interface()
}

// copyRecursive defined TODO
func copyRecursive(src, dst reflect.Value) {
	if src.CanInterface() {
		if copier, ok := src.Interface().(DeepCopyInterface); ok {
			dst.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	switch src.Kind() {
	case reflect.Ptr:
		originalValue := src.Elem()

		if !originalValue.IsValid() {
			return
		}
		dst.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, dst.Elem())

	case reflect.Interface:
		if src.IsNil() {
			return
		}
		originalValue := src.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		dst.Set(copyValue)

	case reflect.Struct:
		t, ok := src.Interface().(time.Time)
		if ok {
			dst.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < src.NumField(); i++ {
			if src.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(src.Field(i), dst.Field(i))
		}

	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			copyRecursive(src.Index(i), dst.Index(i))
		}

	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			originalValue := src.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			copyKey := Copy(key.Interface())
			dst.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		dst.Set(src)
	}
}

// ISBlank defined TODO
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

// ViperSetDefault defined TODO
func ViperSetDefault(key string, value interface{}, standby ...interface{}) {
	isBlank := ISBlank(reflect.ValueOf(value))
	if isBlank {
		value = standby[0]
	}
	viper.SetDefault(key, value)
}

// WalkFileInDirWithSuffix defined TODO
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

// SearchFileInDirWithSuffix defined TODO
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

// RemoveFileWithSuffix defined TODO
func RemoveFileWithSuffix(wd string, suffix string) {
	files, _ := WalkFileInDirWithSuffix(wd, suffix)
	funk.ForEach(files, func(file string) {
		os.Remove(file)
	})
}

// SetFormatter defined TODO
func SetFormatter(isTerminal bool) {
	if !isTerminal {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: timeFormat,
		})
	}
}

// SetLevel defined TODO
func SetLevel(cmd *cobra.Command) {
	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// OpenFile defiend TODO
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if err := os.MkdirAll(path.Dir(name), os.ModePerm); err != nil {
		return nil, err
	}
	return os.OpenFile(name, flag, perm)
}

// NetWorkStatus defined TODO
func NetWorkStatus() bool {
	if err := exec.Command("ping", "baidu.com", "-c", "1", "-W", "5").Run(); err != nil {
		return false
	}
	return true
}

// InstallPackages defined TODO
func InstallPackages(pkgs ...string) error {
	for i := range pkgs {
		var stderr bytes.Buffer
		cmd := exec.Command("go", "get", pkgs[i])
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("%v", stderr.String())
		}
	}
	return nil
}

// HasBin defined TODO
func HasBin(bins ...string) bool {
	for i := range bins {
		if err := exec.Command(bins[i]).Run(); err != nil && strings.Contains(err.Error(), exec.ErrNotFound.Error()) {
			return false
		}
	}
	return true
}

// EnsureLeft defined TODO
func EnsureLeft(left interface{}, err error) interface{} {
	if err != nil {
		panic(fmt.Errorf("%v", string(errors.Wrap(err, 3).Stack())))
	}
	return left
}

// EnsureRight defined TODO
func EnsureRight(err error, right interface{}) interface{} {
	if err != nil {
		panic(fmt.Errorf("%v", string(errors.Wrap(err, 3).Stack())))
	}
	return right
}

// Ensure defined TODO
func Ensure(err error) {
	if err != nil {
		panic(fmt.Errorf("%v", string(errors.Wrap(err, 3).Stack())))
	}
}

// FileNameTrimSuffix defined TODO
func FileNameTrimSuffix(fp string) string {
	extension, filename := filepath.Ext(fp), filepath.Base(fp)
	filename = filename[0 : len(filename)-len(extension)]
	return filename
}

// HasSuffix defined suffix of string
func HasSuffix(s string, suffix ...string) bool {
	for _, v := range suffix {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}
