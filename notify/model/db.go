package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Notify{},
		&NotifyLog{},
	)
}
