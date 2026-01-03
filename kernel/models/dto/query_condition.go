package dto

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
)

func JsonQueryCondition(c *gin.Context, result *models.Result) (*TrQueryCondition, bool) {
	ret := &TrQueryCondition{
		Offset:           -1,
		Limit:            -1,
		TsRange:          make([]int64, 0),
		TransactionTypes: make([]string, 0),
		CategoryTags:     make(map[string][]string),
	}
	if err := c.BindJSON(ret); nil != err {
		result.Code = -1
		result.Msg = fmt.Sprintf("解析消费记录查询条件失败: %v", err)
		return nil, false
	}
	return ret, true
}

type TrQueryCondition struct {
	LedgerID         string              `json:"ledgerId"`
	Offset           int                 `json:"offset"`
	Limit            int                 `json:"limit"`
	TsRange          []int64             `json:"tsRange"`
	TransactionTypes []string            `json:"transactionTypes"`
	CategoryTags     map[string][]string `json:"categoryTags"`
}

func (qc *TrQueryCondition) Validate(result *models.Result) bool {
	if len(qc.LedgerID) == 0 {
		result.Code = -1
		result.Msg = fmt.Sprintf("账本Id不可为空: %s", qc.LedgerID)
		return false
	}
	return true
}
