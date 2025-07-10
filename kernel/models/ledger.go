package models

import "time"

type Ledger struct {
	ID        string    `gorm:"primaryKey;comment:账本UUID" json:"id"`
	Name      string    `gorm:"not null;comment:账本名称" json:"name"`
	CreatedAt time.Time `gorm:"type:datetime;autoCreateTime;not null;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;autoUpdateTime;not null;comment:更新时间" json:"updated_at"`
}

func (l *Ledger) TableName() string {
	return "tbl_billadm_ledger"
}
