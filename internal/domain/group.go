package domain

type Group struct {
	ID        int64  `json:"id"`
	GroupName string `json:"groupName"`
	Owner     User   `json:"owner"`
	Users     []User `json:"users"`
}
