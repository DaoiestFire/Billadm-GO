package models

type Category struct {
	Name  string `gorm:"primaryKey;comment:消费类型" json:"name"`
	Scope string `gorm:"not null;comment:作用域" json:"scope"`
}

func (tr *Category) TableName() string {
	return "tbl_billadm_category"
}
