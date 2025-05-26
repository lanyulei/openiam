package sql

import (
	"openops/cmd/migrate/sql/statements/openops"
	"openops/common/models"

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
		ListSQL = openops.ListSQL
	}
}
