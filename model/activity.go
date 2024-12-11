package model

const (
	TableNameActivity       = "activity"
	TableNameActivitySignup = "activity_signup"
)

type Activity struct {
	Base
	SponsorID           int64   `gorm:"column:sponsor_id;type:bigint;index:idx_sponsor_id;not null;comment:发起人ID" json:"sponsorId"`
	GroupID             int64   `gorm:"column:group_id;type:bigint;not null;comment:群聊id" json:"groupId"`
	Title               string  `gorm:"column:title;type:varchar(255);index:idx_title;not null;default:'';comment:活动标题" json:"title"`
	Desc                string  `gorm:"column:desc;type:text;not null;comment:活动描述" json:"desc"`
	Media               Strings `gorm:"column:media;type:text;not null;comment:活动图片或视频" json:"media"`
	AgeRestrict         uint    `gorm:"column:age_restrict;type:tinyint;not null;default:0;comment:最大年龄限制" json:"ageRestrict"`
	GenderRestrict      uint    `gorm:"column:gender_restrict;type:tinyint;not null;default:0;comment:性别限制 男|女|不限" json:"genderRestrict"`
	CostRestrict        uint    `gorm:"column:cost_restrict;type:tinyint;not null;default:0;comment:费用支付方式" json:"CostRestrict"`
	Visibility          uint    `gorm:"column:visibility;type:tinyint;not null;default:0;comment:报名可见度" json:"visibility"`
	MaxPeopleNumber     int64   `gorm:"column:max_people_number;type:tinyint;not null;default:0;comment:最大报名人数" json:"maxPeopleNumber"`
	CurrentPeopleNumber int64   `gorm:"column:current_people_number;type:tinyint;not null;default:0;comment:当前报名人数" json:"CurrentPeopleNumber"`
	Address             string  `gorm:"column:address;type:varchar(255);not null;comment:获取地点" json:"address"`
	Category            uint    `gorm:"column:category;type:int;not null;default:0;comment:活动类型" json:"category"`
	StartTime           uint    `gorm:"column:start_time;type:int;not null;default:0;comment:活动开始时间" json:"startTime"`
	DeadlineTime        uint    `gorm:"column:deadline_time;type:int;not null;default:0;comment:活动报名截止时间" json:"deadlineTime"`
	Status              uint    `gorm:"column:status;type:tinyint;not null;default:0;comment:活动状态" json:"status"`
}

func (Activity) TableName() string {
	return TableNameActivity
}

type ActivitySignup struct {
	Base
	ActivityID  int64 `gorm:"column:activity_id;type:bigint;uniqueIndex:udx_activity_sponsor_applicant;not null;comment:活动id" json:"activityID"`
	SponsorID   int64 `gorm:"column:sponsor_id;type:bigint;uniqueIndex:udx_activity_sponsor_applicant;not null;comment:活动发起人id" json:"sponsorID"`
	ApplicantID int64 `gorm:"column:applicant_id;type:bigint;uniqueIndex:udx_activity_sponsor_applicant;not null;comment:申请人id" json:"applicantID"`
	ReviewTime  uint  `gorm:"column:review_time;type:bigint;not null;comment:审核时间" json:"reviewTime"`
	Status      uint  `gorm:"column:status;type:tinyint;not null;comment:审核状态" json:"status"`
}

func (ActivitySignup) TableName() string {
	return TableNameActivitySignup
}
