package common

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CreateConfig(ak, sk string) *openapi.Config {
	config := &openapi.Config{}

	if ak != "" && sk != "" {
		config.AccessKeyId = tea.String(ak)
		config.AccessKeySecret = tea.String(sk)
	}

	return config
}
