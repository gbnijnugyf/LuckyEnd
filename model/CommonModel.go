package model

import (
	"github.com/shawu21/test/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = mysql.MySqlDb

func init() {
	db.AutoMigrate(
		&User{},
		&Desire{},
	)
}
