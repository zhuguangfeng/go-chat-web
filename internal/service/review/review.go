package review

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	activityEvent "github.com/zhuguangfeng/go-chat/internal/event/activity"
	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

type ReviewService interface {
	ImplementReview(ctx context.Context, review domain.Review) error
	ReviewDetail(ctx context.Context, uuid string) (domain.Review, error)
	ReviewList(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]domain.Review, int64, error)
}

type reviewService struct {
	logger        logger.Logger
	reviewRepo    repository.ReviewRepository
	activityRepo  repository.ActivityRepository
	activityEvent activityEvent.Producer
}

func NewReviewService(logger logger.Logger, reviewRepo repository.ReviewRepository, activityRepo repository.ActivityRepository, activityEvent activityEvent.Producer) ReviewService {
	return &reviewService{
		logger:        logger,
		reviewRepo:    reviewRepo,
		activityRepo:  activityRepo,
		activityEvent: activityEvent,
	}
}

func (svc *reviewService) ImplementReview(ctx context.Context, review domain.Review) error {
	rvw, err := svc.reviewRepo.DetailReview(ctx, review.UUID)
	if err != nil {
		return err
	}

	if rvw.Status != common.ReviewStatusPendingReview.Uint() {
		return errorx.NewBizError(common.ReviewNotReview)
	}

	err = svc.reviewRepo.ChangeReview(ctx, review)
	if err != nil {
		return err
	}

	switch rvw.Biz {
	case "activity":
		activity, err := svc.activityRepo.DetailActivity(ctx, rvw.BizID)
		if err != nil {
			svc.logger.Error("[Review.Service.ImplementReview]获取活动详情失败")
			return nil
		}
		if review.Status == common.ActivityStatusSignUp.Uint() {
			err := svc.activityEvent.ProducerSyncActivityEvent(ctx, activityEvent.ToEvent(activity))
			if err != nil {
				svc.logger.Error("[review.service.implementReview]发送同步es消息失败",
					logger.Int64("activityID", activity.ID),
					logger.Error(err))
				return errorx.NewBizError(common.SystemInternalError).WithError(err)
			}
		}

	}
	return nil
}

func (svc *reviewService) ReviewDetail(ctx context.Context, uuid string) (domain.Review, error) {
	review, err := svc.reviewRepo.DetailReview(ctx, uuid)
	if err != nil {
		return domain.Review{}, err
	}

	switch review.Biz {
	case "activity":
		activity, err := svc.activityRepo.DetailActivity(ctx, review.ID)
		if err != nil {
			return domain.Review{}, err
		}
		review.Activity = activity
	}

	return review, err
}

func (svc *reviewService) ReviewList(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]domain.Review, int64, error) {
	reviews, count, err := svc.reviewRepo.ListReview(ctx, pageNum, pageSize, biz, status)
	if err != nil {
		svc.logger.Error("[ReviewService.ReviewList]获取审核列表失败", logger.Error(err))
		return nil, 0, err
	}
	return reviews, count, err
}
