package util

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/billadm/kernel/constant"
)

// GetRootDir billadm的后端二进制目录的服务即软件的根目录
func GetRootDir() string {
	if constant.Mode == gin.DebugMode {
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
