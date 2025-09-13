package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/api"
	"github.com/billadm/logger"
	"github.com/billadm/server"
	"github.com/billadm/util"
	"github.com/billadm/workspace"
)

func main() {
	var err error
	// 解析命令行选项
	err = util.NewBilladmConfigFromFlags()
	if err != nil {
		logrus.Fatalf("解析命令行选项 %v", err)
	}
	// 设置日志级别
	err = logger.Init(util.Config.LogLevel)
	if err != nil {
		logrus.Fatalf("初始化日志模块失败 %v", err)
	}
	// 初始化工作空间
	err = workspace.Manager.InitFromConfig(util.Config.HomeDir)
	if err != nil {
		logrus.Fatalf("初始化工作空间失败 %v", err)
	}
	// 启动优雅退出
	server.NewExitManager().Start()
	// 启动服务器
	logrus.Info("--------- 启动Billadm ---------")
	gin.SetMode(util.Config.Mode)
	ginServer := server.NewGinServer()
	// 注册接口
	api.ServeAPI(ginServer)
	if err := ginServer.Run("127.0.0.1:" + util.Config.Port); err != nil {
		logrus.Errorf("启动Billadm失败 %v", err)
		return
	}
}
