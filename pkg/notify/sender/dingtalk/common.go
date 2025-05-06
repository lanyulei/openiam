package dingtalk

import (
	"encoding/json"
	"fmt"
	"openiam/pkg/notify/commom"
	"time"

	"github.com/guonaihong/gout"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

// GetAccountToken
// @Description: get dingtalk account token
// @return err
func GetAccountToken() (at string, err error) {
	var (
		result []byte
	)

	if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
		err = gout.GET(GetAccessTokenURL).SetQuery(gout.H{
			"appkey":    viper.GetString("notify.dingTalkNotify.appKey"),
			"appsecret": viper.GetString("notify.dingTalkNotify.appSecret"),
		}).BindBody(&result).Do()
		if err != nil {
			err = fmt.Errorf("failed to get access token, err:%s\n", err.Error())
			return
		}

		err = json.Unmarshal(result, &accessToken)
		if err != nil {
			err = fmt.Errorf("failed to get access token, err:%s\n", err.Error())
			return
		}

		if errCode, ok := accessToken["errcode"]; ok && int(errCode.(float64)) != 0 {
			err = fmt.Errorf("failed to get dingtalk access token, err:%s\n", accessToken["errmsg"].(string))
			return
		}

		accessToken["expires_time"] = time.Now().Unix() + int64(accessToken["expires_in"].(float64))
	}

	at = accessToken["access_token"].(string)

	return
}

// GetDingtalkUserId
// @Description: 通过手机号获取钉钉用户的 userId，企业应用需拥有此接口权限，https://open.dingtalk.com/document/orgapp/obtain-the-userid-of-your-mobile-phone-number
// @param mobile 手机号
// @return res 返回结果
// @return err 错误信息
func GetDingtalkUserId(mobile string) (res UserIdResponse, err error) {
	var (
		accessTokenValue string
	)

	accessTokenValue, err = GetAccountToken()
	if err != nil {
		logger.Errorf("failed to get account token, err:%s", err.Error())
		return
	}

	err = gout.POST(GetUserIdURL).
		SetQuery(gout.H{
			"access_token": accessTokenValue,
		}).
		SetHeader(gout.H{
			"Accept": "application/json",
		}).
		SetJSON(gout.H{
			"mobile": mobile,
		}).
		BindJSON(&res).Do()
	if err != nil {
		logger.Errorf("failed to get user id, err:%s", err.Error())
		return
	}

	return
}

func FormatMarkdown(title string, message *commom.Message) (res string) {
	return fmt.Sprintf("### %s  \n  > 标题：%s  \n  > 优先级：%s  \n  > 申请人：%s  \n  > 申请时间：%s  \n  > 最近处理时间：%s",
		title,
		message.Title,
		message.Priority,
		message.Creator,
		message.CreatedAt,
		message.UpdatedAt,
	)
}
