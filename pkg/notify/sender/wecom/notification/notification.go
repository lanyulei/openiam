package notification

import (
	"openiam/pkg/notify/sender/wecom"

	"github.com/lanyulei/toolkit/logger"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Send(content map[string]interface{}) (result map[string]interface{}, err error) {
	var (
		accessToken string
	)

	accessToken, err = wecom.GetAccountToken()
	if err != nil {
		logger.Errorf("failed to get access token, err:%s\n", err.Error())
		return
	}

	err = gout.POST(SendMessageURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": accessToken}).
		SetJSON(content).
		BindJSON(&result).
		Do()
	if err != nil {
		logger.Errorf("failed to send message, err:%s\n", err.Error())
		return
	}

	return
}
