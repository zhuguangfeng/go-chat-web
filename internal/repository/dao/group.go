package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/model"
	"gorm.io/gorm"
)

type GroupDao interface {
}

type GormGroupDao struct {
	db *gorm.DB
}

func NewGroupDao(db *gorm.DB) GroupDao {
	return &GormGroupDao{
		db: db,
	}
}

// InsertGroup 创建群聊
func (dao *GormGroupDao) InsertGroup(ctx context.Context, tx *gorm.DB, group model.Group) error {
	return tx.WithContext(ctx).Create(&group).Error
}
