package model

import (
	"test/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = mysql.MySqlDb

