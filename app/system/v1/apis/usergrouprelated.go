package apis

import (
	"openiam/app/system/models"
	"openiam/pkg/tools/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

func CreateUserGroupRelated(c *gin.Context) {
	var (
		err              error
		userGroupRelated struct {
			UserGroup []*models.UserGroupRelated `json:"user_group"`
		}
	)

	err = c.ShouldBindJSON(&userGroupRelated)
	if err != nil {
		response.Error(c, err, respstatus.CreateUserGroupRelatedError)
		return
	}

	err = db.Orm().Create(&userGroupRelated.UserGroup).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateUserGroupRelatedError)
		return
	}

	response.OK(c, userGroupRelated, "ok")
}

func DeleteUserGroupRelated(c *gin.Context) {
	var (
		err              error
		userGroupRelated struct {
			UserGroup []*models.UserGroupRelated `json:"user_group"`
		}
	)

	err = c.ShouldBindJSON(&userGroupRelated)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	dbConn := db.Orm()

	for _, userGroupRelatedId := range userGroupRelated.UserGroup {
		dbConn = dbConn.Or("user_id = ? and group_id = ?", userGroupRelatedId.UserId, userGroupRelatedId.GroupId)
	}

	err = dbConn.Unscoped().Delete(&models.UserGroupRelated{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteUserGroupRelatedError)
		return
	}

	response.OK(c, userGroupRelated.UserGroup, "ok")
}
