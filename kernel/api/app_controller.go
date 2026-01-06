package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
	"github.com/billadm/workspace"
)

func exitApp(c *gin.Context) {
	ret := models.NewResult()
	defer c.JSON(http.StatusOK, ret)
	go func() {
		workspace.Manager.Close()
		time.Sleep(500 * time.Millisecond)
		os.Exit(0)
	}()
}
