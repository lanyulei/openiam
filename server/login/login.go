package login

import (
	"encoding/base64"
	"fmt"
	"openiam/app/system/models"
	"openiam/pkg/ldap"
	"openiam/server/system"

	"github.com/lanyulei/toolkit/logger"

	"github.com/lanyulei/toolkit/db"

	ldapV3 "github.com/go-ldap/ldap/v3"

	"golang.org/x/crypto/bcrypt"
)

/*
  @Author : lanyulei
  @Desc :
*/

func LoginHandler(username, password string, isLdap bool) (user models.UserRequest, err error) {
	var (
		userInfo        *ldapV3.Entry
		addUserInfo     models.User
		decodedPassword []byte
	)

	defer func() {
		err := recover()
		if err != nil {
			logger.Infof("login panic, error: %v", err)
		}
	}()

	if isLdap {
		// ldap登陆
		userInfo, err = ldap.Login(username, password)
		if err != nil {
			err = fmt.Errorf("ldap login error: %v", err)
			return
		}
	}

	// 查询用户信息
	err = db.Orm().Model(&models.User{}).Where("username = ?", username).Find(&user).Error
	if err != nil {
		err = fmt.Errorf("query user error: %v", err)
		return
	}

	if user.Username == "" {
		if isLdap {
			addUserInfo, err = ldap.FieldsMap(userInfo)
			if err != nil {
				err = fmt.Errorf("ldap fields map error: %v", err)
				return
			}

			addUserInfo.Username = username
			addUserInfo.Status = true

			// 用户保存到本地
			err = db.Orm().Create(&addUserInfo).Error
			if err != nil {
				err = fmt.Errorf("create user error: %v", err)
				return
			}

			user.User = addUserInfo
		} else {
			err = fmt.Errorf("user not exist")
			return
		}
	}

	if !user.Status {
		err = fmt.Errorf("user is disabled")
		return
	}

	if !isLdap {
		decodedPassword, err = base64.StdEncoding.DecodeString(password)
		if err != nil {
			err = fmt.Errorf("decode password error: %v", err)
			return
		}

		password = string(decodedPassword)[:len(string(decodedPassword))-system.SaltNumber]

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			err = fmt.Errorf("password error: %v", err)
			return
		}
	}

	return
}
