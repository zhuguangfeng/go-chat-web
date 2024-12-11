package dao

import (
	"context"
	"errors"
	"github.com/google/uuid"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
	"gorm.io/gorm"
)

var (
	ErrActivityNotFound = errors.New("activity not found")
)

type ActivityDao interface {
	InsertActivity(ctx context.Context, activity model.Activity) error
	UpdateActivity(ctx context.Context, activity model.Activity, review model.Review) error
	DeleteActivity(ctx context.Context, id int64) error
	DetailActivity(ctx context.Context, id int64) (model.Activity, error)
	ListActivity(ctx context.Context, req dtoV1.SearchActivityReq) ([]model.Activity, error)
	FindActivityCount(ctx context.Context, req dtoV1.SearchActivityReq) (int64, error)
}

type GormActivityDao struct {
	db *gorm.DB
}

func NewActivityDao(db *gorm.DB) ActivityDao {
	return &GormActivityDao{
		db: db,
	}
}

// InsertActivity 插入活动并插入审核
func (dao *GormActivityDao) InsertActivity(ctx context.Context, activity model.Activity) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&activity).Error
		if err == nil {
			return tx.Create(&model.Review{
				UUID:   uuid.NewString(),
				Biz:    common.ReviewBizActivity,
				BizID:  activity.ID,
				Status: common.ReviewStatusPendingReview.Uint(),
			}).Error
		}
		return err
	})

}

// UpdateActivity 修改活动并修改审核信息
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
	if utils.IsRecordNotFoundError(err) {
		return model.Activity{}, ErrActivityNotFound
	}
	return res, err
}

// ListActivity 活动列表
func (dao *GormActivityDao) ListActivity(ctx context.Context, req dtoV1.SearchActivityReq) ([]model.Activity, error) {
	var res = make([]model.Activity, 0)
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
		WithLike("title", req.SearchKey).
		WithLike("desc", req.SearchKey).
		WithEqual("age_restrict", req.AgeRestrict).
		WithEqual("gender_restrict", req.GenderRestrict).
		WithEqual("cost_restrict", req.CostRestrict).
		WithEqual("visibility", req.Visibility).
		WithEqual("Address", req.Address).
		WithEqual("category", req.Category).
		WithGte("start_time", req.StartTime).
		WithLte("start_time", req.EndTime).
		WithEqual("visibility", req.Status).
		WithPagination(req.PageNum, req.PageSize).DB.Find(&res).Error
	return res, err
}

// FindActivityCount 获取总条数
func (dao *GormActivityDao) FindActivityCount(ctx context.Context, req dtoV1.SearchActivityReq) (int64, error) {
	var count int64
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
		WithLike("title", req.SearchKey).
		WithLike("desc", req.SearchKey).
		WithEqual("age_restrict", req.AgeRestrict).
		WithEqual("gender_restrict", req.GenderRestrict).
		WithEqual("cost_restrict", req.CostRestrict).
		WithEqual("visibility", req.Visibility).
		WithEqual("Address", req.Address).
		WithEqual("category", req.Category).
		WithGte("start_time", req.StartTime).
		WithLte("start_time", req.EndTime).
		WithEqual("visibility", req.Status).
		DB.Count(&count).Error
	return count, err
}
