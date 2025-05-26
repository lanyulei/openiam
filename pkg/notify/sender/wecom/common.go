package wecom

import (
	"fmt"
	"openops/pkg/notify/commom"
	"time"

	"github.com/lanyulei/toolkit/logger"

	"github.com/guonaihong/gout"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

// GetAccountToken
// @Description: get wecom account token
// @return err
func GetAccountToken() (at string, err error) {
	var (
		result map[string]interface{}
	)

	if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
		if accessToken == nil {
			accessToken = make(map[string]interface{})
		}

		err = gout.GET(GetAccountTokenURL).
			SetQuery(gout.H{
				"corpid":     viper.GetString("notify.weComNotify.corpId"),
				"corpsecret": viper.GetString("notify.weComNotify.appSecret"),
			}).
			BindJSON(&result).
			Do()
		if err != nil {
			logger.Errorf("failed to get access token, err:%s\n", err.Error())
			return
		}

		if int(result["errcode"].(float64)) != 0 {
			err = fmt.Errorf("failed to get wecom access token, err:%s\n", result["errmsg"])
			return
		}

		accessToken["expires_time"] = time.Now().Unix() + int64(result["expires_in"].(float64))
		accessToken["access_token"] = result["access_token"]
	}

	at = accessToken["access_token"].(string)

	return
}

func GetWeComUserId(mobile string) (result string, err error) {
	var (
		accessToken string
		req         map[string]interface{}
	)

	accessToken, err = GetAccountToken()
	if err != nil {
		logger.Errorf("failed to get wecom account token, err:%s\n", err.Error())
		return
	}

	err = gout.POST("https://qyapi.weixin.qq.com/cgi-bin/user/getuserid").
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": accessToken}).
		SetJSON(map[string]interface{}{
			"mobile": mobile,
		}).
		BindJSON(&req).
		Do()
	if err != nil {
		logger.Errorf("failed to get wecom user id, err:%s\n", err.Error())
		return
	}

	if int(req["errcode"].(float64)) != 0 {
		err = fmt.Errorf("failed to get wecom user id, err:%s\n", req["errmsg"])
		return
	}

	result = req["userid"].(string)
	return
}

func FormatMarkdown(title string, message *commom.Message) (content string) {
	return fmt.Sprintf(`### %s
><font color="comment">标题:</font> %s
><font color="comment">优先级:</font> %s
><font color="comment">申请人:</font> %s
><font color="comment">申请时间:</font> %s
><font color="comment">最近处理时间:</font> %s`,
		title,
		message.Title,
		message.Priority,
		message.Creator,
		message.CreatedAt,
		message.UpdatedAt,
	)
}
