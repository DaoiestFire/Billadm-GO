package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
	"github.com/billadm/server"
)

func exitApp(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)
	server.NewExitManager().Exit()
}
