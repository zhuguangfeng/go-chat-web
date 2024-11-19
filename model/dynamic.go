package model

const TableNameDynamic = "dynamic"

type Dynamic struct {
	Base
	UserID      int64    `gorm:"column:user_id" json:"userId"`
	Title       string   `gorm:"column:title" json:"title"`
	Media       []string `gorm:"media" json:"media"`
	Tags        []int64  `gorm:"tags" json:"tags"`
	Visibility  int64    `gorm:"visibility" json:"visibility"`
	DynamicType int64    `gorm:"dynamic_type" json:"dynamicType"`
	Status      int64    `gorm:"status" json:"status"`
}

func (Dynamic) TableName() string {
	return TableNameDynamic
}
