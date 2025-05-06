package dingtalk

/*
  @Author : lanyulei
  @Desc :
*/

var (
	accessToken map[string]interface{}
)

const (
	GetAccessTokenURL = "https://oapi.dingtalk.com/gettoken"
	GetUserIdURL      = "https://oapi.dingtalk.com/topapi/v2/user/getbymobile"
)

type UserIdResponse struct {
	RequestID string `json:"request_id"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Result    struct {
		UserID                     string   `json:"userid"`
		ExclusiveAccountUserIDList []string `json:"exclusive_account_userid_list"`
	} `json:"result"`
}
