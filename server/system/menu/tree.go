package menu

import (
	"encoding/json"
	"errors"
	"openiam/app/system/models"
	"openiam/common/middleware/permission"
	commonModels "openiam/common/models"
	"openiam/pkg/route"
	"openiam/pkg/tools/common"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Tree(c *gin.Context, all bool) (res map[string]interface{}, err error) {
	var (
		menus            []*models.Menu
		result           []*models.MenuValue
		m                models.MenuTree
		buttonList       []*models.Menu
		groups           [][]string
		roles            []string
		roleIds, menuIds []int
		path             string
		appValue         models.App
		appId            string
	)

	appId = c.DefaultQuery("appId", "")
	path = c.DefaultQuery("path", "")
	if path == "" && appId == "" {
		err = errors.New("path or app_id is empty")
		logger.Errorf(err.Error())
		return
	}

	// get the application id
	if appId != "" {
		err = db.Orm().Model(&models.App{}).Where("id = ?", appId).First(&appValue).Error
		if err != nil {
			logger.Errorf("get app info failed with, error %v", err)
			return
		}
	} else {
		appValue, err = getAppInfo(path)
		if err != nil {
			logger.Errorf("get app info failed with, error %v", err)
			return
		}
	}

	roleIds = make([]int, 0)

	dbConn := db.Orm().Model(&models.Menu{})
	bottomConn := db.Orm().Model(&models.Menu{})
	if !all && !c.GetBool("isAdmin") {
		groups, err = permission.Enforcer().GetFilteredNamedGroupingPolicy("g", 0, c.GetString("username"))
		if err != nil {
			logger.Errorf("get user role failed with, error %v", err)
			return
		}

		if len(groups) > 0 {
			roles = make([]string, 0, len(groups))
			for _, g := range groups {
				roles = append(roles, g[1])
			}
		}

		keyValue := "\"key\""
		if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
			keyValue = "`key`"
		}
		err = db.Orm().Model(&models.Role{}).Where("? in (?)", keyValue, roles).Pluck("id", &roleIds).Error
		if err != nil {
			logger.Errorf("query role id failed with, error %v", err)
			return
		}

		dbConn = dbConn.
			Select("distinct system_menu.*").
			Joins(common.AddQuotesToSQLTableNames("left join system_role_menu on system_role_menu.menu = system_menu.id")).
			Where(`system_role_menu.role in (?)`, roleIds)

		bottomConn = bottomConn.
			Select("distinct system_menu.*").
			Joins(common.AddQuotesToSQLTableNames("left join system_role_menu on system_role_menu.menu = system_menu.id")).
			Where(`system_role_menu.role in (?)`, roleIds)
	}

	err = dbConn.Where("system_menu.type in (1, 2) and app_id = ?", appValue.Id).
		Order("system_menu.sort, system_menu.id").
		Find(&menus).Error
	if err != nil {
		logger.Errorf("query menu failed with, error %v", err)
		return
	}

	for _, m := range menus {
		menuIds = append(menuIds, m.Id)
	}

	m = models.MenuTree{Menus: menus}
	result = m.GetMenuTree()

	// 查询所有页面对应的按钮列表
	err = bottomConn.
		Where("system_menu.type = 3 and system_menu.parent in (?)", menuIds).
		Order("system_menu.sort, system_menu.id").
		Find(&buttonList).Error
	if err != nil {
		logger.Errorf("query menu button failed with, error %v", err)
		return
	}

	buttonMap := make(map[int][]*models.Menu)
	for _, b := range buttonList {
		if _, ok := buttonMap[b.Parent]; ok {
			buttonMap[b.Parent] = append(buttonMap[b.Parent], b)
		} else {
			buttonMap[b.Parent] = []*models.Menu{b}
		}
	}

	return map[string]interface{}{
		"menu":   result,
		"button": buttonMap,
	}, nil
}

func getAppInfo(path string) (app models.App, err error) {
	var (
		menu     map[string]interface{}
		menus    []models.Menu
		menuMaps []map[string]interface{}
		menuByte []byte
	)

	err = db.Orm().Model(&models.Menu{}).Where("path = ?", path).Find(&menu).Error
	if err != nil {
		logger.Errorf("query menu failed with, error %v", err)
		return
	}

	if _, ok := menu["id"]; !ok {
		// 获取所有带有路径参数的菜单
		err = db.Orm().Model(&models.Menu{}).Where("path like ?", "%/:%").Find(&menus).Error
		if err != nil {
			logger.Errorf("query menu failed with, error %v", err)
			return
		}

		menuByte, err = json.Marshal(menus)
		if err != nil {
			return
		}

		err = json.Unmarshal(menuByte, &menuMaps)
		if err != nil {
			return
		}

		menu, ok = route.MatchRouter(menuMaps, path, "")
		if !ok {
			err = errors.New("menu not found")
			logger.Errorf(err.Error())
			return
		}
	}

	if _, ok := menu["app_id"]; ok {
		err = db.Orm().Model(&models.App{}).Where("id = ?", menu["app_id"]).First(&app).Error
		if err != nil {
			logger.Errorf("query application failed with, error %v", err)
			return
		}
	}

	return
}
