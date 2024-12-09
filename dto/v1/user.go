package v1

import "time"

type UserPwdLoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserSmsLoginReq struct {
	Phone string `json:"phone" json:"phone"`
	Code  string `json:"code"`
}

type User struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	Age           uint      `json:"age,omitempty"`
	Gender        uint      `json:"gender,omitempty"`
	IsRealName    bool      `json:"isRealName,omitempty"`
	LoginIp       string    `json:"loginIp,omitempty"`
	LastLoginTime time.Time `json:"lastTime,omitempty"`
	Status        uint      `json:"status,omitempty"`
	CreatedTime   time.Time `json:"createdTime,omitempty"`
	UpdatedTime   time.Time `json:"updatedTime,omitempty"`
}
