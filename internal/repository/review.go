package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"time"
)

type ReviewRepository interface {
	DetailReview(ctx context.Context, id int64) (domain.Review, error)
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

func (repo *reviewRepository) DetailReview(ctx context.Context, id int64) (domain.Review, error) {
	review, err := repo.reviewDao.DetailReview(ctx, id)
	if err != nil {
		return domain.Review{}, err
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
		UUID:   review.UUID,
		Biz:    review.Biz,
		BizID:  review.BizID,
		Status: review.Status,
	}
}

func (repo *reviewRepository) toDomain(review model.Review) domain.Review {
	return domain.Review{
		ID:         review.ID,
		UUID:       review.UUID,
		Biz:        review.Biz,
		BizID:      review.BizID,
		Status:     review.Status,
		ReviewTime: time.Unix(int64(review.ReviewTime), 0),
		CreateTime: time.Unix(int64(review.CreatedAt), 0),
		UpdateTime: time.Unix(int64(review.UpdatedAt), 0),
	}
}
