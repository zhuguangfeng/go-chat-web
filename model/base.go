package model

import (
	"gorm.io/plugin/soft_delete"
)

type Base struct {
	ID        int64                 `gorm:"primarykey"         json:"id"`
	CreatedAt uint                  `gorm:"column:created_at"  json:"createdAt"`
	UpdatedAt uint                  `gorm:"column:updated_at;autoUpdateTime;"  json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at"  json:"deletedAt"`
}
