package model

import (
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		//&Dynamic{},
		&Activity{},
		&Review{},
		&ActivitySignup{},
		&Group{},
		&GroupUserMap{},
	)
}
