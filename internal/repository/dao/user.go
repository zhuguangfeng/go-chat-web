package dao

import (
	"context"
	"errors"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicate = errors.New("用户已存在")
	ErrUserNotFound  = errors.New("用户不存在")
)

type UserDao interface {
	InsertUser(ctx context.Context, user model.User) error
	FindUserByPhone(ctx context.Context, phone string) (model.User, error)
	FindUserByID(ctx context.Context, id int64) (model.User, error)
	FindUserByIDs(ctx context.Context, ids []int64) ([]model.User, error)
}

type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GormUserDao{
		db: db,
	}
}

// InsertUser 创建用户
func (dao *GormUserDao) InsertUser(ctx context.Context, user model.User) error {
	err := dao.db.WithContext(ctx).Create(&user).Error
	if utils.IsDuplicateKeyError(err) {
		return ErrUserDuplicate
	}
	return err
}

// FindUserByPhone 根据手机号码查找用户
func (dao *GormUserDao) FindUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Model(model.User{}).Where("phone = ?", phone).First(&u).Error
	if utils.IsRecordNotFoundError(err) {
		return model.User{}, ErrUserNotFound
	}
	return u, err
}

// FindUserByID 根据用户id查找用户
func (dao *GormUserDao) FindUserByID(ctx context.Context, id int64) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Model(model.User{}).Where("id = ?", id).First(&u).Error
	if utils.IsRecordNotFoundError(err) {
		return model.User{}, ErrUserNotFound
	}
	return u, err
}

func (dao *GormUserDao) FindUserByIDs(ctx context.Context, ids []int64) ([]model.User, error) {
	var res = make([]model.User, 0)
	err := dao.db.WithContext(ctx).Where("id in (?)", ids).Find(&res).Error
	return res, err
}
