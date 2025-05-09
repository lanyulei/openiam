package sql

import (
	"openiam/cmd/migrate/sql/statements/openiam"
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
	if viper.GetString("db.type") == string(models.DBTypePostgres) {
		ListSQL = openiam.ListSQL
	}
}
