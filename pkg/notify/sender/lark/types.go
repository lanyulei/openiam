package lark

/*
  @Author : lanyulei
  @Desc :
*/

var (
	accessToken map[string]interface{}
)

const (
	GetTenantAccountTokenURL  = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	GetLarkUserIDByMobilesURL = "https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id"

	CardMessageType = "interactive"
)
