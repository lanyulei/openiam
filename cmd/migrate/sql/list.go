package sql

import (
	"openiam/cmd/migrate/sql/statements/openiam/mysql"
	"openiam/cmd/migrate/sql/statements/openiam/postgres"
	"openiam/common/models"

	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

var (
	ListSQL = make([]map[string]string, 0)
)

func Init() {
	if viper.GetString("db.type") == string(models.DBTypeMySQL) {
		ListSQL = mysql.ListSQL
	} else if viper.GetString("db.type") == string(models.DBTypePostgres) {
		ListSQL = postgres.ListSQL
	}
}
