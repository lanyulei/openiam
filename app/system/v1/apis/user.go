package apis

import (
	"fmt"
	"openiam/app/system/models"
	"openiam/common/middleware/permission"
	commonModels "openiam/common/models"
	"openiam/pkg/tools/common"
	"openiam/pkg/tools/respstatus"
	"openiam/server/system"
	password2 "openiam/server/system/password"
	"strings"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
	"golang.org/x/crypto/bcrypt"
)

// UserList 用户列表
func UserList(c *gin.Context) {
	var (
		err      error
		userList []*struct {
			models.User
			DepartmentName string `json:"department_name"`
		}
		result interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	systemUserField := `"system_user"`
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		systemUserField = "`system_user`"
	}
	dbConn := db.Orm().Model(&models.User{}).
		Select(fmt.Sprintf("system_department.name as department_name, %s.*", systemUserField)).
		Joins(common.AddQuotesToSQLTableNames("left join system_department on system_department.id = system_user.department"))

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &userList, SearchParams)
	if err != nil {
		response.Error(c, err, respstatus.UserListError)
		return
	}

	response.OK(c, result, "")
}

// UserInfo 用户详情
func UserInfo(c *gin.Context) {
	var (
		err  error
		user struct {
			models.User
			DepartmentInfo models.Department `gorm:"-" json:"department_info"`
			Page           []string          `gorm:"-" json:"page"`
			Button         []string          `gorm:"-" json:"button"`
		}
		groups     [][]string
		roleIds    []int
		department models.Department
	)

	err = db.Orm().Model(&models.User{}).Where("username = ?", c.GetString("username")).Scan(&user).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserInfoError)
		return
	}

	// 获取部门信息
	err = db.Orm().Model(&models.Department{}).Where("id = ?", user.Department).Find(&department).Error
	if err != nil {
		response.Error(c, err, respstatus.GetDepartmentError)
		return
	}
	user.DepartmentInfo = department

	groups, err = permission.Enforcer().GetFilteredNamedGroupingPolicy("g", 0, user.Username)
	if err != nil {
		response.Error(c, err, respstatus.GetRoleError)
		return
	}

	if len(groups) > 0 {
		roles := make([]string, 0, len(groups))
		for _, g := range groups {
			roles = append(roles, g[1])
		}
		user.Role = roles
	}

	if !user.IsAdmin {
		keyValue := "\"key\""
		if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
			keyValue = "`key`"
		}

		// 查询角色ID
		err = db.Orm().Model(&models.Role{}).Where("? in (?)", keyValue, user.Role).Pluck("id", &roleIds).Error
		if err != nil {
			response.Error(c, err, respstatus.GetRoleError)
			return
		}

		// 查询菜单权限
		err = db.Orm().Model(&models.RoleMenu{}).
			Joins(common.AddQuotesToSQLTableNames("left join system_menu on system_menu.id = system_role_menu.menu")).
			Select("distinct UNNEST(system_menu.auth) as auth").
			Where(`system_role_menu.role in (?) and system_role_menu."type" = 1`, roleIds).
			Pluck("auth", &user.Page).
			Error
		if err != nil {
			response.Error(c, err, respstatus.GetRoleMenuError)
			return
		}

		// 查询按钮权限
		err = db.Orm().Model(&models.RoleMenu{}).
			Joins(common.AddQuotesToSQLTableNames("left join system_menu on system_menu.id = system_role_menu.menu")).
			Select("distinct UNNEST(system_menu.auth) as auth").
			Where(`system_role_menu.role in (?) and system_role_menu."type" = 2`, roleIds).
			Pluck("auth", &user.Button).
			Error
		if err != nil {
			response.Error(c, err, respstatus.GetRoleMenuError)
			return
		}
	} else {
		user.Page = []string{}
		user.Button = []string{}
	}
	response.OK(c, user, "")
}

// UserInfoById 通过ID获取用户详情
func UserInfoById(c *gin.Context) {
	var (
		err    error
		user   models.User
		groups [][]string
		userId string
	)

	userId = c.Param("id")

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).Find(&user).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserInfoError)
		return
	}

	groups, err = permission.Enforcer().GetFilteredNamedGroupingPolicy("g", 0, user.Username)
	if err != nil {
		response.Error(c, err, respstatus.GetRoleError)
		return
	}

	if len(groups) > 0 {
		roles := make([]string, 0, len(groups))
		for _, g := range groups {
			roles = append(roles, g[1])
		}
		user.Role = roles
	}

	response.OK(c, user, "")
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var (
		err       error
		user      models.UserRequest
		userCount int64
		groups    [][]string
	)

	err = c.ShouldBind(&user)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 判断用户是否存在
	err = db.Orm().Model(&models.User{}).Where("username = ?", user.Username).Count(&userCount).Error
	if err != nil {
		response.Error(c, err, respstatus.QueryUserError)
		return
	}
	if userCount > 0 {
		response.Error(c, err, respstatus.UserExistError)
		return
	}

	user.Password, err = password2.EncryptionPassword(user.Password)
	if err != nil {
		response.Error(c, err, respstatus.EncryptPasswordError)
		return
	}

	// 创建用户
	tx := db.Orm().Begin()
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		response.Error(c, err, respstatus.CreateUserError)
		return
	}

	if len(user.Role) > 0 {
		for _, role := range user.Role {
			groups = append(groups, []string{user.Username, role})
		}
		_, err = permission.Enforcer().AddNamedGroupingPolicies("g", groups)
		if err != nil {
			tx.Rollback()
			response.Error(c, err, respstatus.CreateUserRoleError)
			return
		}

		err = permission.Sync()
		if err != nil {
			tx.Rollback()
			response.Error(c, err, respstatus.UpdatePermissionCacheError)
			return
		}
	}

	tx.Commit()

	response.OK(c, "", "")
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	var (
		err           error
		userId        string
		user          models.User
		groups        [][]string
		currentGroups [][]string
		groupMap      map[string]struct{}
		deleteGroups  [][]string
		createGroups  [][]string
	)

	userId = c.Param("id")

	err = c.ShouldBind(&user)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 更新用户
	tx := db.Orm().Begin()
	err = tx.Model(&models.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"username":   user.Username,
		"nickname":   user.Nickname,
		"avatar":     user.Avatar,
		"department": user.Department,
		"tel":        user.Tel,
		"email":      user.Email,
		"status":     user.Status,
		"is_admin":   user.IsAdmin,
		"remark":     user.Remark,
	}).Error
	if err != nil {
		tx.Rollback()
		response.Error(c, err, respstatus.UpdateUserError)
		return
	}

	groupMap = make(map[string]struct{})
	for _, role := range user.Role {
		groups = append(groups, []string{user.Username, role})
		groupMap[role] = struct{}{}
	}

	// 查询现有的角色关联
	deleteGroups = make([][]string, 0)
	currentGroups, err = permission.Enforcer().GetFilteredNamedGroupingPolicy("g", 0, user.Username)
	if err != nil {
		tx.Rollback()
		response.Error(c, err, respstatus.GetRoleError)
		return
	}

	for _, g := range currentGroups {
		if _, ok := groupMap[g[1]]; !ok {
			deleteGroups = append(deleteGroups, g)
			delete(groupMap, g[1])
		}
	}

	if len(deleteGroups) > 0 {
		// 删除用户角色关联
		_, err = permission.Enforcer().RemoveNamedGroupingPolicies("g", deleteGroups)
		if err != nil {
			tx.Rollback()
			response.Error(c, err, respstatus.DeleteUserRoleError)
			return
		}

		err = permission.Sync()
		if err != nil {
			tx.Rollback()
			response.Error(c, err, respstatus.UpdatePermissionCacheError)
			return
		}
	}

	if len(groupMap) > 0 {
		createGroups = make([][]string, 0)
		for k, _ := range groupMap {
			createGroups = append(deleteGroups, []string{user.Username, k})
		}

		// 保存用户角色关联
		_, err = permission.Enforcer().AddNamedGroupingPolicies("g", createGroups)
		if err != nil {
			tx.Rollback()
			response.Error(c, err, respstatus.CreateUserRoleError)
			return
		}

		err = permission.Sync()
		if err != nil {
			tx.Rollback()
			response.Error(c, err, respstatus.UpdatePermissionCacheError)
			return
		}
	}
	tx.Commit()

	response.OK(c, "", "")
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	var (
		err    error
		userId string
		user   models.User
	)

	userId = c.Param("id")

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).Find(&user).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserInfoError)
		return
	}

	groups, err := permission.Enforcer().GetFilteredNamedGroupingPolicy("g", 0, user.Username)
	if err != nil {
		response.Error(c, err, respstatus.GetRoleError)
		return
	}

	if len(groups) > 0 {
		response.Error(c, err, respstatus.UserRoleError)
		return
	}

	err = db.Orm().Delete(&models.User{}, userId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteUserError)
		return
	}

	response.OK(c, "", "")
}

// UpdateUserAvatar 更新用户头像
func UpdateUserAvatar(c *gin.Context) {
	var (
		err    error
		avatar struct {
			Avatar string `json:"avatar"`
		}
		userId string
	)

	userId = c.Param("id")

	err = c.ShouldBind(&avatar)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).Update("avatar", avatar.Avatar).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateUserError)
		return
	}

	response.OK(c, "", "")
}

// UpdateUserInfo
// @Description: 更新用户基本信息
// @param c
func UpdateUserInfo(c *gin.Context) {
	var (
		err    error
		userId int
		user   models.User
	)

	userId = c.GetInt("userId")

	err = c.ShouldBind(&user)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"email":    user.Email,
		"tel":      user.Tel,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateUserError)
		return
	}

	response.OK(c, "", "")
}

func InitPassword(c *gin.Context) {
	var (
		err      error
		password []byte
		userId   string
	)

	userId = c.Param("id")

	// 加密密码
	password, err = bcrypt.GenerateFromPassword([]byte(system.InitPassword), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, err, respstatus.EncryptPasswordError)
		return
	}

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).Update("password", string(password)).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateUserError)
		return
	}

	response.OK(c, "", "")
}

func UpdatePassword(c *gin.Context) {
	var (
		err         error
		userInfo    models.User
		oldPassword string
		params      struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}
	)

	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 获取当前用户信息
	err = db.Orm().Model(&models.User{}).Where("id = ?", c.GetInt("userId")).Find(&userInfo).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserInfoError)
		return
	}

	// 解码密码
	oldPassword, err = password2.DecodePassword(params.OldPassword)
	if err != nil {
		response.Error(c, err, respstatus.DecodePasswordError)
		return
	}

	// 验证旧密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(oldPassword))
	if err != nil {
		response.Error(c, err, respstatus.IncorrectPasswordError)
		return
	}

	// 加密密码
	params.NewPassword, err = password2.EncryptionPassword(params.NewPassword)
	if err != nil {
		response.Error(c, err, respstatus.EncryptPasswordError)
		return
	}

	err = db.Orm().Model(&models.User{}).Where("id = ?", c.GetInt("userId")).Update("password", params.NewPassword).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateUserError)
		return
	}

	response.OK(c, "", "")
}

func UserListById(c *gin.Context) {
	var (
		err   error
		query struct {
			Username string `form:"username"`
			Nickname string `form:"nickname"`
			Ids      string `form:"ids"`
		}
		ids   []string
		users []models.User
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	ids = strings.Split(query.Ids, ",")

	dbConn := db.Orm().Model(&models.User{})
	if len(ids) > 0 && query.Ids != "" {
		dbConn = dbConn.Where("id in (?)", ids)
	}

	if query.Username != "" {
		dbConn = dbConn.Where("username like ?", "%"+query.Username+"%")
	}

	if query.Nickname != "" {
		dbConn = dbConn.Where("nickname like ?", "%"+query.Nickname+"%")
	}

	err = dbConn.Order("id desc").Find(&users).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserError)
		return
	}

	response.OK(c, users, "")
}

func UserListByGroupId(c *gin.Context) {
	var (
		err   error
		query struct {
			Username string `form:"username"`
			Nickname string `form:"nickname"`
			GroupId  string `form:"group_id" binding:"required"`
		}
		userIds []int
		users   []models.User
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	groupIds := strings.Split(query.GroupId, ",")

	err = db.Orm().Model(&models.UserGroupRelated{}).
		Select("distinct user_id as user_id").
		Where("group_id in (?)", groupIds).
		Pluck("user_id", &userIds).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserError)
		return
	}

	if len(userIds) > 0 {
		dbConn := db.Orm().Model(&models.User{}).Where("id in (?)", userIds)

		if query.Username != "" {
			dbConn = dbConn.Where("username like ?", "%"+query.Username+"%")
		}

		if query.Nickname != "" {
			dbConn = dbConn.Where("nickname like ?", "%"+query.Nickname+"%")
		}

		err = dbConn.Order("id desc").Find(&users).Error
		if err != nil {
			response.Error(c, err, respstatus.GetUserError)
			return
		}
	}

	response.OK(c, users, "")
}

func UserInfoByUsername(c *gin.Context) {
	var (
		err      error
		userInfo models.UserRequest
		username string
	)

	username = c.Param("username")

	err = db.Orm().Model(&models.User{}).Where("username = ?", username).Find(&userInfo).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserInfoError)
		return
	}

	response.OK(c, userInfo, "")
}
