package api

import "github.com/gin-gonic/gin"

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
