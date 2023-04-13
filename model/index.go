package model

import (
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// create a database connection
// then export it into global var
func OpenConnection() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_fiber_restapi?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

// run migration for model 'book'
func RunMigration(db *gorm.DB) {
	// run database migration
	db.AutoMigrate(&domain.Book{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.ApiKey{})
}
