package dto

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
)

func JsonQueryCondition(c *gin.Context, result *models.Result) (*QueryCondition, bool) {
	ret := &QueryCondition{
		Offset:  -1,
		Limit:   -1,
		TsRange: make([]int64, 0),
	}
	if err := c.BindJSON(ret); nil != err {
		result.Code = -1
		result.Msg = fmt.Sprintf("parses request failed: %v", err)
		return nil, false
	}
	return ret, true
}

type QueryCondition struct {
	LedgerID string  `json:"ledger_id"`
	Offset   int     `json:"offset"`
	Limit    int     `json:"limit"`
	TsRange  []int64 `json:"ts_range"`
}

func (qc *QueryCondition) Validate(result *models.Result) bool {
	if len(qc.LedgerID) == 0 {
		result.Code = -1
		result.Msg = fmt.Sprintf("invalid ledger id: %s", qc.LedgerID)
		return false
	}
	return true
}
