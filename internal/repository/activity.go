package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
)

type ActivityRepository interface {
	CreateActivity(ctx context.Context, activity domain.Activity) error
	UpdateActivity(ctx context.Context, activity domain.Activity) error
	DeleteActivity(ctx context.Context, id int64) error
	DetailActivity(ctx context.Context, id int64) (domain.Activity, error)
	ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error)
}

type activityRepository struct {
	dao dao.ActivityDao
}

func NewActivityRepository(dao dao.ActivityDao) ActivityRepository {
	return &activityRepository{
		dao: dao,
	}
}

func (repo *activityRepository) CreateActivity(ctx context.Context, activity domain.Activity) error {
	return repo.dao.InsertActivity(ctx, repo.toEntity(activity))
}

func (repo *activityRepository) UpdateActivity(ctx context.Context, activity domain.Activity) error {
	return repo.dao.UpdateActivity(ctx, repo.toEntity(activity))
}

func (repo *activityRepository) DeleteActivity(ctx context.Context, id int64) error {
	return repo.dao.DeleteActivity(ctx, id)
}

func (repo *activityRepository) DetailActivity(ctx context.Context, id int64) (domain.Activity, error) {
	activity, err := repo.dao.DetailActivity(ctx, id)
	if err != nil {
		return domain.Activity{}, err
	}
	return repo.toDomain(activity), nil
}

func (repo *activityRepository) ListActivity(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error) {
	activitys, err := repo.dao.ListActivity(ctx, pageNum, pageSize, searchKey)
	if err == nil {
		count, err := repo.dao.FindActivityCount(ctx, searchKey)
		if err == nil {
			//TODO
		}

	}
	return nil, 0, err
}

func (repo *activityRepository) toEntity(activity domain.Activity) model.Activity {
	return model.Activity{
		Base: model.Base{
			ID: activity.ID,
		},
		SponsorID: activity.Sponsor.ID,
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
