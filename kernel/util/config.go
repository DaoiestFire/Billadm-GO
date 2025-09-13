package util

import (
	"flag"
	"strings"

	"github.com/sirupsen/logrus"
)

type BilladmConfig struct {
	Port     string // 服务器监听端口
	LogLevel string // 日志级别
	Mode     string // 运行模式
	HomeDir  string
}

var Config = BilladmConfig{
	Port:     "31943",
	LogLevel: "debug",
	Mode:     "debug",
	HomeDir:  "",
}

// NewBilladmConfigFromFlags 解析命令行标志并返回一个配置对象
func NewBilladmConfigFromFlags() error {
	portPtr := flag.String("port", "31943", "服务器监听端口 (默认: 31943)")
	logLevelPtr := flag.String("log_level", "info", "日志级别 (debug, info, warn, warning, error) (默认: info)")
	modePtr := flag.String("mode", "debug", "billadm的运行模式 (debug, release) (默认: debug)")
	homeDirPtr := flag.String("home_dir", "", "用户的家目录路径")

	flag.Parse()

	Config = BilladmConfig{
		Port:     *portPtr,
		LogLevel: strings.ToLower(*logLevelPtr),
		Mode:     *modePtr,
		HomeDir:  *homeDirPtr,
	}

	logrus.Warnf("port: %s, log_level: %s, mode: %s, home_dir: %s", Config.Port, Config.LogLevel, Config.Mode, Config.HomeDir)

	return nil
}
