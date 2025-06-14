package models

import "openops/app/resource/models"

var ResourceModels = []interface{}{
	&models.ModelGroup{},
	&models.Model{},
	&models.Field{},
	&models.FieldGroup{},
	&models.ModelRelation{},
	&models.ModelUnique{},
	&models.Data{},
	&models.CloudAccount{},
	&models.CloudRegion{},
	&models.LogicResource{},
	&models.LogicHandle{},
	&models.CloudModels{},
}
