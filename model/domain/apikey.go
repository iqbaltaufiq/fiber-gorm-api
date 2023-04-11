package domain

import (
	"time"

	"gorm.io/gorm"
)

type ApiKey struct {
	Id        int64          `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(50)" json:"username"`
	ApiKey    string         `gorm:"type:varchar(255)" json:"api_key"`
	Expires   time.Time      `gorm:"type:datetime" json:"expires"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
