package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
)

func MenuList(c *gin.Context) {
	response.OK(c, "", "")
}
