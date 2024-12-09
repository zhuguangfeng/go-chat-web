package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/model"
	"gorm.io/gorm"
)

type GroupUserDao interface {
}

type GormGroupUserDao struct {
	db *gorm.DB
}

func NewGroupUserDao(db *gorm.DB) GroupUserDao {
	return &GormGroupUserDao{
		db: db,
	}
}

// InsertGroupUser 插入群聊用户关联关系
func (dao *GormGroupDao) InsertGroupUser(ctx context.Context, group model.GroupUserMap) error {
	return dao.db.WithContext(ctx).Create(&group).Error
}
