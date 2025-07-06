package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/logger"
	"github.com/billadm/server"
	"github.com/billadm/util"
	"github.com/billadm/util/db"
)

func main() {
	// 初始化配置管理器
	if err := util.InitConfigManager(); err != nil {
		logrus.Fatalf("初始化配置管理器失败 %v", err)
	}
	// 初始化日志配置
	logLevel := util.GetConfigManager().Get(util.LogLevel, "info")
	logFile := util.GetConfigManager().Get(util.LogFile, "billadm.log")
	if err := logger.Init(logLevel, util.GetLogDir(), logFile); err != nil {
		logrus.Fatalf("初始化日志模块失败 %v", err)
	}
	// 初始化数据库连接
	var dbName, dbPath string
	dbName = util.GetConfigManager().Get(util.DbName, "billadm.db")
	if util.IsDebugMode() {
		dbPath = filepath.Join(util.GetTestDir(), dbName)
	} else {
		dbPath = filepath.Join(util.GetRootDir(), dbName)
	}
	if err := db.Init(dbPath); err != nil {
		logrus.Fatalf("初始化数据库连接失败 %v", err)
	}
	// 启动服务器
	logrus.Info("--------- start billadm ---------")
	port := util.GetConfigManager().Get(util.Port, "31943")
	ginServer := server.NewGinServer()
	if !util.IsDebugMode() {
		gin.SetMode(gin.ReleaseMode)
	}
	err := ginServer.Run(":" + port)
	if err != nil {
		logrus.Errorf("start billadm err: %v", err)
		return
	}
	logrus.Info("--------- end billadm ---------")
}
