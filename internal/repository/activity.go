package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
)

type ActivityRepository interface {
	CreateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error
	UpdateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error
	DeleteActivity(ctx context.Context, id int64) error
	DetailActivity(ctx context.Context, id int64) (domain.Activity, error)
	ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error)
}

type activityRepository struct {
	activityDao dao.ActivityDao
	reviewDao   dao.ReviewDao
}

func NewActivityRepository(activityDao dao.ActivityDao, reviewDao dao.ReviewDao) ActivityRepository {
	return &activityRepository{
		activityDao: activityDao,
		reviewDao:   reviewDao,
	}
}

func (repo *activityRepository) CreateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error {
	return repo.activityDao.InsertActivity(ctx, repo.toActivityEntity(activity), repo.toReviewEntity(review))
}

func (repo *activityRepository) UpdateActivity(ctx context.Context, activity domain.Activity, review domain.Review) error {
	return repo.activityDao.UpdateActivity(ctx, repo.toActivityEntity(activity), repo.toReviewEntity(review))
}

func (repo *activityRepository) DeleteActivity(ctx context.Context, id int64) error {
	return repo.activityDao.DeleteActivity(ctx, id)
}

func (repo *activityRepository) DetailActivity(ctx context.Context, id int64) (domain.Activity, error) {
	activity, err := repo.activityDao.DetailActivity(ctx, id)
	if err != nil {
		return domain.Activity{}, err
	}
	return repo.toDomain(activity), nil
}

func (repo *activityRepository) ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error) {
	activitys, err := repo.activityDao.ListActivity(ctx, pageNum, pageSize, searchKey)
	if err == nil {
		count, err := repo.activityDao.FindActivityCount(ctx, searchKey)
		if err != nil {
			//TODO

		}
		return slice.Map(activitys, func(idx int, src model.Activity) domain.Activity {
			return repo.toDomain(src)
		}), count, nil

	}
	return nil, 0, err
}

func (repo *activityRepository) toActivityEntity(activity domain.Activity) model.Activity {
	return model.Activity{
		Base: model.Base{
			ID: activity.ID,
		},
		SponsorID: activity.Sponsor.ID,
	}
}

func (repo *activityRepository) toReviewEntity(review domain.Review) model.Review {
	return model.Review{
		Base: model.Base{
			ID: review.ID,
		},
		UUID:   review.UUID,
		Biz:    review.Biz,
		BizID:  review.BizID,
		Status: review.Status,
	}
}

func (repo *activityRepository) toDomain(activity model.Activity) domain.Activity {
	return domain.Activity{
		ID: activity.ID,
		Sponsor: domain.User{
			ID: activity.ID,
		},
	}
}
