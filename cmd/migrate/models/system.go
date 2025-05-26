package models

import "openops/app/system/models"

var SystemModels = []interface{}{
	&models.Migrate{},
	&models.Token{},
	&models.User{},
	&models.Menu{},
}
