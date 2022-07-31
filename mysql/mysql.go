package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MySqlDb *gorm.DB
var MySqlDbErr error

// func init() {
// 	dbDSN := "root:237156@(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
// 	MySqlDb, MySqlDbErr = gorm.Open("mysql", dbDSN)

// 	if MySqlDbErr != nil {
// 		panic("database open error" + MySqlDbErr.Error())
// 	}
// }