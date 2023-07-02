package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Con *gorm.DB

func ConnectDatabase() {

	connection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	Con = connection
}
