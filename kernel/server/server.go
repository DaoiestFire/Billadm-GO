package server

import (
	"github.com/gin-gonic/gin"

	"github.com/billadm/kernel/api"
	"github.com/billadm/kernel/constant"
	"github.com/billadm/kernel/logger"
)

func NewGinServer() *gin.Engine {
	logger.Info("start server with mode: %s", constant.Mode)
	gin.SetMode(constant.Mode)
	server := gin.Default()
	api.ServeAPI(server)
	return server
}
