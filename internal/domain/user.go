package domain

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	IDCard   string `json:"idCard"`
	Age      int64  `json:"age"`
	Gender   int64  `json:"gender"`
}
