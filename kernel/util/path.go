package util

import (
	"os"
	"path/filepath"
)

func GetRootDir() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

func GetDistDir() string {
	return filepath.Join(GetRootDir(), "dist")
}
