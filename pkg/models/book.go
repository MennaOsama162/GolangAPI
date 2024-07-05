package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            uint           `gorm:"primaryKey"`
	Title         string         `gorm:"not null"`
	ISBN          string         `gorm:"not null;unique"`
	PublishedDate time.Time      `gorm:"type:date"`
	AuthorID      uint           `gorm:"not null"`
	Author        Author         `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE;"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func MigrateBooks(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
