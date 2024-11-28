package repository

import (
	"context"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/ekit/slice"
)

type DynamicRepository interface {
	CreateDynamic(ctx context.Context, dynamic domain.Dynamic) error
	DeleteDynamic(ctx context.Context, id int64, uid int64) error
	ChangeDynamic(ctx context.Context, dynamic domain.Dynamic) error
	DetailDynamic(ctx context.Context, id int64) (domain.Dynamic, error)
	ListDynamic(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Dynamic, int64, error)
}

type dynamicRepository struct {
	dynamicDao dao.DynamicDao
}

func NewDynamicRepository(dynamicDao dao.DynamicDao) DynamicRepository {
	return &dynamicRepository{
		dynamicDao: dynamicDao,
	}
}

// CreateDynamic 创建动态
func (d *dynamicRepository) CreateDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	return d.dynamicDao.InsertDynamic(ctx, d.toEntity(dynamic))
}

// DeleteDynamic 删除动态
func (d *dynamicRepository) DeleteDynamic(ctx context.Context, id int64, uid int64) error {
	return d.dynamicDao.DeleteDynamic(ctx, id, uid)
}

// ChangeDynamic 修改动态
func (d *dynamicRepository) ChangeDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	return d.dynamicDao.UpdateDynamic(ctx, d.toEntity(dynamic))
}

// DetailDynamic 动态详情
func (d *dynamicRepository) DetailDynamic(ctx context.Context, id int64) (domain.Dynamic, error) {
	dynamic, err := d.dynamicDao.DetailDynamic(ctx, id)
	if err != nil {
		return domain.Dynamic{}, err
	}
	return d.toDomain(dynamic), nil
}

// ListDynamic 动态列表
func (d *dynamicRepository) ListDynamic(ctx context.Context, pageNum, pageSize int, searchKey string) ([]domain.Dynamic, int64, error) {
	dynamics, err := d.dynamicDao.ListDynamic(ctx, pageNum, pageSize, searchKey)
	if err == nil {
		count, err := d.dynamicDao.FindDynamicCount(ctx, searchKey)
		if err != nil {
			//可以返回错误 也可以降级 返回dynamic数据 不返回总条数
			// TODO 日志
		}
		return slice.Map(dynamics, func(idx int, src model.Dynamic) domain.Dynamic {
			return d.toDomain(src)
		}), count, nil
	}
	return nil, 0, err
}

// toEntity dynamic转换为model实体
func (d *dynamicRepository) toEntity(dynamic domain.Dynamic) model.Dynamic {
	return model.Dynamic{
		Base: model.Base{
			ID: dynamic.ID,
		},
		UserID:      dynamic.User.ID,
		Title:       dynamic.Title,
		Media:       dynamic.Media,
		Tags:        dynamic.Tags,
		Visibility:  dynamic.Visibility,
		DynamicType: dynamic.DynamicType,
		Status:      dynamic.Status,
	}
}

// toDomain dynamic实体转换为domain
func (d *dynamicRepository) toDomain(dynamic model.Dynamic) domain.Dynamic {
	return domain.Dynamic{
		ID: dynamic.ID,
		User: domain.User{
			ID: dynamic.UserID,
		},
		Title:       dynamic.Title,
		Media:       dynamic.Media,
		Tags:        dynamic.Tags,
		Visibility:  dynamic.Visibility,
		DynamicType: dynamic.DynamicType,
		Status:      dynamic.Status,
	}
}
