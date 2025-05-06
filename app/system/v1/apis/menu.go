package apis

import (
	"openiam/app/system/models"
	"openiam/pkg/tools/common"
	"openiam/pkg/tools/respstatus"
	serverMenu "openiam/server/system/menu"
	"strconv"

	"github.com/gin-gonic/gin"
	dbConn "github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

// MenuTree 菜单树
func MenuTree(c *gin.Context) {
	var (
		err    error
		result map[string]interface{}
	)

	result, err = serverMenu.Tree(c, true)
	if err != nil {
		response.Error(c, err, respstatus.GetMenuButtonError)
		return
	}

	response.OK(c, result, "")
}

// UserMenuTree 用户拥有权限的菜单
func UserMenuTree(c *gin.Context) {
	var (
		err    error
		result map[string]interface{}
	)

	result, err = serverMenu.Tree(c, false)
	if err != nil {
		response.Error(c, err, respstatus.GetMenuButtonError)
		return
	}

	response.OK(c, result, "")
}

// SaveMenu 保存菜单
func SaveMenu(c *gin.Context) {
	var (
		err               error
		menu, currentMenu models.Menu
	)

	err = c.ShouldBind(&menu)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	db := dbConn.Orm().Model(&models.Menu{})

	if menu.Id != 0 {
		db = db.Where("id = ?", menu.Id)

		err = dbConn.Orm().Where("id = ?", menu.Id).Find(&currentMenu).Error
		if err != nil {
			response.Error(c, err, respstatus.GetMenuError)
			return
		}
		menu.CreatedAt = currentMenu.CreatedAt
	}

	err = db.Save(&menu).Error
	if err != nil {
		response.Error(c, err, respstatus.SaveMenuError)
		return
	}

	response.OK(c, menu, "")
}

// DeleteMenu 删除菜单
func DeleteMenu(c *gin.Context) {
	var (
		err       error
		menuId    string
		menuCount int64
	)

	menuId = c.Param("id")

	// 确认是否有菜单节点，若有，则不允许删除
	err = dbConn.Orm().Model(&models.Menu{}).Where("parent = ?", menuId).Count(&menuCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuError)
		return
	}

	if menuCount > 0 {
		response.Error(c, err, respstatus.SubmenuExistsError)
		return
	}

	err = dbConn.Orm().Delete(&models.Menu{}, menuId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteMenuError)
		return
	}

	response.OK(c, "", "")
}

// BatchDeleteMenu 批量删除菜单
func BatchDeleteMenu(c *gin.Context) {
	var (
		err     error
		menuIds struct {
			MenuIds []int `json:"menu_ids"`
		}
		menuCount int64
	)

	err = c.ShouldBind(&menuIds)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 确认是否存在子节点
	err = dbConn.Orm().Model(&models.Menu{}).Where("parent in (?)", menuIds.MenuIds).Count(&menuCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuError)
		return
	}
	if menuCount > 0 {
		response.Error(c, err, respstatus.SubmenuExistsError)
		return
	}

	err = dbConn.Orm().Delete(&models.Menu{}, menuIds.MenuIds).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteMenuError)
		return
	}

	response.OK(c, "", "")
}

// MenuButton 查询菜单下的所有按钮
func MenuButton(c *gin.Context) {
	var (
		err        error
		menuId     string
		buttonList []*models.Menu
	)

	menuId = c.Param("id")

	err = dbConn.Orm().Model(&models.Menu{}).
		Where("parent = ?", menuId).
		Order("sort, id").
		Find(&buttonList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuButtonError)
		return
	}

	response.OK(c, buttonList, "")
}

// MenuBindApi 菜单绑定API
func MenuBindApi(c *gin.Context) {
	var (
		err    error
		params struct {
			Type int   `json:"type"` // 1 绑定， 0 解绑
			Apis []int `json:"apis"`
		}
		existApis []int
		apiMaps   map[int]struct{}
		menuApi   []models.MenuApi
		menuId    string
	)

	menuId = c.Param("id")
	menuIdInt, _ := strconv.Atoi(menuId)

	err = c.ShouldBind(&params)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	if params.Type == 1 {
		err = dbConn.Orm().Model(&models.MenuApi{}).Where("menu = ? and api in (?)", menuIdInt, params.Apis).Pluck("api", &existApis).Error
		if err != nil {
			response.Error(c, err, respstatus.GetMenuApiError)
			return
		}
		apiMaps = make(map[int]struct{})
		if len(existApis) != 0 {
			for _, i := range existApis {
				apiMaps[i] = struct{}{}
			}
		}

		menuApi = make([]models.MenuApi, 0)
		for _, i := range params.Apis {
			if _, ok := apiMaps[i]; !ok {
				menuApi = append(menuApi, models.MenuApi{
					Menu: menuIdInt,
					Api:  i,
				})
			}
		}

		if len(menuApi) > 0 {
			err = dbConn.Orm().Create(&menuApi).Error
			if err != nil {
				response.Error(c, err, respstatus.MenuBindApiError)
				return
			}
		}
	} else if params.Type == 0 {
		err = dbConn.Orm().Unscoped().Where("menu = ? and api in (?)", menuIdInt, params.Apis).Delete(&models.MenuApi{}).Error
		if err != nil {
			response.Error(c, err, respstatus.MenuUnBindApiError)
			return
		}
	}

	response.OK(c, "", "")
}

// MenuApis 查询菜单绑定的API
func MenuApis(c *gin.Context) {
	var (
		err     error
		menuId  string
		apiList []int
	)

	menuId = c.Param("id")

	err = dbConn.Orm().Model(&models.MenuApi{}).
		Select("distinct api").
		Where("menu = ?", menuId).
		Pluck("api", &apiList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuApiError)
		return
	}

	response.OK(c, apiList, "")
}

// MenuApiList 查询菜单绑定的API列表
func MenuApiList(c *gin.Context) {
	var (
		err     error
		menuId  string
		apiIds  []int
		apiList []*struct {
			models.Api
			GroupName string `json:"group_name"`
		}
	)

	menuId = c.Param("id")

	err = dbConn.Orm().Model(&models.MenuApi{}).
		Select("distinct api").
		Where("menu = ?", menuId).
		Pluck("api", &apiIds).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuApiError)
		return
	}

	err = dbConn.Orm().Model(&models.Api{}).
		Select("system_api_group.name as group_name, system_api.*").
		Joins(common.AddQuotesToSQLTableNames("left join system_api_group on system_api_group.id = system_api.group")).
		Where("system_api.id in (?)", apiIds).
		Find(&apiList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuApiError)
		return
	}

	response.OK(c, apiList, "")
}
