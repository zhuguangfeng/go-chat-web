package model

const TableNameUser = "user"

type User struct {
	Base
	Username      string `gorm:"column:username;type:varchar(32);not null;default:'';comment:用户名称" json:"username"`
	Phone         string `gorm:"column:phone;type:char(11);not null;default:'';comment:手机号码" json:"phone"`
	Password      string `gorm:"column:password;type:varchar(128);not null;default:'';comment:用户密码" json:"password"`
	Age           uint   `gorm:"column:age;type:tinyint;not null;default:0;comment:年龄" json:"age"`
	Gender        uint   `gorm:"column:six;type:tinyint;not null;default:0;comment:性别" json:"gender"`
	IsRealName    bool   `gorm:"column:is_real_name;type:tinyint;not null;default:0;comment:是否实名认证" json:"isRealName"`
	IDCard        string `gorm:"column:id_card;type:char(18);not null;default:'';comment:身份证" json:"idCard"`
	Name          string `gorm:"column:name;type:varchar(32);not null;default:'';comment:真是姓名" json:"name"`
	LastLoginIp   string `gorm:"column:login_ip;type:varchar(32);not null;default:'';comment:登录的ip地址" json:"login_ip"`
	LastLoginTime uint   `gorm:"column:last_login_time;type:int;not null;default:0;comment:最后一次登录时间" json:"last_login_time"`
	Status        uint   `gorm:"column:status;type:tinyint;not null;default:0;comment:账号状态" json:"status"`
}

func (User) TableName() string {
	return TableNameUser
}
