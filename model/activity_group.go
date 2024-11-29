package model

const (
	TableNameGroup        = "group"
	TableNameGroupUserMap = "group_user_map"
)

type Group struct {
	Base
	ActivityID   int64  `gorm:"column:activity_id;type:bigint;not null;default:0;comment:活动ID" json:"activityId"`
	GroupName    string `gorm:"column:group_name;type:varchar(255);not null;default:'';comment:活动群名" json:"name"`
	PeopleNumber int64  `gorm:"column:six;type:int;not null;default:0;comment:群人数" json:"peopleNumber"`
	Status       uint8  `gorm:"column:six;type:tinyint;not null;default:0;comment:群聊状态" json:"status"`
}

func (Group) TableName() string {
	return TableNameGroup
}

type ActivityGroupUserMap struct {
	Base
	GroupID int64 `gorm:"column:group_id;type:bigint;not null;default:0;comment:群聊ID" json:"groupId"`
	UserID  int64 `gorm:"column:user_id;type:bigint;not null;default:0;comment:用户ID" json:"userId"`
	Status  uint8 `gorm:"column:status;type:tinyint;not null;default:0;comment:状态" json:"status"`
}

func (ActivityGroupUserMap) TableName() string {
	return TableNameGroupUserMap
}
