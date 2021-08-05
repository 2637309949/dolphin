package file

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

const (
	// NotExist represents that the file or directory does not exist.
	NotExist = iota
	// FileType represents a file.
	FileType
	// DirType represents a directory.
	DirType
)

// Getwd returns a rooted path name corresponding to the current directory
func Getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

// Type decides the type of a file.
//
// It returns FileType, DirType, or NotExist.
func Type(name string) uint8 {
	fi, err := os.Stat(name)
	if err == nil {
		if fi.IsDir() {
			return DirType
		}
		return FileType
	}
	if os.IsNotExist(err) {
		return NotExist
	}
	return FileType
}

// IsExist returns true if the file or directory exists, or return false.
func IsExist(filename string) bool {
	if Type(filename) == NotExist {
		return false
	}
	return true
}

// IsFile returns true if it's a file, or return false.
func IsFile(filename string) bool {
	if Type(filename) == FileType {
		return true
	}
	return false
}

// IsDir returns true if it's a directory, or return false.
func IsDir(filename string) bool {
	if Type(filename) == DirType {
		return true
	}
	return false
}

// EnsureDir create dir if not exist
func EnsureDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
	}
}

// MustHash cac file hash
func MustHash(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	h := sha256.New()
	_, err = io.Copy(h, br)

	if err != nil {
		panic(err)
	}
	return []byte(fmt.Sprintf("%x", h.Sum(nil)))
}

// RemoveExt defined
func RemoveExt(filePath string) string {
	return filePath[0:strings.Index(filePath, path.Ext(filePath))]
}
