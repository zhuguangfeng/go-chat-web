package domain

type Review struct {
	ID         int64    `json:"id"`
	UUID       string   ` json:"uuid"`
	Biz        string   `json:"biz"`
	BizID      int64    `json:"bizId"`
	Status     uint     `json:"status"`
	Reviewer   User     `json:"reviewer"`
	Opinion    string   `json:"opinion"`
	Activity   Activity `json:"activity"`
	ReviewTime uint     `json:"reviewTime"`
	CreateTime uint     `json:"createTime"`
	UpdateTime uint     `json:"updateTime"`
}
