package models

import "openiam/app/system/models"

var SystemModels = []interface{}{
	&models.Migrate{},
	&models.Token{},
	&models.User{},
}
