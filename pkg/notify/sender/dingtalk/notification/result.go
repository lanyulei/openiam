package notification

import (
	"encoding/json"
	"fmt"
	"openops/pkg/notify/sender/dingtalk"

	"github.com/guonaihong/gout"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

func GetResult(taskId int) (res []byte, err error) {
	var (
		data        []byte
		accessToken string
	)

	accessToken, err = dingtalk.GetAccountToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"agent_id": viper.GetString("notify.dingTalkNotify.agentId"),
		"task_id":  taskId,
	}

	data, err = json.Marshal(params)
	if err != nil {
		err = fmt.Errorf("json serialization failed, %s", err.Error())
		return
	}

	err = gout.POST(NotifyResultURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": accessToken}).
		SetBody(data).
		BindBody(&res).
		Do()
	if err != nil {
		err = fmt.Errorf("failed to get notification details by task id, %s", err.Error())
		return
	}

	return
}
