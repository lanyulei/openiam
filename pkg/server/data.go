package server

import (
	"encoding/json"
	"fmt"
	"openops/app/resource/models"

	"github.com/lanyulei/toolkit/db"
)

func VerifyData(data *models.Data) (ok bool, err error) {
	var (
		fieldList []*models.Field
		fieldMaps = make(map[string]*models.Field)
		dataMaps  map[string]interface{}
	)

	// 1. 查询模型 ID 对应的字段列表
	err = db.Orm().Model(&models.Field{}).Where("model_id = ?", data.ModelId).Find(&fieldList).Error
	if err != nil {
		return
	}

	for _, field := range fieldList {
		fieldMaps[field.Key] = field
	}

	err = json.Unmarshal(data.Data, &dataMaps)
	if err != nil {
		return
	}

	for key, value := range dataMaps {
		fmt.Println(key, value)
	}

	return
}
