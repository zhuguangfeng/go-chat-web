package user

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"golang.org/x/crypto/bcrypt"

	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/pkg/common"
	"github.com/zhuguangfeng/zgf-base-toolbox/utils"
)

type UserService interface {
	UserLoginPwd(ctx context.Context, phone, password string) (domain.User, common.ErrorCode, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// UserLoginPwd 密码登录
func (u *userService) UserLoginPwd(ctx context.Context, phone, password string) (domain.User, common.ErrorCode, error) {
	user, err := u.userRepo.GetUserByPhone(ctx, phone)
	if err != nil {
		if utils.IsRecordNotFoundError(err) {
			return domain.User{}, common.UserNotFound, err
		}
	}

	//密码校验
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, common.UserInvalidPassword, err
	}

	return user, common.NoErr, nil
}

func (u *userService) LoginSms(ctx context.Context, phone, code string) (domain.User, common.ErrorCode, error) {

	return domain.User{}, common.NoErr, nil
}
