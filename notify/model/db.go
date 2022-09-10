package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        uint                  `gorm:"primarykey"`
	CreatedAt int                   `gorm:"not null"`
	UpdatedAt int                   `gorm:"not null"`
	DeletedAt soft_delete.DeletedAt `gorm:"not null;index"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Channel{},
		&Notify{},
		&NotifyLog{},
	)
}
