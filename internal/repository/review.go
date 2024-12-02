package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
)

type ReviewRepository interface {
	ChangeReview(ctx context.Context, review domain.Review) error
	DetailReview(ctx context.Context, uuid string) (domain.Review, error)
	ListReview(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]domain.Review, int64, error)
}

type reviewRepository struct {
	reviewDao dao.ReviewDao
}

func NewReviewRepository(reviewDao dao.ReviewDao) ReviewRepository {
	return &reviewRepository{
		reviewDao: reviewDao,
	}
}

func (repo *reviewRepository) ChangeReview(ctx context.Context, review domain.Review) error {
	err := repo.reviewDao.UpdateReview(ctx, repo.toEntity(review))
	if err != nil {
		return errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return nil
}

func (repo *reviewRepository) DetailReview(ctx context.Context, uuid string) (domain.Review, error) {
	review, err := repo.reviewDao.DetailReview(ctx, uuid)
	if err != nil {
		return domain.Review{}, errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return repo.toDomain(review), nil
}

func (repo *reviewRepository) ListReview(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]domain.Review, int64, error) {
	reviews, err := repo.reviewDao.ListReview(ctx, pageNum, pageSize, biz, status)
	if err == nil {

		count, err := repo.reviewDao.FindReviewCount(ctx, biz, status)
		if err != nil {
			//TODO
		}

		return slice.Map(reviews, func(idx int, review model.Review) domain.Review {
			return repo.toDomain(review)
		}), count, nil
	}

	return nil, 0, err
}

func (repo *reviewRepository) toEntity(review domain.Review) model.Review {
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

func (repo *reviewRepository) toDomain(review model.Review) domain.Review {
	return domain.Review{
		ID:     review.ID,
		UUID:   review.UUID,
		Biz:    review.Biz,
		BizID:  review.BizID,
		Status: review.Status,
		Reviewer: domain.User{
			ID: review.ReviewerID,
		},
		ReviewTime: review.ReviewTime,
		CreateTime: review.CreatedAt,
		UpdateTime: review.UpdatedAt,
	}
}
