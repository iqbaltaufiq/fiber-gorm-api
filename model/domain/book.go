package domain

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id          int64          `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Author      string         `gorm:"type:varchar(100)" json:"author"`
	PublishDate string         `gorm:"type:date" json:"publish_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
