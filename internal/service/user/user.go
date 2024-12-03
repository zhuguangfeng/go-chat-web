package user

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository"
)

type UserService interface {
	UserLoginPwd(ctx context.Context, phone, password string) (domain.User, error)
	GetUserInfo(ctx context.Context, userID int64) (domain.User, error)
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
func (svc *userService) UserLoginPwd(ctx context.Context, phone, password string) (domain.User, error) {
	user, err := svc.userRepo.GetUserByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}

	//密码校验
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//if err != nil {
	//	return domain.User{}, errorx.NewBizError(common.UserInvalidPassword).WithError(err)
	//}

	return user, nil
}

func (svc *userService) LoginSms(ctx context.Context, phone, code string) (domain.User, error) {

	return domain.User{}, nil
}

func (svc *userService) GetUserInfo(ctx context.Context, userID int64) (domain.User, error) {
	return svc.userRepo.GetUserByID(ctx, userID)
}
