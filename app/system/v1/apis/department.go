package apis

import (
	"errors"
	"fmt"
	"openiam/app/system/models"
	commonModels "openiam/common/models"
	"openiam/pkg/tools/common"
	"openiam/pkg/tools/respstatus"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

func DepartmentList(c *gin.Context) {
	var (
		err            error
		departmentList []*struct {
			models.Department
			LeaderUsername string `json:"leader_username"`
			LeaderNickname string `json:"leader_nickname"`
		}
		result interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	systemDepartmentField := `"system_department"`
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		systemDepartmentField = "`system_department`"
	}

	systemUserField := `"system_user"`
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		systemUserField = "`system_user`"
	}

	dbConn := db.Orm().Model(&models.Department{}).
		Select(fmt.Sprintf("%s.username as leader_username, %s.nickname as leader_nickname, %s.*", systemUserField, systemUserField, systemDepartmentField)).
		Joins(common.AddQuotesToSQLTableNames("left join system_user on system_user.id = system_department.leader"))

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &departmentList, SearchParams)
	if err != nil {
		response.Error(c, err, respstatus.GetDepartmentListError)
		return
	}

	response.OK(c, result, "")
}

func CreateDepartment(c *gin.Context) {
	var (
		err        error
		department models.Department
		count      int64
	)

	err = c.ShouldBindJSON(&department)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	err = db.Orm().Model(&models.Department{}).Where("name = ? and parent = ?", department.Name, department.Parent).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetDepartmentListError)
		return
	}
	if count > 0 {
		response.Error(c, errors.New("部门名称已存在"), respstatus.CreateDepartmentError)
		return
	}

	err = db.Orm().Create(&department).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateDepartmentError)
		return
	}

	response.OK(c, department, "")
}

func UpdateDepartment(c *gin.Context) {
	var (
		err          error
		department   models.Department
		departmentId = c.Param("id")
		count        int64
	)

	err = c.ShouldBindJSON(&department)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	err = db.Orm().Model(&models.Department{}).Where("name = ? and parent = ? and id != ?", department.Name, department.Parent, departmentId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetDepartmentListError)
		return
	}
	if count > 0 {
		response.Error(c, errors.New("部门名称已存在"), respstatus.UpdateDepartmentError)
		return
	}

	err = db.Orm().Model(&models.Department{}).Where("id = ?", departmentId).Updates(map[string]interface{}{
		"name":    department.Name,
		"leader":  department.Leader,
		"parent":  department.Parent,
		"remarks": department.Remarks,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateDepartmentError)
		return
	}

	response.OK(c, department, "")
}

func DeleteDepartment(c *gin.Context) {
	var (
		err          error
		departmentId = c.Param("id")
		count        int64
	)

	// 若有用户绑定部门，则部门无法删除
	err = db.Orm().Model(&models.User{}).Where("department = ?", departmentId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserError)
		return
	}
	if count > 0 {
		response.Error(c, errors.New("已有用户绑定部门，无法直接删除"), respstatus.DeleteDepartmentError)
		return
	}

	err = db.Orm().Where("id = ?", departmentId).Delete(&models.Department{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteDepartmentError)
		return
	}

	response.OK(c, departmentId, "")
}

func DepartmentTree(c *gin.Context) {
	var (
		err          error
		treeList     []*models.DepartmentTree
		topValueList []*models.DepartmentTree
		temp         = make(map[int]*models.DepartmentTree)
		result       = make(map[int]*models.DepartmentTree)
	)

	systemDepartmentField := `"system_department"`
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		systemDepartmentField = "`system_department`"
	}

	systemUserField := `"system_user"`
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		systemUserField = "`system_user`"
	}

	dbConn := db.Orm().Model(&models.Department{}).
		Select(fmt.Sprintf("%s.username as leader_username, %s.nickname as leader_nickname, %s.*", systemUserField, systemUserField, systemDepartmentField)).
		Joins(common.AddQuotesToSQLTableNames("left join system_user on system_user.id = system_department.leader"))

	// 查询所有部门
	err = dbConn.Find(&treeList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetDepartmentListError)
		return
	}

	// 查询顶级部门列表
	err = dbConn.Where("parent = 0").Find(&topValueList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetDepartmentListError)
		return
	}

	for _, v := range treeList {
		temp[v.Id] = v
	}

	for _, node := range temp {
		if temp[node.Parent] == nil {
			result[node.Id] = node
		} else {
			temp[node.Parent].Children = append(temp[node.Parent].Children, node)
		}
	}

	for _, v := range topValueList {
		v.Children = result[v.Id].Children
	}

	response.OK(c, topValueList, "")
}
