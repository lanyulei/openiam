package server

import (
	"openops/app/system/models"

	"github.com/lanyulei/toolkit/db"
)

func MenuTree(isVerify bool) (result []*models.MenuTree, err error) {
	var (
		menus    []*models.Menu
		menuList []*models.MenuTree
	)

	// 查询所有菜单
	err = db.Orm().Model(&models.Menu{}).
		Where("is_verify = ?", isVerify).
		Order("sort").
		Find(&menus).Error
	if err != nil {
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
	return
}
