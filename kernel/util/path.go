package util

import (
	"os"
	"path/filepath"
)

func GetRootDir() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}

	return filepath.Dir(exePath)
}

func GetDistDir() string {
	return filepath.Join(GetRootDir(), "dist")
}
