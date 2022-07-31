package model

import (
	"test/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = mysql.MySqlDb

