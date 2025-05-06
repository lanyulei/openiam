package notification

import (
	"openiam/pkg/notify/sender/lark"

	"github.com/lanyulei/toolkit/logger"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Send
// @Description: send lark job notification
// @param string
// @param content
// @return result
// @return err
func Send(mobiles []string, content map[string]interface{}) (result map[string]interface{}, err error) {
	var (
		accessToken         string
		receiveIds          []string
		larkUserIdsResponse map[string]interface{}
	)

	// get account token
	accessToken, err = lark.GetTenantAccountToken()
	if err != nil {
		logger.Errorf("get account token failed, %s", err.Error())
		return
	}

	// get user id
	larkUserIdsResponse, err = lark.GetLarkUserIDByMobiles(mobiles)
	if err != nil {
		logger.Errorf("get user id failed, %s", err.Error())
		return
	}

	if d, ok := larkUserIdsResponse["data"]; ok {
		if u, ok := d.(map[string]interface{})["user_list"]; ok {
			for _, v := range u.([]interface{}) {
				if userId, ok := v.(map[string]interface{})["user_id"]; ok {
					receiveIds = append(receiveIds, userId.(string))
				}
			}
		}
	}

	data := map[string]interface{}{
		"msg_type": lark.CardMessageType,
		"user_ids": receiveIds,
		"card":     content,
	}

	err = gout.POST(NotifyBaseURL).
		SetHeader(gout.H{"Content-Type": "application/json", "Authorization": "Bearer " + accessToken}).
		SetJSON(data).
		BindJSON(&result).
		Do()
	if err != nil {
		logger.Errorf("send notification failed, %s", err.Error())
		return
	}
	if code, ok := result["code"]; ok && int(code.(float64)) == 0 {
		logger.Infof("send notification success, result: %s", result["msg"])
	} else {
		logger.Errorf("send notification failed, result: %s", result["msg"])
	}

	return
}
