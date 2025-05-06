package ldap

import (
	"encoding/json"
	"errors"
	"fmt"
	"openiam/app/system/models"
	commonModels "openiam/common/models"

	"github.com/go-ldap/ldap/v3"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
*/

func getLdapFields() (ldapFields []map[string]string, err error) {
	var (
		settingsValue models.Settings
		contentList   []map[string]string
	)

	keyValue := "\"key\""
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		keyValue = "`key`"
	}

	err = db.Orm().Model(&settingsValue).Where("? = 'ldap'", keyValue).Find(&settingsValue).Error
	if err != nil {
		return
	}

	err = json.Unmarshal(settingsValue.Content, &contentList)
	if err != nil {
		return
	}

	for _, v := range contentList {
		if v["ldap_field_name"] != "" {
			ldapFields = append(ldapFields, v)
		}
	}
	return
}

func searchRequest(username string) (userInfo *ldap.Entry, err error) {
	var (
		ldapFields       []map[string]string
		cur              *ldap.SearchResult
		ldapFieldsFilter = []string{
			"dn",
		}
	)
	ldapFields, err = getLdapFields()
	for _, v := range ldapFields {
		ldapFieldsFilter = append(ldapFieldsFilter, v["ldap_field_name"])
	}
	// 用来获取查询权限的用户。如果 ldap 禁止了匿名查询，那我们就需要先用这个帐户 bind 以下才能开始查询
	if !viper.GetBool("ldap.anonymousQuery") {
		err = conn.Bind(
			viper.GetString("ldap.bindUserDn"),
			viper.GetString("ldap.bindPwd"))
		if err != nil {
			logger.Error("用户或密码错误。", err)
			return
		}
	}

	sql := ldap.NewSearchRequest(
		viper.GetString("ldap.baseDn"),
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		fmt.Sprintf("(%v=%v)", viper.GetString("ldap.userField"), username),
		ldapFieldsFilter,
		nil)

	if cur, err = conn.Search(sql); err != nil {
		err = errors.New(fmt.Sprintf("在Ldap搜索用户失败, %v", err))
		logger.Error(err)
		return
	}

	if len(cur.Entries) == 0 {
		err = errors.New("未查询到对应的用户信息。")
		logger.Error(err)
		return
	}

	userInfo = cur.Entries[0]

	return
}
