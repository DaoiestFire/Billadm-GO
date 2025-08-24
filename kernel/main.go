package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/constant"
	"github.com/billadm/logger"
	"github.com/billadm/server"
	"github.com/billadm/util"
	"github.com/billadm/util/db"
)

func main() {
	// 解析命令行选项
	err := util.NewBilladmConfigFromFlags()
	if err != nil {
		logrus.Fatalf("解析命令行选项 %v", err)
	}
	// 初始化日志配置
	if err := logger.Init(util.Config.LogLevel, util.GetRootDir(), util.Config.LogFile); err != nil {
		logrus.Fatalf("初始化日志模块失败 %v", err)
	}
	// 初始化数据库连接
	dbPath := filepath.Join(util.GetRootDir(), constant.DbName)
	if err := db.Init(dbPath); err != nil {
		logrus.Fatalf("初始化数据库连接失败 %v", err)
	}
	// 启动服务器
	logrus.Info("--------- start billadm ---------")
	ginServer := server.NewGinServer()
	gin.SetMode(util.Config.Mode)
	if err := ginServer.Run("127.0.0.1:" + util.Config.Port); err != nil {
		logrus.Errorf("start billadm err: %v", err)
		return
	}
}
