package domain

import "time"

type Review struct {
	ID         int64     `json:"id"`
	UUID       string    ` json:"uuid"`
	Biz        string    `json:"biz"`
	BizID      int64     `json:"bizId"`
	Status     uint      `json:"status"`
	ReviewTime time.Time `json:"reviewTime"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
