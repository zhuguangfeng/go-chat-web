package domain

type User struct {
	ID            int64  `json:"id"`
	UserName      string `json:"username"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Age           uint   `json:"age"`
	Gender        uint   `json:"gender"`
	IsRealName    bool   `json:"isRealName"`
	Name          string `json:"name"`
	IDCard        string `json:"idCard"`
	LastLoginIp   string `json:"loginIp"`
	LastLoginTime uint   `json:"lastTime"`
	Status        uint   `json:"status"`
	CreatedTime   uint   `json:"createdTime"`
	UpdatedTime   uint   `json:"updatedTime"`
}
