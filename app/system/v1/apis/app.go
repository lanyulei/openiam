package apis

import (
	"fmt"
	"openiam/app/system/models"
	"openiam/common/middleware/permission"
	commonModels "openiam/common/models"
	"openiam/pkg/tools/common"
	"openiam/pkg/tools/respstatus"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

func AppList(c *gin.Context) {
	var (
		err  error
		list []*struct {
			models.App
			AppGroupName string `json:"app_group_name"`
		}
		result interface{}
		params struct {
			SearchKey   string `form:"search_key"`
			SearchValue string `form:"search_value"`
			IsExternal  string `form:"is_external"`
		}
	)

	dbConn := db.Orm().Model(&models.App{}).
		Select("system_app_group.name as app_group_name, system_app.*").
		Joins(common.AddQuotesToSQLTableNames("left join system_app_group on system_app_group.id = system_app.app_group_id"))

	err = c.ShouldBindQuery(&params)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	if params.SearchKey != "" && params.SearchValue != "" {
		tableName := "system_app"
		if params.SearchKey == "group" {
			tableName = "system_app_group"
		}
		dbConn = dbConn.Where(fmt.Sprintf("%s.name like ?", tableName), "%"+params.SearchValue+"%")
	}

	if params.IsExternal != "" {
		v := false
		if params.IsExternal == "true" {
			v = true
		}
		dbConn = dbConn.Where("system_app.is_external = ?", v)
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.GetAppError)
		return
	}

	response.OK(c, result, "")
}

func AppListByGroup(c *gin.Context) {
	var (
		err       error
		appGroups []*struct {
			models.AppGroup
			Children []*models.App `gorm:"-" json:"children"`
		}
		apps                     []*models.App
		appMaps                  map[int][]*models.App
		appName                  string
		appGroupList             []interface{}
		groups                   [][]string
		roles                    []string
		menuIds, roleIds, appIds []int
	)
	// 获取所有的应用组
	if err = db.Orm().Model(&models.AppGroup{}).
		Order(`sort, id`).
		Find(&appGroups).
		Error; err != nil {
		response.Error(c, err, respstatus.GetAppGroupError)
		return
	}

	// 获取所有的应用
	appName = c.DefaultQuery("name", "")
	appConn := db.Orm().Model(&models.App{})
	if appName != "" {
		appConn = appConn.Where("name like ?", "%"+appName+"%")
	}

	if !c.GetBool("isAdmin") {
		// 获取当前用户的角色
		groups, err = permission.Enforcer().GetFilteredNamedGroupingPolicy("g", 0, c.GetString("username"))
		if err != nil {
			response.Error(c, err, respstatus.GetRoleError)
			return
		}

		if len(groups) > 0 {
			roles = make([]string, 0, len(groups))
			for _, g := range groups {
				roles = append(roles, g[1])
			}
		}

		// 获取当前用户关联的角色
		keyValue := "\"key\""
		if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
			keyValue = "`key`"
		}
		if err = db.Orm().Model(&models.Role{}).Where("? in (?)", keyValue, roles).Pluck("id", &roleIds).Error; err != nil {
			response.Error(c, err, respstatus.GetRoleError)
			return
		}

		// 获取当前用户关联的菜单
		if err = db.Orm().Model(&models.RoleMenu{}).Select("distinct menu").Where("role in (?)", roleIds).Pluck("menu", &menuIds).Error; err != nil {
			response.Error(c, err, respstatus.GetMenuError)
			return
		}

		// 获取当前用户关联的应用
		if err = db.Orm().Model(&models.Menu{}).Select("distinct app_id").Where("id in (?)", menuIds).Pluck("app_id", &appIds).Error; err != nil {
			response.Error(c, err, respstatus.GetAppError)
			return
		}

		appConn = appConn.Where("id in (?)", appIds)
	}

	if err = appConn.Order(`"sort", id`).Find(&apps).Error; err != nil {
		response.Error(c, err, respstatus.GetAppError)
		return
	}

	// 将应用按照应用组分组
	appMaps = make(map[int][]*models.App)
	for _, app := range apps {
		appMaps[app.AppGroupId] = append(appMaps[app.AppGroupId], app)
	}

	// 将应用组下的应用放入应用组中
	for _, appGroup := range appGroups {
		if _, ok := appMaps[appGroup.Id]; ok && len(appMaps[appGroup.Id]) > 0 {
			appGroup.Children = appMaps[appGroup.Id]
			appGroupList = append(appGroupList, appGroup)
		}
	}

	response.OK(c, appGroupList, "")
}

func AppListByGroupId(c *gin.Context) {
	var (
		err        error
		list       []*models.App
		appGroupId string
	)

	appGroupId = c.Param("id")

	if err = db.Orm().Model(&models.App{}).Where("app_group_id = ?", appGroupId).Find(&list).Error; err != nil {
		response.Error(c, err, respstatus.GetAppError)
		return
	}

	response.OK(c, list, "")
}

func CreateApp(c *gin.Context) {
	var (
		err   error
		app   models.App
		count int64
	)

	if err = c.ShouldBindJSON(&app); err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 名称存在则不创建
	if err = db.Orm().Model(&models.App{}).Where("name = ?", app.Name).Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetAppError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.AppExistError)
		return
	}

	if err = db.Orm().Create(&app).Error; err != nil {
		response.Error(c, err, respstatus.CreateAppError)
		return
	}

	response.OK(c, app, "")
}

func UpdateApp(c *gin.Context) {
	var (
		err        error
		app        models.App
		count      int64
		appId      string
		currentApp models.App
		appMap     = make(map[string]interface{})
	)

	appId = c.Param("id")

	if err = c.ShouldBindJSON(&app); err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 获取当前应用
	if err = db.Orm().First(&currentApp, "id = ?", appId).Error; err != nil {
		response.Error(c, err, respstatus.GetAppError)
		return
	}

	// models.App{}
	appMap = map[string]interface{}{
		"icon":         app.Icon,
		"is_external":  app.IsExternal,
		"link":         app.Link,
		"is_blank":     app.IsBlank,
		"app_group_id": app.AppGroupId,
		"sort":         app.Sort,
		"remarks":      app.Remarks,
	}

	if currentApp.Name != app.Name {
		appMap["name"] = app.Name

		// 名称存在则不更新
		if err = db.Orm().Model(&models.App{}).Where("name = ?", app.Name).Count(&count).Error; err != nil {
			response.Error(c, err, respstatus.GetAppError)
			return
		}

		if count > 0 {
			response.Error(c, err, respstatus.AppExistError)
			return
		}
	}

	if err = db.Orm().Model(&models.App{}).Where("id = ?", appId).Updates(appMap).Error; err != nil {
		response.Error(c, err, respstatus.UpdateAppError)
		return
	}

	response.OK(c, app, "")
}

func DeleteApp(c *gin.Context) {
	var (
		err   error
		appId string
		count int64
	)

	appId = c.Param("id")

	// 判断有没有菜单绑定
	if err = db.Orm().Model(&models.Menu{}).Where("app_id = ?", appId).Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetAppError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.AppHasMenuError)
		return
	}

	// 删除应用
	if err = db.Orm().Delete(&models.App{}, "id = ?", appId).Error; err != nil {
		response.Error(c, err, respstatus.DeleteAppError)
		return
	}

	response.OK(c, "", "")
}
