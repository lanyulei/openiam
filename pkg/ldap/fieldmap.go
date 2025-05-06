package ldap

import (
	"openiam/app/system/models"

	"github.com/go-ldap/ldap/v3"
)

/*
  @Author : lanyulei
*/

func FieldsMap(ldapUserInfo *ldap.Entry) (userInfo models.User, err error) {
	var (
		ldapFields []map[string]string
	)

	ldapFields, err = getLdapFields()
	if err != nil {
		return
	}

	for _, v := range ldapFields {
		switch v["local_field_name"] {
		case "nickname":
			userInfo.Nickname = ldapUserInfo.GetAttributeValue(v["ldap_field_name"])
		case "tel":
			userInfo.Tel = ldapUserInfo.GetAttributeValue(v["ldap_field_name"])
		case "avatar":
			userInfo.Avatar = ldapUserInfo.GetAttributeValue(v["ldap_field_name"])
		case "email":
			userInfo.Email = ldapUserInfo.GetAttributeValue(v["ldap_field_name"])
		case "remark":
			userInfo.Remark = ldapUserInfo.GetAttributeValue(v["ldap_field_name"])
		}
	}

	return
}
