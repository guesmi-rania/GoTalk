package user

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string
    Email string
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&User{})
}
