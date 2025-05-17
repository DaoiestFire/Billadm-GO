package main

import (
	"fmt"
	"github.com/billadm/kernel/util/db"
	"path/filepath"

	"github.com/billadm/kernel/logger"
	"github.com/billadm/kernel/server"
	"github.com/billadm/kernel/util"
)

func main() {
	// 初始化配置管理器
	if err := util.InitConfigManager(); err != nil {
		fmt.Printf("初始化配置管理器失败 %v\n", err)
	}
	// 初始化日志配置
	logLevel := util.GetConfigManager().Get("log_level", "info")
	logFile := util.GetConfigManager().Get("log_file", "billadm.log")
	if err := logger.Init(logLevel, util.GetLogDir(), logFile); err != nil {
		fmt.Printf("初始化日志模块失败 %v\n", err)
	}
	// 初始化数据库连接
	dbName := util.GetConfigManager().Get("db_name", "billadm.db")
	var dbPath string
	if util.IsDebugMode() {
		dbPath = filepath.Join(util.GetTestDir(), dbName)
	} else {
		dbPath = filepath.Join(util.GetRootDir(), dbName)
	}
	if err := db.Init(dbPath); err != nil {
		fmt.Printf("初始化数据库连接失败 %v\n", err)
	}
	logger.Info("--------- start billadm ---------")
	port := util.GetConfigManager().Get("port", "31943")
	ginServer := server.NewGinServer()
	err := ginServer.Run(":" + port)
	if err != nil {
		logger.Error("start billadm err: %v", err)
		return
	}
	logger.Info("--------- end billadm ---------")
}
