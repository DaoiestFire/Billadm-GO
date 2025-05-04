package server

import (
	"github.com/gin-gonic/gin"

	"github.com/billadm/kernel/api"
)

func NewGinServer() *gin.Engine {
	server := gin.Default()
	api.ServeAPI(server)
	return server
}
