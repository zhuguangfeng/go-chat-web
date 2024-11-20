package dynamic

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/pkg/mysqlx"
)

type DynamicService interface {
	CreateDynamic(ctx context.Context, dynamic domain.Dynamic) error
	DeleteDynamic(ctx context.Context, id int64, uid int64) error
	ChangeDynamic(ctx context.Context, dynamic domain.Dynamic) error
	DetailDynamic(ctx context.Context, id int64) (domain.Dynamic, error)
	ListDynamic(ctx context.Context, pageSize, pageNum int, conditions []mysqlx.Condition) ([]domain.Dynamic, int64, error)
}

type dynamicService struct {
	dynamicRepo repository.DynamicRepository
}

func NewDynamicService(dynamicRepo repository.DynamicRepository) DynamicService {
	return &dynamicService{
		dynamicRepo: dynamicRepo,
	}
}

func (svc *dynamicService) CreateDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	return svc.dynamicRepo.CreateDynamic(ctx, dynamic)
}

func (svc *dynamicService) DeleteDynamic(ctx context.Context, id int64, uid int64) error {
	return svc.dynamicRepo.DeleteDynamic(ctx, id, uid)
}

func (svc *dynamicService) ChangeDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	return svc.dynamicRepo.ChangeDynamic(ctx, dynamic)
}

func (svc *dynamicService) DetailDynamic(ctx context.Context, id int64) (domain.Dynamic, error) {
	return svc.dynamicRepo.DetailDynamic(ctx, id)
}

func (svc *dynamicService) ListDynamic(ctx context.Context, pageSize, pageNum int, conditions []mysqlx.Condition) ([]domain.Dynamic, int64, error) {
	return svc.dynamicRepo.ListDynamic(ctx, pageSize, pageNum, conditions)
}
