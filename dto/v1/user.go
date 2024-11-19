package v1

type UserPwdLoginReq struct {
	Phone    string `json:"phone" `
	Password string `json:"password"`
}

type UserSmsLoginReq struct {
	Phone string `json:"phone" json:"phone"`
	Code  string `json:"code"`
}
