package file

import "os"

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
