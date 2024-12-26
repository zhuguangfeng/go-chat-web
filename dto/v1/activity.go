package v1

type CreateActivityReq struct {
	UserID          int64    `json:"userId" dc:"活动发起人"`
	Title           string   `json:"title" dc:"活动标题"`
	Desc            string   `json:"desc" dc:"活动描述"`
	Media           []string `json:"media" dc:"资源 视频或图片"`
	AgeRestrict     uint     `json:"ageRestrict" dc:"最大年龄"`
	GenderRestrict  uint     `json:"genderRestrict" dc:"性别限制"`
	CostRestrict    uint     `json:"costRestrict" dc:"费用限制"`
	Visibility      uint     `json:"visibility" dc:"可见度"`
	MaxPeopleNumber int64    `json:"maxPeopleNumber" dc:"最大报名人数"`
	Address         string   `json:"address" dc:"活动地址"`
	Category        uint     `json:"category" dc:"活动分类"`
	StartTime       uint     `json:"startTime" dc:"活动开始时间"`
	DeadlineTime    uint     `json:"deadlineTime" dc:"活动截止时间"`
}

type ChangeActivityReq struct {
	ID int64 `json:"id" dc:"活动ID"`
	CreateActivityReq
}

type Activity struct {
	UserID          int64    `json:"userId,omitempty" dc:"用户id"`
	Username        string   `json:"username,omitempty" dc:""`
	Avatar          string   `json:"avatar,omitempty"`
	Title           string   `json:"title,omitempty" dc:"活动标题"`
	Desc            string   `json:"desc,omitempty" dc:"活动描述"`
	Media           []string `json:"media,omitempty" dc:"资源 视频或图片"`
	AgeRestrict     uint     `json:"ageRestrict,omitempty" dc:"最大年龄"`
	GenderRestrict  uint     `json:"genderRestrict,omitempty" dc:"性别限制"`
	CostRestrict    uint     `json:"costRestrict,omitempty" dc:"费用限制"`
	Visibility      uint     `json:"visibility,omitempty" dc:"可见度"`
	MaxPeopleNumber int64    `json:"maxPeopleNumber,omitempty" dc:"最大报名人数"`
	Address         string   `json:"address,omitempty" dc:"活动地址"`
	Category        uint     `json:"category,omitempty" dc:"活动分类"`
	StartTime       uint     `json:"startTime,omitempty" dc:"活动开始时间"`
	DeadlineTime    uint     `json:"deadlineTime,omitempty" dc:"活动截止时间"`
	Status          uint     `json:"status,omitempty" dc:"活动截止时间"`
}

type ActivityListReq struct {
	BaseListReq
	AgeRestrict    uint   `json:"ageRestrict"`
	GenderRestrict uint   `json:"genderRestrict"`
	CostRestrict   uint   `json:"CostRestrict"`
	Visibility     uint   `json:"visibility"`
	Address        string `json:"address"`
	Category       uint   `json:"category"`
	StartTime      uint   `json:"startTime"`
	EndTime        uint   `json:"EndTime"`
	Status         uint   `json:"status"`
}

type SignUpActivityReq struct {
	ActivityID int64 `json:"activityId" dc:"活动id"`
}

type CancelSignUpActivityReq struct {
	ActivityID int64 `json:"activityId" dc:"活动id"`
}

type ReviewSignupReq struct {
	SignupID   int64 `json:"signupId" dc:"报名id"`
	ActivityID int64 `json:"activityId" dc:"活动id"`
	Status     uint  `json:"status" dc:"状态"`
}

type SignUpListReq struct {
	BaseListReq
	UID        int64
	ActivityID int64 `json:"activityId"`
}
