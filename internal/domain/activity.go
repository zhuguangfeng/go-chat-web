package domain

type Activity struct {
	ID                  int64    `json:"id" dc:"主键ID"`
	Sponsor             User     `json:"user" dc:"活动发起人"`
	Title               string   `json:"title" dc:"活动标题"`
	Desc                string   `json:"desc" dc:"活动描述"`
	Media               []string `json:"media" dc:"资源 视频或图片"`
	AgeRestrict         uint     `json:"ageRestrict" dc:"最大年龄"`
	GenderRestrict      uint     `json:"genderRestrict" dc:"性别限制"`
	CostRestrict        uint     `json:"costRestrict" dc:"费用限制"`
	Visibility          uint     `json:"visibility" dc:"可见度"`
	MaxPeopleNumber     int64    `json:"maxPeopleNumber" dc:"最大报名人数"`
	CurrentPeopleNumber int64    `json:"currentPeopleNumber" dc:"当前报名人数"`
	Address             string   `json:"address" dc:"活动地址"`
	Category            uint     `json:"category" dc:"活动分类"`
	StartTime           uint     `json:"startTime" dc:"活动开始时间"`
	DeadlineTime        uint     `json:"deadlineTime" dc:"活动截止时间"`
	Status              uint     `json:"status" dc:"活动状态"`
}
