package util

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/billadm/constant"
)

func IsDebugMode() bool {
	return constant.Mode == gin.DebugMode
}

// GetRootDir billadm的后端二进制目录的服务即软件的根目录
func GetRootDir() string {
	if IsDebugMode() {
		_, filename, _, _ := runtime.Caller(0)
		dir := filepath.Dir(filepath.Dir(filename))
		return dir
	}
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

func GetTestDir() string {
	return filepath.Join(GetRootDir(), "test")
}

// GetConfDir 根目录下的配置目录conf
func GetConfDir() string {
	return filepath.Join(GetRootDir(), "conf")
}

// GetLogDir 根目录下的日志目录log
func GetLogDir() string {
	return filepath.Join(GetRootDir(), "log")
}

func GetDistDir() string {
	if IsDebugMode() {
		return filepath.Join(filepath.Dir(GetRootDir()), "app", "dist")
	}
	return filepath.Join(GetRootDir(), "dist")
}
