package domain

type Group struct {
	ID              int64  `json:"id"`
	GroupName       string `json:"groupName"`
	Owner           User   `json:"owner"`
	PeopleNumber    int64  `json:"peopleNumber"`
	MaxPeopleNumber int64  `json:"maxPeopleNumber"`
	Category        uint   `json:"category"`
	Status          uint   `json:"status"`
	Users           []User `json:"users"`
}
