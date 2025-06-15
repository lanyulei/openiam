package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

func GetRegions(c *gin.Context) {
	var (
		err     error
		regions []models.CloudRegion
		query   struct {
			CloudAccountId int    `form:"cloud_account_id" binding:"required"`
			RegionId       string `form:"region_id"`
			Name           string `form:"name"`
		}
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	dbConn := db.Orm().Model(&models.CloudRegion{}).Where("cloud_account_id = ?", query.CloudAccountId)

	if query.Name != "" {
		dbConn = dbConn.Where("name like ?", "%"+query.Name+"%")
	}

	if query.RegionId != "" {
		dbConn = dbConn.Where("region_id like ?", "%"+query.RegionId+"%")
	}

	err = dbConn.Find(&regions).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudRegionError)
		return
	}

	response.OK(c, regions, "")
}

func CreateRegion(c *gin.Context) {
	var (
		err    error
		region models.CloudRegion
		count  int64
	)

	err = c.ShouldBindJSON(&region)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.CloudRegion{}).
		Where("cloud_account_id = ? AND region_id = ? AND name = ?",
			region.CloudAccountId,
			region.RegionId,
			region.Name,
		).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudRegionError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.CloudRegionExistError)
		return
	}

	err = db.Orm().Create(&region).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateCloudRegionError)
		return
	}

	response.OK(c, "", "")
}

func UpdateRegion(c *gin.Context) {
	var (
		err      error
		region   models.CloudRegion
		regionId = c.Param("id")
		count    int64
	)

	err = c.ShouldBindJSON(&region)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.CloudRegion{}).
		Where("cloud_account_id = ? AND region_id = ? AND name = ? AND id != ?",
			region.CloudAccountId,
			region.RegionId,
			region.Name,
			regionId,
		).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudRegionError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.CloudRegionExistError)
		return
	}

	err = db.Orm().Model(&models.CloudRegion{}).Where("id = ?", regionId).Updates(&region).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateCloudRegionError)
		return
	}

	response.OK(c, "", "")
}

func DeleteRegion(c *gin.Context) {
	regionId := c.Param("id")

	err := db.Orm().Delete(&models.CloudRegion{}, "id = ?", regionId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteCloudRegionError)
		return
	}

	response.OK(c, "", "")
}
