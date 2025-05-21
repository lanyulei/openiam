package api

import (
	"openiam/app/system/models"
	"openiam/pkg/respstatus"
	"openiam/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
)

func GetConstantRoutes(c *gin.Context) {
	var (
		err    error
		result []*models.MenuTree
	)

	result, err = server.MenuTree(false)
	if err != nil {
		response.Error(c, err, respstatus.GetMenuTreeError)
		return
	}

	response.OK(c, result, "")
}
