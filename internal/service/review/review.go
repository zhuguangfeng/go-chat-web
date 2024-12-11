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
	ReviewCreateActivity(ctx context.Context, review domain.Review) error
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

// ImplementReview 审核
func (svc *reviewService) ReviewCreateActivity(ctx context.Context, review domain.Review) error {
	//验证是否审核过 可以前端来审核
	rvw, err := svc.reviewRepo.DetailReview(ctx, review.UUID)
	if err != nil {
		return err
	}

	if rvw.Status != common.ReviewStatusPendingReview.Uint() {
		return errorx.NewBizError(common.ReviewNotReview)
	}

	//获取activity
	activity, err := svc.activityRepo.DetailActivity(ctx, rvw.BizID)
	if err != nil {
		svc.logger.Error("[ReviewService.ImplementReview]获取活动详情失败 并且未将活动数据写入到es")
		return nil
	}

	err = svc.reviewRepo.ReviewActivity(ctx, review, domain.Group{
		GroupName: activity.Title + "_活动群",
		Owner: domain.User{
			ID: activity.Sponsor.ID,
		},
		MaxPeopleNumber: activity.MaxPeopleNumber,
		PeopleNumber:    0,
		Status:          common.GroupStatusNormal.Uint(),
	})
	if err != nil {
		return err
	}

	//如果审核通过 发送写入到es事件
	if review.Status == common.ReviewStatusSuccess.Uint() {
		go func() {
			activity, err := svc.activityRepo.DetailActivity(ctx, rvw.BizID)
			if err != nil {
				svc.logger.Error("[ReviewService.ImplementReview]获取活动详情失败 并且未将活动数据写入到es")
				return
			}
			if review.Status == common.ActivityStatusSignUp.Uint() {
				err := svc.activityEvent.ProducerSyncActivityEvent(ctx, activityEvent.ToEvent(activity))
				if err != nil {
					svc.logger.Error("[ReviewService.ImplementReview]发送同步es消息失败",
						logger.Int64("activityID", activity.ID),
						logger.Error(err))
				}
			}
		}()
	}

	return nil
}

func (svc *reviewService) ReviewDetail(ctx context.Context, uuid string) (domain.Review, error) {
	review, err := svc.reviewRepo.DetailReview(ctx, uuid)
	if err != nil {
		return domain.Review{}, err
	}

	switch review.Biz {
	case common.ReviewBizActivity:
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
