package util

import (
	"os"
	"path/filepath"
)

func GetRootDir() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

// GetConfDir 根目录下的配置目录conf
func GetConfDir() string {
	return filepath.Join(GetRootDir(), "conf")
}

func GetDistDir() string {
	return filepath.Join(GetRootDir(), "dist")
}
