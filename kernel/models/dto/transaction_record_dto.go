package dto

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/billadm/models"
)

func JsonTransactionRecordDto(c *gin.Context, result *models.Result) (*TransactionRecordDto, bool) {
	ret := &TransactionRecordDto{}
	if err := c.BindJSON(ret); nil != err {
		result.Code = -1
		result.Msg = fmt.Sprintf("parses request failed: %v", err)
		return nil, false
	}
	return ret, true
}

type TransactionRecordDto struct {
	LedgerID      string `json:"ledgerId"`
	TransactionID string `json:"transactionId"`

	// 交易核心信息
	Price           float64 `json:"price"`
	TransactionType string  ` json:"transactionType"`

	// 分类与描述
	Category    string   ` json:"category"`
	Description string   ` json:"description"`
	Tags        []string ` json:"tags"`

	// 时间信息
	TransactionAt int64 `json:"transactionAt"`
}

func (dto *TransactionRecordDto) Validate(result *models.Result) bool {
	// TODO:校验账本ID是否合法
	// 校验交易类型是否合法
	if dto.TransactionType != models.Income && dto.TransactionType != models.Expense && dto.TransactionType != models.Transfer {
		result.Code = -1
		result.Msg = fmt.Sprintf("invalid transaction type: %s", dto.TransactionType)
		return false
	}
	// TODO: 校验类型ID是否合法
	return true
}

func (dto *TransactionRecordDto) ToTransactionRecord() *models.TransactionRecord {
	tr := &models.TransactionRecord{}
	tr.TransactionID = dto.TransactionID
	tr.LedgerID = dto.LedgerID
	tr.Price = dto.Price
	tr.TransactionType = dto.TransactionType
	tr.Category = dto.Category
	tr.Description = dto.Description
	tr.TransactionAt = dto.TransactionAt
	return tr
}

func (dto *TransactionRecordDto) FromTransactionRecord(tr *models.TransactionRecord) {
	dto.LedgerID = tr.LedgerID
	dto.TransactionID = tr.TransactionID
	dto.Price = tr.Price
	dto.TransactionType = tr.TransactionType
	dto.Category = tr.Category
	dto.Description = tr.Description
	dto.Tags = make([]string, 0)
	dto.TransactionAt = tr.TransactionAt
}
