package repository

import (
	"context"
	"errors"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

var (
	ErrActivityNotFound = dao.ErrActivityNotFound
)

type ActivityRepository interface {
	CreateActivity(ctx context.Context, activity domain.Activity) error
	UpdateActivity(ctx context.Context, activity domain.Activity) error
	CancelActivity(ctx context.Context, activity domain.Activity) error
	DeleteActivity(ctx context.Context, id int64) error
	GetActivity(ctx context.Context, id int64) (domain.Activity, error)
	ActivityList(ctx context.Context, req dtoV1.ActivityListReq) ([]domain.Activity, int64, error)
	InputActivity(ctx context.Context, activity domain.Activity) error
}

type activityRepository struct {
	logger        logger.Logger
	activityDao   dao.ActivityDao
	reviewDao     dao.ReviewDao
	activityEsDao dao.ActivityEsDao
}

func NewActivityRepository(logger logger.Logger, activityDao dao.ActivityDao, reviewDao dao.ReviewDao, activityEsDao dao.ActivityEsDao) ActivityRepository {
	return &activityRepository{
		logger:        logger,
		activityDao:   activityDao,
		reviewDao:     reviewDao,
		activityEsDao: activityEsDao,
	}
}

// CreateActivity 创建活动
func (repo *activityRepository) CreateActivity(ctx context.Context, activity domain.Activity) error {
	return repo.activityDao.InsertActivity(ctx, repo.toActivityEntity(activity))
}

// UpdateActivity 修改活动
func (repo *activityRepository) UpdateActivity(ctx context.Context, activity domain.Activity) error {
	return repo.activityDao.UpdateActivity(ctx, repo.toActivityEntity(activity))
}

// CancelActivity 取消活动
func (repo *activityRepository) CancelActivity(ctx context.Context, activity domain.Activity) error {
	return repo.activityDao.CancelActivity(ctx, repo.toActivityEntity(activity))
}

// DeleteActivity 删除活动
func (repo *activityRepository) DeleteActivity(ctx context.Context, id int64) error {
	return repo.activityDao.DeleteActivity(ctx, id)
}

// GetActivity 获取活动信息
func (repo *activityRepository) GetActivity(ctx context.Context, id int64) (domain.Activity, error) {
	activity, err := repo.activityDao.FindActivityByID(ctx, id)
	if err != nil && errors.Is(err, dao.ErrActivityNotFound) {
		return domain.Activity{}, ErrActivityNotFound
	}
	return repo.toActivityDomain(activity), err
}

// ActivityList 活动列表
func (repo *activityRepository) ActivityList(ctx context.Context, req dtoV1.ActivityListReq) ([]domain.Activity, int64, error) {
	activityEsResp, err := repo.activityEsDao.SearchActivity(ctx, req)
	if err == nil {
		return slice.Map(activityEsResp, func(idx int, src model.ActivityEs) domain.Activity {
			return repo.esToActivityDomain(src)
		}), 0, nil
	}

	repo.logger.Error("从es获取活动记录失败", logger.Error(err))

	activitys, count, err := repo.activityDao.ActivityList(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return slice.Map(activitys, func(idx int, src model.Activity) domain.Activity {
		return repo.toActivityDomain(src)
	}), count, nil

}

// InputActivity 同步活动
func (repo *activityRepository) InputActivity(ctx context.Context, activity domain.Activity) error {
	return repo.activityEsDao.InputActivity(ctx, repo.toActivityEs(activity))
}

func (repo *activityRepository) toActivityEntity(activity domain.Activity) model.Activity {
	return model.Activity{
		Base: model.Base{
			ID: activity.ID,
		},
		SponsorID:           activity.Sponsor.ID,
		GroupID:             activity.Group.ID,
		Title:               activity.Title,
		Desc:                activity.Desc,
		Media:               activity.Media,
		AgeRestrict:         activity.AgeRestrict,
		GenderRestrict:      activity.GenderRestrict,
		CostRestrict:        activity.CostRestrict,
		Visibility:          activity.Visibility,
		MaxPeopleNumber:     activity.MaxPeopleNumber,
		CurrentPeopleNumber: activity.CurrentPeopleNumber,
		Address:             activity.Address,
		Category:            activity.Category.Uint(),
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status.Uint(),
	}
}

func (repo *activityRepository) toReviewEntity(review domain.Review) model.Review {
	return model.Review{
		Base: model.Base{
			ID: review.ID,
		},
		UUID:       review.UUID,
		Biz:        review.Biz.String(),
		BizID:      review.BizID,
		ReviewerID: review.Reviewer.ID,
		Status:     review.Status.Uint(),
		ReviewTime: review.ReviewTime,
	}
}

func (repo *activityRepository) toActivityDomain(activity model.Activity) domain.Activity {
	return domain.Activity{
		ID: activity.ID,
		Sponsor: domain.User{
			ID: activity.SponsorID,
		},
		Group: domain.Group{
			ID: activity.GroupID,
		},
		Title:               activity.Title,
		Desc:                activity.Desc,
		Media:               activity.Media,
		AgeRestrict:         activity.AgeRestrict,
		GenderRestrict:      activity.GenderRestrict,
		CostRestrict:        activity.CostRestrict,
		Visibility:          activity.Visibility,
		MaxPeopleNumber:     activity.MaxPeopleNumber,
		CurrentPeopleNumber: activity.CurrentPeopleNumber,
		Address:             activity.Address,
		Category:            domain.ActivityCategory(activity.Category),
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              domain.ActivityStatus(activity.Status),
		CreatedTime:         activity.CreatedAt,
		UpdatedTime:         activity.UpdatedAt,
	}
}

func (repo *activityRepository) toActivityEs(activity domain.Activity) model.ActivityEs {
	return model.ActivityEs{
		ID:                  activity.ID,
		SponsorID:           activity.Sponsor.ID,
		Title:               activity.Title,
		Desc:                activity.Desc,
		Media:               activity.Media,
		AgeRestrict:         activity.AgeRestrict,
		GenderRestrict:      activity.GenderRestrict,
		CostRestrict:        activity.CostRestrict,
		Visibility:          activity.Visibility,
		MaxPeopleNumber:     activity.MaxPeopleNumber,
		CurrentPeopleNumber: activity.CurrentPeopleNumber,
		Address:             activity.Address,
		Category:            activity.Category.Uint(),
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status.Uint(),
		CreatedTime:         activity.CreatedTime,
		UpdatedTime:         activity.UpdatedTime,
	}
}

func (repo *activityRepository) esToActivityDomain(activity model.ActivityEs) domain.Activity {
	return domain.Activity{
		ID: activity.ID,
		Sponsor: domain.User{
			ID: activity.SponsorID,
		},
		Title:               activity.Title,
		Desc:                activity.Desc,
		Media:               activity.Media,
		AgeRestrict:         activity.AgeRestrict,
		GenderRestrict:      activity.GenderRestrict,
		CostRestrict:        activity.CostRestrict,
		Visibility:          activity.Visibility,
		MaxPeopleNumber:     activity.MaxPeopleNumber,
		CurrentPeopleNumber: activity.CurrentPeopleNumber,
		Address:             activity.Address,
		Category:            domain.ActivityCategory(activity.Category),
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              domain.ActivityStatus(activity.Status),
		CreatedTime:         activity.CreatedTime,
		UpdatedTime:         activity.UpdatedTime,
	}
}
