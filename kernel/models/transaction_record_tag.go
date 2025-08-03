package models

type TrTag struct {
	TransactionID string `gorm:"not null;comment:交易ID" json:"transaction_id"`
	Tag           string `gorm:"not null;comment:标签名称" json:"tag"`
}

func (tr *TrTag) TableName() string {
	return "tbl_billadm_transaction_record_tag"
}
