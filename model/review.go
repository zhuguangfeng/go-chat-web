package model

const (
	TableNameReview = "review"
)

// Review 审核表
type Review struct {
	Base
	UUID       string `gorm:"column:uuid;type:char(64);not null;default:0;comment:活动ID" json:"uuid"`
	Biz        string `gorm:"column:biz;type:varchar(32);not null;default:'';comment:业务" json:"biz"`
	BizID      int64  `gorm:"column:biz_id;type:bigint;not null;default:0;comment:业务ID" json:"bizId"`
	SponsorID  int64  `gorm:"column:sponsor_id;type:bigint;not null;default:0;comment:发起人ID" json:"sponsorId"`
	ReviewerID int64  `gorm:"column:user_id;type:bigint;not null;default:0;comment:审核人ID" json:"reviewerId"`
	Status     uint   `gorm:"column:status;type:tinyint;not null;default:0;comment:审核状态" json:"status"`
	ReviewTime uint   `gorm:"column:review_time;type:uint;not null;default:0;comment:审核时间"  json:"reviewTime"`
}

func (Review) TableName() string {
	return TableNameReview
}
