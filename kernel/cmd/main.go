package main

import (
	"github.com/billadm/kernel/logger"
	"github.com/billadm/kernel/server"
	"github.com/billadm/kernel/util"
)

func main() {
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
