package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB global databsae instance
var DB *gorm.DB

// InitialSqliteDatabase function
func InitialSqliteDatabase() error {
	var connectionErr error
	DB, connectionErr = gorm.Open("sqlite3", "./models/user.db")
	autoMigrate()

	return connectionErr
}

func autoMigrate() {
	DB.AutoMigrate(&User{})
}
