package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
)

func ServeAPI(ginServer *gin.Engine) {
	// ledger
	ginServer.POST("/api/v1/ledger/get", getLedger)
	ginServer.POST("/api/v1/ledger/post", createLedger)
	ginServer.POST("/api/v1/ledger/put", updateLedger)
	ginServer.POST("/api/v1/ledger/delete", deleteLedger)
	// transaction record
	ginServer.POST("/api/v1/tr/get", getTransactionRecord)
	ginServer.POST("/api/v1/tr/post", createTransactionRecord)
	ginServer.POST("/api/v1/tr/put", updateTransactionRecord)
	ginServer.POST("/api/v1/tr/delete", deleteTransactionRecord)
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
