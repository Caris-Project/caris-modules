package utils

import (
	"os"
	"path"
)

// InitFolder ...
func InitFolder() {
	dir, _ := os.Getwd()
	uploadFilePath := path.Join(dir + "/assets")
	createFolder(uploadFilePath)
}

// createFolder ...
func createFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		_ = os.MkdirAll(path, os.ModePerm)
	}
}