package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"gorm.io/gorm"
)

type DynamicDao interface {
	InsertDynamic(ctx context.Context, dynamic model.Dynamic) error
	DeleteDynamic(ctx context.Context, id int64, uid int64) error
	UpdateDynamic(ctx context.Context, dynamic model.Dynamic) error
	DetailDynamic(ctx context.Context, id int64) (model.Dynamic, error)
	ListDynamic(ctx context.Context, pageNum, pageSize int, searchKey string) ([]model.Dynamic, error)
	FindDynamicCount(ctx context.Context, searchKey string) (int64, error)
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
func (dao *GormDynamicDao) DeleteDynamic(ctx context.Context, id int64, uid int64) error {
	return dao.db.WithContext(ctx).Where("id = ? and uid = ?", id, uid).Delete(model.Dynamic{}).Error
}

// UpdateDynamic 修改动态
func (dao *GormDynamicDao) UpdateDynamic(ctx context.Context, dynamic model.Dynamic) error {
	return dao.db.WithContext(ctx).Where("id = ?", dynamic.ID).Updates(&dynamic).Error
}

// DetailDynamic 动态详情
func (dao *GormDynamicDao) DetailDynamic(ctx context.Context, id int64) (model.Dynamic, error) {
	var res model.Dynamic
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	return res, err
}

// ListDynamic 获取动态列表
func (dao *GormDynamicDao) ListDynamic(ctx context.Context, pageNum, pageSize int, searchKey string) ([]model.Dynamic, error) {
	var res = make([]model.Dynamic, 0)
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
		WithLike("title", searchKey).
		WithPagination((pageNum-1)*pageSize, pageSize).DB.Scan(&res).Error
	return res, err
}

// FindDynamicCount 获取动态总条数
func (dao *GormDynamicDao) FindDynamicCount(ctx context.Context, searchKey string) (int64, error) {
	var count int64
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).WithLike("title", searchKey).DB.Count(&count).Error
	return count, err
}
