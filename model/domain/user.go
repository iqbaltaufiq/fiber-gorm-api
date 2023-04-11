package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int64          `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(50);unique" json:"username"`
	Password  string         `gorm:"type:varchar(255)" json:"password"`
	Role      string         `gorm:"type:varchar(20)" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
