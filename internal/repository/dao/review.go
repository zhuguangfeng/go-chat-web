package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"gorm.io/gorm"
)

type ReviewDao interface {
	DetailReview(ctx context.Context, uuid string) (model.Review, error)
	ListReview(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]model.Review, error)
	FindReviewCount(ctx context.Context, biz string, status uint) (int64, error)
	ReviewActivity(ctx context.Context, review model.Review, group model.Group) error
}

type GormReviewDao struct {
	db *gorm.DB
}

func NewReviewDao(db *gorm.DB) ReviewDao {
	return &GormReviewDao{
		db: db,
	}
}

// DetailReview 审核详情
func (dao *GormReviewDao) DetailReview(ctx context.Context, uuid string) (model.Review, error) {
	var res model.Review
	err := dao.db.WithContext(ctx).Where("uuid = ?", uuid).First(&res).Error
	if err != nil {
		return model.Review{}, err
	}
	return res, nil
}

// ListReview 审核列表
func (dao *GormReviewDao) ListReview(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]model.Review, error) {
	var res = make([]model.Review, 0)
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
		WithEqual("biz", biz).
		WithEqual("status", status).
		WithPagination(pageNum, pageSize).DB.Find(&res).Error
	return res, err
}

// FindReviewCount 获取总条数
func (dao *GormReviewDao) FindReviewCount(ctx context.Context, biz string, status uint) (int64, error) {
	var count int64
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx).Model(&model.Review{})).
		WithEqual("biz", biz).
		WithEqual("status", status).DB.Count(&count).Error
	return count, err
}

// ReviewActivity 审核活动
func (dao *GormReviewDao) ReviewActivity(ctx context.Context, review model.Review, group model.Group) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := dao.updateReview(tx, review)
		if err != nil {
			return err
		}
		if review.Status == common.ReviewStatusSuccess.Uint() {
			return tx.Create(&group).Error
		}
		return nil
	})
}

// updateReview 修改审核
func (dao *GormReviewDao) updateReview(tx *gorm.DB, review model.Review) error {
	return tx.Where("uuid = ?", review.UUID).Updates(&review).Error
}
