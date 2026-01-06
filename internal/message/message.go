package message

import "gorm.io/gorm"

type Message struct {
    gorm.Model
    Content string
    UserID  uint
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&Message{})
}
