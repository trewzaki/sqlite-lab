package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitialSqliteDatabase() error {
	var connectionErr error
	DB, connectionErr = gorm.Open("sqlite3", "./models/user.db")
	autoMigrate()

	return connectionErr
}

func autoMigrate() {
	DB.AutoMigrate(&User{})
}
