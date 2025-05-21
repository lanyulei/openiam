package api

import (
	"fmt"
	"openiam/app/system/models"
	"openiam/pkg/respstatus"
	"openiam/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// MenuList 获取菜单列表
func MenuList(c *gin.Context) {
	var (
		err      error
		menuList []*models.Menu
		result   interface{}
	)

	dbConn := db.Orm().Model(&models.Menu{})

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &menuList)
	if err != nil {
		response.Error(c, err, respstatus.GetMenuListError)
		return
	}

	response.OK(c, result, "")
}

func MenuTree(c *gin.Context) {
	var (
		err    error
		result []*models.MenuTree
	)

	result, err = server.MenuTree(true)
	if err != nil {
		response.Error(c, err, respstatus.GetMenuTreeError)
		return
	}

	response.OK(c, result, "")
}

// CreateMenu 创建菜单
func CreateMenu(c *gin.Context) {
	var (
		err   error
		menu  models.Menu
		count int64
	)

	err = c.ShouldBindJSON(&menu)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// path 和 name 不能重复
	err = db.Orm().Model(&models.Menu{}).
		Where("path = ? OR name = ?", menu.Path, menu.Name).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuError)
		return
	}

	if count > 0 {
		response.Error(c, fmt.Errorf("path already exists"), respstatus.PathAlreadyExistsError)
		return
	}

	err = db.Orm().Create(&menu).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateMenuError)
		return
	}

	response.OK(c, menu, "")
}

// MenuDetailByMenuId 获取菜单详情
func MenuDetailByMenuId(c *gin.Context) {
	var (
		err  error
		id   = c.Param("id")
		menu models.Menu
	)

	err = db.Orm().Model(&models.Menu{}).
		Where("id = ?", id).
		First(&menu).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuDetailsError)
		return
	}

	response.OK(c, menu, "")
}

// UpdateMenu 更新菜单
func UpdateMenu(c *gin.Context) {
	var (
		err   error
		id    = c.Param("id")
		menu  models.Menu
		count int64
	)

	err = c.ShouldBindJSON(&menu)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// path 和 name 不能重复, 排除自己
	err = db.Orm().Model(&models.Menu{}).
		Where("(path =? OR name =?) AND id !=?", menu.Path, menu.Name, id).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetMenuError)
		return
	}

	if count > 0 {
		response.Error(c, fmt.Errorf("path already exists"), respstatus.PathAlreadyExistsError)
		return
	}

	err = db.Orm().Model(&models.Menu{}).
		Where("id = ?", id).
		Updates(&menu).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateMenuError)
		return
	}

	response.OK(c, menu, "")
}

// DeleteMenu 删除菜单
func DeleteMenu(c *gin.Context) {
	var (
		err error
		id  = c.Param("id")
	)

	err = db.Orm().Model(&models.Menu{}).
		Where("id = ?", id).
		Delete(&models.Menu{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteMenuError)
		return
	}

	response.OK(c, "", "")
}
