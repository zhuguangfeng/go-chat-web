package review

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
)

type ReviewService interface {
	ImplementReview(ctx context.Context, review domain.Review) error
	ReviewDetail(ctx context.Context, uuid string) (domain.Review, error)
	ReviewList(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]domain.Review, int64, error)
}

type reviewService struct {
	reviewRepo   repository.ReviewRepository
	activityRepo repository.ActivityRepository
}

func NewReviewService(reviewRepo repository.ReviewRepository, activityRepo repository.ActivityRepository) ReviewService {
	return &reviewService{
		reviewRepo:   reviewRepo,
		activityRepo: activityRepo,
	}
}

func (svc *reviewService) ImplementReview(ctx context.Context, review domain.Review) error {
	review, err := svc.reviewRepo.DetailReview(ctx, review.UUID)
	if err != nil {
		return err
	}

	if review.Status != common.ReviewStatusPendingReview.Uint() {
		return errorx.NewBizError(common.ReviewNotReview)
	}

	return svc.reviewRepo.ChangeReview(ctx, review)
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
		return nil, 0, err
	}
	return reviews, count, err
}
