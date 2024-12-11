package repository

import (
	"context"
	"errors"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

type ActivityRepository interface {
	CreateActivity(ctx context.Context, activity domain.Activity) error
	UpdateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error
	DeleteActivity(ctx context.Context, id int64) error
	DetailActivity(ctx context.Context, id int64) (domain.Activity, error)
	ListActivity(ctx context.Context, req dtoV1.SearchActivityReq) ([]domain.Activity, int64, error)
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

func (repo *activityRepository) CreateActivity(ctx context.Context, activity domain.Activity) error {
	err := repo.activityDao.InsertActivity(ctx, repo.toActivityEntity(activity))
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
		if errors.Is(err, dao.ErrActivityNotFound) {
			return domain.Activity{}, errorx.NewBizError(common.ActivityNotFound).WithError(err)
		}
		return domain.Activity{}, errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return repo.toActivityDomain(activity), nil
}

func (repo *activityRepository) ListActivity(ctx context.Context, req dtoV1.SearchActivityReq) ([]domain.Activity, int64, error) {

	activityEsResp, err := repo.activityEsDao.SearchActivity(ctx, req)
	if err == nil {
		return slice.Map(activityEsResp, func(idx int, src model.ActivityEs) domain.Activity {
			return repo.esToActivityDomain(src)
		}), 0, nil
	}

	repo.logger.Error("[ActivityRepository.ListActivity]从es获取活动记录失败", logger.Error(err))

	activitys, err := repo.activityDao.ListActivity(ctx, req)
	if err == nil {
		count, err := repo.activityDao.FindActivityCount(ctx, req)
		if err != nil {
			repo.logger.Error("[activity.repository.list] 获取活动列表总条数失败",
				logger.String("searchKey", req.SearchKey),
				logger.Error(err),
			)
		}
		return slice.Map(activitys, func(idx int, src model.Activity) domain.Activity {
			return repo.toActivityDomain(src)
		}), count, nil

	}
	return nil, 0, errorx.NewBizError(common.SystemInternalError).WithError(err)
}

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
		Category:            activity.Category,
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status,
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
		Category:            activity.Category,
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status,
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
		Category:            activity.Category,
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status,
		CreatedTime:         activity.CreatedTime,
		UpdatedTime:         activity.UpdatedTime,
	}
}
