package common

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CreateConfig(ak, sk string) *common.Credential {
	var (
		credential common.Credential
	)

	if ak != "" && sk != "" {
		credential.SecretId = ak
		credential.SecretKey = sk
	}

	return &credential
}
