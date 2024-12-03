package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

type ReviewRepository interface {
	ChangeReview(ctx context.Context, review domain.Review) error
	DetailReview(ctx context.Context, uuid string) (domain.Review, error)
	ListReview(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]domain.Review, int64, error)
}

type reviewRepository struct {
	logger    logger.Logger
	reviewDao dao.ReviewDao
}

func NewReviewRepository(logger logger.Logger, reviewDao dao.ReviewDao) ReviewRepository {
	return &reviewRepository{
		logger:    logger,
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
			repo.logger.Error("[ReviewRepository.ListReview] 获取review总条数失败", logger.Error(err))
		}

		return slice.Map(reviews, func(idx int, review model.Review) domain.Review {
			return repo.toDomain(review)
		}), count, nil
	}

	return nil, 0, errorx.NewBizError(common.SystemInternalError).WithError(err)
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
