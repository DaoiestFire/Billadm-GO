// Package logger
package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())[:4]
	file := entry.Data["file"]
	line := entry.Data["line"]
	msg := entry.Message
	fl := fmt.Sprintf("[%s:%v]", file, line)

	logLine := fmt.Sprintf("%s [%s] %s %s\n",
		timestamp,
		level,
		fl,
		msg,
	)
	return []byte(logLine), nil
}

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&CustomFormatter{})
}

// Init 初始化日志配置
func Init(level, path, name string) error {
	// 设置日志级别
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(logLevel)

	// 设置输出目标
	if path != "" && name != "" {
		fullPath := filepath.Join(path, name)
		file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		log.SetOutput(file)
	}
	return nil
}

// 获取调用者信息
func getCaller(skip int) (file string, line int) {
	_, filePath, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0
	}
	return filepath.Base(filePath), line
}

func Debug(format string, args ...interface{}) {
	if log.IsLevelEnabled(logrus.DebugLevel) {
		file, line := getCaller(2)
		log.WithFields(logrus.Fields{
			"file": file,
			"line": line,
		}).Debugf(format, args...)
	}
}

func Info(format string, args ...interface{}) {
	if log.IsLevelEnabled(logrus.InfoLevel) {
		file, line := getCaller(2)
		log.WithFields(logrus.Fields{
			"file": file,
			"line": line,
		}).Infof(format, args...)
	}
}

func Warn(format string, args ...interface{}) {
	if log.IsLevelEnabled(logrus.WarnLevel) {
		file, line := getCaller(2)
		log.WithFields(logrus.Fields{
			"file": file,
			"line": line,
		}).Warnf(format, args...)
	}
}

func Error(format string, args ...interface{}) {
	if log.IsLevelEnabled(logrus.ErrorLevel) {
		file, line := getCaller(2)
		log.WithFields(logrus.Fields{
			"file": file,
			"line": line,
		}).Errorf(format, args...)
	}
}
