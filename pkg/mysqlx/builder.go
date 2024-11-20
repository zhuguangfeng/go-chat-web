package mysqlx

import (
	"fmt"
	"gorm.io/gorm"
)

type Condition struct {
	Key   string
	Where string
	Val   string
}

type Builder struct {
	DB *gorm.DB
}

func NewDaoBuilder(db *gorm.DB) *Builder {
	return &Builder{
		DB: db,
	}
}

func (b *Builder) WithPagination(pageNum, pageSize int) *Builder {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset, limit := (pageNum-1)*pageSize, pageSize
	b.DB = b.DB.Offset(offset).Limit(limit)
	return b
}

func (b *Builder) WithQuery(conditions []Condition) *Builder {
	for _, condition := range conditions {
		if condition.Val != "" {
			b.DB = b.DB.Where(fmt.Sprintf("%s %s %s", condition.Key, condition.Where, condition.Val))
		}
	}
	return b
}
