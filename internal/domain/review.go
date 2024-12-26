package domain

type Review struct {
	ID         int64        `json:"id"`
	UUID       string       ` json:"uuid"`
	Biz        ReviewBiz    `json:"biz"`
	BizID      int64        `json:"bizId"`
	Status     ReviewStatus `json:"status"`
	Sponsor    User         `json:"sponsor"`
	Reviewer   User         `json:"reviewer"`
	Activity   Activity     `json:"activity"`
	Opinion    string       `json:"opinion"`
	ReviewTime uint         `json:"reviewTime"`
	CreateTime uint         `json:"createTime"`
	UpdateTime uint         `json:"updateTime"`
}

type ReviewBiz string

func (r ReviewBiz) String() string {
	return string(r)
}

const (
	ReviewBizActivity       ReviewBiz = "activity"
	ReviewBizSignUpActivity ReviewBiz = "activity_sign_up"
)

type ReviewStatus uint

func (r ReviewStatus) Uint() uint {
	return uint(r)
}

const (
	ReviewStatusPendingReview ReviewStatus = iota + 1 //待审核
	ReviewStatusReviewCancel                          //审核取消
	ReviewStatusSuccess                               //审核通过
	ReviewStatusPass                                  //审核拒绝
)
