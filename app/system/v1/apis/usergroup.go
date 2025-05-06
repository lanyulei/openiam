package apis

import (
	"openiam/app/system/models"
	"openiam/pkg/tools/respstatus"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

func UserGroups(c *gin.Context) {
	var (
		err        error
		userGroups []*models.UserGroup
		query      struct {
			Name     string `form:"name"`
			GroupIds string `form:"group_ids"`
			App      string `form:"app"`
		}
		ids []string
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	ids = strings.Split(query.GroupIds, ",")

	dbConn := db.Orm()

	if query.Name != "" {
		dbConn = dbConn.Where("name like ?", "%"+query.Name+"%")
	}

	if query.App != "" {
		dbConn = dbConn.Where("app = ?", query.App)
	}

	if len(ids) > 0 && query.GroupIds != "" {
		dbConn = dbConn.Where("id in (?)", ids)
	}

	err = dbConn.Order("id desc").Find(&userGroups).Error
	if err != nil {
		response.Error(c, err, respstatus.UserGroupListError)
		return
	}

	response.OK(c, userGroups, "ok")
}

func UserGroupList(c *gin.Context) {
	var (
		err           error
		userGroupList []*struct {
			models.UserGroup
			Users []*models.User `gorm:"-" json:"users"`
		}
		result  interface{}
		userIds []int
		query   struct {
			Name string `form:"name"`
			App  string `form:"app"`
		}
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	dbConn := db.Orm().Model(&models.UserGroup{})

	if query.Name != "" {
		dbConn = dbConn.Where("name like ?", "%"+query.Name+"%")
	}

	if query.App != "" {
		dbConn = dbConn.Where("app like ?", "%"+query.App+"%")
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &userGroupList)
	if err != nil {
		response.Error(c, err, respstatus.UserGroupListError)
		return
	}

	for _, userGroup := range userGroupList {
		err = db.Orm().Model(&models.UserGroupRelated{}).Where("group_id = ?", userGroup.Id).Pluck("user_id", &userIds).Error
		if err != nil {
			response.Error(c, err, respstatus.UserIdsError)
			return
		}

		err = db.Orm().Model(&models.User{}).Where("id in (?)", userIds).Find(&userGroup.Users).Error
		if err != nil {
			response.Error(c, err, respstatus.GetUserError)
			return
		}
	}

	response.OK(c, result, "ok")
}

func CreateUserGroup(c *gin.Context) {
	var (
		err       error
		userGroup models.UserGroup
	)

	err = c.ShouldBindJSON(&userGroup)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	err = db.Orm().Create(&userGroup).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateUserGroupError)
		return
	}

	response.OK(c, userGroup, "ok")
}

func UpdateUserGroup(c *gin.Context) {
	var (
		err         error
		userGroup   models.UserGroup
		userGroupId string
	)

	userGroupId = c.Param("id")

	err = c.ShouldBindJSON(&userGroup)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	err = db.Orm().Model(&userGroup).Where("id = ?", userGroupId).Updates(map[string]interface{}{
		"name": userGroup.Name,
		"app":  userGroup.App,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateUserGroupError)
		return
	}

	response.OK(c, userGroup, "ok")
}

func DeleteUserGroup(c *gin.Context) {
	var (
		err         error
		userGroupId string
	)

	userGroupId = c.Param("id")

	// 判断是否绑定数据，若绑定，则无法删除
	err = db.Orm().Where("group_id = ?", userGroupId).First(&models.UserGroupRelated{}).Error
	if err == nil {
		response.Error(c, err, respstatus.UserGroupRelatedExistError)
		return
	}

	err = db.Orm().Where("id = ?", userGroupId).Delete(&models.UserGroup{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteUserGroupError)
		return
	}

	response.OK(c, nil, "ok")
}
