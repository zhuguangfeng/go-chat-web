package activity

import (
	"context"
	"errors"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type ActivityService interface {
	CreateActivity(ctx context.Context, activity domain.Activity) error
	ChangeActivity(ctx context.Context, activity domain.Activity) error
	CancelActivity(ctx context.Context, activityID int64) error
	ActivityDetail(ctx context.Context, id int64) (domain.Activity, error)
	ActivityList(ctx context.Context, req dtoV1.SearchActivityReq) ([]domain.Activity, int64, error)
	SignUpActivity(ctx context.Context, activitySuDomain domain.ActivitySignup) error
	CancelSignup(ctx context.Context) error
	ReviewSignup(ctx context.Context, signup domain.ActivitySignup) error
}

type activityService struct {
	logger             logger.Logger
	activityRepo       repository.ActivityRepository
	activitySignupRepo repository.ActivitySignupRepository1
	userRepo           repository.UserRepository
	reviewRepo         repository.ReviewRepository
}

func NewActivityService(logger logger.Logger,
	activityRepo repository.ActivityRepository,
	userRepo repository.UserRepository,
	reviewRepo repository.ReviewRepository,
	activitySignupRepo repository.ActivitySignupRepository1,
) ActivityService {
	return &activityService{
		logger:             logger,
		activityRepo:       activityRepo,
		activitySignupRepo: activitySignupRepo,
		userRepo:           userRepo,
		reviewRepo:         reviewRepo,
	}

}

// CreateActivity 创建活动
func (svc *activityService) CreateActivity(ctx context.Context, activity domain.Activity) error {
	return svc.activityRepo.CreateActivity(ctx, activity)
}

// ChangeActivity 修改活动信息
func (svc *activityService) ChangeActivity(ctx context.Context, activity domain.Activity) error {
	activity, err := svc.activityRepo.DetailActivity(ctx, activity.ID)
	if err != nil {
		return err
	}

	if activity.Status != common.ActivityStatusPendingReview.Uint() {
		return errorx.NewBizError(common.ActivityNotChange).WithError(errors.New("已经审核通过的活动不能修改"))
	}

	return svc.activityRepo.UpdateActivity(ctx, activity, domain.Review{
		Biz:   "activity",
		BizID: activity.ID,
	})
}

// 取消活动
func (svc *activityService) CancelActivity(ctx context.Context, activityID int64) error {
	activity, err := svc.activityRepo.DetailActivity(ctx, activityID)
	if err != nil {
		return err
	}

	if activity.Status != common.ActivityStatusPendingReview.Uint() && activity.Status != common.ActivityStatusSignUp.Uint() {
		return errorx.NewBizError(common.ActivityNotChange).WithError(nil)
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
func (svc *activityService) ActivityList(ctx context.Context, req dtoV1.SearchActivityReq) ([]domain.Activity, int64, error) {

	activitys, count, err := svc.activityRepo.ListActivity(ctx, req)
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

// SignUpActivity 活动报名
func (svc *activityService) SignUpActivity(ctx context.Context, activitySuDomain domain.ActivitySignup) error {
	return svc.activitySignupRepo.CreateActivitySignup(ctx, activitySuDomain)
}

// CancelSignup 取消活动报名
func (svc *activityService) CancelSignup(ctx context.Context) error {
	return nil
}

// ReviewSignup 审核报名
func (svc *activityService) ReviewSignup(ctx context.Context, signup domain.ActivitySignup) error {
	var (
		eg       errgroup.Group
		su       domain.ActivitySignup
		activity domain.Activity
		err      error
	)

	switch signup.Status {
	case common.ReviewStatusSuccess.Uint():
		//审核通过  假如群聊

		eg.Go(func() error {
			su, err = svc.activitySignupRepo.GetActivitySignupByID(ctx, signup.ID)
			return err
		})
		eg.Go(func() error {
			activity, err = svc.activityRepo.DetailActivity(ctx, signup.ID)
			return err
		})
		if err = eg.Wait(); err != nil {
			return err
		}

		signup.Activity.Group.ID = activity.Group.ID
		signup.Applicant.ID = su.Applicant.ID

		return svc.activitySignupRepo.UpdateActivitySignup(ctx, signup)
	case common.ReviewStatusPass.Uint():
		//申请拒绝 修改状态
		err = svc.activitySignupRepo.UpdateActivitySignup(ctx, signup)
		if err != nil {
			return err
		}
	default:

	}

	return nil
}
