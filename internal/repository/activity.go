package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

type ActivityRepository interface {
	CreateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error
	UpdateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error
	DeleteActivity(ctx context.Context, id int64) error
	DetailActivity(ctx context.Context, id int64) (domain.Activity, error)
	ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error)
}

type activityRepository struct {
	logger      logger.Logger
	activityDao dao.ActivityDao
	reviewDao   dao.ReviewDao
}

func NewActivityRepository(logger logger.Logger, activityDao dao.ActivityDao, reviewDao dao.ReviewDao) ActivityRepository {
	return &activityRepository{
		logger:      logger,
		activityDao: activityDao,
		reviewDao:   reviewDao,
	}
}

func (repo *activityRepository) CreateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error {
	err := repo.activityDao.InsertActivity(ctx, repo.toActivityEntity(activity), repo.toReviewEntity(review))
	if err != nil {
		return errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return nil
}

func (repo *activityRepository) UpdateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error {
	err := repo.activityDao.UpdateActivity(ctx, repo.toActivityEntity(activity), repo.toReviewEntity(review))
	if err != nil {
		return errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return nil
}

func (repo *activityRepository) DeleteActivity(ctx context.Context, id int64) error {
	err := repo.activityDao.DeleteActivity(ctx, id)
	if err != nil {
		return errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return nil
}

func (repo *activityRepository) DetailActivity(ctx context.Context, id int64) (domain.Activity, error) {
	activity, err := repo.activityDao.DetailActivity(ctx, id)
	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			return domain.Activity{}, errorx.NewBizError(common.ActivityNotFound).WithError(err)
		}
		return domain.Activity{}, errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return repo.toActivityDomain(activity), nil
}

func (repo *activityRepository) ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error) {
	activitys, err := repo.activityDao.ListActivity(ctx, pageNum, pageSize, searchKey)
	if err == nil {
		count, err := repo.activityDao.FindActivityCount(ctx, searchKey)
		if err != nil {
			repo.logger.Error("[activity.repository.list] 获取活动列表总条数失败",
				logger.String("searchKey", searchKey),
				logger.Error(err),
			)
		}
		return slice.Map(activitys, func(idx int, src model.Activity) domain.Activity {
			return repo.toActivityDomain(src)
		}), count, nil

	}
	return nil, 0, errorx.NewBizError(common.SystemInternalError).WithError(err)
}

func (repo *activityRepository) toActivityEntity(activity domain.Activity) model.Activity {
	return model.Activity{
		Base: model.Base{
			ID: activity.ID,
		},
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
		Category:            activity.Category,
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status,
	}
}

func (repo *activityRepository) toReviewEntity(review domain.Review) model.Review {
	return model.Review{
		Base: model.Base{
			ID: review.ID,
		},
		UUID:       review.UUID,
		Biz:        review.Biz,
		BizID:      review.BizID,
		ReviewerID: review.Reviewer.ID,
		Status:     review.Status,
		ReviewTime: review.ReviewTime,
	}
}

func (repo *activityRepository) toActivityDomain(activity model.Activity) domain.Activity {
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
		Category:            activity.Category,
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status,
	}
}
