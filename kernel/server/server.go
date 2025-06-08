package server

import (
	"github.com/billadm/kernel/api"
	"github.com/billadm/kernel/constant"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewGinServer() *gin.Engine {
	logrus.Infof("start server with mode: %s", constant.Mode)
	gin.SetMode(constant.Mode)
	server := gin.Default()
	api.ServeAPI(server)
	return server
}
