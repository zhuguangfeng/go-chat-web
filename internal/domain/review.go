package domain

type Review struct {
	ID         int64    `json:"id"`
	UUID       string   ` json:"uuid"`
	Biz        string   `json:"biz"`
	BizID      int64    `json:"bizId"`
	Status     uint     `json:"status"`
	Sponsor    User     `json:"sponsor"`
	Reviewer   User     `json:"reviewer"`
	Activity   Activity `json:"activity"`
	Opinion    string   `json:"opinion"`
	ReviewTime uint     `json:"reviewTime"`
	CreateTime uint     `json:"createTime"`
	UpdateTime uint     `json:"updateTime"`
}
