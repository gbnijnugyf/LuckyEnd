package mysql

import (
	"fmt"

	"github.com/shawu21/test/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlDb *gorm.DB
var MySqlDbErr error

func init() {
	dbConfig := config.GetDbConfig()
	dbDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Hostname,
		dbConfig.Port,
		dbConfig.Dbname,
		dbConfig.Charset,
		dbConfig.PareseTime,
		dbConfig.Local,
	)
	MySqlDb, MySqlDbErr = gorm.Open(mysql.Open(dbDSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if MySqlDbErr != nil {
		panic("database open error" + MySqlDbErr.Error())
	}
}