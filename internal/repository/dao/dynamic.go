package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/model"
	"gorm.io/gorm"
)

type DynamicDao interface {
	InsertDynamic(ctx context.Context, dynamic model.Dynamic) error
	DeleteDynamic(ctx context.Context, id int64) error
	UpdateDynamic(ctx context.Context, dynamic model.Dynamic) error
	ListDynamic(ctx context.Context, pageNum, pageSize int, query []query) ([]model.Dynamic, error)
	FindDynamicCount(ctx context.Context, query []query) (int64, error)
}

type GormDynamicDao struct {
	db *gorm.DB
}

func NewDynamicDao(db *gorm.DB) DynamicDao {
	return &GormDynamicDao{
		db: db,
	}
}

// InsertDynamic 插入动态到db
func (dao *GormDynamicDao) InsertDynamic(ctx context.Context, dynamic model.Dynamic) error {
	return dao.db.Create(&dynamic).Error
}

// DeleteDynamic 从db根据id删除动态
func (dao *GormDynamicDao) DeleteDynamic(ctx context.Context, id int64) error {
	return dao.db.WithContext(ctx).Where("id = ?", id).Delete(model.Dynamic{}).Error
}

// UpdateDynamic 修改动态
func (dao *GormDynamicDao) UpdateDynamic(ctx context.Context, dynamic model.Dynamic) error {
	return dao.db.WithContext(ctx).Where("id = ?", dynamic.ID).Updates(&dynamic).Error
}

// ListDynamic 获取动态列表
func (dao *GormDynamicDao) ListDynamic(ctx context.Context, pageNum, pageSize int, query []query) ([]model.Dynamic, error) {
	var res = make([]model.Dynamic, 0)
	err := NewDaoBuilder(dao.db.WithContext(ctx)).WithQuery(query).WithPagination((pageNum-1)*pageSize, pageSize).db.Scan(&res).Error
	return res, err
}

// FindDynamicCount 获取动态总条数
func (dao *GormDynamicDao) FindDynamicCount(ctx context.Context, query []query) (int64, error) {
	var count int64
	err := NewDaoBuilder(dao.db.WithContext(ctx)).WithQuery(query).db.Count(&count).Error
	return count, err
}
