package activity

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository"
)

type ActivityService interface {
	CreateActivity(ctx context.Context, activity domain.Activity) error
	ChangeActivity(ctx context.Context, activity domain.Activity) error
	CancelActivity(ctx context.Context, activityID int64) error
	ActivityDetail(ctx context.Context, id int64) (domain.Activity, error)
	ActivityList(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error)
}

type activityService struct {
	activityRepo repository.ActivityRepository
	userRepo     repository.UserRepository
}

func NewActivityService(activityRepo repository.ActivityRepository, userRepo repository.UserRepository) ActivityService {
	return &activityService{
		activityRepo: activityRepo,
		userRepo:     userRepo,
	}

}

// CreateActivity 创建活动
func (svc *activityService) CreateActivity(ctx context.Context, activity domain.Activity) error {
	return svc.activityRepo.CreateActivity(ctx, activity, domain.Review{
		UUID:   uuid.New().String(),
		Biz:    "activity",
		Status: common.ReviewStatusPendingReview.Uint(),
	})
}

// ChangeActivity 修改活动信息
func (svc *activityService) ChangeActivity(ctx context.Context, activity domain.Activity) error {
	activity, err := svc.activityRepo.DetailActivity(ctx, activity.ID)
	if err != nil {
		return err
	}

	if activity.Status != common.ActivityStatusPendingReview.Uint() {
		return errors.New("已经审核通过的活动不能修改")
	}

	return svc.activityRepo.UpdateActivity(ctx, activity, domain.Review{
		Biz:    "activity",
		BizID:  activity.ID,
		Status: common.ReviewStatusReviewCancel.Uint(),
	})
}

// 取消活动
func (svc *activityService) CancelActivity(ctx context.Context, activityID int64) error {
	activity, err := svc.activityRepo.DetailActivity(ctx, activityID)
	if err != nil {
		return err
	}

	if activity.Status != common.ActivityStatusReviewPass.Uint() && activity.Status != common.ActivityStatusSignUp.Uint() {
		return errors.New("该活动暂不能取消")
	}

	return svc.activityRepo.UpdateActivity(ctx, domain.Activity{
		ID:     activityID,
		Status: common.ActivityStatusCancel.Uint(),
	}, domain.Review{
		Biz:    "activity",
		BizID:  activityID,
		Status: common.ReviewStatusReviewCancel.Uint(),
	})
}

// ActivityDetail 活动详情
func (svc *activityService) ActivityDetail(ctx context.Context, id int64) (domain.Activity, error) {
	return svc.activityRepo.DetailActivity(ctx, id)
}

// ActivityList 活动列表
func (svc *activityService) ActivityList(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Activity, int64, error) {

	activitys, count, err := svc.activityRepo.ListActivity(ctx, pageNum, pageSize, searchKey)
	if err != nil {
		return nil, 0, err
	}

	if len(activitys) > 0 {

		userIds := make([]int64, len(activitys))
		for _, activity := range activitys {
			userIds = append(userIds, activity.Sponsor.ID)
		}

		userMap, err := svc.userRepo.GetUsersByIDs(ctx, userIds)
		if err != nil {
			return nil, 0, err
		}

		for i, activity := range activitys {
			activitys[i].Sponsor = userMap[activity.Sponsor.ID]
		}
	}

	return activitys, count, nil
}
