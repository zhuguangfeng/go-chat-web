package model

const (
	TableNameGroup                    = "group"
	TableNameGroupUserMap             = "group_user_map"
	TableNameUserApplyJoinGroupRecord = "user_apply_join_group_record"
)

type Group struct {
	Base
	OwnerID int64 `gorm:"column:owner_id;type:bigint;not null;default:0;comment:群主ID" json:"ownerId"`
	//ActivityID   int64  `gorm:"column:activity_id;type:bigint;not null;default:0;comment:活动ID" json:"activityId"`
	GroupName    string `gorm:"column:group_name;type:varchar(255);not null;default:'';comment:活动群名" json:"name"`
	PeopleNumber int64  `gorm:"column:six;type:int;not null;default:0;comment:群人数" json:"peopleNumber"`
	Category     uint   `gorm:"column:category;type:tinyint;not null;default:0;comment:群聊类型" json:"category"`
	Status       uint   `gorm:"column:status;type:tinyint;not null;default:0;comment:群聊状态" json:"status"`
}

func (Group) TableName() string {
	return TableNameGroup
}

type GroupUserMap struct {
	Base
	GroupID int64 `gorm:"column:group_id;type:bigint;not null;default:0;comment:群聊ID" json:"groupId"`
	UserID  int64 `gorm:"column:user_id;type:bigint;not null;default:0;comment:用户ID" json:"userId"`
	Status  uint8 `gorm:"column:status;type:tinyint;not null;default:0;comment:状态" json:"status"`
}

func (GroupUserMap) TableName() string {
	return TableNameGroupUserMap
}

//type UserApplyJoinGroupRecord struct {
//	Base
//	UserID  int64 `gorm:"column:user_id;type:bigint;not null;default:0;comment:用户id" json:"userId"`
//	GroupID int64 `gorm:"column:group_id;type:bigint;not null;default:0;comment:群聊id" json:"groupId"`
//	OwnerID int64 `gorm:"column:owner_id;type:bigint;not null;default:0;comment:群主id" json:"ownerId"`
//	Status  uint  `gorm:"column:status;type:tinyint;not null;default:0;comment:装填" json:"status"`
//}
//
//func (UserApplyJoinGroupRecord) TableName() string {
//	return TableNameUserApplyJoinGroupRecord
//}
