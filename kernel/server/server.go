package server

import "github.com/gin-gonic/gin"

func NewGinServer() *gin.Engine {
	server := gin.Default()
	return server
}
