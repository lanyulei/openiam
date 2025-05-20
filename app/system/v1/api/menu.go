package api

import (
	"openiam/app/system/models"
	"openiam/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

func MenuList(c *gin.Context) {
	response.OK(c, "", "")
}

func MenuTree(c *gin.Context) {
	var (
		menus            []*models.Menu
		menuList, result []*models.MenuTree
		err              error
	)

	// 查询所有菜单
	err = db.Orm().Model(&models.Menu{}).
		Order("sort").
		Find(&menus).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuError)
		return
	}

	for _, menu := range menus {
		menuValue := &models.MenuTree{
			Id:        menu.Id,
			Name:      menu.Name,
			Path:      menu.Path,
			Component: menu.Component,
			ParentId:  menu.ParentId,
			Redirect:  menu.Redirect,
			Meta: models.MenuMeta{
				Title:       menu.Title,
				Hyperlink:   menu.Hyperlink,
				IsHide:      menu.IsHide,
				IsKeepAlive: menu.IsKeepAlive,
				IsAffix:     menu.IsAffix,
				IsIframe:    menu.IsIframe,
				Icon:        menu.Icon,
			},
		}
		menuList = append(menuList, menuValue)
	}

	// 构建菜单树
	menuMap := make(map[string]*models.MenuTree)
	for _, menu := range menuList {
		menuMap[menu.Id] = menu
	}

	for _, menu := range menuList {
		if menu.ParentId != "" {
			if parent, ok := menuMap[menu.ParentId]; ok {
				if parent.Children == nil {
					parent.Children = make([]*models.MenuTree, 0)
				}
				parent.Children = append(parent.Children, menu)
			}
		} else {
			result = append(result, menu)
		}
	}

	response.OK(c, result, "")
}
