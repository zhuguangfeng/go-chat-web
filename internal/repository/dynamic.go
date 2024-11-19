package repository

import (
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
)

type DynamicRepository interface {
}

type dynamicRepository struct {
	dynamicDao dao.DynamicDao
}

func NewDynamicRepository(dynamicDao dao.DynamicDao) DynamicRepository {
	return &dynamicRepository{
		dynamicDao: dynamicDao,
	}
}

func (d *dynamicRepository) CreateDynamic() error {

}
