package v1

type UserPwdLoginReq struct {
	Phone    string `json:"phone" json:"phone"`
	Password string `json:"password" json:"password"`
}
