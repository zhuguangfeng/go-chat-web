package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"gorm.io/gorm"
)

type ActivityDao interface {
	InsertActivity(ctx context.Context, activity model.Activity, review model.Review) error
	UpdateActivity(ctx context.Context, activity model.Activity, review model.Review) error
	DeleteActivity(ctx context.Context, id int64) error
	DetailActivity(ctx context.Context, id int64) (model.Activity, error)
	ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]model.Activity, error)
	FindActivityCount(ctx context.Context, searchKey string) (int64, error)
}

type GormActivityDao struct {
	db *gorm.DB
}

func NewActivityDao(db *gorm.DB) ActivityDao {
	return &GormActivityDao{
		db: db,
	}
}

// InsertActivity 插入活动
func (dao *GormActivityDao) InsertActivity(ctx context.Context, activity model.Activity, review model.Review) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&activity).Error
		if err == nil {
			review.BizID = activity.ID
			return tx.Create(&review).Error
		}
		return err
	})
}

// UpdateActivity 修改活动
func (dao *GormActivityDao) UpdateActivity(ctx context.Context, activity model.Activity, review model.Review) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", activity.ID).Updates(&activity).Error
		if err == nil {
			return tx.Where("biz = ? and biz_id = ?", "activity", activity.ID).Updates(&review).Error
		}
		return err
	})
}

// DeleteActivity 删除活动
func (dao *GormActivityDao) DeleteActivity(ctx context.Context, id int64) error {
	return dao.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Activity{}).Error
}

// DetailActivity 活动详情
func (dao *GormActivityDao) DetailActivity(ctx context.Context, id int64) (model.Activity, error) {
	var res model.Activity
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	return res, err
}

// ListActivity 活动列表
func (dao *GormActivityDao) ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]model.Activity, error) {
	var res = make([]model.Activity, 0)

	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
		WithLike("title", searchKey).
		WithLike("desc", searchKey).
		WithPagination(pageNum, pageNum).DB.Find(&res).Error
	return res, err
}

func (dao *GormActivityDao) FindActivityCount(ctx context.Context, searchKey string) (int64, error) {
	var count int64
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
		WithLike("title", searchKey).
		WithLike("desc", searchKey).
		DB.Count(&count).Error
	return count, err
}
