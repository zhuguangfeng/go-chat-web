package domain

type Activity struct {
	ID                  int64            `json:"id" dc:"主键ID"`
	Sponsor             User             `json:"user" dc:"活动发起人"`
	Group               Group            `json:"group" dc:"群聊"`
	Title               string           `json:"title" dc:"活动标题"`
	Desc                string           `json:"desc" dc:"活动描述"`
	Media               []string         `json:"media" dc:"资源 视频或图片"`
	AgeRestrict         uint             `json:"ageRestrict" dc:"最大年龄"`
	GenderRestrict      uint             `json:"genderRestrict" dc:"性别限制"`
	CostRestrict        uint             `json:"costRestrict" dc:"费用限制"`
	Visibility          uint             `json:"visibility" dc:"可见度"`
	MaxPeopleNumber     int64            `json:"maxPeopleNumber" dc:"最大报名人数"`
	CurrentPeopleNumber int64            `json:"currentPeopleNumber" dc:"当前报名人数"`
	Address             string           `json:"address" dc:"活动地址"`
	Category            ActivityCategory `json:"category" dc:"活动分类"`
	StartTime           uint             `json:"startTime" dc:"活动开始时间"`
	DeadlineTime        uint             `json:"deadlineTime" dc:"活动截止时间"`
	Status              ActivityStatus   `json:"status" dc:"活动状态"`
	CreatedTime         uint             `json:"createdTime" dc:"创建时间"`
	UpdatedTime         uint             `json:"updatedTime" dc:"修改时间"`
}

// ActivityCategory 活动类型
type ActivityCategory uint

const (
	ActivityCategoryUnknown ActivityCategory = iota
	ActivityCategoryStudy
)

func (a ActivityCategory) Uint() uint {
	return uint(a)
}

// ActivityStatus 活动状态
type ActivityStatus uint

func (a ActivityStatus) Uint() uint {
	return uint(a)
}

const (
	ActivityStatusPendingReview ActivityStatus = iota + 1 //待审核
	ActivityStatusReviewPass                              //审核失败
	ActivityStatusSignUp                                  //报名中 == 审核通过
	ActivityStatusCancel                                  //已取消
	ActivityStatusStart                                   //已开始
	ActivityStatusEnd                                     //已结束
)

type ActivitySignup struct {
	ID         int64                `json:"id"`
	Activity   Activity             `json:"activity"`
	Applicant  User                 `json:"applicant"`
	ReviewTime uint                 `json:"reviewTime"`
	Status     ActivitySignupStatus `json:"status"`
}

// ActivitySignupStatus 活动报名状态
type ActivitySignupStatus uint

func (a ActivitySignupStatus) Uint() uint {
	return uint(a)
}

const (
	ActivitySignupStatusUnknown       ActivitySignupStatus = iota
	ActivitySignupStatusPendingReview                      //待审核
	ActivitySignupStatusCancelReview                       //取消审核
	ActivitySignupStatusReviewPass                         //审核拒绝
	ActivitySignupStatusReviewSuccess                      //审核通过
)
