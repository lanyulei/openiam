package apis

import (
	"openiam/pkg/route"
	"openiam/pkg/tools/respstatus"

	//"openiam/common/router"
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

// CheckRegisterRoute 验证路由是否注册
func CheckRegisterRoute(c *gin.Context) {
	var (
		err    error
		routes struct {
			Routes []*route.Route `json:"routes"`
		}
		result map[string][]*route.Route
	)

	err = c.ShouldBindJSON(&routes)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	result, err = route.CheckRegisterRoute(routes.Routes)
	if err != nil {
		response.Error(c, err, respstatus.CheckRegisterRouteError)
		return
	}

	response.OK(c, result, "")
}
