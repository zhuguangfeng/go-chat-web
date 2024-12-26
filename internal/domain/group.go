package domain

type Group struct {
	ID              int64         `json:"id"`
	GroupName       string        `json:"groupName"`
	Owner           User          `json:"owner"`
	PeopleNumber    int64         `json:"peopleNumber"`
	MaxPeopleNumber int64         `json:"maxPeopleNumber"`
	Category        GroupCategory `json:"category"`
	Status          GroupStatus   `json:"status"`
	Users           []User        `json:"users"`
}

type GroupCategory uint

func (g GroupCategory) Uint() uint {
	return uint(g)
}

const (
	GroupCategoryUnknown  GroupCategory = iota
	GroupCategoryActivity               //活动群
	GroupCategoryChat                   //聊天群
)

type GroupStatus uint

func (g GroupStatus) Uint() uint {
	return uint(g)
}

const (
	GroupStatusUnknown GroupStatus = iota
	GroupStatusNormal              //正常
	GroupStatusDisband             //解散
	GroupStatusBan                 //封禁
)

type GroupUserMap struct {
	ID       int64
	Group    Group
	User     User
	Position uint
	Status   GroupUserMapStatus
}

type GroupUserMapStatus uint

func (g GroupUserMapStatus) Uint() uint {
	return uint(g)
}

const (
	GroupUserMapStatusUnknown GroupUserMapStatus = iota
	GroupUserMapStatusJoin                       //加入
	GroupUserMapStatusQuit                       //推出
)
