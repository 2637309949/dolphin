package utils

import "path/filepath"

// FileNameTrimSuffix defined
func FileNameTrimSuffix(fp string) string {
	extension, filename := filepath.Ext(fp), filepath.Base(fp)
	filename = filename[0 : len(filename)-len(extension)]
	return filename
}
