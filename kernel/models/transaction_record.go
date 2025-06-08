package models

import (
	"time"
)

// TransactionType 交易类型枚举
type TransactionType string

const (
	Income   TransactionType = "income"
	Expense  TransactionType = "expense"
	Transfer TransactionType = "transfer"
)

// TransactionRecord 消费记录结构体
type TransactionRecord struct {
	TransactionID string `gorm:"primaryKey;comment:交易UUID" json:"transaction_id"`
	LedgerID      string `gorm:"not null;index;comment:关联账本ID" json:"ledger_id"`

	// 交易核心信息
	Price           float64         `gorm:"not null;comment:交易金额" json:"price"`
	TransactionType TransactionType `gorm:"not null;index;comment:交易类型" json:"transaction_type"`

	// 分类与描述
	CategoryID  string      `gorm:"not null;index;comment:分类ID" json:"category_id"`
	Description string      `gorm:"comment:交易描述" json:"description"`
	Tags        StringSlice `gorm:"comment:标签列表" json:"tags"`

	// 时间信息
	TransactionAt time.Time `gorm:"type:datetime;not null;index;comment:交易时间" json:"transaction_at"`
	CreatedAt     time.Time `gorm:"type:datetime;autoCreateTime;not null;comment:创建时间" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:datetime;autoUpdateTime;not null;comment:更新时间" json:"updated_at"`
}

func (tr *TransactionRecord) TableName() string {
	return "tbl_billadm_transaction_record"
}
