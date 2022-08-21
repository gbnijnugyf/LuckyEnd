package model

import (
	"gorm.io/gorm"
	"test/mysql"
)

var db *gorm.DB = mysql.MySqlDb
