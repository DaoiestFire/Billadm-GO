package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
)

func ServeAPI(ginServer *gin.Engine) {
	// app
	ginServer.POST("/api/v1/app/exit_app", exitApp)
	// ledger
	ginServer.POST("/api/v1/ledger/query_all_ledger", queryAllLedgers)
	ginServer.POST("/api/v1/ledger/post", createLedger)
	ginServer.POST("/api/v1/ledger/put", updateLedger)
	ginServer.POST("/api/v1/ledger/delete", deleteLedger)
	// transaction record
	ginServer.POST("/api/v1/tr/query_all_tr", queryAllTrs)
	ginServer.POST("/api/v1/tr/query_trs_on_condition", queryTrsOnCondition)
	ginServer.POST("/api/v1/tr/query_tr_count_on_condition", queryTrCountOnCondition)
	ginServer.POST("/api/v1/tr/create_tr", createTransactionRecord)
	ginServer.POST("/api/v1/tr/put", updateTransactionRecord)
	ginServer.POST("/api/v1/tr/delete_tr_by_id", deleteTransactionRecord)
	// category
	ginServer.GET("/api/v1/category/query_category/:type", queryCategory)
	// tag
	ginServer.GET("/api/v1/tag/query_tag/:category", queryTags)
	// workspace
	ginServer.POST("/api/v1/workspace/open_workspace", openWorkspace)
	ginServer.POST("/api/v1/workspace/is_opened", hasOpenedWorkspace)
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
