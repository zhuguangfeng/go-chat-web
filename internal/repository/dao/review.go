package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"gorm.io/gorm"
)

type ReviewDao interface {
	DetailReview(ctx context.Context, id int64) (model.Review, error)
	ListReview(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]model.Review, error)
	FindReviewCount(ctx context.Context, biz string, status uint) (int64, error)
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
func (dao *GormReviewDao) DetailReview(ctx context.Context, id int64) (model.Review, error) {
	var res model.Review
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
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
		WithEqual("status", status).DB.Find(&res).Error
	return res, err
}

func (dao *GormReviewDao) FindReviewCount(ctx context.Context, biz string, status uint) (int64, error) {
	var count int64
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
		WithEqual("biz", biz).
		WithEqual("status", status).DB.Count(&count).Error
	return count, err
}
