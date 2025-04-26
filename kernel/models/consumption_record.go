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
	TransactionID string `gorm:"primaryKey;type:char(36);comment:交易UUID"`
	LedgerID      string `gorm:"type:char(36);not null;index;comment:关联账本ID"`

	// 交易核心信息
	Price           float64         `gorm:"type:double;not null;comment:交易金额"`
	TransactionType TransactionType `gorm:"type:varchar(16);check:transaction_type IN ('income','expense','transfer');not null;index;comment:交易类型"`

	// 分类与描述
	CategoryID  string `gorm:"type:char(36);not null;index;comment:分类ID"`
	Description string `gorm:"type:text;comment:交易描述"`
	Tags        JSON   `gorm:"type:json;comment:标签列表"`

	// 时间信息
	TransactionAt time.Time `gorm:"type:datetime;not null;index;comment:交易时间"`
	CreatedAt     time.Time `gorm:"type:datetime;autoCreateTime;not null;comment:创建时间"`
	UpdatedAt     time.Time `gorm:"type:datetime;autoUpdateTime;not null;comment:更新时间"`
}
