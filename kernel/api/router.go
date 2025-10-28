package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
)

func ServeAPI(ginServer *gin.Engine) {
	// app
	ginServer.POST("/api/v1/app/exit-app", exitApp)
	// ledger
	ginServer.POST("/api/v1/ledger/query-all", queryAllLedgers)
	ginServer.POST("/api/v1/ledger/post", createLedger)
	ginServer.POST("/api/v1/ledger/put", updateLedger)
	ginServer.POST("/api/v1/ledger/delete", deleteLedger)
	// transaction record
	ginServer.POST("/api/v1/tr/query", queryTrOnCondition)
	ginServer.POST("/api/v1/tr/query-count", queryTrCountOnCondition)
	ginServer.POST("/api/v1/tr/create", createTransactionRecord)
	ginServer.POST("/api/v1/tr/put", updateTransactionRecord)
	ginServer.POST("/api/v1/tr/delete-by-id", deleteTransactionRecordById)
	// category
	ginServer.GET("/api/v1/category/query/:type", queryCategoryByType)
	// tag
	ginServer.GET("/api/v1/tag/query/:category", queryTagsByCategory)
	// workspace
	ginServer.POST("/api/v1/workspace/open", openWorkspace)
	ginServer.POST("/api/v1/workspace/is-opened", hasOpenedWorkspace)
	// statistics
	ginServer.POST("/api/v1/statistics/tr", getTrStatistics)
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
