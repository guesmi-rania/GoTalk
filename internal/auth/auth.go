package user

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
