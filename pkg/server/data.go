package server

import (
	"encoding/json"
	"fmt"
	"openops/app/resource/models"
	"regexp"
	"time"

	"github.com/lanyulei/toolkit/db"
)

func VerifyData(data *models.Data) (err error) {
	var (
		fieldList           []*models.Field
		fieldMap            = make(map[string]*models.Field)
		dataMap, oldDataMap map[string]interface{}
		oldData             *models.Data
	)

	err = db.Orm().Model(&models.Field{}).Where("model_id = ?", data.ModelId).Find(&fieldList).Error
	if err != nil {
		return
	}

	// data.Id 不等于空则更新，反之则创建
	if data.Id != "" {
		err = db.Orm().Model(&models.Data{}).Where("id = ?", data.Id).Find(&oldData).Error
		if err != nil {
			return
		}

		if oldData != nil {
			err = json.Unmarshal(oldData.Data, &oldDataMap)
			if err != nil {
				return
			}
		}
	}

	for _, field := range fieldList {
		fieldMap[field.Key] = field
	}

	err = json.Unmarshal(data.Data, &dataMap)
	if err != nil {
		return
	}

	for key, value := range dataMap { // dataMap 为新的数据
		field, exists := fieldMap[key]
		if !exists {
			err = fmt.Errorf("field %s does not exist in model %s", key, data.ModelId)
			return
		}

		if field.IsRequired && value == nil {
			err = fmt.Errorf("field %s is required", key)
			return
		}

		//FieldTypeShortString FieldType = "shortString"
		//FieldTypeNumber      FieldType = "number"
		//FieldTypeFloat       FieldType = "float"
		//FieldTypeEnum        FieldType = "enum"
		//FieldTypeEnumMulti   FieldType = "enumMulti"
		//FieldTypeDate        FieldType = "date"
		//FieldTypeTime        FieldType = "time"
		//FieldTypeDateTime    FieldType = "dateTime"
		//FieldTypeLongString  FieldType = "longString"
		//FieldTypeUser        FieldType = "user"
		//FieldTypeTimeZone    FieldType = "timeZone"
		//FieldTypeBoolean     FieldType = "boolean"
		//FieldTypeList        FieldType = "list"
		//FieldTypeTable       FieldType = "table"

		switch field.Type {
		case models.FieldTypeShortString, models.FieldTypeLongString:
			if field.IsRequired && value.(string) == "" {
				err = fmt.Errorf("field %s is required", key)
				return
			}

			if data.Id != "" { // data.Id 不等于空则是更新
				if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
					if value.(string) != oldDataMap[key].(string) {
						err = fmt.Errorf("field %s cannot be edited", key)
						return
					}
				}
			}

			if len(field.Options) > 0 {
				var options models.StringOptions
				err = json.Unmarshal(field.Options, &options)
				if err != nil {
					return
				}

				regexpValue := options.Regexp
				if regexpValue != "" {
					re, reErr := regexp.Compile(regexpValue)
					if reErr != nil {
						err = fmt.Errorf("invalid regexp for field %s: %v", key, reErr)
						return
					}
					if !re.MatchString(value.(string)) {
						err = fmt.Errorf("field %s does not match regexp", key)
						return
					}
				}
			}
		case models.FieldTypeNumber, models.FieldTypeFloat:
			if data.Id != "" { // data.Id 不等于空则是更新
				if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
					var newValue, oldValue float64
					switch v := value.(type) {
					case float64:
						newValue = v
					case int:
						newValue = float64(v)
					}
					switch v := oldDataMap[key].(type) {
					case float64:
						oldValue = v
					case int:
						oldValue = float64(v)
					}
					if newValue != oldValue {
						err = fmt.Errorf("field %s cannot be edited", key)
						return
					}
				}
			}

			if len(field.Options) > 0 {
				options := make(map[string]float64)
				err = json.Unmarshal(field.Options, &options)
				if err != nil {
					return
				}

				var v float64
				switch val := value.(type) {
				case float64:
					v = val
				case int:
					v = float64(val)
				}

				minValue, minOk := options["min"]
				if minOk && v < minValue {
					err = fmt.Errorf("field %s must be greater than or equal to %v", key, minValue)
					return
				}

				maxValue, maxOk := options["max"]
				if maxOk && v > maxValue {
					err = fmt.Errorf("field %s must be less than or equal to %v", key, maxValue)
					return
				}
			}
		case models.FieldTypeEnum, models.FieldTypeEnumMulti:
			if len(field.Options) > 0 {
				var (
					options       models.EnumOptions
					enumOptionMap = make(map[string]struct{})
				)
				err = json.Unmarshal(field.Options, &options)
				if err != nil {
					return
				}

				for _, v := range options.Options {
					enumOptionMap[v.ID] = struct{}{}
				}

				if field.Type == models.FieldTypeEnum {
					if field.IsRequired && value.(string) == "" {
						err = fmt.Errorf("field %s is required", key)
						return
					}

					if data.Id != "" { // data.Id 不等于空则是更新
						if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
							if value.(string) != oldDataMap[key].(string) {
								err = fmt.Errorf("field %s cannot be edited", key)
								return
							}
						}
					}

					if _, ok := enumOptionMap[value.(string)]; !ok {
						err = fmt.Errorf("field %s value %s is not in options", key, value)
						return
					}
				} else if field.Type == models.FieldTypeEnumMulti {
					values, ok := value.([]interface{})
					if !ok {
						err = fmt.Errorf("field %s value must be an array for enumMulti type", key)
						return
					}

					if field.IsRequired && len(values) == 0 {
						err = fmt.Errorf("field %s is required", key)
						return
					}

					if data.Id != "" { // data.Id 不等于空则是更新
						if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
							oldValues, ok := oldDataMap[key].([]interface{})
							if !ok {
								err = fmt.Errorf("field %s old value must be an array for enumMulti type", key)
								return
							}

							if len(values) != len(oldValues) {
								err = fmt.Errorf("field %s cannot be edited, length mismatch", key)
								return
							}

							valueMap := make(map[string]struct{})
							for _, v := range values {
								valueMap[v.(string)] = struct{}{}
							}

							for _, v := range oldValues {
								if _, ok := valueMap[v.(string)]; !ok {
									err = fmt.Errorf("field %s cannot be edited, value mismatch", key)
									return
								}
							}
						}
					}

					for _, v := range values {
						if _, ok := enumOptionMap[v.(string)]; !ok {
							err = fmt.Errorf("field %s value %s is not in options", key, v.(string))
							return
						}
					}
				}
			}
		case models.FieldTypeDate, models.FieldTypeTime, models.FieldTypeDateTime:
			if field.IsRequired && value.(string) == "" {
				err = fmt.Errorf("field %s is required", key)
				return
			}

			if data.Id != "" { // data.Id 不等于空则是更新
				if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
					if value.(string) != oldDataMap[key].(string) {
						err = fmt.Errorf("field %s cannot be edited", key)
						return
					}
				}
			}

			switch field.Type {
			case models.FieldTypeDate:
				// 校验是不是日期格式
				_, err = time.Parse("2006-01-02", value.(string))
				if err != nil {
					err = fmt.Errorf("field %s value %s is not a valid date", key, value)
					return
				}
			case models.FieldTypeTime:
				// 校验是不是时间格式
				_, err = time.Parse("15:04:05", value.(string))
				if err != nil {
					err = fmt.Errorf("field %s value %s is not a valid time", key, value)
					return
				}
			case models.FieldTypeDateTime:
				// 校验是不是日期时间格式
				_, err = time.Parse("2006-01-02 15:04:05", value.(string))
				if err != nil {
					err = fmt.Errorf("field %s value %s is not a valid dateTime", key, value)
					return
				}
			}
		case models.FieldTypeUser, models.FieldTypeTimeZone:
			if field.IsRequired && value.(string) == "" {
				err = fmt.Errorf("field %s is required", key)
				return
			}

			if data.Id != "" { // data.Id 不等于空则是更新
				if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
					if value.(string) != oldDataMap[key].(string) {
						err = fmt.Errorf("field %s cannot be edited", key)
						return
					}
				}
			}
		case models.FieldTypeBoolean:
			if data.Id != "" { // data.Id 不等于空则是更新
				if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
					if value.(bool) != oldDataMap[key].(bool) {
						err = fmt.Errorf("field %s cannot be edited", key)
						return
					}
				}
			}
		case models.FieldTypeList:
			if field.IsRequired && len(value.([]interface{})) == 0 {
				err = fmt.Errorf("field %s is required", key)
				return
			}

			if data.Id != "" { // data.Id 不等于空则是更新
				if !field.IsEdit { // 如果字段不可编辑，则需要验证旧数据
					oldValues, ok := oldDataMap[key].([]interface{})
					if !ok {
						err = fmt.Errorf("field %s old value must be an array for list type", key)
						return
					}

					if len(value.([]interface{})) != len(oldValues) {
						err = fmt.Errorf("field %s cannot be edited, length mismatch", key)
						return
					}

					valueMap := make(map[string]struct{})
					for _, v := range value.([]interface{}) {
						valueMap[v.(string)] = struct{}{}
					}

					for _, v := range oldValues {
						if _, ok := valueMap[v.(string)]; !ok {
							err = fmt.Errorf("field %s cannot be edited, value mismatch", key)
							return
						}
					}
				}
			}
		case models.FieldTypeTable:
			// 暂无需进行数据校验
		}
	}

	return
}

func VerifyDataHandler() {

}
