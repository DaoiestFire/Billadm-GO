package models

import "time"

type Ledger struct {
	ID        string    `gorm:"primaryKey;type:char(36);comment:账本UUID"`
	UserID    string    `gorm:"type:char(36);not null;comment:用户UUID"`
	Name      string    `gorm:"type:varchar(255);not null;comment:账本名称"`
	CreatedAt time.Time `gorm:"type:datetime;autoCreateTime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;autoUpdateTime;not null;comment:更新时间"`
}

func (l *Ledger) TableName() string {
	return "tbl_billadm_ledger"
}
