package route

import (
	"openiam/app/system/models"

	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CheckRegisterRoute(routes []*Route) (result map[string][]*Route, err error) {
	var (
		apis, createApis []*models.Api
		apiMap           = make(map[string]struct{})
	)

	result = make(map[string][]*Route)

	err = db.Orm().Model(&models.Api{}).Find(&apis).Error
	if err != nil {
		logger.Errorf("failed to get interface list, error: %s", err.Error())
		return
	}

	// get unregistered routes
	for _, a := range apis {
		apiMap[a.Method+a.URL] = struct{}{}
	}
	for _, r := range routes {
		if _, ok := apiMap[r.Method+r.Path]; !ok {
			result[Unregistered] = append(result[Unregistered], r)
			createApis = append(createApis, &models.Api{
				Method: r.Method,
				URL:    r.Path,
			})
		}
	}

	// create unregistered routes
	if len(createApis) > 0 {
		err = db.Orm().Model(&models.Api{}).Create(&createApis).Error
		if err != nil {
			logger.Errorf("failed to create unregistered routes, error: %s", err.Error())
			for _, v := range createApis {
				logger.Warnf("unregistered route: %s %s", v.Method, v.URL)
			}
			return
		}
	}

	return
}
