package password

import (
	"encoding/base64"
	"github.com/lanyulei/toolkit/logger"
	"golang.org/x/crypto/bcrypt"
)

/*
  @Author : lanyulei
  @Desc :
*/

func DecodePassword(password string) (result string, err error) {
	var (
		passwd []byte
	)

	passwd, err = base64.StdEncoding.DecodeString(password)
	if err != nil {
		logger.Errorf("base64 decode password error: %v", err)
		return
	}

	result = string(passwd)

	return
}

// EncryptionPassword 密码加密
func EncryptionPassword(password string) (result string, err error) {
	var (
		passwd []byte
	)

	// 加密密码
	passwd, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("bcrypt generate password error: %v", err)
		return
	}

	result = string(passwd)
	return
}
