package model

import (
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// create a database connection
// then export it into global var
func SetupDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_fiber_restapi?parseTime=true"))
	if err != nil {
		panic(err)
	}

	DB = db
}

// run migration for model 'book'
func RunMigration() {
	// run database migration
	DB.AutoMigrate(&domain.Book{})
	DB.AutoMigrate(&domain.Admin{})
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.ApiKey{})
}
