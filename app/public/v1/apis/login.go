package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
)

// Login 登陆
func Login(c *gin.Context) {
	response.OK(c, "", "")
}
