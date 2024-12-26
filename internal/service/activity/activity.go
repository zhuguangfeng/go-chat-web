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

var (
	ErrActivityNotFound  = repository.ErrActivityNotFound
	ErrActivityNotChange = errors.New("activity not change")
)

type ActivityService interface {
	CreateActivity(ctx context.Context, activity domain.Activity) error
	ChangeActivity(ctx context.Context, activity domain.Activity) error
	CancelActivity(ctx context.Context, activityID int64) error
	ActivityDetail(ctx context.Context, id int64) (domain.Activity, error)
	ActivityList(ctx context.Context, req dtoV1.ActivityListReq) ([]domain.Activity, int64, error)
	SignUpActivity(ctx context.Context, activitySuDomain domain.ActivitySignup) error
	CancelSignup(ctx context.Context, signup domain.ActivitySignup) error
	ReviewSignup(ctx context.Context, signup domain.ActivitySignup) error
	SignupList(ctx context.Context, req dtoV1.SignUpListReq) ([]domain.ActivitySignup, int64, error)
}

type activityService struct {
	logger             logger.Logger
	activityRepo       repository.ActivityRepository
	activitySignupRepo repository.ActivitySignupRepository1
	userRepo           repository.UserRepository
	reviewRepo         repository.ReviewRepository
}

func NewActivityService(
	logger logger.Logger,
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
	atv, err := svc.activityRepo.GetActivity(ctx, activity.ID)
	if err != nil {
		return err
	}

	if atv.Status != domain.ActivityStatusPendingReview {
		return errorx.NewBizError(common.ActivityNotChange)
	}

	return svc.activityRepo.UpdateActivity(ctx, activity)
}

// 取消活动
func (svc *activityService) CancelActivity(ctx context.Context, activityID int64) error {
	activity, err := svc.activityRepo.GetActivity(ctx, activityID)
	if err != nil {
		return err
	}

	//TODO 可以扣取保证金或者添加其他逻辑

	return svc.activityRepo.CancelActivity(ctx, activity)

}

// ActivityDetail 活动详情
func (svc *activityService) ActivityDetail(ctx context.Context, id int64) (domain.Activity, error) {
	return svc.activityRepo.GetActivity(ctx, id)
}

// ActivityList 活动列表
func (svc *activityService) ActivityList(ctx context.Context, req dtoV1.ActivityListReq) ([]domain.Activity, int64, error) {
	activitys, count, err := svc.activityRepo.ActivityList(ctx, req)
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
func (svc *activityService) CancelSignup(ctx context.Context, signup domain.ActivitySignup) error {
	su, err := svc.activitySignupRepo.GetActivitySignup(ctx, signup.ID)
	if err != nil {
		return err
	}

	if su.Status == domain.ActivitySignupStatusReviewPass {
		return nil
	}

	if su.Status == domain.ActivitySignupStatusReviewSuccess {
		return svc.activitySignupRepo.CancelReviewSuccessSignup(ctx, su.ID, su.Activity.Group.ID)

	}
	return svc.activitySignupRepo.UpdateActivitySignup(ctx, domain.ActivitySignup{
		ID:     signup.ID,
		Status: domain.ActivitySignupStatusCancelReview,
	})

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
	case domain.ActivitySignupStatusReviewSuccess:
		//审核通过  假如群聊
		eg.Go(func() error {
			su, err = svc.activitySignupRepo.GetActivitySignup(ctx, signup.ID)
			return err
		})
		eg.Go(func() error {
			activity, err = svc.activityRepo.GetActivity(ctx, signup.ID)
			return err
		})
		if err = eg.Wait(); err != nil {
			return err
		}

		signup.Activity.Group.ID = activity.Group.ID
		signup.Applicant.ID = su.Applicant.ID

		return svc.activitySignupRepo.ReviewSignupSuccess(ctx, signup)
	case domain.ActivitySignupStatusReviewPass:
		//申请拒绝 修改状态
		return svc.activitySignupRepo.UpdateActivitySignup(ctx, signup)
	default:

	}

	return nil
}

func (svc *activityService) SignupList(ctx context.Context, req dtoV1.SignUpListReq) ([]domain.ActivitySignup, int64, error) {
	signups, count, err := svc.activitySignupRepo.GetSignupList(ctx, req)
	if err != nil {
		svc.logger.Error("活动报名列表获取报名信息失败", logger.Error(err))
		return nil, 0, err
	}

	var userIDs = make([]int64, len(signups))
	for i := range signups {
		userIDs = append(userIDs, signups[i].Applicant.ID)
	}

	users, err := svc.userRepo.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		svc.logger.Error("活动报名列表获取用户信息失败", logger.Error(err))
		return nil, 0, err
	}

	for i := range signups {
		signups[i].Applicant = users[signups[i].Applicant.ID]
	}

	return signups, count, nil
}
