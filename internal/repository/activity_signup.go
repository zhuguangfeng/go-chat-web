package repository

import (
	"context"
	"errors"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
)

var (
	ErrActivitySignupExists   = errors.New("活动报名记录已存在")
	ErrActivitySignupNotFound = errors.New("活动报名记录不存在")
)

type ActivitySignupRepository1 interface {
	CreateActivitySignup(ctx context.Context, su domain.ActivitySignup) error
	GetActivitySignup(ctx context.Context, id int64) (domain.ActivitySignup, error)
	UpdateActivitySignup(ctx context.Context, su domain.ActivitySignup) error
	ReviewSignupSuccess(ctx context.Context, su domain.ActivitySignup) error
	CancelReviewSuccessSignup(ctx context.Context, signupID, groupID int64) error
	GetSignupList(ctx context.Context, req dtoV1.SignUpListReq) ([]domain.ActivitySignup, int64, error)
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
	if err != nil && errors.Is(err, dao.ErrActivitySignUpDuplicate) {
		return ErrActivitySignupExists
	}
	return err
}

func (repo *activitySignupRepository) CancelReviewSuccessSignup(ctx context.Context, signupID, groupID int64) error {
	return repo.suDao.CancelActivitySignUpWithQuitGroup(ctx, signupID, groupID)
}

// GetActivitySignup 获取活动报名信息详情
func (repo *activitySignupRepository) GetActivitySignup(ctx context.Context, id int64) (domain.ActivitySignup, error) {
	su, err := repo.suDao.FindActivitySignUpByID(ctx, id)
	if err != nil {
		if errors.Is(err, dao.ErrActivitySignUpNotFount) {
			return domain.ActivitySignup{}, ErrActivitySignupNotFound
		}
		return domain.ActivitySignup{}, err
	}
	return repo.toSuDomain(su), nil
}

// ReviewSignupSuccess 审核报名通过
func (repo *activitySignupRepository) ReviewSignupSuccess(ctx context.Context, su domain.ActivitySignup) error {
	return repo.suDao.UpdateActivitySignupStatusSuccess(ctx, repo.toSuEntity(su), model.GroupUserMap{
		GroupID: su.Activity.Group.ID,
		UserID:  su.Applicant.ID,
		Status:  domain.GroupUserMapStatusJoin.Uint(),
	})
}

// UpdateActivitySignup 审核报名
func (repo *activitySignupRepository) UpdateActivitySignup(ctx context.Context, su domain.ActivitySignup) error {
	return repo.suDao.UpdateActivitySignUp(ctx, repo.toSuEntity(su))
}

// GetSignupList 获取报名列表
func (repo *activitySignupRepository) GetSignupList(ctx context.Context, req dtoV1.SignUpListReq) ([]domain.ActivitySignup, int64, error) {
	signup, count, err := repo.suDao.ActivitySignUpList(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return slice.Map(signup, func(idx int, src model.ActivitySignup) domain.ActivitySignup {
		return repo.toSuDomain(src)
	}), count, nil
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
		Status:      su.Status.Uint(),
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
		Status:     domain.ActivitySignupStatus(su.Status),
	}
}
