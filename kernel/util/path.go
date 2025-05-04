package util

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetTestDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filepath.Dir(filename))
	return filepath.Join(dir, "test")
}

// GetRootDir billadm的后端二进制目录的服务即软件的根目录
func GetRootDir() string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// 解析符号链接以获取真实路径
	realPath, err := filepath.EvalSymlinks(exe)
	if err != nil {
		panic(err)
	}

	return filepath.Dir(filepath.Dir(realPath))
}

// GetConfDir 根目录下的配置目录conf
func GetConfDir() string {
	return filepath.Join(GetRootDir(), "conf")
}
