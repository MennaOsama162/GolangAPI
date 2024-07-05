package models

import (
	"gorm.io/gorm"
)

type Author struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"not null;unique"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func MigrateAuthors(db *gorm.DB) {
	db.AutoMigrate(&Author{})
}
