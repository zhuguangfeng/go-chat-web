package dao

import (
	"context"
	"errors"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
	"gorm.io/gorm"
)

var (
	ErrActivitySignUpDuplicate = errors.New("报名记录重复")
	ErrActivitySignUpNotFount  = errors.New("报名记录不存在")
)

type ActivitySignUpDao interface {
	InsertActivitySignUp(ctx context.Context, signUp model.ActivitySignup) error
	FindActivitySignUpByID(ctx context.Context, id int64) (model.ActivitySignup, error)
	UpdateActivitySignUp(ctx context.Context, signUp model.ActivitySignup) error
	UpdateActivitySignupStatusSuccess(ctx context.Context, signup model.ActivitySignup, gu model.GroupUserMap) error
	CancelActivitySignUpWithQuitGroup(ctx context.Context, signUpID, groupID int64) error
	ActivitySignUpList(ctx context.Context, req dtoV1.SignUpListReq) ([]model.ActivitySignup, int64, error)
}

type GormActivitySignUpDao struct {
	ActivityDao
	db *gorm.DB
}

func NewActivitySignUp(db *gorm.DB) ActivitySignUpDao {
	return &GormActivitySignUpDao{
		db: db,
	}
}

func (dao *GormActivitySignUpDao) FindActivitySignUpByID(ctx context.Context, id int64) (model.ActivitySignup, error) {
	var res model.ActivitySignup
	err := dao.db.WithContext(ctx).Where("id=?", id).Find(&res).Error
	if utils.IsRecordNotFoundError(err) {
		return model.ActivitySignup{}, ErrActivitySignUpNotFount
	}
	return res, nil
}

// InsertActivitySignUp 插入报名记录
func (dao *GormActivitySignUpDao) InsertActivitySignUp(ctx context.Context, signUp model.ActivitySignup) error {
	err := dao.db.WithContext(ctx).Create(&signUp).Error
	if utils.IsDuplicateKeyError(err) {
		return ErrActivitySignUpDuplicate
	}
	return err
}

// CancelActivitySignUpWithQuitGroup 取消报名推出群聊
func (dao *GormActivitySignUpDao) CancelActivitySignUpWithQuitGroup(ctx context.Context, signUpID, groupID int64) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := dao.updateActivitySignUp(ctx, tx, model.ActivitySignup{
			Base: model.Base{
				ID: signUpID,
			},
			Status: domain.ActivitySignupStatusCancelReview.Uint(),
		})
		if err != nil {
			return err
		}
		tx.Model(model.Group{}).Where("id = ?", groupID).Updates(model.Group{})
		return tx.Model(model.GroupUserMap{}).Where("group_id = ? and user_id = ?", groupID).Update("status = ?", domain.GroupUserMapStatusQuit).Error
	})
}

func (dao *GormActivitySignUpDao) UpdateActivitySignUp(ctx context.Context, signUp model.ActivitySignup) error {
	return dao.updateActivitySignUp(ctx, nil, signUp)
}

func (dao *GormActivitySignUpDao) UpdateActivitySignupStatusSuccess(ctx context.Context, signup model.ActivitySignup, gu model.GroupUserMap) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := dao.updateActivitySignUp(ctx, tx, model.ActivitySignup{
			ActivityID: signup.ActivityID,
			Status:     domain.ActivitySignupStatusReviewSuccess.Uint(),
		})
		if err != nil {
			return err
		}

		return tx.Create(&gu).Error

	})
}

func (dao *GormActivitySignUpDao) updateActivitySignUp(ctx context.Context, tx *gorm.DB, signUp model.ActivitySignup) error {
	if tx != nil {
		dao.db = tx
	}
	return dao.db.WithContext(ctx).Where("id = ?", signUp.ID).Updates(signUp).Error
}

// ListActivitySignUp 报名列表
func (dao *GormActivitySignUpDao) ActivitySignUpList(ctx context.Context, req dtoV1.SignUpListReq) ([]model.ActivitySignup, int64, error) {
	var (
		res   = make([]model.ActivitySignup, 0)
		count int64
	)
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx).Where("activity_id = ? and sponsor_id = ?", req.ActivityID, req.UID)).
		WithPagination(req.PageNum, req.PageSize).
		DB.Find(&res).Limit(-1).Offset(-1).Count(&count).Error

	return res, count, err
}
