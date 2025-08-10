package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/billadm/api"
	"github.com/billadm/constant"
)

func NewGinServer() *gin.Engine {
	logrus.Infof("start server with mode: %s", constant.Mode)
	gin.SetMode(constant.Mode)
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                // 允许的源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},         // 允许的头部信息
		ExposeHeaders:    []string{"Content-Length"},                                   // 暴露的头部信息
		AllowCredentials: true,                                                         // 是否允许发送Cookie
		MaxAge:           12 * time.Hour,                                               // 预检请求的有效期
	}))
	api.ServeAPI(server)
	return server
}
