package repository

import (
	"context"
	"errors"
	"github.com/zhuguangfeng/go-chat/internal/common"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/errorx"
)

type ActivitySignupRepository1 interface {
	CreateActivitySignup(ctx context.Context, su domain.ActivitySignup) error
	GetActivitySignupByID(ctx context.Context, id int64) (domain.ActivitySignup, error)
	UpdateActivitySignup(ctx context.Context, su domain.ActivitySignup) error
}

type activitySignupRepository struct {
	suDao dao.ActivitySignUpDao
}

func NewActivitySignupRepository(suDao dao.ActivitySignUpDao) ActivitySignupRepository1 {
	return &activitySignupRepository{
		suDao: suDao,
	}
}

// CreateActivitySignup 创建活动报名信息
func (repo *activitySignupRepository) CreateActivitySignup(ctx context.Context, su domain.ActivitySignup) error {
	err := repo.suDao.InsertActivitySignUp(ctx, repo.toSuEntity(su))
	if err != nil {
		if errors.Is(err, dao.ErrActivitySignUpDuplicate) {
			return errorx.NewBizError(common.ActivitySignupIsExists).WithError(err)
		}
		return errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return nil
}

// GetActivitySignupByID 获取活动报名信息详情
func (repo *activitySignupRepository) GetActivitySignupByID(ctx context.Context, id int64) (domain.ActivitySignup, error) {
	su, err := repo.suDao.FindActivitySignUpByID(ctx, id)
	if err != nil {
		if errors.Is(err, dao.ErrActivitySignUpNotFount) {
			return domain.ActivitySignup{}, errorx.NewBizError(common.ActivitySignupNotFound).WithError(err)
		}
		return domain.ActivitySignup{}, errorx.NewBizError(common.SystemInternalError).WithError(err)
	}
	return repo.toSuDomain(su), nil
}

// UpdateActivitySignup 修改活动报名信息
func (repo *activitySignupRepository) UpdateActivitySignup(ctx context.Context, su domain.ActivitySignup) error {
	return repo.suDao.UpdateActivitySignUp(ctx, repo.toSuEntity(su), model.GroupUserMap{
		UserID:   su.Applicant.ID,
		GroupID:  su.Activity.Group.ID,
		Status:   su.Status,
		Position: 1,
	})
}

func (repo *activitySignupRepository) toSuEntity(su domain.ActivitySignup) model.ActivitySignup {
	return model.ActivitySignup{
		Base: model.Base{
			ID: su.ID,
		},
		ActivityID:  su.Activity.ID,
		SponsorID:   su.Activity.Sponsor.ID,
		ApplicantID: su.Applicant.ID,
		ReviewTime:  su.ReviewTime,
		Status:      su.Status,
	}
}

func (repo *activitySignupRepository) toSuDomain(su model.ActivitySignup) domain.ActivitySignup {
	return domain.ActivitySignup{
		ID: su.ID,
		Activity: domain.Activity{
			ID: su.ActivityID,
		},
		Applicant: domain.User{
			ID: su.ApplicantID,
		},
		ReviewTime: su.ReviewTime,
		Status:     su.Status,
	}
}
