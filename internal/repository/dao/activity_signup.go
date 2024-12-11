package dao

import (
	"context"
	"errors"
	"github.com/zhuguangfeng/go-chat/internal/common"
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
	//UpdateActivitySignUp(ctx context.Context, signUp model.ActivitySignup) error
	FindActivitySignUpByID(ctx context.Context, id int64) (model.ActivitySignup, error)
	UpdateActivitySignUp(ctx context.Context, signUp model.ActivitySignup, ugMap model.GroupUserMap) error
	ListActivitySignUp(ctx context.Context, pageNum, pageSize int, activityID, userID int64, status uint) ([]model.ActivitySignup, int64, error)
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

func (dao *GormActivitySignUpDao) UpdateActivitySignUp(ctx context.Context, signUp model.ActivitySignup, ugMap model.GroupUserMap) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := dao.updateActivitySignUp(tx, signUp)
		if err == nil {
			if signUp.Status == common.ReviewStatusSuccess.Uint() {
				return dao.db.Create(&ugMap).Error
			}
		}
		return err
	})
}

func (dao *GormActivitySignUpDao) updateActivitySignUp(tx *gorm.DB, signUp model.ActivitySignup) error {
	return tx.Updates(&signUp).Error
}

// ListActivitySignUp 报名列表
func (dao *GormActivitySignUpDao) ListActivitySignUp(ctx context.Context, pageNum, pageSize int, activityID, userID int64, status uint) ([]model.ActivitySignup, int64, error) {
	var (
		res   = make([]model.ActivitySignup, 0)
		count int64
	)
	err := mysqlx.NewDaoBuilder(dao.db.WithContext(ctx).Where("activity_id = ? and sponsor_id = ?", activityID, userID)).
		WithEqual("status", status).
		WithPagination(pageNum, pageSize).
		DB.Find(&res).Limit(-1).Offset(-1).Count(&count).Error

	return res, count, err
}
