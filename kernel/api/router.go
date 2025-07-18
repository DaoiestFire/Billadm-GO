package api

import (
	"fmt"
	
	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
)

func ServeAPI(ginServer *gin.Engine) {
	// ledger
	ginServer.GET("/api/v1/ledger", getLedger)
	ginServer.POST("/api/v1/ledger", createLedger)
	ginServer.PUT("/api/v1/ledger", updateLedger)
	ginServer.DELETE("/api/v1/ledger", deleteLedger)
	// transaction record
	ginServer.GET("/api/v1/tr", getTransactionRecord)
	ginServer.POST("/api/v1/tr", createTransactionRecord)
	ginServer.PUT("/api/v1/tr", updateTransactionRecord)
	ginServer.DELETE("/api/v1/tr", deleteTransactionRecord)
}

func JsonArg(c *gin.Context, result *models.Result) (arg map[string]interface{}, ok bool) {
	arg = map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = fmt.Sprintf("parses request failed: %v", err)
		return
	}

	ok = true
	return
}
