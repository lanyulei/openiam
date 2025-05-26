package common

import (
	"openops/common/models"
	"regexp"

	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

// AddQuotesToSQLTableNames 函数会将 SQL 语句中的表名加上双引号
func AddQuotesToSQLTableNames(sql string) string {
	// 定义正则表达式，匹配 SQL 中的表名
	re := regexp.MustCompile(`\b([a-zA-Z0-9_]+)\b`)

	// 使用正则表达式替换表名，加上双引号
	return re.ReplaceAllStringFunc(sql, func(match string) string {
		// 排除 SQL 关键字，如 left, join, on, =
		if match == "left" || match == "join" || match == "on" || match == "=" || match == "as" {
			return match
		}

		result := `"` + match + `"`
		if viper.GetString("db.type") == string(models.DBTypeMySQL) {
			return "`" + match + "`"
		}

		// 为每个匹配的表名加上双引号
		return result
	})
}
