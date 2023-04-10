package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDBConnection() {
	// create a new connection
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_fiber_restapi?parseTime=true"))

	if err != nil {
		panic(err)
	}

	// run database migration
	db.AutoMigrate(&Book{})

	// export the conn object
	DB = db
}
