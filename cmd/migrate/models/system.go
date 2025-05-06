package models

import "openiam/app/system/models"

var SystemModels = []interface{}{
	&models.User{},
	&models.Department{},
	&models.Role{},
	&models.Menu{},
	&models.RoleMenu{},
	&models.Api{},
	&models.MenuApi{},
	&models.ApiGroup{},
	&models.LoginLog{},
	&models.Settings{},
	&models.Audit{},
	&models.Migrate{},
	&models.App{},
	&models.AppGroup{},
	&models.UserGroup{},
	&models.UserGroupRelated{},
}
