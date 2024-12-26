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
	ReviewActivity(ctx context.Context, review domain.Review, group domain.Group) error
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

// ReviewActivity 审核活动
func (repo *reviewRepository) ReviewActivity(ctx context.Context, review domain.Review, group domain.Group) error {
	err := repo.reviewDao.ReviewActivity(ctx, repo.toReviewEntity(review), repo.toGroupEntity(group))
	if err != nil {
		return errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return nil
}

// DetailReview 审核详情
func (repo *reviewRepository) DetailReview(ctx context.Context, uuid string) (domain.Review, error) {
	review, err := repo.reviewDao.DetailReview(ctx, uuid)
	if err != nil {
		return domain.Review{}, errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return repo.toReviewDomain(review), nil
}

// ListReview 审核列表
func (repo *reviewRepository) ListReview(ctx context.Context, pageNum, pageSize int, biz string, status uint) ([]domain.Review, int64, error) {
	reviews, err := repo.reviewDao.ListReview(ctx, pageNum, pageSize, biz, status)
	if err == nil {
		count, err := repo.reviewDao.FindReviewCount(ctx, biz, status)
		if err != nil {
			repo.logger.Error("[ReviewRepository.ListReview] 获取review总条数失败", logger.Error(err))
		}

		return slice.Map(reviews, func(idx int, review model.Review) domain.Review {
			return repo.toReviewDomain(review)
		}), count, nil
	}

	return nil, 0, errorx.NewBizError(common.SystemInternalError).WithError(err)
}

func (repo *reviewRepository) toReviewEntity(review domain.Review) model.Review {
	return model.Review{
		Base: model.Base{
			ID: review.ID,
		},
		SponsorID:  review.Sponsor.ID,
		UUID:       review.UUID,
		Biz:        review.Biz.String(),
		BizID:      review.BizID,
		ReviewerID: review.Reviewer.ID,
		Status:     review.Status.Uint(),
		ReviewTime: review.ReviewTime,
	}
}

func (repo *reviewRepository) toReviewDomain(review model.Review) domain.Review {
	return domain.Review{
		ID:     review.ID,
		UUID:   review.UUID,
		Biz:    domain.ReviewBiz(review.Biz),
		BizID:  review.BizID,
		Status: domain.ReviewStatus(review.Status),
		Sponsor: domain.User{
			ID: review.SponsorID,
		},
		Reviewer: domain.User{
			ID: review.ReviewerID,
		},
		ReviewTime: review.ReviewTime,
		CreateTime: review.CreatedAt,
		UpdateTime: review.UpdatedAt,
	}
}

func (repo *reviewRepository) toGroupEntity(group domain.Group) model.Group {
	return model.Group{
		Base: model.Base{
			ID: group.ID,
		},
		GroupName:       group.GroupName,
		OwnerID:         group.Owner.ID,
		PeopleNumber:    group.PeopleNumber,
		MaxPeopleNumber: group.MaxPeopleNumber,
		Category:        group.Category.Uint(),
		Status:          group.Status.Uint(),
	}
}

func (repo *reviewRepository) toGroupDomain(group model.Group) domain.Group {
	return domain.Group{
		ID:        group.ID,
		GroupName: group.GroupName,
		Owner: domain.User{
			ID: group.OwnerID,
		},
		PeopleNumber:    group.PeopleNumber,
		MaxPeopleNumber: group.MaxPeopleNumber,
		Category:        domain.GroupCategory(group.Category),
		Status:          domain.GroupStatus(group.Status),
	}
}
