package dao

import (
	"fmt"
	"gorm.io/gorm"
)

type DaoBuilder struct {
	db *gorm.DB
}

type query struct {
	Key   string
	where string
	Val   string
}

func NewDaoBuilder(db *gorm.DB) *DaoBuilder {
	return &DaoBuilder{
		db: db,
	}
}

func (d *DaoBuilder) WithPagination(pageNum, pageSize int) *DaoBuilder {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset, limit := (pageNum-1)*pageSize, pageSize
	d.db = d.db.Offset(offset).Limit(limit)
	return d
}

func (d *DaoBuilder) WithQuery(querys []query) *DaoBuilder {
	for _, query := range querys {
		if query.Val != "" {
			d.db = d.db.Where(fmt.Sprintf("%s %s %s", query.Key, query.where, query.Val))
		}
	}
	return d
}
