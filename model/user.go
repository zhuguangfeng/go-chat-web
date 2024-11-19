package model

import (
	"time"
)

const TableNameUser = "user"

type User struct {
	Base
	UserName      string    `gorm:"column:username;type:varchar(32);not null;default:'';comment:用户名称" json:"username"`
	Phone         string    `gorm:"column:phone;type:char(11);not null;default:'';comment:手机号码" json:"phone"`
	Password      string    `gorm:"column:password;type:varchar(128);not null;default:'';comment:用户密码" json:"password"`
	Age           uint      `gorm:"column:age;type:tinyint;not null;default:0;comment:性别" json:"age"`
	Six           uint      `gorm:"column:six;type:tinyint;not null;default:0;comment:年龄" json:"six"`
	LoginIp       string    `gorm:"column:login_ip;type:varchar(32);not null;default:'';comment:登录的ip地址" json:"login_ip"`
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime;not null;default:'';comment:最后一次登录时间" json:"last_login_time"`
	Status        uint      `gorm:"column:status;type:tinyint;not null;default:0;comment:账号状态" json:"status"`
}

func (User) TableName() string {
	return TableNameUser
}
