package dao

import (
	"context"
	"errors"
	"github.com/google/uuid"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/domain"
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
	UpdateActivity(ctx context.Context, activity model.Activity) error
	CancelActivity(ctx context.Context, activity model.Activity) error
	DeleteActivity(ctx context.Context, id int64) error
	FindActivityByID(ctx context.Context, id int64) (model.Activity, error)
	ActivityList(ctx context.Context, req dtoV1.ActivityListReq) ([]model.Activity, int64, error)
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
				SponsorID: activity.SponsorID,
				UUID:      uuid.NewString(),
				Biz:       domain.ReviewBizActivity.String(),
				BizID:     activity.ID,
				Status:    domain.ReviewStatusPendingReview.Uint(),
			}).Error
		}
		return err
	})
}

// UpdateActivity 修改活动信息
func (dao *GormActivityDao) UpdateActivity(ctx context.Context, activity model.Activity) error {
	return dao.db.WithContext(ctx).Where("id = ?", activity.ID).Updates(&activity).Error
}

// CancelActivity 取消活动
func (dao *GormActivityDao) CancelActivity(ctx context.Context, activity model.Activity) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(model.Activity{}).Where("id = ?", activity.ID).Update("status = ?", domain.ActivityStatusCancel).Error
		if err != nil {
			return err
		}
		err = tx.Model(model.Review{}).Where("biz = ? and biz_id = ?", domain.ReviewBizActivity, activity.ID).Update("status = ?", domain.ReviewStatusReviewCancel).Error
		if err != nil {
			return err
		}
		if activity.GroupID > 0 {
			return tx.Where(model.Group{}).Where("id = ?", activity.GroupID).Update("status = ?", domain.GroupStatusDisband).Error
		}
		return nil
	})
}

// DeleteActivity 删除活动
func (dao *GormActivityDao) DeleteActivity(ctx context.Context, id int64) error {
	return dao.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Activity{}).Error
}

// FindActivityByID 根据id获取活动信息
func (dao *GormActivityDao) FindActivityByID(ctx context.Context, id int64) (model.Activity, error) {
	var res model.Activity
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if utils.IsRecordNotFoundError(err) {
		return model.Activity{}, ErrActivityNotFound
	}
	return res, err
}

// ActivityList 活动列表
func (dao *GormActivityDao) ActivityList(ctx context.Context, req dtoV1.ActivityListReq) ([]model.Activity, int64, error) {
	var (
		res   = make([]model.Activity, 0)
		count int64
	)

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
		WithPagination(req.PageNum, req.PageSize).DB.Find(&res).Offset(0).Limit(-1).Count(&count).Error
	return res, count, err
}

// FindActivityCount 获取总条数
//func (dao *GormActivityDao) FindActivityCount(ctx context.Context, req dtoV1.SearchActivityReq) (int64, error) {
//	var count int64
//	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx)).
//		WithLike("title", req.SearchKey).
//		WithLike("desc", req.SearchKey).
//		WithEqual("age_restrict", req.AgeRestrict).
//		WithEqual("gender_restrict", req.GenderRestrict).
//		WithEqual("cost_restrict", req.CostRestrict).
//		WithEqual("visibility", req.Visibility).
//		WithEqual("Address", req.Address).
//		WithEqual("category", req.Category).
//		WithGte("start_time", req.StartTime).
//		WithLte("start_time", req.EndTime).
//		WithEqual("visibility", req.Status).
//		DB.Count(&count).Error
//	return count, err
//}
