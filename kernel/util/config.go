package util

import (
	"flag"
	"strings"
)

type BilladmConfig struct {
	Port     string // 服务器监听端口
	LogLevel string // 日志级别
	LogFile  string // 日志文件路径
	Mode     string // 运行模式
}

var Config BilladmConfig

// NewBilladmConfigFromFlags 解析命令行标志并返回一个配置对象
func NewBilladmConfigFromFlags() error {
	portPtr := flag.String("port", "31943", "服务器监听端口 (默认: 31943)")
	logLevelPtr := flag.String("log_level", "info", "日志级别 (debug, info, warn, warning, error) (默认: info)")
	logFilePtr := flag.String("log_file", "", "日志文件路径 (默认: 标准输出)")
	modePtr := flag.String("mode", "debug", "billadm的运行模式 (debug, release) (默认: debug)")

	flag.Parse()

	Config = BilladmConfig{
		Port:     *portPtr,
		LogLevel: strings.ToLower(*logLevelPtr),
		LogFile:  *logFilePtr,
		Mode:     *modePtr,
	}

	return nil
}
